/**
 * @openfluke/lucy 0.1.0
 *
 * Board types aligned with Go `github.com/openfluke/lucy/lucy`.
 * Measuring runtime (wasm / native) lands at 0.2 — do not reimplement LPD in TS.
 */

export const VERSION = "0.1.0";
export const KEEP_FLOOR = 0.7;
export const GOLD_KEEP = 0.8;
export const LEAN_KEEP = 0.95;
export const GOLD_RAM = 0.2;
export const NEAR_RAM = 0.5;
export const SHRINK_CAP = 32;

export type Sample = {
  tide?: string;
  id: string;
  mode?: string;
  dtype?: string;
  format?: string;
  arch?: string;
  score: number;
  soft?: number;
  acc: number;
  thru: number;
  avail: number;
  ramKiB: number;
};

export type LPDRow = {
  id: string;
  band: string;
  q: number;
  lpd: number;
  relAcc: number;
  shrink: number;
  ramKiB: number;
};

export type Board = {
  formula: string;
  champId?: string;
  accChampId?: string;
  goldStdId?: string;
  top: LPDRow[];
  trapIds?: string[];
};

/**
 * Rank samples for Lucy Pareto density.
 * 0.1: types only — call Go BuildLPD or wait for wasm at 0.2.
 */
export function buildLPD(_samples: Sample[]): Board {
  throw new Error(
    "@openfluke/lucy 0.1.0: use Go github.com/openfluke/lucy/lucy.BuildLPD; wasm/native lands at 0.2",
  );
}
