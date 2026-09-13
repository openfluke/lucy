import json
import sys
import unittest
from pathlib import Path

sys.path.insert(0, str(Path(__file__).resolve().parents[1] / "src"))

from lucy import __version__, board_records, build_lpd, chart_svg, version


class TestBuildLPD(unittest.TestCase):
    def setUp(self):
        root = Path(__file__).resolve().parents[2]
        self.g = json.loads((root / "testdata" / "goldens_lpd_v0.3.json").read_text())

    def test_build_lpd_matches_golden(self):
        self.assertEqual(__version__, "0.3.0")
        self.assertEqual(version(), "0.3.0")
        resp = build_lpd(self.g["samples"])
        self.assertEqual(resp["version"], "0.3.0")
        self.assertEqual(resp["board"]["top"][0]["id"], self.g["expect"]["top"][0]["id"])

    def test_board_records_and_chart_svg(self):
        resp = build_lpd(self.g["samples"])
        rows = board_records(resp)
        self.assertGreaterEqual(len(rows), 1)
        self.assertIn("lpd", rows[0])
        svg = chart_svg(self.g["samples"], kind="radar")
        self.assertIn("<svg", svg)
        self.assertIn("Consciousness", svg)


if __name__ == "__main__":
    unittest.main()
