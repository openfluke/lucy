import { test } from "node:test";
import assert from "node:assert/strict";
import { readFile, mkdtemp } from "node:fs/promises";
import { tmpdir } from "node:os";
import path from "node:path";
import { fileURLToPath } from "node:url";
import { writeCSVNative, floorsNative, createLucyClient, VERSION } from "../src/index.js";
import { spawn } from "node:child_process";
import { resolveLucyBinary } from "../src/native.js";

const root = path.resolve(path.dirname(fileURLToPath(import.meta.url)), "../..");

test("floorsNative + writeCSVNative", async () => {
  assert.equal(VERSION, "1.0.1");
  const floors = await floorsNative();
  assert.equal(floors.version, "1.0.1");
  assert.equal(floors.keep_floor, 0.7);
  const g = JSON.parse(await readFile(path.join(root, "testdata/goldens_lpd_v1.0.json"), "utf8"));
  const dir = await mkdtemp(path.join(tmpdir(), "lucy-csv-"));
  const out = path.join(dir, "board.csv");
  const csv = await writeCSVNative(g.samples, out);
  assert.ok(csv.startsWith("id,band,lpd,"));
  assert.ok(csv.includes("int8"));
  assert.equal(await readFile(out, "utf8"), csv);
});

test("lucy serve floors + csv", async () => {
  const bin = await resolveLucyBinary();
  const g = JSON.parse(await readFile(path.join(root, "testdata/goldens_lpd_v1.0.json"), "utf8"));
  const child = spawn(bin, ["serve", "127.0.0.1:17476"], { stdio: ["ignore", "ignore", "pipe"] });
  try {
    await new Promise((r) => setTimeout(r, 400));
    const client = createLucyClient("http://127.0.0.1:17476");
    const floors = await client.floors();
    assert.equal(floors.version, "1.0.1");
    const csv = await client.csv(g.samples);
    assert.ok(csv.startsWith("id,band,lpd,"));
  } finally {
    child.kill("SIGTERM");
  }
});
