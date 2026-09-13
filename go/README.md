# go/ — Lucy measuring core

**Version:** `0.2.0` (`lucy.Version`)

Source of truth for **Score**, **Q**, **LPD**, gold / lean / trap.

```bash
cd go && go test ./lucy/
./scripts/build-artifacts.sh   # from repo: go/scripts/…
```

```go
import "github.com/openfluke/lucy/lucy"

board := lucy.BuildLPD(samples)
board = lucy.BuildLPDWithOptions(samples, lucy.DensityOptions{KeepFloor: 0.7})
resp, err := lucy.BuildFromJSON(raw) // stable host JSON schema
```

## Artifacts

| Target | Command |
|--------|---------|
| Native CLI | `cmd/lucy` — `lucy version` / `lucy build-lpd` |
| Wasm | `cmd/lucywasm` — `lucyBuildLPD` / `lucyVersion` globals |

Goldens: [`../testdata/goldens_lpd_v0.2.json`](../testdata/goldens_lpd_v0.2.json)
