# Hosts, npm, PyPI

## Go host (Tide pattern)

```go
import "github.com/openfluke/lucy/lucy"
board := lucy.BuildLPD(samples)
```

```
require github.com/openfluke/lucy v1.0.1
replace github.com/openfluke/lucy => ../lucy/go  // monorepo
```

## npm (published shape)

```bash
npm install @openfluke/lucy
```

```js
import { buildLPD, writeSitePDFNative, LucyBoard } from "@openfluke/lucy";
import { registerLucyElements } from "@openfluke/lucy/elements";
```

Until publish: point at this repo’s `js/` (see `examples/npm/README.md`).

## Angular CE

```html
<lucy-board [attr.board-json]="board | json"></lucy-board>
```

See `js/examples/angular.md` and `examples/browser/index.html`.

## PyPI

```bash
pip install openfluke-lucy
```

Local: `PYTHONPATH=python/src` + binary from `python/src/lucy/bin/lucy`.

## Portable rule

Do **not** reimplement LPD in TS/Python. See [`PORTABLE.md`](../../PORTABLE.md).
