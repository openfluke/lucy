# charts/

**0.3.0:** SVG generation lives in Go (`lucy.RadarSVG` / `ScatterSVG` / `BarsSVG`)
and JS canvas helpers (`@openfluke/lucy/charts`). CLI:

```bash
lucy chart-radar < request.json > radar.svg
```

**JPG / PDF parity** with Tide/River full report set → targeted for **1.0**.
