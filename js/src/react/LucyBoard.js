import { createElement as h, useEffect, useRef } from "react";
import { drawRadar, consciousnessSeries, densitySeries } from "../charts/radar.js";
import { drawScatter, lpdScatterPoints } from "../charts/scatter.js";
import { lpdTableHTML } from "../charts/table.js";

function Radar({ board, mode = "consciousness", title }) {
  const ref = useRef(null);
  useEffect(() => {
    if (!ref.current || !board) return;
    const series =
      mode === "density" ? densitySeries(board) : consciousnessSeries(board);
    drawRadar(ref.current, series, {
      title:
        title ||
        (mode === "density" ? "Memory density radar" : "Consciousness radar"),
    });
  }, [board, mode, title]);
  return h("canvas", { ref, width: 960, height: 480, className: "lucy-radar" });
}

function Scatter({ board, title = "Q% vs RAM" }) {
  const ref = useRef(null);
  useEffect(() => {
    if (!ref.current || !board) return;
    drawScatter(ref.current, lpdScatterPoints(board), {
      title,
      xLabel: "RAM KiB",
      yLabel: "Q %",
    });
  }, [board, title]);
  return h("canvas", { ref, width: 960, height: 440, className: "lucy-scatter" });
}

function Table({ board }) {
  if (!board) return null;
  return h("div", {
    className: "lucy-table-wrap",
    dangerouslySetInnerHTML: { __html: lpdTableHTML(board) },
  });
}

/** Tide/Ocean-style LPD board: radars + scatter + table. */
export function LucyBoard({ board, showDensity = true }) {
  if (!board) return h("div", { className: "lucy-board empty" }, "no board");
  return h(
    "div",
    { className: "lucy-board" },
    h("div", { className: "lucy-board-charts" },
      h(Radar, { board, mode: "consciousness" }),
      showDensity ? h(Radar, { board, mode: "density" }) : null,
      h(Scatter, { board }),
    ),
    h(Table, { board }),
  );
}

export { Radar as ConsciousnessRadar, Radar as DensityRadar, Scatter as LPDScatter, Table as LPDTable };
