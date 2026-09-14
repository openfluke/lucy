# Lucy examples (1.0)

All demos use [`shared/samples.json`](shared/samples.json) (v1 goldens).

| Area | Path | Run |
|------|------|-----|
| **Why / tutorials** | [`../docs/tutorials/`](../docs/tutorials/) | read 00→05 |
| **CLI** | [`cli/`](cli/) | `./examples/cli/run.sh` |
| **Go** | [`go/`](go/) | `cd examples/go && go run ./01_build_lpd` |
| **Node wasm / native / HTTP** | [`node/`](node/) | `node 01_wasm_build.mjs` … |
| **Browser CE** | [`browser/`](browser/) | `python3 -m http.server` → `/examples/browser/` |
| **Python** | [`python/`](python/) | `PYTHONPATH=python/src python3 …` |
| **Jupyter** | [`python/lucy_tutorial.ipynb`](python/lucy_tutorial.ipynb) | open in Jupyter |
| **npm how-to** | [`npm/`](npm/) | install / link guide |

**Smoke all runnable paths:**

```bash
./scripts/run-examples.sh
```

Outputs land in [`out/`](out/) (gitignored).
