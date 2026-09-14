import { test } from "node:test";
import assert from "node:assert/strict";
import { readFile, mkdtemp } from "node:fs/promises";
import { tmpdir } from "node:os";
import path from "node:path";
import { fileURLToPath } from "node:url";
import { writePDFNative, createLucyClient } from "../src/index.js";
import { spawn } from "node:child_process";
import { resolveLucyBinary } from "../src/native.js";

const root = path.resolve(path.dirname(fileURLToPath(import.meta.url)), "../..");

test("writePDFNative produces PDF", async () => {
  const g = JSON.parse(await readFile(path.join(root, "testdata/goldens_lpd_v0.7.json"), "utf8"));
  const dir = await mkdtemp(path.join(tmpdir(), "lucy-pdf-"));
  const out = path.join(dir, "board.pdf");
  await writePDFNative(g.samples, out);
  const buf = await readFile(out);
  assert.ok(buf.toString("utf8", 0, 8).startsWith("%PDF-1."));
});

test("lucy serve /api/pdf", async () => {
  const bin = await resolveLucyBinary();
  const g = JSON.parse(await readFile(path.join(root, "testdata/goldens_lpd_v0.7.json"), "utf8"));
  const child = spawn(bin, ["serve", "127.0.0.1:17475"], { stdio: ["ignore", "ignore", "pipe"] });
  try {
    await new Promise((r) => setTimeout(r, 400));
    const client = createLucyClient("http://127.0.0.1:17475");
    const pdf = await client.pdf(g.samples);
    assert.ok(Buffer.isBuffer(pdf) || pdf instanceof Uint8Array);
    assert.equal(String.fromCharCode(pdf[0], pdf[1], pdf[2], pdf[3]), "%PDF");
  } finally {
    child.kill("SIGTERM");
  }
});
