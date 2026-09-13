# `@openfluke/lucy`

**Version:** `0.1.0` (unpublished)

npm package for **web / Node / Bun**: board **types** now; measuring via Go
**wasm** + **native binaries** at **0.2+**. UI boards (React / Angular / vanilla)
target **0.3**.

0.1 does **not** reimplement LPD in TypeScript — call Go `BuildLPD` or wait for wasm.

```ts
import { VERSION, KEEP_FLOOR, type Sample } from "@openfluke/lucy";
```
