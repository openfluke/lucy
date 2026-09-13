# Host import guide (0.2)

Lucy does not train models. A host feeds finished cells and gets a board back.

## Go (Tide / Ocean / River / labs)

```go
import "github.com/openfluke/lucy/lucy"

board := lucy.BuildLPD(samples)
// or tunable:
board = lucy.BuildLPDWithOptions(samples, lucy.DensityOptions{KeepFloor: 0.7})
```

Replace `github.com/openfluke/welvet/lucy` when you are ready — same types.

## JSON (any language)

```json
{
  "samples": [
    {"id": "f32", "avg_accuracy": 90, "throughput": 200, "availability": 40, "score": 100, "ram_kib": 1000}
  ],
  "options": {"keep_floor": 0.7}
}
```

```bash
lucy build-lpd < request.json
```

## Node / Bun

```js
import { buildLPD } from "@openfluke/lucy";
const { board } = await buildLPD(samples);
```

## Python

```python
from lucy import build_lpd
resp = build_lpd(samples)
```

Rebuild artifacts after Go changes: `./go/scripts/build-artifacts.sh`
