# python/ — `openfluke-lucy`

**Version:** `0.4.0`

```python
from lucy import build_lpd, chart_svg, chart_png, chart_pack

resp = build_lpd(samples)
open("radar.png", "wb").write(chart_png(samples, "radar"))
pack = chart_pack(samples)  # svg + png_b64
```
