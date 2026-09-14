import { drawRadar, consciousnessSeries, densitySeries } from "../charts/radar.js";
import { drawScatter, lpdScatterPoints } from "../charts/scatter.js";
import { lpdTableHTML } from "../charts/table.js";
import { registerLucyTableElement } from "./lucy-table.js";

const Base =
  typeof HTMLElement !== "undefined"
    ? HTMLElement
    : class {};

export class LucyBoardElement extends Base {
  #board = null;
  set board(v) {
    this.#board = v;
    this.render();
  }
  get board() {
    return this.#board;
  }
  connectedCallback() {
    this.render();
  }
  render() {
    if (typeof document === "undefined") return;
    this.innerHTML = "";
    const root = document.createElement("div");
    root.className = "lucy-board";
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
    root.append(live, dens, scat, table);
    this.append(root);
    if (!this.#board) {
      table.textContent = "no board";
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
