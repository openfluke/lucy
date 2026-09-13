import { readFile } from "node:fs/promises";
import { buildLPD, KEEP_FLOOR } from "../src/index.js";

const path = process.argv[2];
if (!path) {
  console.error("usage: node examples/byo-poll.mjs samples.json");
  process.exit(2);
}
const raw = JSON.parse(await readFile(path, "utf8"));
const samples = raw.samples || raw;
const { board } = await buildLPD(samples, { keep_floor: KEEP_FLOOR });
console.log("top", board.top?.[0]?.id, "LPD", board.top?.[0]?.lpd, "band", board.top?.[0]?.band);
