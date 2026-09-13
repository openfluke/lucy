import json
import sys
import unittest
from pathlib import Path

sys.path.insert(0, str(Path(__file__).resolve().parents[1] / "src"))

from lucy import __version__, build_lpd, version


class TestBuildLPD(unittest.TestCase):
    def test_build_lpd_matches_golden(self):
        root = Path(__file__).resolve().parents[2]
        g = json.loads((root / "testdata" / "goldens_lpd_v0.2.json").read_text())
        self.assertEqual(__version__, "0.2.0")
        self.assertEqual(version(), "0.2.0")
        resp = build_lpd(g["samples"])
        self.assertEqual(resp["version"], "0.2.0")
        self.assertEqual(resp["board"]["top"][0]["id"], g["expect"]["top"][0]["id"])
        self.assertAlmostEqual(resp["board"]["top"][0]["lpd"], g["expect"]["top"][0]["lpd"])


if __name__ == "__main__":
    unittest.main()
