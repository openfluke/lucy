#!/usr/bin/env bash
# Cut a GitHub release for the VERSION in-repo.
#
#   ./scripts/release.sh              # dry-run (print plan)
#   ./scripts/release.sh --push       # tag + gh release + upload artifacts
#   ./scripts/release.sh --push --skip-examples
#
# Requires: git, gh (authenticated), go. Optional: node/python for example smoke.
set -euo pipefail
ROOT="$(cd "$(dirname "$0")/.." && pwd)"
cd "$ROOT"

PUSH=0
SKIP_EXAMPLES=0
SKIP_BUILD=0
for arg in "$@"; do
  case "$arg" in
    --push) PUSH=1 ;;
    --skip-examples) SKIP_EXAMPLES=1 ;;
    --skip-build) SKIP_BUILD=1 ;;
    -h|--help)
      sed -n '2,12p' "$0"
      exit 0
      ;;
    *)
      echo "unknown arg: $arg" >&2
      exit 2
      ;;
  esac
done

V="$(tr -d '[:space:]' < VERSION)"
TAG="v${V}"
NOTES_FILE="$(mktemp)"
trap 'rm -f "$NOTES_FILE"' EXIT

echo "== Lucy release $TAG =="
bash scripts/publish-check.sh

if [[ "$SKIP_BUILD" -eq 0 ]]; then
  echo "== build artifacts =="
  bash go/scripts/build-artifacts.sh
fi

if [[ "$SKIP_EXAMPLES" -eq 0 ]]; then
  echo "== example smoke =="
  bash scripts/run-examples.sh
fi

# Collect example proof lines for release body
TOP_LINE=""
if [[ -x python/src/lucy/bin/lucy ]]; then
  TOP_LINE="$(python/src/lucy/bin/lucy build-lpd < examples/shared/samples.json 2>/dev/null \
    | python3 -c "import sys,json; b=json.load(sys.stdin); t=b['board']['top'][0]; print(f\"{t['id']} LPD={t['lpd']:.4g} band={t['band']}\")" 2>/dev/null || true)"
fi

cat > "$NOTES_FILE" << MD
## Lucy ${TAG}

Portable **live-fit measuring + boards** — the ruler Tide / Ocean / River import.
**Not** a rewrite of those apps (they still train/serve/UI); Lucy owns Score / Q / LPD / exports.

### What this release is

| Is | Isn’t |
|----|-------|
| One Go core → wasm / CLI / npm / PyPI | Tide permute / live runner |
| Boards · charts · CSV · site PDF · \`lucy serve\` | River store / full dashboard |
| Floors tunable without forking | A drop-in Ocean/Tide product |

See [\`PORTABLE.md\`](https://github.com/openfluke/lucy/blob/${TAG}/PORTABLE.md) · [\`docs/tutorials\`](https://github.com/openfluke/lucy/tree/${TAG}/docs/tutorials).

### Quick proof (shared goldens)

\`\`\`bash
./go/scripts/build-artifacts.sh
./scripts/run-examples.sh
lucy build-lpd < examples/shared/samples.json
# expect top: int8 (gold), bin trap at LPD 0
\`\`\`

$(if [[ -n "$TOP_LINE" ]]; then echo "**Smoke top row:** \`${TOP_LINE}\`"; fi)

### Try it

| Surface | How |
|---------|-----|
| **CLI** | \`lucy site-pdf out.pdf < examples/shared/samples.json\` |
| **Go** | \`import "github.com/openfluke/lucy/lucy"\` → \`BuildLPD\` / \`SitePDF\` |
| **npm** | \`npm i @openfluke/lucy\` · see \`examples/node/\` · \`examples/npm/\` |
| **Python** | \`pip install openfluke-lucy\` · \`examples/python/\` · Jupyter \`lucy_tutorial.ipynb\` |
| **Browser** | \`examples/browser/\` + \`python3 -m http.server\` |
| **HTTP** | \`lucy serve :7474\` → \`/api/lpd\` · \`/api/site-pdf\` · \`/api/floors\` |

Tutorials: [\`docs/tutorials/01-quickstart.md\`](https://github.com/openfluke/lucy/blob/${TAG}/docs/tutorials/01-quickstart.md)

### Artifacts attached

Cross binaries + \`lucy.wasm\` from \`./go/scripts/build-artifacts.sh\`.

### Channels

| Channel | Version |
|---------|---------|
| Monorepo / Go module | \`${TAG}\` |
| npm \`@openfluke/lucy\` | \`${V}\` (publish separately if needed) |
| PyPI \`openfluke-lucy\` | \`${V}\` (publish separately if needed) |
| Goldens | \`testdata/goldens_lpd_v${V%.*}.json\` |

MD

echo
echo "---- release notes preview ----"
cat "$NOTES_FILE"
echo "---- end preview ----"
echo

ASSETS=()
for f in artifacts/lucy-linux-amd64 artifacts/lucy-linux-arm64 \
         artifacts/lucy-darwin-amd64 artifacts/lucy-darwin-arm64 \
         artifacts/lucy-windows-amd64.exe artifacts/lucy.wasm; do
  [[ -f "$f" ]] && ASSETS+=("$f")
done

if [[ "$PUSH" -eq 0 ]]; then
  echo "Dry-run only. Re-run with --push to tag + create GitHub release."
  echo "Would: git tag ${TAG} && git push origin ${TAG}"
  echo "Would: gh release create ${TAG} with ${#ASSETS[@]} assets"
  exit 0
fi

# Ensure clean-ish tree warning
if [[ -n "$(git status --porcelain)" ]]; then
  echo "WARNING: working tree not clean. Continue anyway in 3s…" >&2
  sleep 3
fi

if git rev-parse "$TAG" >/dev/null 2>&1; then
  echo "Tag ${TAG} already exists locally."
else
  git tag -a "$TAG" -m "Lucy ${TAG}"
fi

git push origin "refs/tags/${TAG}"

if gh release view "$TAG" >/dev/null 2>&1; then
  echo "Release ${TAG} already exists — uploading assets / editing notes…"
  gh release upload "$TAG" "${ASSETS[@]}" --clobber
  gh release edit "$TAG" --notes-file "$NOTES_FILE"
else
  gh release create "$TAG" "${ASSETS[@]}" \
    --title "Lucy ${TAG}" \
    --notes-file "$NOTES_FILE"
fi

echo
echo "OK → $(gh release view "$TAG" --json url -q .url)"
