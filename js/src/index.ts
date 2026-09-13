/**
 * @openfluke/lucy — placeholder entry.
 *
 * Planned surface:
 * - measuring API (wasm / native) mirroring Go: Score, Q, BuildLPD, bands
 * - board helpers for Node / Bun backends
 * - React / framework-agnostic chart components (Tide/Ocean/River style)
 *
 * Hosts supply samples or poll their own API — nothing Tide-hardcoded.
 */

export type Sample = {
  id: string;
  mode?: string;
  dtype?: string;
  arch?: string;
  score: number;
  soft?: number;
  acc: number;
  thru: number;
  avail: number;
  ramKiB: number;
};

export type Board = {
  formula: string;
  top: Array<{ id: string; q: number; lpd: number }>;
};

/** Placeholder — will call Go wasm/binary once wired. */
export function buildLPD(_samples: Sample[]): Board {
  return {
    formula:
      "placeholder — Score = T×Avail×Acc/10_000; LPD = Q×shrink if RelAcc≥KeepFloor else 0",
    top: [],
  };
}

export const KEEP_FLOOR = 0.7;
