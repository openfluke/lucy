#!/usr/bin/env python3
"""LucyClient against lucy serve."""
from __future__ import annotations

import json
import subprocess
import sys
import time
from pathlib import Path

ROOT = Path(__file__).resolve().parents[2]
sys.path.insert(0, str(ROOT / "python" / "src"))

from lucy import LucyClient
from lucy._bin import default_binary

req = json.loads((ROOT / "examples/shared/samples.json").read_text())
addr = "127.0.0.1:17491"
proc = subprocess.Popen(
    [str(default_binary()), "serve", addr],
    stdout=subprocess.DEVNULL,
    stderr=subprocess.DEVNULL,
)
try:
    time.sleep(0.5)
    c = LucyClient(f"http://{addr}")
    print("version", c.version())
    print("floors keep", c.floors()["keep_floor"])
    board = c.build_lpd(req["samples"], req.get("options"))["board"]
    print("top", board["top"][0]["id"])
    pdf = c.site_pdf(req["samples"], req.get("options"))
    csv = c.csv(req["samples"], req.get("options"))
    print("pdf", len(pdf), "csv lines", len(csv.strip().splitlines()))
    print("OK python serve client")
finally:
    proc.terminate()
    proc.wait(timeout=5)
