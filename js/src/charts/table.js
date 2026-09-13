/** Build an HTML table string for LPD top rows (vanilla / Angular-friendly). */
export function lpdTableHTML(board, { max = 40 } = {}) {
  const rows = (board?.top || []).slice(0, max);
  const body = rows
    .map((r) => {
      const q = ((r.q ?? 0) * 100).toFixed(0);
      const ram = ((r.ram_frac ?? 0) * 100).toFixed(0);
      const lpd = (r.lpd ?? 0).toFixed(2);
      const kib = (r.ram_kib ?? 0).toFixed(1);
      return `<tr data-band="${r.band || ""}"><td>${escapeHtml(r.band || "—")}</td><td>${escapeHtml(r.id)}</td><td class="num">${q}</td><td class="num">${ram}</td><td class="num">${lpd}</td><td class="num">${kib}</td></tr>`;
    })
    .join("");
  return `<table class="lucy-lpd-table"><thead><tr><th>band</th><th>cell</th><th class="num">Q%</th><th class="num">RAM%</th><th class="num">LPD</th><th class="num">KiB</th></tr></thead><tbody>${body}</tbody></table>`;
}

function escapeHtml(s) {
  return String(s)
    .replaceAll("&", "&amp;")
    .replaceAll("<", "&lt;")
    .replaceAll(">", "&gt;")
    .replaceAll('"', "&quot;");
}
