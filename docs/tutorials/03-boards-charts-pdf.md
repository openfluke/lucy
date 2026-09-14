# Boards, charts, CSV, site PDF

After `BuildLPD` you get a board: champs, `top[]` with `lpd` / `q` / `band`, gold/lean/trap.

## Charts

| Output | Go / CLI | JS | Python |
|--------|----------|----|--------|
| SVG | `lucy chart-radar` | canvas helpers | `chart_svg` |
| PNG/JPG | `chart-*-png` | `chartPNGNative` | `chart_png` / `chart_jpg` |
| Pack | `chart-pack` | `chartPackNative` | `chart_pack` |

## CSV

```bash
lucy csv < samples.json > board.csv
```

## Site PDF (River parity)

cover · compare(mode×dtype×arch) · Acc-keep · LPD · thru · bands · chart pages:

```bash
lucy site-pdf out/site.pdf < samples.json
lucy report out/report < samples.json   # HTML + csv + pdf
```

JS: `writeSitePDFNative` · Python: `write_site_pdf`.

Next: [04-serve-http](04-serve-http.md).
