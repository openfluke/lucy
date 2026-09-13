"""Lucy measuring (PyPI 0.1.0).

Go core owns Score / Q / LPD. This package will wrap Go binaries at 0.2+;
until then the API surface matches the board schema for host wiring.
"""

from __future__ import annotations

from typing import Any, TypedDict

__version__ = "0.1.0"

KEEP_FLOOR = 0.70


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


def build_lpd(samples: list[Sample]) -> dict[str, Any]:
    """Rank samples for Lucy Pareto density.

    0.1: API placeholder — call Go ``lucy.BuildLPD`` (or wait for 0.2 binaries).
    Does not reimplement formulas in Python.
    """
    _ = samples
    raise NotImplementedError(
        "openfluke-lucy 0.1.0: use Go github.com/openfluke/lucy/lucy.BuildLPD; "
        "PyPI binary wrapper lands at 0.2"
    )
