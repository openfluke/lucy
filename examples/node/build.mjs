import { buildLPDNative, writeCSVNative, writePDFNative, floorsNative } from "../../js/src/index.js";
import { readFile } from "node:fs/promises";
import path from "node:path";
import { fileURLToPath } from "node:url";

const root = path.resolve(path.dirname(fileURLToPath(import.meta.url)), "../..");
const g = JSON.parse(await readFile(path.join(root, "testdata/goldens_lpd_v0.7.json"), "utf8"));
console.log("floors", await floorsNative());
const { board } = await buildLPDNative(g.samples);
console.log("top", board.top[0].id, board.top[0].lpd);
await writeCSVNative(g.samples, "board.csv");
await writePDFNative(g.samples, "board.pdf");
console.log("wrote board.csv board.pdf");
