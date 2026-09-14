package lucy

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// WriteReportDir writes an HTML + assets report for a board under dir.
// Layout: index.html, board.json, *.svg, *.png, *.jpg
func WriteReportDir(dir string, board LPD, max int) error {
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	pack := BuildBoardCharts(board, max)
	write := func(name string, b []byte) error {
		return os.WriteFile(filepath.Join(dir, name), b, 0o644)
	}
	files := []struct {
		name string
		b    []byte
	}{
		{"consciousness.svg", []byte(pack.ConsciousnessSVG)},
		{"density.svg", []byte(pack.DensitySVG)},
		{"scatter.svg", []byte(pack.ScatterSVG)},
		{"bars.svg", []byte(pack.BarsSVG)},
		{"consciousness.png", pack.ConsciousnessPNG},
		{"density.png", pack.DensityPNG},
		{"scatter.png", pack.ScatterPNG},
		{"bars.png", pack.BarsPNG},
		{"consciousness.jpg", pack.ConsciousnessJPG},
		{"density.jpg", pack.DensityJPG},
		{"scatter.jpg", pack.ScatterJPG},
		{"bars.jpg", pack.BarsJPG},
	}
	for _, f := range files {
		if len(f.b) == 0 {
			continue
		}
		if err := write(f.name, f.b); err != nil {
			return err
		}
	}
	bj, err := json.MarshalIndent(board, "", "  ")
	if err != nil {
		return err
	}
	if err := write("board.json", bj); err != nil {
		return err
	}
	html := fmt.Sprintf(`<!DOCTYPE html>
<html lang="en"><head><meta charset="utf-8"/><title>Lucy board report</title>
<style>
body{font-family:system-ui,sans-serif;background:#0d1216;color:#c5d0d8;margin:1.5rem}
h1,h2{font-weight:600} img,object{max-width:100%%;background:#0d1216;margin:.5rem 0}
.grid{display:grid;gap:1rem;grid-template-columns:1fr}
a{color:#3dd6c6}
</style></head><body>
<h1>Lucy board report <small style="color:#8aa0ad">%s</small></h1>
<p><a href="board.json">board.json</a></p>
<h2>Consciousness</h2><img src="consciousness.png" alt="consciousness"/>
<h2>Memory density</h2><img src="density.png" alt="density"/>
<h2>Q%% vs RAM</h2><img src="scatter.png" alt="scatter"/>
<h2>Top LPD</h2><img src="bars.png" alt="bars"/>
</body></html>
`, Version)
	return write("index.html", []byte(html))
}
