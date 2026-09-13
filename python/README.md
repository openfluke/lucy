# python/ — `openfluke-lucy`

**Version:** `0.3.0`

```python
from lucy import build_lpd, board_records, chart_svg, to_dataframe

resp = build_lpd(samples, options={"keep_floor": 0.7})
rows = board_records(resp)       # list[dict] for notebooks
svg = chart_svg(samples, "radar")  # Go SVG
# df = to_dataframe(resp)        # needs pandas
```
