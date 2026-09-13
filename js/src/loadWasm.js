import { readFile } from "node:fs/promises";
import { fileURLToPath } from "node:url";
import path from "node:path";
import vm from "node:vm";
import crypto from "node:crypto";
import { performance } from "node:perf_hooks";

const __dirname = path.dirname(fileURLToPath(import.meta.url));
const wasmDir = path.resolve(__dirname, "../wasm");

let ready;

/** Load Go wasm once. Registers globalThis.lucyBuildLPD / lucyVersion. */
export async function ensureWasm() {
  if (ready) return ready;
  ready = (async () => {
    // polyfills wasm_exec.js expects on globalThis in Node
    globalThis.crypto ??= crypto;
    globalThis.performance ??= performance;

    const execPath = path.join(wasmDir, "wasm_exec.js");
    const code = await readFile(execPath, "utf8");
    vm.runInThisContext(code, { filename: "wasm_exec.js" });
    if (typeof globalThis.Go !== "function") {
      throw new Error("wasm_exec.js did not define globalThis.Go");
    }

    const go = new globalThis.Go();
    const buf = await readFile(path.join(wasmDir, "lucy.wasm"));
    const result = await WebAssembly.instantiate(buf, go.importObject);
    const run = go.run(result.instance);
    // exports are set during init before the parked select{}
    for (let i = 0; i < 50 && typeof globalThis.lucyBuildLPD !== "function"; i++) {
      await new Promise((r) => setTimeout(r, 10));
    }
    if (typeof globalThis.lucyBuildLPD !== "function") {
      // keep run promise from being GC'd oddly
      void run;
      throw new Error("lucy wasm loaded but lucyBuildLPD is missing");
    }
    void run;
  })();
  return ready;
}
