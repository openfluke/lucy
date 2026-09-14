#!/usr/bin/env bash
# Verify version alignment across channels (publish readiness).
set -euo pipefail
ROOT="$(cd "$(dirname "$0")/.." && pwd)"
V="$(tr -d '[:space:]' < "$ROOT/VERSION")"
fail=0
check() {
  local label=$1 got=$2
  if [[ "$got" != "$V" ]]; then
    echo "FAIL $label: got $got want $V"
    fail=1
  else
    echo "OK   $label = $V"
  fi
}
check VERSION "$V"
check go "$(grep -oE 'Version = "[^"]+"' "$ROOT/go/lucy/version.go" | head -1 | cut -d'"' -f2)"
check npm "$(node -p "require('$ROOT/js/package.json').version")"
check py "$(python3 -c "import tomllib; print(tomllib.load(open('$ROOT/python/pyproject.toml','rb'))['project']['version'])")"
G="$ROOT/testdata/goldens_lpd_v${V%.*}.json"
# goldens file uses major.minor from version like 0.8.0 → v0.8
majmin="${V%.*}"
G="$ROOT/testdata/goldens_lpd_v${majmin}.json"
if [[ -f "$G" ]]; then
  gv="$(python3 -c "import json; print(json.load(open('$G'))['version'])")"
  check "goldens $majmin" "$gv"
else
  echo "FAIL missing $G"; fail=1
fi
if command -v npm >/dev/null; then
  (cd "$ROOT/js" && npm publish --dry-run >/tmp/lucy-npm-dry.txt 2>&1) && echo "OK   npm publish --dry-run" || {
    echo "WARN npm dry-run failed (see /tmp/lucy-npm-dry.txt)"; cat /tmp/lucy-npm-dry.txt | tail -20
  }
fi
exit $fail
