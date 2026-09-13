function bandColor(band) {
  switch (band) {
    case "gold": return "#e0a458";
    case "near": return "#6ea8fe";
    case "trap": return "#f07178";
    case "keep": return "#3dd6c6";
    default: return "#5a7a8a";
  }
}

export function lpdScatterPoints(board) {
  return (board?.top || []).map((r) => ({
    id: r.id,
    x: r.ram_kib ?? 0,
    y: (r.q ?? 0) * 100,
    band: r.band,
  }));
}

export function drawScatter(canvas, pts, { title = "", xLabel = "X", yLabel = "Y" } = {}) {
  const w = canvas.width || 960;
  const h = canvas.height || 440;
  const ctx = canvas.getContext("2d");
  ctx.fillStyle = "#0d1216";
  ctx.fillRect(0, 0, w, h);
  ctx.fillStyle = "#8aa0ad";
  ctx.font = "13px sans-serif";
  if (title) ctx.fillText(title, 12, 16);
  if (!pts?.length) {
    ctx.fillStyle = "#5a7a8a";
    ctx.fillText("no points", w / 2 - 30, h / 2);
    return;
  }
  const padL = 48, padR = 16, padT = 12, padB = 36;
  let xmin = Infinity, xmax = -Infinity, ymin = Infinity, ymax = -Infinity;
  for (const p of pts) {
    xmin = Math.min(xmin, p.x); xmax = Math.max(xmax, p.x);
    ymin = Math.min(ymin, p.y); ymax = Math.max(ymax, p.y);
  }
  if (xmax <= xmin) xmax = xmin + 1;
  if (ymax <= ymin) ymax = ymin + 1;
  const X = (v) => padL + ((w - padL - padR) * (v - xmin)) / (xmax - xmin);
  const Y = (v) => h - padB - ((h - padT - padB) * (v - ymin)) / (ymax - ymin);
  ctx.strokeStyle = "#1d3342";
  ctx.strokeRect(padL, padT, w - padL - padR, h - padT - padB);
  for (const p of pts) {
    ctx.fillStyle = bandColor(p.band);
    ctx.fillRect(X(p.x) - 3, Y(p.y) - 3, 6, 6);
  }
  ctx.fillStyle = "#8aa0ad";
  ctx.font = "12px sans-serif";
  ctx.textAlign = "center";
  ctx.fillText(xLabel, w / 2, h - 8);
  ctx.save();
  ctx.translate(14, h / 2);
  ctx.rotate(-Math.PI / 2);
  ctx.fillText(yLabel, 0, 0);
  ctx.restore();
}
