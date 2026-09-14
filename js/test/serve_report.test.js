import { test } from "node:test";
import assert from "node:assert/strict";
import { readFile, mkdtemp } from "node:fs/promises";
import { tmpdir } from "node:os";
import path from "node:path";
import { fileURLToPath } from "node:url";
import { chartJPGNative, writeReportNative, createLucyClient } from "../src/index.js";
import { spawn } from "node:child_process";
import { resolveLucyBinary } from "../src/native.js";

const root = path.resolve(path.dirname(fileURLToPath(import.meta.url)), "../..");

test("jpg + report dir", async () => {
  const g = JSON.parse(await readFile(path.join(root, "testdata/goldens_lpd_v0.8.json"), "utf8"));
  const jpg = await chartJPGNative(g.samples, "radar");
  assert.equal(jpg[0], 0xff);
  assert.equal(jpg[1], 0xd8);
  const dir = await mkdtemp(path.join(tmpdir(), "lucy-report-"));
  const out = await writeReportNative(g.samples, dir);
  assert.ok(out.includes(dir) || out === dir);
  await readFile(path.join(dir, "index.html"), "utf8");
});

test("lucy serve /api/lpd", async () => {
  const bin = await resolveLucyBinary();
  const g = JSON.parse(await readFile(path.join(root, "testdata/goldens_lpd_v0.8.json"), "utf8"));
  const child = spawn(bin, ["serve", "127.0.0.1:17474"], { stdio: ["ignore", "ignore", "pipe"] });
  try {
    await new Promise((r) => setTimeout(r, 400));
    const client = createLucyClient("http://127.0.0.1:17474");
    const ver = await client.version();
    assert.ok(ver);
    const resp = await client.buildLPD(g.samples);
    assert.equal(resp.board.top[0].id, "int8");
  } finally {
    child.kill("SIGTERM");
  }
});
