# npm `@openfluke/lucy` — how to consume

## Published (after `npm publish`)

```bash
npm install @openfluke/lucy
```

```js
import {
  buildLPD,           // wasm
  buildLPDNative,     // Node binary
  writeSitePDFNative,
  createLucyClient,
  VERSION,
} from "@openfluke/lucy";
import { LucyBoard } from "@openfluke/lucy/react"; // peer: react
import { registerLucyElements } from "@openfluke/lucy/elements";
```

## Local monorepo (this repo, pre-publish)

Point Node at the package root:

```js
import { buildLPD } from "../../js/src/index.js";
```

Or:

```bash
cd js && npm link
cd /your/app && npm link @openfluke/lucy
```

Runnable demos: [`../node/`](../node/) · browser CE: [`../browser/`](../browser/).
