/** HTTP client for `lucy serve` (BYO frontend → local Go API). */

export function createLucyClient(baseURL = "http://127.0.0.1:7474") {
  const root = baseURL.replace(/\/$/, "");
  return {
    async version() {
      const r = await fetch(`${root}/api/version`);
      if (!r.ok) throw new Error(`version ${r.status}`);
      return (await r.json()).version;
    },
    async buildLPD(samples, options) {
      const r = await fetch(`${root}/api/lpd`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ samples, options }),
      });
      if (!r.ok) throw new Error(await r.text());
      return r.json();
    },
    async chartPack(samples, options) {
      const r = await fetch(`${root}/api/chart-pack`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ samples, options }),
      });
      if (!r.ok) throw new Error(await r.text());
      return r.json();
    },
  };
}
