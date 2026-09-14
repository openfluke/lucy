# Portable claim (Lucy 1.0)

> **Not a Tide/Ocean/River rewrite.** Lucy is the measuring ruler + boards/exports.
> Hosts still own training, serving, and product UI — they **import** this package.

**No host owns a fork of the ruler.** Tide / Ocean / River / labs **import**
`github.com/openfluke/lucy/lucy` (or `@openfluke/lucy` / `openfluke-lucy`).
Formulas live here. Floors are knobs, not copy-paste.

## One measuring core

| Runtime | How you consume |
|---------|-----------------|
| Go | `import "github.com/openfluke/lucy/lucy"` → `BuildLPD` / `BuildLPDWithOptions` |
| Browser / Node wasm | `@openfluke/lucy` → `buildLPD` |
| Node / Bun native | `@openfluke/lucy` → `buildLPDNative` / CLI binary |
| Python | `openfluke-lucy` → wraps the same Go CLI (no Python re-impl) |
| HTTP | `lucy serve` → `/api/lpd` · `/api/site-pdf` · `/api/floors` · … |

## Artifacts (rebuild anytime)

```bash
./go/scripts/build-artifacts.sh
```

Ships into `artifacts/` + `js/wasm/` + `python/src/lucy/bin/`:

| Artifact | Use |
|----------|-----|
| `lucy-linux-amd64` / `arm64` | CLI + PyPI / native Node |
| `lucy-darwin-amd64` / `arm64` | macOS CLI |
| `lucy-windows-amd64.exe` | Windows CLI |
| `lucy.wasm` + `wasm_exec.js` | browser / Node wasm |

Inventory helper: `lucy artifacts` (lists expected names + this version).

## Floors without forks

```go
lucy.BuildLPDWithOptions(samples, lucy.DensityOptions{KeepFloor: 0.75})
```

```js
await buildLPD(samples, { keep_floor: 0.75 });
```

```python
build_lpd(samples, options={"keep_floor": 0.75})
```

Defaults: `lucy floors` / `GET /api/floors` / [`FloorsMap()`](go/lucy/options.go).

## Host cutover

Tide already imports this module (`replace => ../lucy/go` for local monorepos).
See [`MIGRATE.md`](MIGRATE.md). Do **not** reintroduce `welvet/lucy` copies.

## Stability

`testdata/goldens_lpd_v1.0.json` is the **v1 golden freeze**. Patch releases
`1.0.x` must keep these vectors unless a documented formula change ships a new
minor/major.
