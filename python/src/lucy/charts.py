"""Chart helpers — call Go CLI for SVG (JPG path later)."""

from __future__ import annotations

import json
import subprocess
from pathlib import Path
from typing import Any, Mapping, Optional, Sequence, Union

from ._bin import default_binary


def chart_svg(
    samples: Sequence[Mapping[str, Any]],
    kind: str = "radar",
    options: Optional[Mapping[str, float]] = None,
    *,
    binary: Optional[Union[str, Path]] = None,
) -> str:
    """Return SVG string: kind in radar|scatter|bars."""
    cmd = {
        "radar": "chart-radar",
        "scatter": "chart-scatter",
        "bars": "chart-bars",
    }.get(kind, kind)
    bin_path = Path(binary) if binary else default_binary()
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
    return proc.stdout.decode()
