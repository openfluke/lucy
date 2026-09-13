# python/ — `openfluke-lucy`

**Version:** `0.2.0`

Wraps the Go **`lucy` CLI** (no Python reimplementation of LPD).

```python
from lucy import build_lpd, version

print(version())
board = build_lpd(samples, options={"keep_floor": 0.7})
```

Binary resolution: `LUCY_BIN` → `lucy/bin/lucy-<os>-<arch>` → `lucy/bin/lucy`.

Rebuild:

```bash
../go/scripts/build-artifacts.sh
```
