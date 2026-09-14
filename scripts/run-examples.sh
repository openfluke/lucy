#!/usr/bin/env bash
# Smoke every language surface.
set -euo pipefail
ROOT="$(cd "$(dirname "$0")/.." && pwd)"
export LUCY_ROOT="$ROOT"
cd "$ROOT"
mkdir -p examples/out

echo "======== CLI ========"
bash examples/cli/run.sh

echo "======== Go ========"
(cd examples/go && go run ./01_build_lpd && go run ./02_site_pdf)

echo "======== Node ========"
(cd examples/node && node 01_wasm_build.mjs && node 02_native_full.mjs && node 03_serve_client.mjs)

echo "======== Python ========"
PYTHONPATH=python/src python3 examples/python/01_build_lpd.py
PYTHONPATH=python/src python3 examples/python/02_full_export.py
PYTHONPATH=python/src python3 examples/python/03_serve_client.py

echo "======== Jupyter cells (no jupyter required) ========"
PYTHONPATH=python/src python3 - <<'PY'
import json, sys, os
from pathlib import Path
ROOT = Path(os.environ["LUCY_ROOT"])
sys.path.insert(0, str(ROOT / "python" / "src"))
from lucy import build_lpd, board_csv, write_site_pdf
nb = json.loads((ROOT / "examples/python/lucy_tutorial.ipynb").read_text())
assert nb["nbformat"] == 4 and len(nb["cells"]) >= 4
req = json.loads((ROOT / "examples/shared/samples.json").read_text())
resp = build_lpd(req["samples"], req.get("options"))
assert resp["board"]["top"][0]["id"] == "int8"
out = ROOT / "examples/out/jupyter"
out.mkdir(parents=True, exist_ok=True)
(out / "board.csv").write_text(board_csv(req["samples"], req.get("options")))
write_site_pdf(req["samples"], out / "site.pdf", req.get("options"))
print("OK notebook logic + wrote", out)
PY

echo "======== browser static check ========"
test -f examples/browser/index.html
grep -q "lucy-board" examples/browser/index.html
echo "OK browser index present (serve with: python3 -m http.server 8765)"

echo
echo "ALL EXAMPLES OK"
find examples/out -type f | sort
