# go/ — Lucy measuring core

**Version:** `0.1.0` (`lucy.Version`)

Source of truth for **Score**, **Q**, **LPD** (Lucy Pareto density), gold / lean /
trap, and related live-fit math. Ported from `welvet/lucy` for a portable module.

```bash
cd go
go test ./lucy/
```

```go
import "github.com/openfluke/lucy/lucy"

board := lucy.BuildLPD(samples)
_ = board.Top[0].LPD
```

Goldens: [`../testdata/goldens_lpd_v0.1.json`](../testdata/goldens_lpd_v0.1.json)

## Later (0.2+)

- wasm + native binaries for `@openfluke/lucy` / PyPI
- optional board → chart / JPG
