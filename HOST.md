# Host import guide (1.0)

Lucy is the portable measuring core. Feed finished cells → board / site PDF / CSV / HTTP.

```go
import "github.com/openfluke/lucy/lucy"

board := lucy.BuildLPDWithOptions(samples, lucy.DensityOptions{KeepFloor: 0.7})
pdf, _ := lucy.SitePDF(board, 8) // compare · near · LPD · thru · bands · charts
```

```bash
lucy floors
lucy artifacts
lucy site-pdf ./site.pdf < request.json
lucy serve :7474
```

Portable claim: [`PORTABLE.md`](PORTABLE.md) · Migrate: [`MIGRATE.md`](MIGRATE.md) · Publish: [`PUBLISH.md`](PUBLISH.md)
