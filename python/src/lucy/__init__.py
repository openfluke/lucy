"""Lucy measuring (PyPI placeholder).

Will wrap Go-built binaries / CLI so Python hosts get the same Score / Q / LPD
ruler as Go and @openfluke/lucy — no reimplemented formulas.
"""

__version__ = "0.0.0"

KEEP_FLOOR = 0.70


def build_lpd(samples):
    """Placeholder — call into Go binary once artifacts exist."""
    _ = samples
    return {
        "formula": (
            "placeholder — Score = T×Avail×Acc/10_000; "
            "LPD = Q×shrink if RelAcc≥KeepFloor else 0"
        ),
        "top": [],
    }
