"""HTTP client for `lucy serve`."""

from __future__ import annotations

import json
import urllib.request
from typing import Any, Mapping, Optional, Sequence


class LucyClient:
    def __init__(self, base_url: str = "http://127.0.0.1:7474"):
        self.base_url = base_url.rstrip("/")

    def version(self) -> str:
        with urllib.request.urlopen(self.base_url + "/api/version") as r:
            return json.loads(r.read().decode())["version"]

    def build_lpd(self, samples: Sequence[Mapping[str, Any]], options: Optional[Mapping[str, float]] = None) -> dict[str, Any]:
        return self._post("/api/lpd", {"samples": list(samples), "options": options})

    def chart_pack(self, samples: Sequence[Mapping[str, Any]], options: Optional[Mapping[str, float]] = None) -> dict[str, Any]:
        return self._post("/api/chart-pack", {"samples": list(samples), "options": options})

    def _post(self, path: str, body: dict[str, Any]) -> dict[str, Any]:
        data = json.dumps({k: v for k, v in body.items() if v is not None}).encode()
        req = urllib.request.Request(
            self.base_url + path,
            data=data,
            headers={"Content-Type": "application/json"},
            method="POST",
        )
        with urllib.request.urlopen(req) as r:
            return json.loads(r.read().decode())
