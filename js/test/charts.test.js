import { test } from "node:test";
import assert from "node:assert/strict";
import { readFile } from "node:fs/promises";
import path from "node:path";
import { fileURLToPath } from "node:url";
import { buildLPD, consciousnessSeries, lpdTableHTML, lpdScatterPoints, buildLPDNative } from "../src/index.js";

const root = path.resolve(path.dirname(fileURLToPath(import.meta.url)), "../..");

test("chart helpers from board", async () => {
  const g = JSON.parse(await readFile(path.join(root, "testdata/goldens_lpd_v0.8.json"), "utf8"));
  const { board } = await buildLPD(g.samples);
  const series = consciousnessSeries(board);
  assert.ok(series.length >= 1);
  assert.equal(series[0].vals.length, 3);
  const html = lpdTableHTML(board);
  assert.match(html, /lucy-lpd-table/);
  assert.match(html, /int8/);
  assert.ok(lpdScatterPoints(board).length >= 1);
});

test("native binary buildLPD matches wasm lead", async () => {
  const g = JSON.parse(await readFile(path.join(root, "testdata/goldens_lpd_v0.8.json"), "utf8"));
  const native = await buildLPDNative(g.samples);
  assert.equal(native.board.top[0].id, "int8");
});
