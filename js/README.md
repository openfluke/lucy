# `@openfluke/lucy`

npm package for **web / Node / Bun**: measuring math (via Go **wasm** + later
**native binaries**) and portable boards/charts for **React, Angular, or plain JS**.

## Role

| Layer | What |
|-------|------|
| Backend (Node/Bun) | `buildLPD`, Score / Q / bands — same ruler as Go |
| Frontend | Tide/Ocean/River-style graphs — BYO data or poll your API |
| Artifacts | wasm first; optional platform binaries later |

You bring the train/serve host. Lucy does not hardcode Tide paths, MNIST, or
Welvet modes — feed `Sample[]` (or a board JSON) and tweak floors/formulas when
you wrap your own duty-cycle experiment.

## Status

Scaffold only. Not published yet.

```bash
cd js
# npm publish — later, after wasm + goldens
```

## Layout (planned)

```
js/
  src/           # TS API + UI entrypoints
  dist/          # build output (+ wasm)
```
