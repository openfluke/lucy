# Samples & floors

## Sample shape (wire JSON)

Hosts map their cells onto:

| Field | Meaning |
|-------|---------|
| `id` | Cell id |
| `avg_accuracy` / `acc` | Hard Acc % |
| `throughput` / `thru` | outputs / s |
| `availability` / `avail` | duty % |
| `score` | Lucy Score (or leave 0 — board recomputes relations) |
| `ram_kib` | weight RAM |
| `mode` `dtype` `arch` | compare grid labels |

CamelCase aliases work in JS (`ramKiB`, `dType`).

## Why floors

| Knob | Default | Effect |
|------|---------|--------|
| `keep_floor` | 0.70 | Acc keep required for LPD > 0 |
| `gold_keep` | 0.80 | pillar keep for gold/near |
| `lean_keep` | 0.95 | Acc keep for lean band |
| `gold_ram` | 0.20 | max RAM frac of Acc champ for gold |
| `near_ram` | 0.50 | near band RAM |
| `shrink_cap` | 32 | max shrink multiplier |

Inspect defaults: `lucy floors` or `GET /api/floors`.

## Tunable call

```go
lucy.BuildLPDWithOptions(samples, lucy.DensityOptions{KeepFloor: 0.75})
```

```js
await buildLPD(samples, { keep_floor: 0.75 });
```

```python
build_lpd(samples, options={"keep_floor": 0.75})
```

Next: [03-boards-charts-pdf](03-boards-charts-pdf.md).
