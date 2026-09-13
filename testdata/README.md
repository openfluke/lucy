# testdata/

Shared golden vectors for Score / Q / LPD / bands.

| File | Version | Notes |
|------|---------|-------|
| [`goldens_lpd_v0.1.json`](goldens_lpd_v0.1.json) | 0.1.0 | Locked from Go `BuildLPD` (f32/int8/bin/fat fixture) |
| [`samples.example.json`](samples.example.json) | — | Informal example |

Regenerate LPD golden:

```bash
cd go && go run ./cmd/gen-goldens/
```

Every runtime (Go, wasm via npm, Python binary wrapper) must match these.
