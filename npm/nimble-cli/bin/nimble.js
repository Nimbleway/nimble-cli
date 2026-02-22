#!/usr/bin/env node
const { execFileSync } = require("child_process");
const path = require("path");
const os = require("os");

const platform = os.platform(); // linux, darwin, win32
const arch = os.arch() === "x64" ? "x64" : "arm64";
const pkg = `nimble-cli-${platform}-${arch}`;

let binary;
try {
  const dir = path.dirname(require.resolve(`${pkg}/package.json`));
  binary = path.join(dir, "bin", platform === "win32" ? "nimble.exe" : "nimble");
} catch {
  console.error(`nimble-cli: no binary found for ${platform}/${arch}`);
  console.error(`  tried package: ${pkg}`);
  process.exit(1);
}

try {
  execFileSync(binary, process.argv.slice(2), { stdio: "inherit" });
} catch (err) {
  process.exit(err.status ?? 1);
}
