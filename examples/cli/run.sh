#!/usr/bin/env bash
set -euo pipefail
ROOT="$(cd "$(dirname "$0")/../.." && pwd)"
OS=$(go env GOOS); ARCH=$(go env GOARCH)
BIN="$ROOT/artifacts/lucy-${OS}-${ARCH}"
[[ -x "$BIN" ]] || BIN="$ROOT/python/src/lucy/bin/lucy"
OUT="$ROOT/examples/out/cli"
mkdir -p "$OUT"
SAMPLES="$ROOT/examples/shared/samples.json"

echo "== version =="; "$BIN" version
echo "== floors =="; "$BIN" floors | head -c 200; echo …
echo "== artifacts =="; "$BIN" artifacts | head -c 200; echo …
echo "== build-lpd =="; "$BIN" build-lpd < "$SAMPLES" | python3 -c "import sys,json; b=json.load(sys.stdin); print('top', b['board']['top'][0]['id'], b['board']['top'][0]['lpd'])"
echo "== csv =="; "$BIN" csv < "$SAMPLES" > "$OUT/board.csv"; head -2 "$OUT/board.csv"
echo "== site-pdf =="; "$BIN" site-pdf "$OUT/site.pdf" < "$SAMPLES"
echo "== report =="; "$BIN" report "$OUT/report" < "$SAMPLES"
echo "OK cli → $OUT"
