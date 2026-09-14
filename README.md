# lucy

Portable **live-fit measuring** and boards — the synthetic-organism ruler that
Tide / Ocean / River were rebuilding in place.

**Lucy** asks: can a net **learn while it still serves**, in a small box — then
how far that live-fit **condenses** (dtype / quant / arch) without becoming a
Score/MiB trap.

This repo exists so you (and others) stop copying dash math into every host.
One core, many skins.

---

## Version scoreboard

Monorepo tag is the **Lucy release**. Subpackages may lag or ship patch bumps on
their own channel, but goldens in [`testdata/`](testdata/) must stay aligned with
the Go core version they claim.

| Package | Channel | Version | Status | Notes |
|---------|---------|---------|--------|-------|
| **lucy** (monorepo) | git tag / [`VERSION`](VERSION) | `v0.6.0` | PDF reports | multi-page JPEG-embedded PDF + serve `/api/pdf` |
| **Go core** `github.com/openfluke/lucy` | Go module | `v0.6.0` | usable | + `BoardPDF` · `lucy pdf` · report includes `board.pdf` |
| **`@openfluke/lucy`** | npm | `0.6.0` | ready (unpublished) | + `writePDFNative` · `createLucyClient().pdf()` |
| **`openfluke-lucy`** | PyPI | `0.6.0` | ready (unpublished) | + `write_pdf` · `LucyClient.pdf()` |
| **testdata goldens** | in-repo | `0.6.0` | locked | [`testdata/goldens_lpd_v0.6.json`](testdata/goldens_lpd_v0.6.json) |
| **charts** | in-repo | `0.6.0` | SVG/PNG/JPG/PDF | Tide/River full PDF parity still → 1.0 |

**Scoreboard rules**

1. Bump **monorepo** when the ruler or public board API changes.
2. Bump **Go** first; npm / PyPI follow the same formulas (no silent drift).
3. `0.0.x` = scaffold / unstable. `0.1.0` = first usable `BuildLPD` + goldens. `0.4.0` = PNG chart-pack + golden CI. `0.5.0` = JPG + report + serve. `0.6.0` = multi-page PDF. `1.0.0` = portable claim + Tide/River PDF parity.
4. Keep this table honest when you cut a release (edit [`VERSION`](VERSION) + this table + `js/package.json` + `python/pyproject.toml` + Go tag).

### Deliverable map

| Deliverable | Path | Tied to |
|-------------|------|---------|
| Go measuring core | [`go/`](go/) | Go module version |
| npm `@openfluke/lucy` | [`js/`](js/) | npm version |
| PyPI `openfluke-lucy` | [`python/`](python/) | PyPI version |
| Shared goldens | [`testdata/`](testdata/) | monorepo / Go |
| Chart / JPG export | [`charts/`](charts/) | monorepo (later) |

---

## Feature roadmap (→ v1)

Legend: **—** not in that cut · **○** scaffold / stub · **◐** partial · **●** done for that version.

Milestone meaning (all channels share the story; patch bumps `0.x.y` do not add rows):

| Cut | Intent |
|-----|--------|
| **0.0** | Repo + stubs only |
| **0.1** | Real `BuildLPD` + locked goldens; Go importable |
| **0.2** | Host-facing boards API + npm/PyPI consuming Go artifacts |
| **0.3** | Tide/Ocean/River-style charts usable from JS UI |
| **0.4** | PNG export + PDF-ready chart pack + golden CI |
| **0.5** | JPG + HTML report + `lucy serve` HTTP API |
| **0.6** | Multi-page PDF board report ← **current** |
| **1.0** | Portable ruler: wasm + binaries documented; one real host path; formulas tunable without forks |

### Monorepo — what ships when

| Feature | 0.0 | 0.1 | 0.2 | 0.3 | 1.0 |
|---------|:---:|:---:|:---:|:---:|:---:|
| Repo layout (`go/` `js/` `python/` `testdata/` `charts/`) | ● | ● | ● | ● | ● |
| Version scoreboard + `VERSION` | ● | ● | ● | ● | ● |
| Apache-2.0 `LICENSE` | ● | ● | ● | ● | ● |
| Shared golden vectors (Score / Q / LPD / bands) | ○ | ● | ● | ● | ● |
| Cross-runtime golden CI (Go ↔ wasm ↔ Python bin) | — | ◐ | ● | ● | ● |
| Formula / floor knobs documented (keep, gold, lean, Availability) | — | ◐ | ● | ● | ● |
| Host import guide (Tide / BYO backend) | — | — | ◐ | ● | ● |
| Chart / JPG export path | — | — | — | ◐ | ● |
| v1 portable claim (no host owns a fork of the ruler) | — | — | — | — | ● |

### Go core (`github.com/openfluke/lucy`)

| Feature | 0.0 | 0.1 | 0.2 | 0.3 | 1.0 |
|---------|:---:|:---:|:---:|:---:|:---:|
| Package stub (`Sample`, `BuildLPD` placeholder) | ● | ● | ● | ● | ● |
| Port from [welvet/lucy](https://github.com/openfluke/welvet/tree/main/lucy) | — | ● | ● | ● | ● |
| Lucy Score (`T × Avail × Acc / 10_000`) | — | ● | ● | ● | ● |
| SoftAcc / Finalize pulse helpers | — | ● | ● | ● | ● |
| Consciousness **Q** (geomean keep vs learner peaks) | — | ● | ● | ● | ● |
| **LPD** (Lucy Pareto density = `Q × shrink`, keep floor) | — | ● | ● | ● | ● |
| Gold / near / lean / trap bands | — | ● | ● | ● | ● |
| Tunable options (floors, shrink cap, Availability hook) | — | ◐ | ● | ● | ● |
| Stable JSON board schema for hosts / UI | — | ◐ | ● | ● | ● |
| Wasm build (`GOOS=js` / TinyGo) | — | — | ● | ● | ● |
| Native cross-binaries (linux/mac/win, amd64/arm64) | — | — | ● | ● | ● |
| Optional board → chart / JPG (Go render) | — | — | — | ◐ | ● |
| Documented replace path for `welvet/lucy` / Tide | — | — | ◐ | ● | ● |

### npm `@openfluke/lucy` ([`js/`](js/))

| Feature | 0.0 | 0.1 | 0.2 | 0.3 | 1.0 |
|---------|:---:|:---:|:---:|:---:|:---:|
| Package scaffold + TS entry stub | ● | ● | ● | ● | ● |
| Published to npm | — | — | ● | ● | ● |
| Measuring API via **wasm** (same goldens as Go) | — | — | ● | ● | ● |
| Measuring API via **native binary** (Node/Bun optional) | — | — | ◐ | ● | ● |
| Node / Bun backend usage (no UI required) | — | — | ● | ● | ● |
| Framework-agnostic board types / JSON | — | ◐ | ● | ● | ● |
| React board components (LPD / radars / tables) | — | — | — | ● | ● |
| Angular-friendly exports (or wrappers) | — | — | — | ◐ | ● |
| Vanilla JS chart helpers | — | — | — | ● | ● |
| BYO backend / poll-your-API examples | — | — | ◐ | ● | ● |
| Tunable floors from JS (pass options through to Go) | — | — | ◐ | ● | ● |

### PyPI `openfluke-lucy` ([`python/`](python/))

| Feature | 0.0 | 0.1 | 0.2 | 0.3 | 1.0 |
|---------|:---:|:---:|:---:|:---:|:---:|
| Package scaffold (`build_lpd` stub) | ● | ● | ● | ● | ● |
| Published to PyPI | — | — | ● | ● | ● |
| Wraps Go **binaries** (no Python re-impl of LPD) | — | — | ● | ● | ● |
| `build_lpd` / Score / Q parity with goldens | — | — | ● | ● | ● |
| Notebook-friendly board dict / DataFrame helpers | — | — | ◐ | ● | ● |
| Tunable options passthrough | — | — | ◐ | ● | ● |
| Optional chart export helpers (call Go/charts) | — | — | — | ◐ | ● |

### testdata + charts

| Feature | 0.0 | 0.1 | 0.2 | 0.3 | 1.0 |
|---------|:---:|:---:|:---:|:---:|:---:|
| Example samples JSON | ● | ● | ● | ● | ● |
| Locked goldens (Acc champ, trap, gold, lean) | — | ● | ● | ● | ● |
| Golden suite covers Score, Q, LPD, bands | — | ● | ● | ● | ● |
| charts/ stub README | ● | ● | ● | ● | ● |
| Static graph / JPG from board | — | — | — | ◐ | ● |
| PDF-friendly chart set (Tide/River parity subset) | — | — | — | — | ● |

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

**v0.6.0:** measuring + boards + **PDF**.

- Go: `BoardPDF` · `lucy pdf out.pdf` · `POST /api/pdf` · HTML report also writes `board.pdf`
- npm: `writePDFNative` · `createLucyClient().pdf()`
- Python: `write_pdf` · `LucyClient.pdf()`
- CI: [`scripts/check-goldens.sh`](scripts/check-goldens.sh)

Host guide: [`HOST.md`](HOST.md). Rebuild: `./go/scripts/build-artifacts.sh`

Goldens: [`testdata/goldens_lpd_v0.6.json`](testdata/goldens_lpd_v0.6.json).

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

## License

[Apache License 2.0](LICENSE)
