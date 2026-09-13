export const VERSION: string;
export const KEEP_FLOOR: number;
export const GOLD_KEEP: number;
export const LEAN_KEEP: number;
export const GOLD_RAM: number;
export const NEAR_RAM: number;
export const SHRINK_CAP: number;

export type Sample = Record<string, unknown> & { id: string };
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

export function buildLPD(samples: Sample[], options?: DensityOptions): Promise<BuildResponse>;
export function buildLPDNative(samples: Sample[], options?: DensityOptions): Promise<BuildResponse>;
export function resolveLucyBinary(): Promise<string>;
export function version(): Promise<string>;
export function drawRadar(canvas: HTMLCanvasElement, series: unknown[], opts?: { title?: string }): void;
export function consciousnessSeries(board: unknown, max?: number): unknown[];
export function densitySeries(board: unknown, max?: number): unknown[];
export function drawScatter(canvas: HTMLCanvasElement, pts: unknown[], opts?: object): void;
export function lpdScatterPoints(board: unknown): unknown[];
export function lpdTableHTML(board: unknown, opts?: { max?: number }): string;
export function registerLucyElements(): void;
export class LucyBoardElement extends HTMLElement { board: unknown }
