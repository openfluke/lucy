#!/usr/bin/env python3
"""Build LPD — same goldens as Go/npm."""
from __future__ import annotations

import json
import sys
from pathlib import Path

ROOT = Path(__file__).resolve().parents[2]
sys.path.insert(0, str(ROOT / "python" / "src"))

from lucy import __version__, build_lpd, floors, version

req = json.loads((ROOT / "examples/shared/samples.json").read_text())
print("package", __version__, "binary", version())
print("floors", {k: floors()[k] for k in ("keep_floor", "gold_keep", "lean_keep")})
resp = build_lpd(req["samples"], req.get("options"))
top = resp["board"]["top"][0]
print(f"top {top['id']} LPD={top['lpd']:.4g} band={top['band']}")
assert top["id"] == "int8", "expected int8 to lead"
print("OK python build_lpd")
