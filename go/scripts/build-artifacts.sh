#!/usr/bin/env bash
# Build native lucy CLI + wasm for npm / PyPI wrappers.
set -euo pipefail
ROOT="$(cd "$(dirname "$0")/../.." && pwd)"
GO_DIR="$ROOT/go"
OUT="$ROOT/artifacts"
JS_WASM="$ROOT/js/wasm"
PY_BIN="$ROOT/python/src/lucy/bin"

mkdir -p "$OUT" "$JS_WASM" "$PY_BIN"
cd "$GO_DIR"

VERSION="$(cat "$ROOT/VERSION" 2>/dev/null || echo dev)"
echo "building lucy $VERSION artifacts…"

# Native CLI — host platform + common cross targets
build_cli() {
  local goos=$1 goarch=$2
  local name="lucy-${goos}-${goarch}"
  if [[ "$goos" == windows ]]; then name="${name}.exe"; fi
  echo "  cli $name"
  GOOS=$goos GOARCH=$goarch CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o "$OUT/$name" ./cmd/lucy
}

build_cli "$(go env GOOS)" "$(go env GOARCH)"
build_cli linux amd64
build_cli linux arm64
build_cli darwin amd64
build_cli darwin arm64
build_cli windows amd64

# Host binary into Python package data
HOST_NAME="lucy-$(go env GOOS)-$(go env GOARCH)"
cp -f "$OUT/$HOST_NAME" "$PY_BIN/lucy"
chmod +x "$PY_BIN/lucy"
# also keep named copy
cp -f "$OUT/$HOST_NAME" "$PY_BIN/$HOST_NAME"
chmod +x "$PY_BIN/$HOST_NAME"

# Wasm
echo "  wasm lucy.wasm"
GOOS=js GOARCH=wasm go build -trimpath -ldflags="-s -w" -o "$OUT/lucy.wasm" ./cmd/lucywasm
cp -f "$OUT/lucy.wasm" "$JS_WASM/lucy.wasm"
cp -f "$(go env GOROOT)/lib/wasm/wasm_exec.js" "$JS_WASM/wasm_exec.js"
# Node helper from Go tree when present
if [[ -f "$(go env GOROOT)/lib/wasm/wasm_exec_node.js" ]]; then
  cp -f "$(go env GOROOT)/lib/wasm/wasm_exec_node.js" "$JS_WASM/wasm_exec_node.js"
fi

echo "done → $OUT"
ls -la "$OUT"
