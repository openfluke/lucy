"""Chart helpers — Go CLI for SVG/PNG (and PDF-ready packs)."""

from __future__ import annotations

import json
import subprocess
from pathlib import Path
from typing import Any, Mapping, Optional, Sequence, Union

from ._bin import default_binary


def _run(bin_path: Path, cmd: str, samples: Sequence[Mapping[str, Any]], options: Optional[Mapping[str, float]]) -> bytes:
    req: dict[str, Any] = {"samples": list(samples)}
    if options:
        req["options"] = dict(options)
    proc = subprocess.run(
        [str(bin_path), cmd],
        input=json.dumps(req).encode(),
        capture_output=True,
        check=False,
    )
    if proc.returncode != 0:
        raise RuntimeError(proc.stderr.decode() or f"lucy exited {proc.returncode}")
    return proc.stdout


def chart_svg(
    samples: Sequence[Mapping[str, Any]],
    kind: str = "radar",
    options: Optional[Mapping[str, float]] = None,
    *,
    binary: Optional[Union[str, Path]] = None,
) -> str:
    cmd = {"radar": "chart-radar", "scatter": "chart-scatter", "bars": "chart-bars"}.get(kind, kind)
    return _run(Path(binary) if binary else default_binary(), cmd, samples, options).decode()


def chart_png(
    samples: Sequence[Mapping[str, Any]],
    kind: str = "radar",
    options: Optional[Mapping[str, float]] = None,
    *,
    binary: Optional[Union[str, Path]] = None,
) -> bytes:
    cmd = {
        "radar": "chart-radar-png",
        "scatter": "chart-scatter-png",
        "bars": "chart-bars-png",
    }.get(kind, kind)
    return _run(Path(binary) if binary else default_binary(), cmd, samples, options)


def chart_pack(
    samples: Sequence[Mapping[str, Any]],
    options: Optional[Mapping[str, float]] = None,
    *,
    binary: Optional[Union[str, Path]] = None,
) -> dict[str, Any]:
    """PDF-friendly pack: SVG strings + PNG base64 fields."""
    raw = _run(Path(binary) if binary else default_binary(), "chart-pack", samples, options)
    return json.loads(raw.decode())
