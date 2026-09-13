"""Notebook-friendly board helpers."""

from __future__ import annotations

from typing import Any, Mapping


def top_rows(board: Mapping[str, Any], max_rows: int = 40) -> list[dict[str, Any]]:
    rows = list(board.get("top") or [])
    return rows[:max_rows]


def board_records(resp_or_board: Mapping[str, Any]) -> list[dict[str, Any]]:
    """Flatten BuildResponse or board into row dicts for DataFrames."""
    board = resp_or_board.get("board", resp_or_board)
    out = []
    for r in top_rows(board):
        out.append(
            {
                "id": r.get("id"),
                "band": r.get("band"),
                "lpd": r.get("lpd"),
                "q": r.get("q"),
                "rel_acc": r.get("rel_acc"),
                "rel_thru": r.get("rel_thru"),
                "rel_avail": r.get("rel_avail"),
                "ram_kib": r.get("ram_kib"),
                "ram_frac": r.get("ram_frac"),
                "shrink": r.get("shrink"),
                "score": r.get("score"),
                "acc": r.get("avg_accuracy", r.get("acc")),
                "thru": r.get("throughput", r.get("thru")),
                "avail": r.get("availability", r.get("avail")),
                "mode": r.get("mode"),
                "dtype": r.get("dtype"),
                "arch": r.get("arch"),
            }
        )
    return out


def to_dataframe(resp_or_board: Mapping[str, Any]):
    """Return a pandas DataFrame if pandas is installed."""
    try:
        import pandas as pd
    except ImportError as e:
        raise ImportError("to_dataframe requires pandas: pip install pandas") from e
    return pd.DataFrame(board_records(resp_or_board))
