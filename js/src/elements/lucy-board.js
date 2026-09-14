import { drawRadar, consciousnessSeries, densitySeries } from "../charts/radar.js";
import { drawScatter, lpdScatterPoints } from "../charts/scatter.js";
import { lpdTableHTML } from "../charts/table.js";
import { registerLucyTableElement } from "./lucy-table.js";

const Base =
  typeof HTMLElement !== "undefined"
    ? HTMLElement
    : class {};

/**
 * <lucy-board board-json='...'> — Tide/Ocean/Angular-friendly custom element.
 *
 * CSS vars: --lucy-bg, --lucy-fg, --lucy-muted, --lucy-gap
 */
export class LucyBoardElement extends Base {
  static get observedAttributes() {
    return ["board-json"];
  }

  #board = null;

  set board(v) {
    this.#board = v;
    this.render();
  }
  get board() {
    return this.#board;
  }

  attributeChangedCallback() {
    this.#syncFromAttrs();
    this.render();
  }

  connectedCallback() {
    this.#syncFromAttrs();
    this.render();
  }

  #syncFromAttrs() {
    if (typeof this.getAttribute !== "function") return;
    const raw = this.getAttribute("board-json");
    if (raw) {
      try {
        this.#board = JSON.parse(raw);
      } catch {
        /* keep previous */
      }
    }
  }

  render() {
    if (typeof document === "undefined") return;
    this.innerHTML = "";
    const root = document.createElement("div");
    root.className = "lucy-board";
    root.style.cssText =
      "display:grid;gap:var(--lucy-gap,1rem);background:var(--lucy-bg,transparent);color:var(--lucy-fg,inherit);";
    const live = document.createElement("canvas");
    live.width = 960;
    live.height = 480;
    const dens = document.createElement("canvas");
    dens.width = 960;
    dens.height = 480;
    const scat = document.createElement("canvas");
    scat.width = 960;
    scat.height = 440;
    const table = document.createElement("div");
    table.className = "lucy-board__table";
    root.append(live, dens, scat, table);
    this.append(root);
    if (!this.#board) {
      table.innerHTML = "<p class=\"lucy-empty\" style=\"color:var(--lucy-muted,#888)\">no board</p>";
      return;
    }
    drawRadar(live, consciousnessSeries(this.#board), { title: "Consciousness radar" });
    drawRadar(dens, densitySeries(this.#board), { title: "Memory density radar" });
    drawScatter(scat, lpdScatterPoints(this.#board), {
      title: "Q% vs RAM",
      xLabel: "RAM KiB",
      yLabel: "Q %",
    });
    table.innerHTML = lpdTableHTML(this.#board);
  }
}

export function registerLucyElements() {
  if (typeof customElements === "undefined") return;
  if (!customElements.get("lucy-board")) {
    customElements.define("lucy-board", LucyBoardElement);
  }
  registerLucyTableElement();
}

export { LucyLPDTableElement, registerLucyTableElement } from "./lucy-table.js";
