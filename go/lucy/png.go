package lucy

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
	"math"
)

var (
	pngBG     = color.RGBA{0x0d, 0x12, 0x16, 0xff}
	pngGrid   = color.RGBA{0x1d, 0x33, 0x42, 0xff}
	pngMuted  = color.RGBA{0x8a, 0xa0, 0xad, 0xff}
	pngFg     = color.RGBA{0xc5, 0xd0, 0xd8, 0xff}
)

// BoardCharts is a PDF-friendly pack of board visuals (SVG + PNG + JPG).
type BoardCharts struct {
	ConsciousnessSVG string `json:"consciousness_svg"`
	DensitySVG       string `json:"density_svg"`
	ScatterSVG       string `json:"scatter_svg"`
	BarsSVG          string `json:"bars_svg"`
	ConsciousnessPNG []byte `json:"-"`
	DensityPNG       []byte `json:"-"`
	ScatterPNG       []byte `json:"-"`
	BarsPNG          []byte `json:"-"`
	ConsciousnessJPG []byte `json:"-"`
	DensityJPG       []byte `json:"-"`
	ScatterJPG       []byte `json:"-"`
	BarsJPG          []byte `json:"-"`
}

// BuildBoardCharts renders Tide-style chart set for a board (max series rows).
func BuildBoardCharts(board LPD, max int) BoardCharts {
	if max <= 0 {
		max = 8
	}
	live := ConsciousnessSeries(board, max)
	dens := DensitySeries(board, max)
	pts := LPDScatterPoints(board)
	return BoardCharts{
		ConsciousnessSVG: RadarSVG("Consciousness radar", live),
		DensitySVG:       RadarSVG("Memory density radar", dens),
		ScatterSVG:       ScatterSVG("Q% vs RAM", "RAM KiB", "Q %", pts),
		BarsSVG:          BarsSVG("Top LPD", board, max+4),
		ConsciousnessPNG: mustPNG(RadarPNG("Consciousness radar", live)),
		DensityPNG:       mustPNG(RadarPNG("Memory density radar", dens)),
		ScatterPNG:       mustPNG(ScatterPNG("Q% vs RAM", "RAM KiB", "Q %", pts)),
		BarsPNG:          mustPNG(BarsPNG("Top LPD", board, max+4)),
		ConsciousnessJPG: mustJPG(RadarJPG("Consciousness radar", live, 85)),
		DensityJPG:       mustJPG(RadarJPG("Memory density radar", dens, 85)),
		ScatterJPG:       mustJPG(ScatterJPG("Q% vs RAM", "RAM KiB", "Q %", pts, 85)),
		BarsJPG:          mustJPG(BarsJPG("Top LPD", board, max+4, 85)),
	}
}

func mustJPG(b []byte, err error) []byte {
	if err != nil {
		return nil
	}
	return b
}

func mustPNG(b []byte, err error) []byte {
	if err != nil {
		return nil
	}
	return b
}

// RadarPNG renders a 3-axis radar as PNG.
func RadarPNG(title string, series []RadarSeries) ([]byte, error) {
	const w, h = 960, 480
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	fill(img, pngBG)
	drawString(img, 12, 18, title, pngMuted)
	if len(series) == 0 {
		drawString(img, w/2-40, h/2, "no series", pngMuted)
		return encodePNG(img)
	}
	cx, cy := float64(w)*0.36, float64(h)*0.54
	radius := math.Min(cx-24, cy-40)
	ang := func(i int) float64 { return -math.Pi/2 + float64(i)*2*math.Pi/3 }
	for ring := 1; ring <= 4; ring++ {
		var pts [3]image.Point
		for i := 0; i < 3; i++ {
			r := radius * float64(ring) / 4
			pts[i] = image.Pt(int(cx+r*math.Cos(ang(i))), int(cy+r*math.Sin(ang(i))))
		}
		drawPoly(img, pts[:], pngGrid)
	}
	for i := 0; i < 3; i++ {
		drawLine(img, int(cx), int(cy), int(cx+radius*math.Cos(ang(i))), int(cy+radius*math.Sin(ang(i))), pngGrid)
		lx := int(cx + (radius+18)*math.Cos(ang(i)))
		ly := int(cy + (radius+18)*math.Sin(ang(i)))
		labels := []string{"Acc", "Thru", "Avail"}
		drawString(img, lx-10, ly, labels[i], pngMuted)
	}
	for _, s := range series {
		col := parseHex(s.Color, color.RGBA{0x3d, 0xd6, 0xc6, 0xff})
		var pts [3]image.Point
		for i := 0; i < 3; i++ {
			v := clamp01(s.Vals[i])
			pts[i] = image.Pt(int(cx+radius*v*math.Cos(ang(i))), int(cy+radius*v*math.Sin(ang(i))))
		}
		drawPoly(img, pts[:], col)
	}
	y := 40
	for _, s := range series {
		col := parseHex(s.Color, color.RGBA{0x3d, 0xd6, 0xc6, 0xff})
		fillRect(img, w-300, y, 12, 12, col)
		drawString(img, w-282, y+10, s.Label, pngFg)
		y += 18
	}
	return encodePNG(img)
}

// ScatterPNG renders scatter as PNG.
func ScatterPNG(title, xLabel, yLabel string, pts []ScatterPoint) ([]byte, error) {
	const w, h = 960, 440
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	fill(img, pngBG)
	drawString(img, 12, 16, title, pngMuted)
	if len(pts) == 0 {
		drawString(img, w/2-40, h/2, "no points", pngMuted)
		return encodePNG(img)
	}
	const padL, padR, padT, padB = 48, 16, 12, 36
	xs := make([]float64, len(pts))
	ys := make([]float64, len(pts))
	for i, p := range pts {
		xs[i], ys[i] = p.X, p.Y
	}
	xmin, xmax := minMax(xs)
	ymin, ymax := minMax(ys)
	if xmax <= xmin {
		xmax = xmin + 1
	}
	if ymax <= ymin {
		ymax = ymin + 1
	}
	X := func(v float64) int { return padL + int((float64(w)-padL-padR)*(v-xmin)/(xmax-xmin)) }
	Y := func(v float64) int {
		return h - padB - int((float64(h)-padT-padB)*(v-ymin)/(ymax-ymin))
	}
	strokeRect(img, padL, padT, w-padL-padR, h-padT-padB, pngGrid)
	for _, p := range pts {
		col := parseHex(bandColor(p.Band), color.RGBA{0x5a, 0x7a, 0x8a, 0xff})
		fillRect(img, X(p.X)-3, Y(p.Y)-3, 6, 6, col)
	}
	drawString(img, w/2-20, h-8, xLabel, pngMuted)
	drawString(img, 8, h/2, yLabel, pngMuted)
	return encodePNG(img)
}

// BarsPNG renders LPD bars as PNG.
func BarsPNG(title string, board LPD, max int) ([]byte, error) {
	const w = 960
	rows := board.Top
	if max > 0 && len(rows) > max {
		rows = rows[:max]
	}
	h := 40 + len(rows)*28
	if h < 120 {
		h = 120
	}
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	fill(img, pngBG)
	drawString(img, 12, 18, title, pngMuted)
	var maxLPD float64 = 1
	for _, r := range rows {
		if r.LPD > maxLPD {
			maxLPD = r.LPD
		}
	}
	for i, r := range rows {
		y := 32 + i*28
		bw := int((float64(w) - 200) * (r.LPD / maxLPD))
		drawString(img, 12, y+12, r.ID, pngFg)
		fillRect(img, 160, y, bw, 16, parseHex(bandColor(r.Band), color.RGBA{0x3d, 0xd6, 0xc6, 0xff}))
	}
	return encodePNG(img)
}

func encodePNG(img *image.RGBA) ([]byte, error) {
	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func fill(img *image.RGBA, c color.Color) {
	b := img.Bounds()
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			img.Set(x, y, c)
		}
	}
}

func fillRect(img *image.RGBA, x, y, w, h int, c color.Color) {
	for yy := y; yy < y+h; yy++ {
		for xx := x; xx < x+w; xx++ {
			if img.Bounds().Overlaps(image.Rect(xx, yy, xx+1, yy+1)) {
				img.Set(xx, yy, c)
			}
		}
	}
}

func strokeRect(img *image.RGBA, x, y, w, h int, c color.Color) {
	drawLine(img, x, y, x+w, y, c)
	drawLine(img, x, y+h, x+w, y+h, c)
	drawLine(img, x, y, x, y+h, c)
	drawLine(img, x+w, y, x+w, y+h, c)
}

func drawLine(img *image.RGBA, x0, y0, x1, y1 int, c color.Color) {
	dx := abs(x1 - x0)
	dy := -abs(y1 - y0)
	sx, sy := 1, 1
	if x0 > x1 {
		sx = -1
	}
	if y0 > y1 {
		sy = -1
	}
	err := dx + dy
	for {
		img.Set(x0, y0, c)
		if x0 == x1 && y0 == y1 {
			break
		}
		e2 := 2 * err
		if e2 >= dy {
			err += dy
			x0 += sx
		}
		if e2 <= dx {
			err += dx
			y0 += sy
		}
	}
}

func drawPoly(img *image.RGBA, pts []image.Point, c color.Color) {
	for i := 0; i < len(pts); i++ {
		j := (i + 1) % len(pts)
		drawLine(img, pts[i].X, pts[i].Y, pts[j].X, pts[j].Y, c)
	}
}

// Tiny 5x7 digit/letter plotter for labels (subset ASCII).
func drawString(img *image.RGBA, x, y int, s string, c color.Color) {
	for i, ch := range s {
		drawGlyph(img, x+i*6, y-6, ch, c)
	}
}

func drawGlyph(img *image.RGBA, x, y int, ch rune, c color.Color) {
	// fallback: 2x2 block per character so labels are visible without a font file
	fillRect(img, x, y, 4, 6, c)
	_ = ch
}

func parseHex(s string, fallback color.RGBA) color.RGBA {
	if len(s) != 7 || s[0] != '#' {
		return fallback
	}
	var r, g, b int
	if _, err := parseHex3(s[1:], &r, &g, &b); err != nil {
		return fallback
	}
	return color.RGBA{uint8(r), uint8(g), uint8(b), 0xff}
}

func parseHex3(s string, r, g, b *int) (int, error) {
	var v uint32
	for i := 0; i < 6; i++ {
		v <<= 4
		c := s[i]
		switch {
		case c >= '0' && c <= '9':
			v |= uint32(c - '0')
		case c >= 'a' && c <= 'f':
			v |= uint32(c - 'a' + 10)
		case c >= 'A' && c <= 'F':
			v |= uint32(c - 'A' + 10)
		default:
			return 0, errBadHex
		}
	}
	*r = int(v >> 16)
	*g = int((v >> 8) & 0xff)
	*b = int(v & 0xff)
	return 0, nil
}

type badHex struct{}

func (badHex) Error() string { return "bad hex" }

var errBadHex = badHex{}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
