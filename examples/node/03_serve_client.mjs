/** HTTP client against lucy serve — BYO frontend pattern. */
import { readFile } from "node:fs/promises";
import path from "node:path";
import { fileURLToPath } from "node:url";
import { spawn } from "node:child_process";
import { createLucyClient, resolveLucyBinary } from "../../js/src/index.js";

const root = path.resolve(path.dirname(fileURLToPath(import.meta.url)), "../..");
const req = JSON.parse(await readFile(path.join(root, "examples/shared/samples.json"), "utf8"));
const bin = await resolveLucyBinary();
const addr = "127.0.0.1:17490";
const child = spawn(bin, ["serve", addr], { stdio: ["ignore", "ignore", "pipe"] });
try {
  await new Promise((r) => setTimeout(r, 500));
  const c = createLucyClient(`http://${addr}`);
  console.log("version", await c.version());
  console.log("floors keep", (await c.floors()).keep_floor);
  const { board } = await c.buildLPD(req.samples, req.options);
  console.log("top", board.top[0].id);
  const pdf = await c.sitePDF(req.samples, req.options);
  const csv = await c.csv(req.samples, req.options);
  console.log("pdf bytes", pdf.length, "csv lines", csv.trim().split("\n").length);
  console.log("OK serve client");
} finally {
  child.kill("SIGTERM");
}
