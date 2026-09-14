#!/usr/bin/env python3
"""CSV, site PDF, report, charts — notebook-backend style."""
from __future__ import annotations

import json
import sys
from pathlib import Path

ROOT = Path(__file__).resolve().parents[2]
sys.path.insert(0, str(ROOT / "python" / "src"))

from lucy import (
    board_csv,
    board_records,
    build_lpd,
    chart_jpg,
    chart_png,
    chart_svg,
    write_report,
    write_site_pdf,
)

req = json.loads((ROOT / "examples/shared/samples.json").read_text())
out = ROOT / "examples" / "out" / "python"
out.mkdir(parents=True, exist_ok=True)
samples, opts = req["samples"], req.get("options")

resp = build_lpd(samples, opts)
print("rows", len(board_records(resp)))
(out / "board.csv").write_text(board_csv(samples, opts))
write_site_pdf(samples, out / "site.pdf", opts)
write_report(samples, out / "report", opts)
(out / "radar.svg").write_text(chart_svg(samples, "radar", opts))
(out / "bars.png").write_bytes(chart_png(samples, "bars", opts))
(out / "radar.jpg").write_bytes(chart_jpg(samples, "radar", opts))
print("OK python exports →", out)
