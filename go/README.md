# go/ — Lucy measuring core

Source of truth for **Score**, **Q**, **LPD** (Lucy Pareto density), gold / lean /
trap, and related live-fit math.

## Why Go first

- One implementation → **wasm** + **native binaries**
- npm (`@openfluke/lucy`) and PyPI wrappers call this core — they do not reimplement formulas
- Tide / Ocean / River become *hosts* that import Lucy instead of forking dash math

## Status

Placeholder. Math still lives in [`welvet/lucy`](https://github.com/openfluke/welvet/tree/main/lucy)
until it moves here with shared [`testdata/`](../testdata) goldens.

## Later

- `GOOS=js GOARCH=wasm` (and/or TinyGo) artifacts for `@openfluke/lucy`
- cross-compiled binaries for Python wheels / CLI
- optional board → chart / JPG generation
