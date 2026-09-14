"""Lucy measuring via Go binaries (openfluke-lucy 1.0.0)."""

from __future__ import annotations

import json
import subprocess
from pathlib import Path
from typing import Any, Mapping, MutableMapping, Optional, Sequence, TypedDict, Union

from ._bin import default_binary
from .board import board_records, to_dataframe, top_rows
from .charts import board_csv, chart_jpg, chart_pack, chart_png, chart_svg, floors, write_pdf, write_report, write_site_pdf
from .client import LucyClient

__version__ = "1.0.0"

KEEP_FLOOR = 0.70
GOLD_KEEP = 0.80
LEAN_KEEP = 0.95
GOLD_RAM = 0.20
NEAR_RAM = 0.50
SHRINK_CAP = 32.0


class Sample(TypedDict, total=False):
    id: str
    mode: str
    dtype: str
    format: str
    arch: str
    score: float
    soft: float
    soft_acc: float
    acc: float
    avg_accuracy: float
    thru: float
    throughput: float
    avail: float
    availability: float
    ram_kib: float
    tide: str


DensityOptions = Mapping[str, float]


def build_lpd(
    samples: Sequence[Mapping[str, Any]],
    options: Optional[DensityOptions] = None,
    *,
    binary: Optional[Union[str, Path]] = None,
) -> dict[str, Any]:
    bin_path = Path(binary) if binary else default_binary()
    req: MutableMapping[str, Any] = {"samples": list(samples)}
    if options:
        req["options"] = dict(options)
    proc = subprocess.run(
        [str(bin_path), "build-lpd"],
        input=json.dumps(req).encode(),
        capture_output=True,
        check=False,
    )
    if proc.returncode != 0:
        raise RuntimeError(proc.stderr.decode() or f"lucy exited {proc.returncode}")
    return json.loads(proc.stdout.decode())


def version(binary: Optional[Union[str, Path]] = None) -> str:
    bin_path = Path(binary) if binary else default_binary()
    proc = subprocess.run([str(bin_path), "version"], capture_output=True, check=True, text=True)
    return proc.stdout.strip()


__all__ = [
    "__version__",
    "KEEP_FLOOR",
    "GOLD_KEEP",
    "LEAN_KEEP",
    "GOLD_RAM",
    "NEAR_RAM",
    "SHRINK_CAP",
    "Sample",
    "build_lpd",
    "version",
    "floors",
    "top_rows",
    "board_records",
    "to_dataframe",
    "chart_svg",
    "chart_png",
    "chart_jpg",
    "chart_pack",
    "board_csv",
    "write_report",
    "write_pdf",
    "write_site_pdf",
    "LucyClient",
]
