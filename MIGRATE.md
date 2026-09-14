# Migrating hosts to Lucy 0.8

## Go (from `welvet/lucy`) — **Tide has switched**

```go
// before
import "github.com/openfluke/welvet/lucy"

// after
import "github.com/openfluke/lucy/lucy"
```

```
require github.com/openfluke/lucy v0.8.0
replace github.com/openfluke/lucy => ../lucy/go   // local monorepo
```

Tune floors without forking:

```go
board := lucy.BuildLPDWithOptions(samples, lucy.DensityOptions{KeepFloor: 0.75})
```

## Node / Python / publish

See [`PUBLISH.md`](PUBLISH.md) and [`examples/`](examples/).

## What 1.0 still owns

- Portable claim + published channels as default
- Optional Tide compare mode×dtype PDF grids inside Lucy

- **0.9** Tide/River full-site PDF parity
- **1.0** portable claim + published channels as the default story
