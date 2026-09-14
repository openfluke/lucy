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
    async pdf(samples, options) {
      const r = await fetch(`${root}/api/pdf`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ samples, options }),
      });
      if (!r.ok) throw new Error(await r.text());
      return Buffer.from(await r.arrayBuffer());
    },
    async sitePDF(samples, options) {
      const r = await fetch(`${root}/api/site-pdf`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ samples, options }),
      });
      if (!r.ok) throw new Error(await r.text());
      return Buffer.from(await r.arrayBuffer());
    },
    async floors() {
      const r = await fetch(`${root}/api/floors`);
      if (!r.ok) throw new Error(`floors ${r.status}`);
      return r.json();
    },
    async csv(samples, options) {
      const r = await fetch(`${root}/api/csv`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ samples, options }),
      });
      if (!r.ok) throw new Error(await r.text());
      return r.text();
    },
  };
}
