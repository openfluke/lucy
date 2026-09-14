import { test } from "node:test";
import assert from "node:assert/strict";
import { readFile } from "node:fs/promises";
import path from "node:path";
import { fileURLToPath } from "node:url";
import { buildLPD, VERSION } from "../src/index.js";

const root = path.resolve(path.dirname(fileURLToPath(import.meta.url)), "../..");

test("buildLPD matches Go golden lead", async () => {
  const raw = await readFile(path.join(root, "testdata/goldens_lpd_v0.5.json"), "utf8");
  const g = JSON.parse(raw);
  const resp = await buildLPD(g.samples);
  assert.equal(VERSION, "0.5.0");
  assert.equal(resp.version, "0.5.0");
  assert.equal(resp.board.top[0].id, g.expect.top[0].id);
  assert.equal(resp.board.top[0].band, g.expect.top[0].band);
  assert.ok(Math.abs(resp.board.top[0].lpd - g.expect.top[0].lpd) < 1e-9);
  assert.equal(resp.board.gold_std.id, g.expect.gold_std_id);
});
