#!/usr/bin/env python3
"""Notebook-style smoke: board → DataFrame → floors (Lucy 1.0)."""
import json
import sys
from pathlib import Path

ROOT = Path(__file__).resolve().parents[2]
sys.path.insert(0, str(ROOT / "python" / "src"))

from lucy import board_records, build_lpd, floors, to_dataframe  # noqa: E402

g = json.loads((ROOT / "testdata" / "goldens_lpd_v1.0.json").read_text())
print("floors", {k: floors()[k] for k in ("version", "keep_floor", "gold_keep")})
resp = build_lpd(g["samples"])
print("top", resp["board"]["top"][0]["id"], resp["board"]["top"][0]["lpd"])
print("records", len(board_records(resp)))
try:
    df = to_dataframe(resp)
    print(df.head())
except ImportError:
    print("(pandas optional — pip install pandas for DataFrame)")
