# Host import guide (0.5)

Lucy does not train models. Feed finished cells → board (+ charts / report / HTTP).

## Go (replace welvet/lucy)

```go
import "github.com/openfluke/lucy/lucy"

board := lucy.BuildLPD(samples)
_ = lucy.WriteReportDir("out", board, 8)
http.ListenAndServe(":7474", lucy.Handler())
```

## CLI

```bash
lucy build-lpd < request.json
lucy chart-radar-png < request.json > radar.png
lucy chart-radar-jpg < request.json > radar.jpg
lucy report ./out < request.json
lucy serve :7474
```

## Node / Bun / React

```js
import { buildLPD, createLucyClient, writeReportNative } from "@openfluke/lucy";
import { LucyBoard } from "@openfluke/lucy/react";

const { board } = await buildLPD(samples);
const client = createLucyClient("http://127.0.0.1:7474");
```

## Python

```python
from lucy import build_lpd, write_report, LucyClient, chart_jpg

resp = build_lpd(samples)
write_report(samples, "./out")
open("radar.jpg", "wb").write(chart_jpg(samples, "radar"))
```

Rebuild: `./go/scripts/build-artifacts.sh`
