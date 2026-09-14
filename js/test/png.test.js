import { test } from "node:test";
import assert from "node:assert/strict";
import { readFile } from "node:fs/promises";
import path from "node:path";
import { fileURLToPath } from "node:url";
import { chartPNGNative, chartPackNative } from "../src/index.js";

const root = path.resolve(path.dirname(fileURLToPath(import.meta.url)), "../..");

test("native PNG + chart-pack", async () => {
  const g = JSON.parse(await readFile(path.join(root, "testdata/goldens_lpd_v0.9.json"), "utf8"));
  const png = await chartPNGNative(g.samples, "radar");
  assert.ok(Buffer.isBuffer(png) || png instanceof Uint8Array);
  assert.equal(png[0], 0x89);
  const pack = await chartPackNative(g.samples);
  assert.ok(pack.consciousness_svg.includes("<svg"));
  assert.ok(pack.consciousness_png_b64.length > 50);
});
