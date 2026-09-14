import { lpdTableHTML } from "../charts/table.js";

const Base =
  typeof HTMLElement !== "undefined"
    ? HTMLElement
    : class {};

/** <lucy-lpd-table> — set .board = LPD board. */
export class LucyLPDTableElement extends Base {
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
    this.innerHTML = this.#board ? lpdTableHTML(this.#board) : "<p>no board</p>";
  }
}

export function registerLucyTableElement() {
  if (typeof customElements === "undefined") return;
  if (!customElements.get("lucy-lpd-table")) {
    customElements.define("lucy-lpd-table", LucyLPDTableElement);
  }
}
