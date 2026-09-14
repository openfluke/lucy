import json
import sys
import unittest
from pathlib import Path

sys.path.insert(0, str(Path(__file__).resolve().parents[1] / "src"))

from lucy import __version__, board_records, build_lpd, chart_pack, chart_png, chart_svg, version


class TestBuildLPD(unittest.TestCase):
    def setUp(self):
        root = Path(__file__).resolve().parents[2]
        self.g = json.loads((root / "testdata" / "goldens_lpd_v0.4.json").read_text())

    def test_build_lpd_matches_golden(self):
        self.assertEqual(__version__, "0.4.0")
        self.assertEqual(version(), "0.4.0")
        resp = build_lpd(self.g["samples"])
        self.assertEqual(resp["version"], "0.4.0")
        self.assertEqual(resp["board"]["top"][0]["id"], self.g["expect"]["top"][0]["id"])

    def test_charts_svg_png_pack(self):
        resp = build_lpd(self.g["samples"])
        self.assertTrue(board_records(resp))
        svg = chart_svg(self.g["samples"], kind="radar")
        self.assertIn("<svg", svg)
        png = chart_png(self.g["samples"], kind="bars")
        self.assertEqual(png[0], 0x89)
        pack = chart_pack(self.g["samples"])
        self.assertIn("consciousness_svg", pack)
        self.assertTrue(len(pack["bars_png_b64"]) > 20)


if __name__ == "__main__":
    unittest.main()
