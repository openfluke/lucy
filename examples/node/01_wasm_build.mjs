/** npm-shaped: buildLPD via wasm (same API as @openfluke/lucy). */
import { readFile } from "node:fs/promises";
import path from "node:path";
import { fileURLToPath } from "node:url";
import { buildLPD, version, VERSION } from "../../js/src/index.js";

const root = path.resolve(path.dirname(fileURLToPath(import.meta.url)), "../..");
const req = JSON.parse(await readFile(path.join(root, "examples/shared/samples.json"), "utf8"));
console.log("package", VERSION, "wasm", await version());
const { board } = await buildLPD(req.samples, req.options);
console.log("top", board.top[0].id, "LPD", board.top[0].lpd, "band", board.top[0].band);
if (board.top[0].id !== "int8") throw new Error("expected int8");
console.log("OK wasm buildLPD");
