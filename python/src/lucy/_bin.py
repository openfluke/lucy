from __future__ import annotations

import os
import platform
from pathlib import Path


def bin_dir() -> Path:
    return Path(__file__).resolve().parent / "bin"


def default_binary() -> Path:
    env = os.environ.get("LUCY_BIN")
    if env:
        return Path(env)
    machine = platform.machine().replace("x86_64", "amd64").replace("aarch64", "arm64")
    named = bin_dir() / f"lucy-{platform.system().lower()}-{machine}"
    if named.exists():
        return named
    plain = bin_dir() / "lucy"
    if plain.exists():
        return plain
    raise FileNotFoundError(
        "lucy binary not found; set LUCY_BIN or run go/scripts/build-artifacts.sh"
    )
