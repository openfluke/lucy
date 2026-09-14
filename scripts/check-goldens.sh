#!/usr/bin/env bash
# Cross-runtime golden check: Go ↔ wasm ↔ Python binary
set -euo pipefail
ROOT="$(cd "$(dirname "$0")/.." && pwd)"
echo "== go =="
(cd "$ROOT/go" && go test ./lucy/ -count=1)
echo "== js (wasm + native + charts) =="
(cd "$ROOT/js" && node --test test/*.test.js)
echo "== python =="
(cd "$ROOT/python" && PYTHONPATH=src python3 tests/test_build_lpd.py)
echo "OK — goldens aligned across Go / npm / PyPI wrappers"
