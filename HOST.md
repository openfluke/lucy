# Host import guide (0.7)

Lucy does not train models. Feed finished cells → board (+ charts / HTML / PDF / CSV / HTTP).

## Go

```go
import "github.com/openfluke/lucy/lucy"

board := lucy.BuildLPD(samples)
_ = lucy.BoardCSV(board)
pdf, _ := lucy.BoardPDF(board, 8) // cover + table + charts
_ = lucy.WriteReportDir("out", board, 8)
http.ListenAndServe(":7474", lucy.Handler())
```

## CLI

```bash
lucy floors
lucy build-lpd < request.json
lucy csv < request.json > board.csv
lucy report ./out < request.json
lucy pdf ./board.pdf < request.json
lucy serve :7474
# GET /api/floors · POST /api/lpd · /api/chart-pack · /api/pdf · /api/csv
```

## Node / Python

```js
import { writeCSVNative, writePDFNative, floorsNative, createLucyClient } from "@openfluke/lucy";
await floorsNative();
await writeCSVNative(samples, "./board.csv");
```

```python
from lucy import board_csv, write_pdf, floors, LucyClient
floors()
board_csv(samples)
```

Migration from `welvet/lucy`: [`MIGRATE.md`](MIGRATE.md).  
Examples: [`examples/`](examples/).  
Rebuild: `./go/scripts/build-artifacts.sh`
