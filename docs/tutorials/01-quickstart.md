# Quickstart (all runtimes)

Shared input: [`examples/shared/samples.json`](../../examples/shared/samples.json).

## CLI (fastest proof)

```bash
cd /path/to/lucy
./go/scripts/build-artifacts.sh   # once
./artifacts/lucy-$(go env GOOS)-$(go env GOARCH) version
./artifacts/lucy-$(go env GOOS)-$(go env GOARCH) floors
./artifacts/lucy-$(go env GOOS)-$(go env GOARCH) build-lpd < examples/shared/samples.json | head
./artifacts/lucy-$(go env GOOS)-$(go env GOARCH) site-pdf examples/out/site.pdf < examples/shared/samples.json
```

Expect top LPD id **`int8`** (gold), **`bin`** at LPD 0 (trap).

## Go

```bash
cd examples/go && go run ./01_build_lpd
```

## Node (wasm — npm-shaped)

```bash
cd examples/node && node 01_wasm_build.mjs
```

## Node (native binary)

```bash
node 02_native_full.mjs
```

## Python

```bash
PYTHONPATH=../../python/src python3 examples/python/01_build_lpd.py
```

## Jupyter

Open [`examples/python/lucy_tutorial.ipynb`](../../examples/python/lucy_tutorial.ipynb).

## Smoke everything

```bash
./scripts/run-examples.sh
```

Next: [02-samples-and-floors](02-samples-and-floors.md).
