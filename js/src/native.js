import { spawn } from "node:child_process";
import { access } from "node:fs/promises";
import path from "node:path";
import { fileURLToPath } from "node:url";
import { constants as fsConstants } from "node:fs";

const __dirname = path.dirname(fileURLToPath(import.meta.url));

function platformName() {
  const os = process.platform;
  const arch = process.arch === "x64" ? "amd64" : process.arch;
  return `lucy-${os}-${arch}${os === "win32" ? ".exe" : ""}`;
}

export async function resolveLucyBinary() {
  if (process.env.LUCY_BIN) return process.env.LUCY_BIN;
  const candidates = [
    path.resolve(__dirname, "../../artifacts", platformName()),
    path.resolve(__dirname, "../../python/src/lucy/bin/lucy"),
    path.resolve(__dirname, "../../python/src/lucy/bin", platformName().replace("win32", "windows")),
  ];
  for (const c of candidates) {
    try {
      await access(c, fsConstants.X_OK);
      return c;
    } catch {
      /* try next */
    }
  }
  throw new Error("lucy binary not found; set LUCY_BIN or run go/scripts/build-artifacts.sh");
}

/** buildLPD via native Go CLI (Node/Bun). Same JSON schema as wasm. */
function wireSample(s) {
  return {
    tide: s.tide,
    id: s.id,
    mode: s.mode ?? "",
    dtype: s.dtype ?? s.dType ?? "",
    format: s.format ?? "",
    arch: s.arch ?? "",
    score: s.score ?? 0,
    soft_acc: s.soft_acc ?? s.soft ?? 0,
    avg_accuracy: s.avg_accuracy ?? s.acc ?? 0,
    throughput: s.throughput ?? s.thru ?? 0,
    availability: s.availability ?? s.avail ?? 0,
    ram_kib: s.ram_kib ?? s.ramKiB ?? 0,
  };
}

export async function buildLPDNative(samples, options) {
  const bin = await resolveLucyBinary();
  const req = JSON.stringify({
    samples: (samples || []).map(wireSample),
    options,
  });
  const out = await new Promise((resolve, reject) => {
    const child = spawn(bin, ["build-lpd"], { stdio: ["pipe", "pipe", "pipe"] });
    let stdout = "";
    let stderr = "";
    child.stdout.on("data", (d) => (stdout += d));
    child.stderr.on("data", (d) => (stderr += d));
    child.on("error", reject);
    child.on("close", (code) => {
      if (code !== 0) reject(new Error(stderr || `lucy exited ${code}`));
      else resolve(stdout);
    });
    child.stdin.write(req);
    child.stdin.end();
  });
  return JSON.parse(out);
}
