import { lpdTableHTML } from "../charts/table.js";

const Base =
  typeof HTMLElement !== "undefined"
    ? HTMLElement
    : class {};

/**
 * <lucy-lpd-table board-json='{"top":[...]}'>
 * or element.board = boardObject
 *
 * Angular-friendly: bind [attr.board-json]="board | json"
 */
export class LucyLPDTableElement extends Base {
  static get observedAttributes() {
    return ["board-json", "max"];
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
    const maxAttr = this.getAttribute?.("max");
    const max = maxAttr ? Number(maxAttr) : undefined;
    this.innerHTML = this.#board
      ? lpdTableHTML(this.#board, max ? { max } : undefined)
      : "<p class=\"lucy-empty\">no board</p>";
  }
}

export function registerLucyTableElement() {
  if (typeof customElements === "undefined") return;
  if (!customElements.get("lucy-lpd-table")) {
    customElements.define("lucy-lpd-table", LucyLPDTableElement);
  }
}
