# Why Lucy?

## The problem

Hosts (Tide, Ocean, River, labs, notebooks) kept **copying** live-fit math:

- Lucy Score = `T × Availability × Acc / 10_000`
- Consciousness **Q** and **LPD** (Lucy Pareto density)
- Gold / lean / trap bands

Every fork drifts. Goldens disagree. PDFs lie.

## The answer

**One measuring core** — `github.com/openfluke/lucy` — consumed as:

| Skin | Package |
|------|---------|
| Go | `github.com/openfluke/lucy/lucy` |
| npm | `@openfluke/lucy` (wasm + native CLI) |
| PyPI | `openfluke-lucy` (wraps the same Go binary) |
| CLI / HTTP | `lucy` / `lucy serve` |

Hosts **feed samples** (finished cells). Lucy returns boards, charts, CSV, site PDFs.

## Why LPD (not Score/MiB)

Tiny dtypes at chance Acc look “dense” on Score/MiB. **LPD is 0** unless Acc keep ≥ keep-floor (default 70%). Goldilocks = live-fit that still fits in a small box.

## Floors without forks

Tune experiments with options — do not fork the formula package:

```json
{ "keep_floor": 0.75, "gold_keep": 0.8, "lean_keep": 0.95 }
```

Next: [01-quickstart](01-quickstart.md).
