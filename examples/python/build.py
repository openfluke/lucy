#!/usr/bin/env python3
import json
import sys
from pathlib import Path

ROOT = Path(__file__).resolve().parents[2]
sys.path.insert(0, str(ROOT / "python" / "src"))

from lucy import board_csv, build_lpd, floors, write_pdf

g = json.loads((ROOT / "testdata" / "goldens_lpd_v0.9.json").read_text())
print("floors", floors())
resp = build_lpd(g["samples"])
print("top", resp["board"]["top"][0]["id"], resp["board"]["top"][0]["lpd"])
Path("board.csv").write_text(board_csv(g["samples"]))
write_pdf(g["samples"], "board.pdf")
print("wrote board.csv board.pdf")
