# Migrating hosts to Lucy 0.7

Lucy is the portable measuring core. Hosts (Tide, Ocean, River, labs) should
**import** it — not copy formulas.

## Go (from `welvet/lucy`)

Change imports:

```go
// before
import "github.com/openfluke/welvet/lucy"

// after
import "github.com/openfluke/lucy/lucy"
```

API surface used by Tide today (`Sample`, `BuildLPD`, `SoftAcc*`, `Finalize`,
floors constants) lives under `github.com/openfluke/lucy/lucy`.

```bash
go get github.com/openfluke/lucy@v0.7.0
```

Tune floors without forking:

```go
board := lucy.BuildLPDWithOptions(samples, lucy.DensityOptions{KeepFloor: 0.75})
```

Or inspect defaults: `lucy floors` / `GET /api/floors`.

## Node / Bun

```js
import { buildLPD, writeCSVNative, writePDFNative } from "@openfluke/lucy";
```

Or point a UI at `lucy serve` via `createLucyClient()`.

## Python

```python
from lucy import build_lpd, board_csv, write_pdf, floors
```

## What 1.0 still owns

- Tide/River **full-site** PDF parity (compare + near + thru packs)
- Documented “no host owns a fork” cutover complete in Tide itself
