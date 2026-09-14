#!/usr/bin/env bash
# Publish @openfluke/lucy to npm (wasm built locally — not committed to git).
#
# Usage:
#   bash publish.sh            # interactive confirm
#   bash publish.sh --dry-run  # build + test + pack only, no publish
#   bash publish.sh --yes      # skip confirm (CI / you already decided)
#
set -euo pipefail

JS_DIR="$(cd "$(dirname "$0")" && pwd)"
ROOT="$(cd "$JS_DIR/.." && pwd)"
cd "$JS_DIR"

DRY_RUN=0
ASSUME_YES=0
for arg in "$@"; do
  case "$arg" in
    --dry-run|-n) DRY_RUN=1 ;;
    --yes|-y) ASSUME_YES=1 ;;
    -h|--help)
      sed -n '2,12p' "$0"
      exit 0
      ;;
    *)
      echo "Unknown arg: $arg" >&2
      exit 2
      ;;
  esac
done

echo "=== @openfluke/lucy — npm publish ==="
echo "cwd: $JS_DIR"
echo ""

NAME=$(node -p "require('./package.json').name")
VERSION=$(node -p "require('./package.json').version")
MONOREPO_V="$(tr -d '[:space:]' < "$ROOT/VERSION")"
echo "Package:  ${NAME}@${VERSION}"
echo "Monorepo: ${MONOREPO_V}"
echo ""

if [[ "$VERSION" != "$MONOREPO_V" ]]; then
  echo "ERROR: js/package.json version (${VERSION}) != repo VERSION (${MONOREPO_V})" >&2
  exit 1
fi

# 1. Build wasm + native artifacts into js/wasm/
echo "→ npm run build  (./go/scripts/build-artifacts.sh)"
npm run build

if [[ ! -f "$JS_DIR/wasm/lucy.wasm" ]]; then
  echo "ERROR: wasm/lucy.wasm missing after build" >&2
  exit 1
fi
if [[ ! -f "$JS_DIR/wasm/wasm_exec.js" ]]; then
  echo "ERROR: wasm/wasm_exec.js missing after build" >&2
  exit 1
fi

WASM_MB=$(du -m "$JS_DIR/wasm/lucy.wasm" | awk '{print $1}')
echo "✓ wasm/lucy.wasm (~${WASM_MB} MB)"
echo ""

# 2. Gate tests
echo "→ npm test"
npm test
echo ""

# 3. Live version check against built WASM
export LUCY_JS_DIR="$JS_DIR"
export LUCY_ROOT="$ROOT"
node --input-type=module <<'NODE'
import { readFile } from "node:fs/promises";
import { pathToFileURL } from "node:url";
import path from "node:path";

const jsDir = process.env.LUCY_JS_DIR;
const root = process.env.LUCY_ROOT;
const { buildLPD, version, VERSION } = await import(
  pathToFileURL(path.join(jsDir, "src/index.js")).href
);

const req = JSON.parse(
  await readFile(path.join(root, "examples/shared/samples.json"), "utf8"),
);
const wasmV = await version();
if (wasmV !== VERSION) {
  console.error(`version mismatch WASM=${wasmV} package=${VERSION}`);
  process.exit(1);
}
const { board } = await buildLPD(req.samples, req.options);
const top = board.top[0];
if (top.id !== "int8") {
  console.error(`expected top int8, got ${top.id}`);
  process.exit(1);
}
console.log(
  `✓ WASM version=${wasmV} top=${top.id} LPD=${Number(top.lpd).toFixed(3)} band=${top.band}`,
);
NODE

echo ""
echo "→ npm pack --dry-run (file list)"
PACK_LIST=$(npm pack --dry-run 2>&1)
echo "$PACK_LIST" | tail -50

if ! echo "$PACK_LIST" | grep -q 'wasm/lucy.wasm'; then
  echo "ERROR: npm pack would not include wasm/lucy.wasm — check package.json files[]" >&2
  exit 1
fi
echo "✓ pack includes wasm/lucy.wasm"

if [[ "$DRY_RUN" -eq 1 ]]; then
  echo ""
  echo "Dry run complete — not publishing."
  echo "Tarball preview: (cd js && npm pack)  → openfluke-lucy-${VERSION}.tgz"
  exit 0
fi

echo ""
if ! npm whoami &>/dev/null; then
  echo "Not logged in to npm."
  echo "  npm login"
  echo "Then re-run: bash $JS_DIR/publish.sh"
  exit 1
fi
echo "Logged in as: $(npm whoami)"
echo ""
echo "This will PUBLISH ${NAME}@${VERSION} (public)."
echo "Wasm is built locally and packed — it is not committed to git."
echo ""

if [[ "$ASSUME_YES" -ne 1 ]]; then
  read -r -p "Publish ${NAME}@${VERSION} to npm? [y/N] " reply
  if [[ ! "$reply" =~ ^[Yy]$ ]]; then
    echo "Cancelled."
    exit 0
  fi
fi

npm publish --access public --ignore-scripts
echo ""
echo "✓ Published ${NAME}@${VERSION}"
echo "  https://www.npmjs.com/package/@openfluke/lucy"
echo ""
echo "Install:"
echo "  npm i ${NAME}@${VERSION}"
