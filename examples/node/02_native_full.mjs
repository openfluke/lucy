/** Native CLI path: board + CSV + site PDF + floors (Node/Bun backend). */
import { readFile, mkdir } from "node:fs/promises";
import path from "node:path";
import { fileURLToPath } from "node:url";
import {
  buildLPDNative,
  writeCSVNative,
  writeSitePDFNative,
  floorsNative,
  chartPackNative,
} from "../../js/src/index.js";

const root = path.resolve(path.dirname(fileURLToPath(import.meta.url)), "../..");
const out = path.join(root, "examples/out/node");
await mkdir(out, { recursive: true });
const req = JSON.parse(await readFile(path.join(root, "examples/shared/samples.json"), "utf8"));
console.log("floors", await floorsNative());
const { board } = await buildLPDNative(req.samples, req.options);
console.log("top", board.top[0].id, board.top[0].lpd);
await writeCSVNative(req.samples, path.join(out, "board.csv"), req.options);
await writeSitePDFNative(req.samples, path.join(out, "site.pdf"), req.options);
const pack = await chartPackNative(req.samples, req.options);
console.log("chart-pack keys", Object.keys(pack).filter((k) => k.includes("svg")).join(", "));
console.log("OK native →", out);
