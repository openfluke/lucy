# `@openfluke/lucy`

**Version:** `1.0.1`

Portable **live-fit measuring** for JS — Score, Q, **LPD**, gold/lean/trap bands,
boards, charts, CSV, and site PDF. Same Go core as Tide / Ocean / River hosts
import; this package is the **npm / wasm / Node** skin.

> Lucy is **not** a train+serve framework. Hosts feed finished cells → Lucy
> returns boards / graphs / PDFs. See [`../PORTABLE.md`](../PORTABLE.md) and
> [`../docs/tutorials/00-why-lucy.md`](../docs/tutorials/00-why-lucy.md).

| Lucy **is** | Lucy **is not** |
|-------------|-----------------|
| Score · Q · LPD · bands | Tide’s permute matrix / live runner |
| Boards, charts, CSV, site PDF, `lucy serve` | River’s results store + full-site UI |
| Wasm (browser/Node) + native CLI wrap | A fork of the formulas per host |

**LPD** = Lucy Pareto density (`Q × shrink`). Zero unless Acc keep ≥ keep-floor
(default **70%**). Gold / lean / trap bands come from the same floors the Go
core exposes.

---

## Install

```bash
npm install @openfluke/lucy
```

Requires **Node ≥ 18**. React is an **optional** peer (`@openfluke/lucy/react`).

The published tarball includes `wasm/lucy.wasm` (built at publish time — not in
git). For local monorepo use, build once:

```bash
# from repo root
./go/scripts/build-artifacts.sh
```

---

## Quick start (wasm)

```js
import { buildLPD, VERSION, KEEP_FLOOR } from "@openfluke/lucy";

const samples = [
  {
    id: "int8",
    avg_accuracy: 0.98,
    soft_acc: 0.97,
    throughput: 1200,
    availability: 99,
    ram_kib: 400,
  },
  // …more cells from your host
];

const { board, version, options } = await buildLPD(samples, {
  keep_floor: KEEP_FLOOR, // 0.7
});

console.log(VERSION, version);
console.log("top", board.top[0].id, "LPD", board.top[0].lpd, "band", board.top[0].band);
```

Sample shape (aliases accepted): `avg_accuracy` / `acc`, `soft_acc` / `soft`,
`throughput` / `thru`, `availability` / `avail`, `ram_kib` / `ramKiB`, plus
optional `mode`, `dtype`, `format`, `arch`, `score`, `tide`.

Golden smoke samples live in
[`../examples/shared/samples.json`](../examples/shared/samples.json)
(top LPD should be **`int8`**).

---

## Runtimes

| Path | Import | When |
|------|--------|------|
| **Wasm** | `buildLPD` | Browser + Node (bundled `lucy.wasm`) |
| **Native** | `buildLPDNative` | Node/Bun with platform CLI (`LUCY_BIN` or artifacts) |
| **HTTP** | `createLucyClient` | BYO UI → `lucy serve` |

```js
import {
  buildLPD,
  buildLPDNative,
  writeCSVNative,
  writeSitePDFNative,
  floorsNative,
  createLucyClient,
  resolveLucyBinary,
} from "@openfluke/lucy";

// Native board + exports (needs built binary)
const floors = await floorsNative();
const { board } = await buildLPDNative(samples, { keep_floor: 0.7 });
await writeCSVNative(samples, "board.csv");
await writeSitePDFNative(samples, "site.pdf");

// HTTP client
const c = createLucyClient("http://127.0.0.1:7474");
await c.version();
await c.buildLPD(samples);
await c.sitePDF(samples); // ArrayBuffer / bytes
```

---

## UI skins

### React

```js
import { LucyBoard } from "@openfluke/lucy/react";
// peerDependency: react >= 18

export function Dash({ board }) {
  return <LucyBoard board={board} showDensity />;
}
```

### Custom elements (vanilla / Angular)

```js
import { registerLucyElements } from "@openfluke/lucy/elements";
registerLucyElements();
```

```html
<lucy-board></lucy-board>
<lucy-lpd-table max="20"></lucy-lpd-table>
<script type="module">
  import { buildLPD, registerLucyElements } from "@openfluke/lucy";
  registerLucyElements();
  const { board } = await buildLPD(samples);
  document.querySelector("lucy-board").board = board;
  document.querySelector("lucy-lpd-table").board = board;
</script>
```

Angular notes: [`examples/angular.md`](examples/angular.md).

### Charts helpers

```js
import {
  drawRadar,
  consciousnessSeries,
  densitySeries,
  drawScatter,
  lpdScatterPoints,
  lpdTableHTML,
} from "@openfluke/lucy";
// or: import { … } from "@openfluke/lucy/charts";
```

---

## Package exports

| Subpath | Purpose |
|---------|---------|
| `@openfluke/lucy` | Measuring + charts + native + CE + serve client |
| `@openfluke/lucy/react` | `LucyBoard` |
| `@openfluke/lucy/charts` | Radar / scatter / table / PNG·JPG·PDF·CSV helpers |
| `@openfluke/lucy/elements` | `<lucy-board>` / `<lucy-lpd-table>` |
| `@openfluke/lucy/serve` | `createLucyClient` |

Constants: `VERSION`, `KEEP_FLOOR`, `GOLD_KEEP`, `LEAN_KEEP`, `GOLD_RAM`,
`NEAR_RAM`, `SHRINK_CAP`.

---

## Examples in this repo

| Demo | Path |
|------|------|
| Wasm build | [`../examples/node/01_wasm_build.mjs`](../examples/node/01_wasm_build.mjs) |
| Native CSV + site PDF | [`../examples/node/02_native_full.mjs`](../examples/node/02_native_full.mjs) |
| HTTP client | [`../examples/node/03_serve_client.mjs`](../examples/node/03_serve_client.mjs) |
| Browser CE | [`../examples/browser/`](../examples/browser/) |
| BYO poll | [`examples/byo-poll.mjs`](examples/byo-poll.mjs) |
| npm how-to | [`../examples/npm/`](../examples/npm/) |
| Tutorials | [`../docs/tutorials/`](../docs/tutorials/) |

```bash
# from repo root
./scripts/run-examples.sh
```

```bash
# local package smoke
cd js
node examples/byo-poll.mjs ../examples/shared/samples.json
```

---

## Publish

```bash
bash publish.sh --dry-run
bash publish.sh             # interactive
bash publish.sh --yes
```

See [`../PUBLISH.md`](../PUBLISH.md). License: Apache-2.0.
