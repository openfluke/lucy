# lucy

Portable **live-fit measuring** and boards — the synthetic-organism ruler that
Tide / Ocean / River were rebuilding in place.

**Lucy** asks: can a net **learn while it still serves**, in a small box — then
how far that live-fit **condenses** (dtype / quant / arch) without becoming a
Score/MiB trap.

This repo exists so you (and others) stop copying dash math into every host.
One core, many skins.


## What Lucy is / isn’t

**Lucy is not Tide, Ocean, or River.** Those hosts still train, serve, store runs,
and render dashboards. Lucy is the **portable measuring + boards layer** they were
rebuilding in place — so no host owns a fork of the ruler.

| Lucy **is** | Lucy **is not** |
|-------------|-----------------|
| Score · Q · **LPD** · gold/lean/trap bands | A serve+train framework |
| Boards, charts, CSV, site PDF, `lucy serve` | Tide’s permute matrix / live runner |
| Go / wasm / npm / PyPI skins over **one** Go core | River’s results store + full-site UI |
| Floors you tune **without** forking formulas | A drop-in replacement for Ocean/Tide apps |

**Hosts feed finished cells → Lucy returns boards / graphs / PDFs.**  
See [`PORTABLE.md`](PORTABLE.md) · [`docs/tutorials/00-why-lucy.md`](docs/tutorials/00-why-lucy.md).

---

## Version scoreboard

Monorepo tag is the **Lucy release**. Subpackages may lag or ship patch bumps on
their own channel, but goldens in [`testdata/`](testdata/) must stay aligned with
the Go core version they claim.

| Package | Channel | Version | Status | Notes |
|---------|---------|---------|--------|-------|
| **lucy** (monorepo) | git tag / [`VERSION`](VERSION) | `v1.0.1` | **stable** | portable claim · golden freeze · compare grid in site PDF |
| **Go core** `github.com/openfluke/lucy` | Go module | `v1.0.1` | stable | `SitePDF` + compare · `lucy artifacts` · [`PORTABLE.md`](PORTABLE.md) |
| **`@openfluke/lucy`** | npm | `1.0.1` | publish-ready | wasm + native + boards · same goldens |
| **`openfluke-lucy`** | PyPI | `1.0.1` | publish-ready | binary wrap · notebook smoke example |
| **testdata goldens** | in-repo | `1.0.1` | **frozen** (v1.0 vectors) | [`testdata/goldens_lpd_v1.0.json`](testdata/goldens_lpd_v1.0.json) |
| **charts** | in-repo | `1.0.1` | site PDF pack | compare · near · LPD · thru · bands · charts |

**Scoreboard rules**

1. Bump **monorepo** when the ruler or public board API changes.
2. Bump **Go** first; npm / PyPI follow the same formulas (no silent drift).
3. Patch bumps `1.0.x` keep the v1 golden freeze unless formulas change. Pre-1.0 minors added roadmap rows.
4. Keep this table honest when you cut a release (edit [`VERSION`](VERSION) + this table + `js/package.json` + `python/pyproject.toml` + Go tag).

### Deliverable map

| Deliverable | Path | Tied to |
|-------------|------|---------|
| Go measuring core | [`go/`](go/) | Go module version |
| npm `@openfluke/lucy` | [`js/`](js/) | npm version |
| PyPI `openfluke-lucy` | [`python/`](python/) | PyPI version |
| Shared goldens | [`testdata/`](testdata/) | monorepo / Go |
| Chart / PDF / CSV export | [`charts/`](charts/) | monorepo |
| Host examples + migrate | [`examples/`](examples/) · [`MIGRATE.md`](MIGRATE.md) | monorepo |
| Portable claim | [`PORTABLE.md`](PORTABLE.md) | monorepo 1.0 |

---

## Feature roadmap (→ v1)

Legend: **—** not in that cut · **○** scaffold / stub · **◐** partial · **●** done for that version · **◇** planned.

All channels share one story. **Current = 1.0** (stable portable ruler).

### Release line (all versions)

| Cut | Status | Intent |
|-----|--------|--------|
| **0.0** | shipped | Repo + stubs only |
| **0.1** | shipped | Real `BuildLPD` + locked goldens; Go importable |
| **0.2** | shipped | Host-facing boards API + npm/PyPI consuming Go artifacts |
| **0.3** | shipped | Tide/Ocean/River-style charts usable from JS UI |
| **0.4** | shipped | PNG export + chart-pack + golden CI |
| **0.5** | shipped | JPG + HTML report + `lucy serve` HTTP API |
| **0.6** | shipped | Multi-page PDF board report (chart pages) |
| **0.7** | shipped | Host cutover kit: cover+table PDF · CSV · floors · examples / MIGRATE |
| **0.8** | shipped | Publish channels + Tide import switch (`github.com/openfluke/lucy/lucy`) |
| **0.9** | shipped | Tide/River PDF parity (near / LPD / thru / bands + charts) |
| **1.0** | **current** | Portable claim: wasm+binaries documented; no host owns a fork; floors tunable without forks |

### Version ↔ channel matrix (what each cut bumped)

| Cut | Monorepo `VERSION` | Go module | npm `@openfluke/lucy` | PyPI `openfluke-lucy` | Goldens |
|-----|--------------------|-----------|------------------------|------------------------|---------|
| 0.1 | `0.1.0` | `0.1.0` | stub | stub | `goldens_lpd_v0.1.json` |
| 0.2 | `0.2.0` | `0.2.0` | `0.2.0` | `0.2.0` | `v0.2` |
| 0.3 | `0.3.0` | `0.3.0` | `0.3.0` | `0.3.0` | `v0.3` |
| 0.4 | `0.4.0` | `0.4.0` | `0.4.0` | `0.4.0` | `v0.4` |
| 0.5 | `0.5.0` | `0.5.0` | `0.5.0` | `0.5.0` | `v0.5` |
| 0.6 | `0.6.0` | `0.6.0` | `0.6.0` | `0.6.0` | `v0.6` |
| 0.7 | `0.7.0` | `0.7.0` | `0.7.0` | `0.7.0` | [`v0.7`](testdata/goldens_lpd_v0.7.json) |
| 0.8 | `0.8.0` | `0.8.0` | `0.8.0` publish-ready | `0.8.0` publish-ready | [`v0.8`](testdata/goldens_lpd_v0.8.json) |
| 0.9 | `0.9.0` | `0.9.0` | `0.9.0` | `0.9.0` | [`v0.9`](testdata/goldens_lpd_v0.9.json) |
| **1.0** | **`1.0.1`** | **`1.0.1`** | **`1.0.1`** | **`1.0.1`** | **[`v1.0`](testdata/goldens_lpd_v1.0.json)** |

### Monorepo — what ships when

| Feature | …0.6 | 0.7 | 0.8 | 0.9 | 1.0 |
|---------|:----:|:---:|:---:|:---:|:---:|
| Repo layout + `VERSION` + Apache-2.0 | ● | ● | ● | ● | ● |
| Shared goldens + cross-runtime CI | ● | ● | ● | ● | ● |
| Host import guide + examples / MIGRATE | ◐ | ● | ● | ● | ● |
| Chart / JPG / PDF / CSV export path | ● | ● | ● | ● | ● |
| npm + PyPI **published** | — | — | ◐ ready | ◐ ready | ◐ ready |
| Tide imports `github.com/openfluke/lucy/lucy` | — | ◐ docs | ● | ● | ● |
| Tide/River full-site PDF packs | — | ◐ board | ◐ | ● | ● |
| v1 portable claim (no host owns a fork) | — | — | ◐ | ◐ | ● |

### Go core (`github.com/openfluke/lucy`)

| Feature | …0.6 | 0.7 | 0.8 | 0.9 | 1.0 |
|---------|:----:|:---:|:---:|:---:|:---:|
| Score · Q · LPD · bands · options · JSON board | ● | ● | ● | ● | ● |
| Wasm + native cross-binaries | ● | ● | ● | ● | ● |
| SVG / PNG / JPG charts · HTML report · `serve` | ● | ● | ● | ● | ● |
| PDF chart pages | ● | ● | ● | ● | ● |
| PDF cover + LPD table · CSV · floors API | — | ● | ● | ● | ● |
| Go module tagged / discoverable for hosts | ◐ | ◐ | ● | ● | ● |
| River-style near / LPD / thru / bands PDF | — | — | — | ● | ● |
| Documented Tide replace path (live in Tide) | ◐ | ◐ | ● | ● | ● |

### npm `@openfluke/lucy` ([`js/`](js/))

| Feature | …0.6 | 0.7 | 0.8 | 0.9 | 1.0 |
|---------|:----:|:---:|:---:|:---:|:---:|
| Wasm + native measuring · React / CE boards | ● | ● | ● | ● | ● |
| PNG / JPG / PDF / report / serve client | ● | ● | ● | ● | ● |
| CSV + floors helpers | — | ● | ● | ● | ● |
| **Published** to npm | — | — | ◐ ready | ◐ ready | ◐ ready |
| Angular CE polish | ◐ | ◐ | ● | ● | ● |
| Tide/River site PDF client | — | ◐ | ◐ | ● | ● |
| Tunable floors documented end-to-end | ◐ | ● | ● | ● | ● |

### PyPI `openfluke-lucy` ([`python/`](python/))

| Feature | …0.6 | 0.7 | 0.8 | 0.9 | 1.0 |
|---------|:----:|:---:|:---:|:---:|:---:|
| Binary wrap · `build_lpd` · chart / report / PDF | ● | ● | ● | ● | ● |
| CSV · floors · DataFrame helpers | ◐ | ● | ● | ● | ● |
| **Published** to PyPI | — | — | ◐ ready | ◐ ready | ◐ ready |
| Notebook examples | ◐ | ◐ | ◐ | ◐ | ● |
| Tide/River site PDF helpers | — | ◐ | ◐ | ● | ● |

### testdata + charts

| Feature | …0.6 | 0.7 | 0.8 | 0.9 | 1.0 |
|---------|:----:|:---:|:---:|:---:|:---:|
| Locked LPD goldens (Score / Q / LPD / bands) | ● | ● | ● | ● | ● |
| SVG / PNG / JPG / board PDF | ● | ● | ● | ● | ● |
| Cover + table PDF · board.csv in reports | — | ● | ● | ● | ● |
| Tide/River site PDF (near/LPD/thru/bands) | — | ◐ | ◐ | ● | ● |
| v1 golden freeze | — | — | — | ◐ | ● |

Update these tables when a cut ships — same honesty rule as the version scoreboard.

---

## What Lucy measures (defaults)

Defaults match the Tide/Ocean story. **Nothing is frozen forever** — floors,
Availability, and board bands are meant to be wrapped and tweaked for your own
duty-cycle experiments.

| Symbol | Default idea |
|--------|----------------|
| **Availability** | `InferMs / (InferMs + TrainMs) × 100` — SGD that blocks serve dies here |
| **Throughput T** | outputs / second while the sweep is live |
| **Acc** | hard argmax (SoftAcc is serve-confidence, **not** this pillar) |
| **Lucy Score** | `T × Availability × Acc / 10_000` — live-fit |
| **Q** | geomean of Acc/Thru/Avail keep vs **learner** peaks |
| **LPD** | **Lucy Pareto density** = `Q × shrink` vs Acc-champ RAM; **0** unless Acc keep ≥ keep-floor (default 70%) |
| **Gold / lean / trap** | trifecta / high-keep compact / tiny+chance Acc |

**Pareto**: improving Acc, RAM, or live duty forces tradeoffs — goldilocks sits
on the undominated edge. **Density**: how much live-fit you keep per byte vs the
Acc champ.

MobileScore (`Score / WeightMiB`) is the binary trap — use **LPD**.

---

## Architecture

```
                 ┌─────────────────────────┐
                 │   go/  (source of truth) │
                 │   Score · Q · BuildLPD  │
                 └───────────┬─────────────┘
           ┌─────────────────┼─────────────────┐
           ▼                 ▼                 ▼
     wasm / native     wasm / native      native binaries
           │                 │                 │
           ▼                 ▼                 ▼
   @openfluke/lucy     @openfluke/lucy    openfluke-lucy
   (Node / Bun API)    (React / Angular    (PyPI)
                        boards / charts)
           │
           └────────── BYO backend / Tide / Ocean / your sweep
```

- **Do not** reimplement LPD in TypeScript or Python.
- **Do** feed `Sample[]` (or board JSON) from any host.
- **Do** override keep floors, Availability, or ranking when you wrap an experiment.

Hosts (Tide, Ocean, River, `lucy-lab`, a FastAPI service, a Bun script) become
thin: train/serve → samples → Lucy → boards / graphs.

---

## Repo layout

```
lucy/
  go/           # module github.com/openfluke/lucy — formulas + later wasm/bin builds
  js/           # @openfluke/lucy — web + Node/Bun (UI + runtime glue)
  python/       # openfluke-lucy — PyPI wrapper around Go binaries
  testdata/     # golden vectors all runtimes must match
  charts/       # optional static graph / JPG generation (later)
```

---

## Relationship to Welvet / Tide

**v1.0.1:** **portable measuring ruler** (stable) — npm README / docs polish; same v1 golden vectors.

- Claim: [`PORTABLE.md`](PORTABLE.md) — one core, many skins; floors without forks
- Go: `SitePDF` (+ compare mode×dtype×arch) · `lucy artifacts` · Tide imports `github.com/openfluke/lucy/lucy`
- npm / PyPI: publish-ready · same goldens as Go
- Golden freeze: [`testdata/goldens_lpd_v1.0.json`](testdata/goldens_lpd_v1.0.json)
- Host: [`HOST.md`](HOST.md) · [`MIGRATE.md`](MIGRATE.md) · [`PUBLISH.md`](PUBLISH.md)
- CI: [`scripts/check-goldens.sh`](scripts/check-goldens.sh) · [`scripts/publish-check.sh`](scripts/publish-check.sh)

Rebuild: `./go/scripts/build-artifacts.sh`

---

## Quick intent

```go
import "github.com/openfluke/lucy/lucy"

board := lucy.BuildLPD(samples)
_ = board.Top // ranked by LPD, traps at 0
```

```js
import { buildLPD } from "@openfluke/lucy";
import { LucyBoard } from "@openfluke/lucy/react"; // peer: react
const { board } = await buildLPD(samples, { keep_floor: 0.7 });
// <LucyBoard board={board} />
```

```python
from lucy import build_lpd
resp = build_lpd(samples, options={"keep_floor": 0.7})
```

```bash
echo '{"samples":[...]}' | lucy build-lpd
```

---


---

## Tutorials & examples

| Start here | |
|------------|---|
| [Why Lucy?](docs/tutorials/00-why-lucy.md) | portable ruler story |
| [Quickstart](docs/tutorials/01-quickstart.md) | CLI · Go · npm · Python · Jupyter |
| [Examples hub](examples/README.md) | runnable matrix + `./scripts/run-examples.sh` |
| [Docs index](docs/README.md) | tutorials + HOST / PORTABLE / PUBLISH |
| [Release](scripts/release.sh) | `./scripts/release.sh` · `--push` for GitHub tag + assets |

## License

[Apache License 2.0](LICENSE)
