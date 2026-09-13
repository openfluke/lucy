# `@openfluke/lucy`

**Version:** `0.2.0`

Node / Bun measuring via **Go wasm**. Board types for React / Angular / vanilla
(UI components land at 0.3).

```bash
cd js && npm test
```

```js
import { buildLPD } from "@openfluke/lucy";

const { board } = await buildLPD(samples, { keep_floor: 0.7 });
console.log(board.top[0].lpd);
```

Rebuild wasm after Go changes:

```bash
npm run build   # runs go/scripts/build-artifacts.sh
```
