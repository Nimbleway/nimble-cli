#!/usr/bin/env node
// Downloads the nimble binary from GitHub Releases on npm install.
const https = require("https");
const fs = require("fs");
const path = require("path");
const os = require("os");
const { spawnSync } = require("child_process");

const PLATFORM_MAP = {
  "linux-x64":    { target: "linux_amd64",   ext: ".tar.gz" },
  "linux-arm64":  { target: "linux_arm64",   ext: ".tar.gz" },
  "darwin-x64":   { target: "macos_amd64",   ext: ".zip" },
  "darwin-arm64": { target: "macos_arm64",   ext: ".zip" },
  "win32-x64":    { target: "windows_amd64", ext: ".zip" },
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

function download(url, dest) {
  return new Promise((resolve, reject) => {
    const file = fs.createWriteStream(dest);
    const get = (url) =>
      https.get(url, (res) => {
        if (res.statusCode === 301 || res.statusCode === 302) {
          file.close();
          return get(res.headers.location);
        }
        if (res.statusCode !== 200) {
          return reject(new Error(`HTTP ${res.statusCode} downloading ${url}`));
        }
        res.pipe(file);
        file.on("finish", () => file.close(resolve));
      }).on("error", (err) => {
        fs.unlink(dest, () => {});
        reject(err);
      });
    get(url);
  });
}

function findBinary(dir, name) {
  for (const entry of fs.readdirSync(dir, { withFileTypes: true })) {
    const full = path.join(dir, entry.name);
    if (entry.isDirectory()) {
      const found = findBinary(full, name);
      if (found) return found;
    } else if (entry.name === name) {
      return full;
    }
  }
  return null;
}

async function main() {
  fs.mkdirSync(binDir, { recursive: true });

  const tmpArchive = path.join(os.tmpdir(), archive);
  const tmpExtract = path.join(os.tmpdir(), `nimble-extract-${Date.now()}`);
  fs.mkdirSync(tmpExtract, { recursive: true });

  console.log(`nimble-cli: downloading ${archive}...`);
  await download(url, tmpArchive);

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

  const found = findBinary(tmpExtract, binaryName);
  if (!found) throw new Error(`binary ${binaryName} not found in archive`);

  fs.copyFileSync(found, binaryDest);
  if (!isWindows) fs.chmodSync(binaryDest, 0o755);

  fs.rmSync(tmpArchive, { force: true });
  fs.rmSync(tmpExtract, { recursive: true, force: true });

  console.log(`nimble-cli: ${binaryName} installed successfully`);
}

main().catch((err) => {
  console.error(`nimble-cli: installation failed — ${err.message}`);
  process.exit(1);
});
