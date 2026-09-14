# Host import guide (0.4)

Lucy does not train models. Feed finished cells → get a board (+ optional charts).

## Replace welvet/lucy / Tide

```go
import "github.com/openfluke/lucy/lucy" // was github.com/openfluke/welvet/lucy

board := lucy.BuildLPD(samples)
svg := lucy.RadarSVG("Consciousness", lucy.ConsciousnessSeries(board, 8))
```

Tide dash/PDF can keep drawing; point measuring at this module when ready.

## JSON / CLI

```bash
lucy build-lpd < request.json
lucy chart-radar < request.json > radar.svg
lucy chart-radar-png < request.json > radar.png
lucy chart-pack < request.json
```

## Node / Bun / React

```js
import { buildLPD, buildLPDNative } from "@openfluke/lucy";
import { LucyBoard } from "@openfluke/lucy/react";
const { board } = await buildLPD(samples, { keep_floor: 0.7 });
```

## Angular

```js
import { registerLucyElements } from "@openfluke/lucy/elements";
registerLucyElements();
// <lucy-board> then el.board = board;  (+ CUSTOM_ELEMENTS_SCHEMA)
```

## Python

```python
from lucy import build_lpd, board_records, chart_svg
resp = build_lpd(samples)
svg = chart_svg(samples, "bars")
```

Rebuild artifacts: `./go/scripts/build-artifacts.sh`
