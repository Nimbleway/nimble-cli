#!/usr/bin/env node
// Downloads the nimble binary from GitHub Releases on npm install.
const fs = require("fs/promises");
const path = require("path");
const os = require("os");
const { spawnSync } = require("child_process");

const PLATFORM_MAP = {
  "linux-x64": { target: "linux_amd64", ext: ".tar.gz" },
  "linux-arm64": { target: "linux_arm64", ext: ".tar.gz" },
  "darwin-x64": { target: "macos_amd64", ext: ".zip" },
  "darwin-arm64": { target: "macos_arm64", ext: ".zip" },
  "win32-x64": { target: "windows_amd64", ext: ".zip" },
};

const platform = os.platform();
const arch = os.arch();
const key = `${platform}-${arch}`;
const entry = PLATFORM_MAP[key];

if (!entry) {
  console.error(`nimble-cli: unsupported platform ${key}`);
  process.exit(1);
}

const { version } = require("./package.json");
const { target, ext } = entry;
const isWindows = platform === "win32";
const binaryName = isWindows ? "nimble.exe" : "nimble";
const archive = `nimble_${version}_${target}${ext}`;
const url = `https://github.com/Nimbleway/nimble-cli/releases/download/v${version}/${archive}`;
const binDir = path.join(__dirname, "bin");
const binaryDest = path.join(binDir, binaryName);

async function download(url, dest) {
  const response = await fetch(url);
  if (!response.ok) {
    throw new Error(`HTTP ${response.status} downloading ${url}`);
  }

  return fs.writeFile(dest, Buffer.from(await response.arrayBuffer()));
  // console.log(`nimble-cli: downloaded ${url} to ${dest}`);

  // return new Promise((resolve, reject) => {

    // const file = fs.createWriteStream(dest);
    // fetch(url)
    //   .then((res) => {
    //     if (!res.ok) {
    //       throw new Error(`HTTP ${res.status} downloading ${url}`);
    //     }
    //     return res.body.pipe(file);
    //   })
    //   .then(() => {
    //     console.log(`nimble-cli: downloaded ${url} to ${dest}`);
    //     resolve();
    //   })
    //   .catch((err) => {
    //     fs.unlink(dest, () => {});
    //     console.error(`nimble-cli: error downloading ${url} — ${err.message}`);
    //     reject(err);
    //   });
  // });
}

async function findBinary(dir, name) {
  for (const entry of await fs.readdir(dir, { withFileTypes: true })) {
    const full = path.join(dir, entry.name);
    if (entry.isDirectory()) {
      const found = await findBinary(full, name);
      if (found) return found;
    } else if (entry.name === name) {
      return full;
    }
  }
  return null;
}

async function main() {
  await fs.mkdir(binDir, { recursive: true });

  const tmpArchive = path.join(os.tmpdir(), archive);
  const tmpExtract = path.join(os.tmpdir(), `nimble-extract-${Date.now()}`);
  await fs.mkdir(tmpExtract, { recursive: true });

  console.log(`nimble-cli: downloading ${url}...`);
  await download(url, tmpArchive);
  console.log(`nimble-cli: download complete, extracting...`);
  if (ext === ".tar.gz") {
    const r = spawnSync("tar", ["-xzf", tmpArchive, "-C", tmpExtract], { stdio: "inherit" });
    if (r.status !== 0) throw new Error("tar extraction failed");
  } else {
    if (isWindows) {
      const r = spawnSync(
        "powershell",
        ["-NoProfile", "-Command", `Expand-Archive -Path '${tmpArchive}' -DestinationPath '${tmpExtract}'`],
        { stdio: "inherit" }
      );
      if (r.status !== 0) throw new Error("Expand-Archive failed");
    } else {
      const r = spawnSync("unzip", ["-o", tmpArchive, "-d", tmpExtract], { stdio: "inherit" });
      if (r.status !== 0) throw new Error("unzip failed");
    }
  }

  const found = await findBinary(tmpExtract, binaryName);
  if (!found) throw new Error(`binary ${binaryName} not found in archive`);

  await fs.copyFile(found, binaryDest);
  if (!isWindows) await fs.chmod(binaryDest, 0o755);

  await fs.rm(tmpArchive, { force: true });
  await fs.rm(tmpExtract, { recursive: true, force: true });

  console.log(`nimble-cli: ${binaryName} installed successfully at ${binaryDest}`);
}

main().catch((err) => {
  console.error(`nimble-cli: installation failed — ${err.message}`);
  console.log(`nimble-cli: installation failed — ${err.message}`);
  process.exit(1);
});
