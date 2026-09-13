package lucy

import (
	"fmt"
	"math"
	"strings"
)

// RadarSeries is one polygon on a 3-axis radar (Acc / Thru / Avail keep or density).
type RadarSeries struct {
	Label string     `json:"label"`
	Color string     `json:"color,omitempty"`
	Vals  [3]float64 `json:"vals"` // Acc, Thru, Avail in [0,1]
}

// ConsciousnessSeries builds radar series from board top rows (keep pillars).
func ConsciousnessSeries(board LPD, max int) []RadarSeries {
	rows := board.Top
	if max > 0 && len(rows) > max {
		rows = rows[:max]
	}
	out := make([]RadarSeries, 0, len(rows))
	cols := []string{"#3dd6c6", "#6ea8fe", "#e0a458", "#c084fc", "#f07178"}
	for i, r := range rows {
		v := r.Consciousness()
		out = append(out, RadarSeries{Label: r.ID, Color: cols[i%len(cols)], Vals: v})
	}
	return out
}

// DensitySeries builds radar series from board top rows (keep × shrink).
func DensitySeries(board LPD, max int) []RadarSeries {
	rows := board.Top
	if max > 0 && len(rows) > max {
		rows = rows[:max]
	}
	out := make([]RadarSeries, 0, len(rows))
	cols := []string{"#3dd6c6", "#6ea8fe", "#e0a458", "#c084fc", "#f07178"}
	for i, r := range rows {
		v := r.MemoryDensity()
		peak := 1.0
		for _, x := range v {
			if x > peak {
				peak = x
			}
		}
		if peak < 1 {
			peak = 1
		}
		out = append(out, RadarSeries{
			Label: r.ID,
			Color: cols[i%len(cols)],
			Vals:  [3]float64{clamp01(v[0] / peak), clamp01(v[1] / peak), clamp01(v[2] / peak)},
		})
	}
	return out
}

func clamp01(x float64) float64 {
	if x < 0 {
		return 0
	}
	if x > 1 {
		return 1
	}
	return x
}

// RadarSVG renders a 3-axis radar chart (Acc / Thru / Avail).
func RadarSVG(title string, series []RadarSeries) string {
	const w, h = 960, 480
	if len(series) == 0 {
		return emptySVG(w, h, title, "no series")
	}
	cx, cy := float64(w)*0.36, float64(h)*0.54
	radius := math.Min(cx-24, cy-40)
	labels := []string{"Acc", "Thru", "Avail"}
	ang := func(i int) float64 { return -math.Pi/2 + float64(i)*2*math.Pi/3 }
	var b strings.Builder
	writeSVGHead(&b, w, h)
	fmt.Fprintf(&b, `<rect width="%d" height="%d" fill="#0d1216"/>`, w, h)
	if title != "" {
		fmt.Fprintf(&b, `<text x="12" y="18" fill="#8aa0ad" font-family="sans-serif" font-size="13">%s</text>`, escSVG(title))
	}
	for ring := 1; ring <= 4; ring++ {
		b.WriteString(`<polygon fill="none" stroke="#1d3342" points="`)
		for i := 0; i < 3; i++ {
			if i > 0 {
				b.WriteByte(' ')
			}
			r := radius * float64(ring) / 4
			fmt.Fprintf(&b, "%.1f,%.1f", cx+r*math.Cos(ang(i)), cy+r*math.Sin(ang(i)))
		}
		b.WriteString(`"/>`)
	}
	for i := 0; i < 3; i++ {
		fmt.Fprintf(&b, `<line x1="%.1f" y1="%.1f" x2="%.1f" y2="%.1f" stroke="#1d3342"/>`,
			cx, cy, cx+radius*math.Cos(ang(i)), cy+radius*math.Sin(ang(i)))
		lx := cx + (radius+22)*math.Cos(ang(i))
		ly := cy + (radius+22)*math.Sin(ang(i))
		fmt.Fprintf(&b, `<text x="%.1f" y="%.1f" fill="#8aa0ad" font-family="sans-serif" font-size="13" font-weight="600" text-anchor="middle">%s</text>`, lx, ly+4, labels[i])
	}
	for _, s := range series {
		col := s.Color
		if col == "" {
			col = "#3dd6c6"
		}
		b.WriteString(`<polygon fill="none" stroke="` + col + `" stroke-width="2" opacity="0.9" points="`)
		for i := 0; i < 3; i++ {
			if i > 0 {
				b.WriteByte(' ')
			}
			v := clamp01(s.Vals[i])
			fmt.Fprintf(&b, "%.1f,%.1f", cx+radius*v*math.Cos(ang(i)), cy+radius*v*math.Sin(ang(i)))
		}
		b.WriteString(`"/>`)
	}
	legX := float64(w) - 300
	y := 40.0
	for _, s := range series {
		col := s.Color
		if col == "" {
			col = "#3dd6c6"
		}
		fmt.Fprintf(&b, `<rect x="%.1f" y="%.1f" width="12" height="12" fill="%s"/>`, legX, y, col)
		fmt.Fprintf(&b, `<text x="%.1f" y="%.1f" fill="#c5d0d8" font-family="sans-serif" font-size="12">%s</text>`, legX+18, y+11, escSVG(s.Label))
		y += 18
	}
	b.WriteString(`</svg>`)
	return b.String()
}

// ScatterPoint is one cell on a 2D board chart.
type ScatterPoint struct {
	ID   string  `json:"id"`
	X    float64 `json:"x"`
	Y    float64 `json:"y"`
	Band string  `json:"band,omitempty"`
}

// LPDScatterPoints maps board top rows to RAM KiB vs Q%.
func LPDScatterPoints(board LPD) []ScatterPoint {
	out := make([]ScatterPoint, 0, len(board.Top))
	for _, r := range board.Top {
		out = append(out, ScatterPoint{ID: r.ID, X: r.RAMKiB, Y: r.Q * 100, Band: r.Band})
	}
	return out
}

// ScatterSVG renders X/Y points (Tide-style dark board).
func ScatterSVG(title, xLabel, yLabel string, pts []ScatterPoint) string {
	const w, h = 960, 440
	if len(pts) == 0 {
		return emptySVG(w, h, title, "no points")
	}
	xs, ys := make([]float64, len(pts)), make([]float64, len(pts))
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
	const padL, padR, padT, padB = 48.0, 16.0, 12.0, 36.0
	X := func(v float64) float64 { return padL + (float64(w)-padL-padR)*(v-xmin)/(xmax-xmin) }
	Y := func(v float64) float64 {
		return float64(h) - padB - (float64(h)-padT-padB)*(v-ymin)/(ymax-ymin)
	}
	var b strings.Builder
	writeSVGHead(&b, w, h)
	fmt.Fprintf(&b, `<rect width="%d" height="%d" fill="#0d1216"/>`, w, h)
	if title != "" {
		fmt.Fprintf(&b, `<text x="12" y="16" fill="#8aa0ad" font-family="sans-serif" font-size="13">%s</text>`, escSVG(title))
	}
	fmt.Fprintf(&b, `<rect x="%.1f" y="%.1f" width="%.1f" height="%.1f" fill="none" stroke="#1d3342"/>`, padL, padT, float64(w)-padL-padR, float64(h)-padT-padB)
	for _, p := range pts {
		col := bandColor(p.Band)
		fmt.Fprintf(&b, `<rect x="%.1f" y="%.1f" width="6" height="6" fill="%s"><title>%s</title></rect>`, X(p.X)-3, Y(p.Y)-3, col, escSVG(p.ID))
	}
	fmt.Fprintf(&b, `<text x="%.1f" y="%.1f" fill="#8aa0ad" font-family="sans-serif" font-size="12" text-anchor="middle">%s</text>`, float64(w)/2, float64(h)-8, escSVG(xLabel))
	fmt.Fprintf(&b, `<text x="14" y="%.1f" fill="#8aa0ad" font-family="sans-serif" font-size="12" transform="rotate(-90 14 %.1f)">%s</text>`, float64(h)/2, float64(h)/2, escSVG(yLabel))
	b.WriteString(`</svg>`)
	return b.String()
}

// BarsSVG renders horizontal LPD bars for top rows.
func BarsSVG(title string, board LPD, max int) string {
	const w = 960
	rows := board.Top
	if max > 0 && len(rows) > max {
		rows = rows[:max]
	}
	h := 40 + len(rows)*28
	if h < 120 {
		h = 120
	}
	var maxLPD float64
	for _, r := range rows {
		if r.LPD > maxLPD {
			maxLPD = r.LPD
		}
	}
	if maxLPD <= 0 {
		maxLPD = 1
	}
	var b strings.Builder
	writeSVGHead(&b, w, h)
	fmt.Fprintf(&b, `<rect width="%d" height="%d" fill="#0d1216"/>`, w, h)
	if title != "" {
		fmt.Fprintf(&b, `<text x="12" y="18" fill="#8aa0ad" font-family="sans-serif" font-size="13">%s</text>`, escSVG(title))
	}
	for i, r := range rows {
		y := 32 + i*28
		bw := (float64(w) - 200) * (r.LPD / maxLPD)
		fmt.Fprintf(&b, `<text x="12" y="%d" fill="#c5d0d8" font-family="sans-serif" font-size="12">%s</text>`, y+12, escSVG(r.ID))
		fmt.Fprintf(&b, `<rect x="160" y="%d" width="%.1f" height="16" fill="%s"/>`, y, bw, bandColor(r.Band))
		fmt.Fprintf(&b, `<text x="%.1f" y="%d" fill="#8aa0ad" font-family="sans-serif" font-size="11">%.2f</text>`, 168+bw, y+12, r.LPD)
	}
	b.WriteString(`</svg>`)
	return b.String()
}

func bandColor(band string) string {
	switch band {
	case "gold":
		return "#e0a458"
	case "near":
		return "#6ea8fe"
	case "trap":
		return "#f07178"
	case "keep":
		return "#3dd6c6"
	default:
		return "#5a7a8a"
	}
}

func writeSVGHead(b *strings.Builder, w, h int) {
	fmt.Fprintf(b, `<svg xmlns="http://www.w3.org/2000/svg" width="%d" height="%d" viewBox="0 0 %d %d">`, w, h, w, h)
}

func emptySVG(w, h int, title, msg string) string {
	var b strings.Builder
	writeSVGHead(&b, w, h)
	fmt.Fprintf(&b, `<rect width="%d" height="%d" fill="#0d1216"/>`, w, h)
	if title != "" {
		fmt.Fprintf(&b, `<text x="12" y="18" fill="#8aa0ad" font-family="sans-serif" font-size="13">%s</text>`, escSVG(title))
	}
	fmt.Fprintf(&b, `<text x="%d" y="%d" fill="#5a7a8a" font-family="sans-serif" font-size="14" text-anchor="middle">%s</text>`, w/2, h/2, escSVG(msg))
	b.WriteString(`</svg>`)
	return b.String()
}

func escSVG(s string) string {
	r := strings.NewReplacer(`&`, "&amp;", `<`, "&lt;", `>`, "&gt;", `"`, "&quot;")
	return r.Replace(s)
}

func minMax(vs []float64) (float64, float64) {
	lo, hi := vs[0], vs[0]
	for _, v := range vs[1:] {
		if v < lo {
			lo = v
		}
		if v > hi {
			hi = v
		}
	}
	return lo, hi
}
