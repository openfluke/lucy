"""Chart helpers — Go CLI for SVG/PNG/JPG + report packs."""

from __future__ import annotations

import json
import subprocess
from pathlib import Path
from typing import Any, Mapping, Optional, Sequence, Union

from ._bin import default_binary


def _run(bin_path: Path, args: list[str], samples: Sequence[Mapping[str, Any]], options: Optional[Mapping[str, float]]) -> bytes:
    req: dict[str, Any] = {"samples": list(samples)}
    if options:
        req["options"] = dict(options)
    proc = subprocess.run(
        [str(bin_path), *args],
        input=json.dumps(req).encode(),
        capture_output=True,
        check=False,
    )
    if proc.returncode != 0:
        raise RuntimeError(proc.stderr.decode() or f"lucy exited {proc.returncode}")
    return proc.stdout


def chart_svg(samples, kind="radar", options=None, *, binary=None) -> str:
    cmd = {"radar": "chart-radar", "scatter": "chart-scatter", "bars": "chart-bars"}.get(kind, kind)
    return _run(Path(binary) if binary else default_binary(), [cmd], samples, options).decode()


def chart_png(samples, kind="radar", options=None, *, binary=None) -> bytes:
    cmd = {"radar": "chart-radar-png", "scatter": "chart-scatter-png", "bars": "chart-bars-png"}.get(kind, kind)
    return _run(Path(binary) if binary else default_binary(), [cmd], samples, options)


def chart_jpg(samples, kind="radar", options=None, *, binary=None) -> bytes:
    cmd = {"radar": "chart-radar-jpg", "scatter": "chart-scatter-jpg", "bars": "chart-bars-jpg"}.get(kind, kind)
    return _run(Path(binary) if binary else default_binary(), [cmd], samples, options)


def chart_pack(samples, options=None, *, binary=None) -> dict[str, Any]:
    raw = _run(Path(binary) if binary else default_binary(), ["chart-pack"], samples, options)
    return json.loads(raw.decode())


def write_report(samples, outdir: Union[str, Path], options=None, *, binary=None) -> str:
    out = _run(Path(binary) if binary else default_binary(), ["report", str(outdir)], samples, options)
    return out.decode().strip() or str(outdir)


def write_pdf(samples, out_path, options=None, *, binary=None) -> str:
    out = _run(Path(binary) if binary else default_binary(), ["pdf", str(out_path)], samples, options)
    return out.decode().strip() or str(out_path)


def board_csv(samples, options=None, *, binary=None) -> str:
    return _run(Path(binary) if binary else default_binary(), ["csv"], samples, options).decode()


def floors(*, binary=None) -> dict[str, Any]:
    bin_path = Path(binary) if binary else default_binary()
    proc = subprocess.run([str(bin_path), "floors"], capture_output=True, check=False)
    if proc.returncode != 0:
        raise RuntimeError(proc.stderr.decode() or f"lucy exited {proc.returncode}")
    return json.loads(proc.stdout.decode())


def write_site_pdf(samples, out_path, options=None, *, binary=None) -> str:
    """River-parity site pack (same as write_pdf since 0.9)."""
    return write_pdf(samples, out_path, options, binary=binary)
