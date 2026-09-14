# Migrating hosts to Lucy 1.0

## Go (from `welvet/lucy`) — **done in Tide**

```go
import "github.com/openfluke/lucy/lucy"
```

```
require github.com/openfluke/lucy v1.0.1
replace github.com/openfluke/lucy => ../lucy/go   // local
```

Tune floors without forking:

```go
board := lucy.BuildLPDWithOptions(samples, lucy.DensityOptions{KeepFloor: 0.75})
```

## Portable claim

See [`PORTABLE.md`](PORTABLE.md). Do not copy formulas into hosts.

## Node / Python

```js
import { buildLPD, writeSitePDFNative } from "@openfluke/lucy";
```

```python
from lucy import build_lpd, write_site_pdf, floors
```
