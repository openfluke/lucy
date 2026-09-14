# Host import guide (0.6)

Lucy does not train models. Feed finished cells → board (+ charts / HTML / PDF / HTTP).

## Go

```go
import "github.com/openfluke/lucy/lucy"

board := lucy.BuildLPD(samples)
pdf, _ := lucy.BoardPDF(board, 8)
_ = lucy.WriteReportDir("out", board, 8) // includes board.pdf
http.ListenAndServe(":7474", lucy.Handler())
```

## CLI

```bash
lucy build-lpd < request.json
lucy report ./out < request.json
lucy pdf ./board.pdf < request.json
lucy serve :7474   # POST /api/lpd · /api/chart-pack · /api/pdf
```

## Node / Python

```js
import { writePDFNative, createLucyClient } from "@openfluke/lucy";
await writePDFNative(samples, "./board.pdf");
```

```python
from lucy import write_pdf, LucyClient
write_pdf(samples, "board.pdf")
```

Rebuild: `./go/scripts/build-artifacts.sh`
