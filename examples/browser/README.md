# Browser custom elements

Serve the repo root (wasm fetch needs HTTP):

```bash
cd /path/to/lucy
python3 -m http.server 8765
# open http://127.0.0.1:8765/examples/browser/
```

Shows `<lucy-board>` with radars + table from wasm `buildLPD`.
