import json
import sys
import tempfile
import unittest
from pathlib import Path

sys.path.insert(0, str(Path(__file__).resolve().parents[1] / "src"))

from lucy import (
    __version__,
    board_csv,
    board_records,
    build_lpd,
    chart_jpg,
    chart_pack,
    chart_png,
    chart_svg,
    floors,
    version,
    write_pdf,
    write_report,
)


class TestBuildLPD(unittest.TestCase):
    def setUp(self):
        root = Path(__file__).resolve().parents[2]
        self.g = json.loads((root / "testdata" / "goldens_lpd_v1.0.json").read_text())

    def test_build_lpd_matches_golden(self):
        self.assertEqual(__version__, "1.0.1")
        self.assertEqual(version(), "1.0.1")
        resp = build_lpd(self.g["samples"])
        self.assertEqual(resp["version"], "1.0.1")
        self.assertEqual(resp["board"]["top"][0]["id"], self.g["expect"]["top"][0]["id"])

    def test_charts_report_pdf(self):
        self.assertTrue(board_records(build_lpd(self.g["samples"])))
        self.assertIn("<svg", chart_svg(self.g["samples"], "radar"))
        self.assertEqual(chart_png(self.g["samples"], "bars")[0], 0x89)
        self.assertEqual(chart_jpg(self.g["samples"], "radar")[0], 0xFF)
        pack = chart_pack(self.g["samples"])
        self.assertIn("bars_jpg_b64", pack)
        self.assertEqual(floors()["version"], "1.0.1")
        csv = board_csv(self.g["samples"])
        self.assertTrue(csv.startswith("id,band,lpd,"))
        with tempfile.TemporaryDirectory() as d:
            out = write_report(self.g["samples"], d)
            self.assertTrue((Path(out) / "index.html").exists())
            self.assertTrue((Path(out) / "board.pdf").exists())
            self.assertTrue((Path(out) / "board.csv").exists())
            pdf_path = Path(d) / "direct.pdf"
            write_pdf(self.g["samples"], pdf_path)
            self.assertTrue(pdf_path.read_bytes().startswith(b"%PDF"))


if __name__ == "__main__":
    unittest.main()
