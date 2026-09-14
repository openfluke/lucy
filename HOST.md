# Host import guide (0.9)

## Go

```go
import "github.com/openfluke/lucy/lucy"

board := lucy.BuildLPD(samples)
pdf, _ := lucy.SitePDF(board, 8) // near · LPD · thru · bands · charts
```

## CLI / HTTP

```bash
lucy site-pdf ./site.pdf < request.json
lucy serve :7474   # POST /api/site-pdf · /api/pdf
```

Rebuild: `./go/scripts/build-artifacts.sh`
