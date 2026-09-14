# `lucy serve` HTTP API

```bash
lucy serve :7474
```

| Method | Path | Body / notes |
|--------|------|----------------|
| GET | `/api/version` | |
| GET | `/api/floors` | defaults + formula |
| POST | `/api/lpd` | BuildRequest → board |
| POST | `/api/chart-pack` | SVG + b64 PNG/JPG |
| POST | `/api/pdf` or `/api/site-pdf` | application/pdf |
| POST | `/api/csv` | text/csv |

JS: `createLucyClient("http://127.0.0.1:7474")`  
Python: `LucyClient("http://127.0.0.1:7474")`

Example: `examples/node/03_serve_client.mjs`, `examples/python/03_serve_client.py`.

Next: [05-hosts-and-npm](05-hosts-and-npm.md).
