# Host import guide (0.8)

Lucy does not train models. Feed finished cells → board (+ charts / HTML / PDF / CSV / HTTP).

## Go (Tide path)

```go
import "github.com/openfluke/lucy/lucy"

board := lucy.BuildLPD(samples)
```

Dev replace in host `go.mod`:

```
require github.com/openfluke/lucy v0.8.0
replace github.com/openfluke/lucy => ../lucy/go
```

See [`MIGRATE.md`](MIGRATE.md) · [`PUBLISH.md`](PUBLISH.md).

## CLI / serve / npm / Python

Same as 0.7 (`floors`, `csv`, `pdf`, `report`, `serve`) plus Angular CE `board-json` attrs.

Rebuild: `./go/scripts/build-artifacts.sh`
