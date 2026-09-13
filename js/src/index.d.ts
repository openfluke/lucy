export const VERSION: string;
export const KEEP_FLOOR: number;
export const GOLD_KEEP: number;
export const LEAN_KEEP: number;
export const GOLD_RAM: number;
export const NEAR_RAM: number;
export const SHRINK_CAP: number;

export type Sample = {
  tide?: string;
  id: string;
  mode?: string;
  dtype?: string;
  format?: string;
  arch?: string;
  score?: number;
  soft?: number;
  soft_acc?: number;
  acc?: number;
  avg_accuracy?: number;
  thru?: number;
  throughput?: number;
  avail?: number;
  availability?: number;
  ramKiB?: number;
  ram_kib?: number;
};

export type DensityOptions = {
  keep_floor?: number;
  gold_keep?: number;
  lean_keep?: number;
  gold_ram?: number;
  near_ram?: number;
  shrink_cap?: number;
};

export type BuildResponse = {
  version: string;
  options: Required<DensityOptions>;
  board: Record<string, unknown>;
};

export function buildLPD(
  samples: Sample[],
  options?: DensityOptions,
): Promise<BuildResponse>;

export function version(): Promise<string>;
