/** Browser: canvas → PNG data URL. */
export function canvasToPNGDataURL(canvas) {
  return canvas.toDataURL("image/png");
}

/** Browser: canvas → PNG Blob. */
export function canvasToPNGBlob(canvas) {
  return new Promise((resolve, reject) => {
    canvas.toBlob((b) => (b ? resolve(b) : reject(new Error("toBlob failed"))), "image/png");
  });
}

/**
 * Node: call native CLI for PNG bytes.
 * kind: radar|scatter|bars
 */
export async function chartPNGNative(samples, kind = "radar", options, spawnFn) {
  const { buildLPDNative } = await import("../native.js");
  // reuse binary path via spawn for chart-*-png
  const { spawn } = await import("node:child_process");
  const { resolveLucyBinary } = await import("../native.js");
  const bin = await resolveLucyBinary();
  const cmd =
    kind === "scatter"
      ? "chart-scatter-png"
      : kind === "bars"
        ? "chart-bars-png"
        : "chart-radar-png";
  const req = JSON.stringify({
    samples,
    options,
  });
  const buf = await new Promise((resolve, reject) => {
    const child = spawn(bin, [cmd], { stdio: ["pipe", "pipe", "pipe"] });
    const chunks = [];
    let stderr = "";
    child.stdout.on("data", (d) => chunks.push(d));
    child.stderr.on("data", (d) => (stderr += d));
    child.on("error", reject);
    child.on("close", (code) => {
      if (code !== 0) reject(new Error(stderr || `lucy exited ${code}`));
      else resolve(Buffer.concat(chunks));
    });
    child.stdin.write(req);
    child.stdin.end();
  });
  void buildLPDNative; // keep import used for tree docs
  void spawnFn;
  return buf;
}

export async function chartPackNative(samples, options) {
  const { spawn } = await import("node:child_process");
  const { resolveLucyBinary } = await import("../native.js");
  const bin = await resolveLucyBinary();
  const req = JSON.stringify({ samples, options });
  const out = await new Promise((resolve, reject) => {
    const child = spawn(bin, ["chart-pack"], { stdio: ["pipe", "pipe", "pipe"] });
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

export async function chartJPGNative(samples, kind = "radar", options) {
  const { spawn } = await import("node:child_process");
  const { resolveLucyBinary } = await import("../native.js");
  const bin = await resolveLucyBinary();
  const cmd =
    kind === "scatter"
      ? "chart-scatter-jpg"
      : kind === "bars"
        ? "chart-bars-jpg"
        : "chart-radar-jpg";
  const req = JSON.stringify({ samples, options });
  return new Promise((resolve, reject) => {
    const child = spawn(bin, [cmd], { stdio: ["pipe", "pipe", "pipe"] });
    const chunks = [];
    let stderr = "";
    child.stdout.on("data", (d) => chunks.push(d));
    child.stderr.on("data", (d) => (stderr += d));
    child.on("error", reject);
    child.on("close", (code) => {
      if (code !== 0) reject(new Error(stderr || `lucy exited ${code}`));
      else resolve(Buffer.concat(chunks));
    });
    child.stdin.write(req);
    child.stdin.end();
  });
}

export async function writeReportNative(samples, outdir, options) {
  const { spawn } = await import("node:child_process");
  const { resolveLucyBinary } = await import("../native.js");
  const bin = await resolveLucyBinary();
  const req = JSON.stringify({ samples, options });
  return new Promise((resolve, reject) => {
    const child = spawn(bin, ["report", outdir], { stdio: ["pipe", "pipe", "pipe"] });
    let stdout = "";
    let stderr = "";
    child.stdout.on("data", (d) => (stdout += d));
    child.stderr.on("data", (d) => (stderr += d));
    child.on("error", reject);
    child.on("close", (code) => {
      if (code !== 0) reject(new Error(stderr || `lucy exited ${code}`));
      else resolve(stdout.trim());
    });
    child.stdin.write(req);
    child.stdin.end();
  });
}

export async function writePDFNative(samples, outPath, options) {
  const { spawn } = await import("node:child_process");
  const { resolveLucyBinary } = await import("../native.js");
  const bin = await resolveLucyBinary();
  const req = JSON.stringify({ samples, options });
  return new Promise((resolve, reject) => {
    const child = spawn(bin, ["pdf", outPath], { stdio: ["pipe", "pipe", "pipe"] });
    let stdout = "";
    let stderr = "";
    child.stdout.on("data", (d) => (stdout += d));
    child.stderr.on("data", (d) => (stderr += d));
    child.on("error", reject);
    child.on("close", (code) => {
      if (code !== 0) reject(new Error(stderr || `lucy exited ${code}`));
      else resolve(stdout.trim() || outPath);
    });
    child.stdin.write(req);
    child.stdin.end();
  });
}

export async function writeCSVNative(samples, outPath, options) {
  const { spawn } = await import("node:child_process");
  const { resolveLucyBinary } = await import("../native.js");
  const { writeFile } = await import("node:fs/promises");
  const bin = await resolveLucyBinary();
  const req = JSON.stringify({ samples, options });
  const csv = await new Promise((resolve, reject) => {
    const child = spawn(bin, ["csv"], { stdio: ["pipe", "pipe", "pipe"] });
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
  if (outPath) await writeFile(outPath, csv);
  return csv;
}

export async function floorsNative() {
  const { spawn } = await import("node:child_process");
  const { resolveLucyBinary } = await import("../native.js");
  const bin = await resolveLucyBinary();
  const out = await new Promise((resolve, reject) => {
    const child = spawn(bin, ["floors"], { stdio: ["ignore", "pipe", "pipe"] });
    let stdout = "";
    let stderr = "";
    child.stdout.on("data", (d) => (stdout += d));
    child.stderr.on("data", (d) => (stderr += d));
    child.on("error", reject);
    child.on("close", (code) => {
      if (code !== 0) reject(new Error(stderr || `lucy exited ${code}`));
      else resolve(stdout);
    });
  });
  return JSON.parse(out);
}
