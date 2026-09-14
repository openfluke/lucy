import { ensureWasm } from "./loadWasm.js";

export const VERSION = "0.8.0";
export const KEEP_FLOOR = 0.7;
export const GOLD_KEEP = 0.8;
export const LEAN_KEEP = 0.95;
export const GOLD_RAM = 0.2;
export const NEAR_RAM = 0.5;
export const SHRINK_CAP = 32;

export * from "./charts/index.js";
export { buildLPDNative, resolveLucyBinary } from "./native.js";
export { registerLucyElements, LucyBoardElement, LucyLPDTableElement, registerLucyTableElement } from "./elements/lucy-board.js";
export { createLucyClient } from "./serve.js";

function wireSample(s) {
  return {
    tide: s.tide,
    id: s.id,
    mode: s.mode ?? "",
    dtype: s.dtype ?? s.dType ?? "",
    format: s.format ?? "",
    arch: s.arch ?? "",
    score: s.score ?? 0,
    soft_acc: s.soft_acc ?? s.soft ?? 0,
    avg_accuracy: s.avg_accuracy ?? s.acc ?? 0,
    throughput: s.throughput ?? s.thru ?? 0,
    availability: s.availability ?? s.avail ?? 0,
    ram_kib: s.ram_kib ?? s.ramKiB ?? 0,
  };
}

/** Rank samples via Go wasm (browser/Node). Pass DensityOptions to tune floors. */
export async function buildLPD(samples, options) {
  await ensureWasm();
  const req = {
    samples: (samples || []).map(wireSample),
    options: options || undefined,
  };
  const raw = globalThis.lucyBuildLPD(JSON.stringify(req));
  const parsed = JSON.parse(raw);
  if (parsed && parsed.error) throw new Error(parsed.error);
  return parsed;
}

export async function version() {
  await ensureWasm();
  return typeof globalThis.lucyVersion === "function"
    ? globalThis.lucyVersion()
    : VERSION;
}
