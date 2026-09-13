const COLORS = ["#3dd6c6", "#6ea8fe", "#e0a458", "#c084fc", "#f07178"];

export function consciousnessSeries(board, max = 8) {
  const rows = (board?.top || []).slice(0, max);
  return rows.map((r, i) => ({
    label: r.id,
    color: COLORS[i % COLORS.length],
    vals: [r.rel_acc ?? 0, r.rel_thru ?? 0, r.rel_avail ?? 0],
  }));
}

export function densitySeries(board, max = 8) {
  const rows = (board?.top || []).slice(0, max);
  return rows.map((r, i) => {
    const vals = [r.dens_acc ?? 0, r.dens_thru ?? 0, r.dens_avail ?? 0];
    const peak = Math.max(1, ...vals);
    return {
      label: r.id,
      color: COLORS[i % COLORS.length],
      vals: vals.map((v) => Math.min(1, Math.max(0, v / peak))),
    };
  });
}

/** Draw 3-axis radar onto a canvas (Acc/Thru/Avail). */
export function drawRadar(canvas, series, { title = "" } = {}) {
  const w = canvas.width || 960;
  const h = canvas.height || 480;
  const ctx = canvas.getContext("2d");
  ctx.fillStyle = "#0d1216";
  ctx.fillRect(0, 0, w, h);
  if (title) {
    ctx.fillStyle = "#8aa0ad";
    ctx.font = "13px sans-serif";
    ctx.fillText(title, 12, 18);
  }
  if (!series?.length) {
    ctx.fillStyle = "#5a7a8a";
    ctx.fillText("no series", w / 2 - 30, h / 2);
    return;
  }
  const cx = w * 0.36;
  const cy = h * 0.54;
  const radius = Math.min(cx - 24, cy - 40);
  const ang = (i) => -Math.PI / 2 + (i * 2 * Math.PI) / 3;
  const labels = ["Acc", "Thru", "Avail"];
  ctx.strokeStyle = "#1d3342";
  for (let ring = 1; ring <= 4; ring++) {
    ctx.beginPath();
    for (let i = 0; i < 3; i++) {
      const r = (radius * ring) / 4;
      const x = cx + r * Math.cos(ang(i));
      const y = cy + r * Math.sin(ang(i));
      if (i === 0) ctx.moveTo(x, y);
      else ctx.lineTo(x, y);
    }
    ctx.closePath();
    ctx.stroke();
  }
  for (let i = 0; i < 3; i++) {
    ctx.beginPath();
    ctx.moveTo(cx, cy);
    ctx.lineTo(cx + radius * Math.cos(ang(i)), cy + radius * Math.sin(ang(i)));
    ctx.stroke();
    ctx.fillStyle = "#8aa0ad";
    ctx.font = "600 13px sans-serif";
    ctx.textAlign = "center";
    ctx.fillText(
      labels[i],
      cx + (radius + 22) * Math.cos(ang(i)),
      cy + (radius + 22) * Math.sin(ang(i)) + 4,
    );
  }
  for (const s of series) {
    ctx.strokeStyle = s.color || "#3dd6c6";
    ctx.lineWidth = 2;
    ctx.beginPath();
    for (let i = 0; i < 3; i++) {
      const v = Math.min(1, Math.max(0, s.vals[i] ?? 0));
      const x = cx + radius * v * Math.cos(ang(i));
      const y = cy + radius * v * Math.sin(ang(i));
      if (i === 0) ctx.moveTo(x, y);
      else ctx.lineTo(x, y);
    }
    ctx.closePath();
    ctx.stroke();
  }
  let ly = 40;
  ctx.textAlign = "left";
  ctx.font = "12px sans-serif";
  for (const s of series) {
    ctx.fillStyle = s.color || "#3dd6c6";
    ctx.fillRect(w - 300, ly, 12, 12);
    ctx.fillStyle = "#c5d0d8";
    ctx.fillText(s.label, w - 282, ly + 11);
    ly += 18;
  }
}
