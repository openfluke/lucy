"""Lucy measuring via Go binaries (PyPI openfluke-lucy 0.2.0)."""

from __future__ import annotations

import json
import os
import platform
import subprocess
from pathlib import Path
from typing import Any, Mapping, MutableMapping, Optional, Sequence, TypedDict, Union

__version__ = "0.2.0"

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
    acc: float
    thru: float
    avail: float
    ram_kib: float
    tide: str


DensityOptions = Mapping[str, float]


def _bin_dir() -> Path:
    return Path(__file__).resolve().parent / "bin"


def _default_binary() -> Path:
    env = os.environ.get("LUCY_BIN")
    if env:
        return Path(env)
    named = _bin_dir() / f"lucy-{platform.system().lower()}-{platform.machine().replace('x86_64', 'amd64').replace('aarch64', 'arm64')}"
    if named.exists():
        return named
    plain = _bin_dir() / "lucy"
    if plain.exists():
        return plain
    raise FileNotFoundError(
        "lucy binary not found; set LUCY_BIN or run go/scripts/build-artifacts.sh"
    )


def build_lpd(
    samples: Sequence[Mapping[str, Any]],
    options: Optional[DensityOptions] = None,
    *,
    binary: Optional[Union[str, Path]] = None,
) -> dict[str, Any]:
    """Rank samples for Lucy Pareto density via the Go CLI."""
    bin_path = Path(binary) if binary else _default_binary()
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
    bin_path = Path(binary) if binary else _default_binary()
    proc = subprocess.run([str(bin_path), "version"], capture_output=True, check=True, text=True)
    return proc.stdout.strip()
