# go/ — Lucy measuring core

**Version:** `0.3.0`

Measuring + SVG boards (Tide-style). JPG export still later.

```bash
cd go && go test ./lucy/
../go/scripts/build-artifacts.sh
```

```go
board := lucy.BuildLPD(samples)
svg := lucy.RadarSVG("Consciousness", lucy.ConsciousnessSeries(board, 8))
```

CLI: `lucy build-lpd` · `lucy chart-radar|chart-scatter|chart-bars`
