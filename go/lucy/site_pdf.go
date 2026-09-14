package lucy

import (
	"fmt"
	"sort"
	"strings"
)

// SitePDF builds the Tide/River-parity board pack:
// cover · compare(mode×dtype×arch) · Acc-keep · LPD · thru · bands · charts.
// Pure board-derived (no Tide Store / gofpdf). Tide Store-only grids remain optional host polish.
func SitePDF(board LPD, maxSeries int) ([]byte, error) {
	if maxSeries <= 0 {
		maxSeries = 8
	}
	pages := []pdfPage{
		coverPage(board),
		compareGridPage(board),
		nearPage(board),
		tablePage(board),
		thruPage(board),
		bandsPage(board),
	}
	pack := BuildBoardCharts(board, maxSeries)
	barH := 40 + minInt(maxSeries+4, maxInt(1, len(board.Top)))*28
	if barH < 120 {
		barH = 120
	}
	for _, p := range []pdfPage{
		{title: "Consciousness radar", jpg: pack.ConsciousnessJPG, w: 960, h: 480},
		{title: "Memory density radar", jpg: pack.DensityJPG, w: 960, h: 480},
		{title: "Q% vs RAM", jpg: pack.ScatterJPG, w: 960, h: 440},
		{title: "Top LPD", jpg: pack.BarsJPG, w: 960, h: barH},
	} {
		if len(p.jpg) > 0 {
			pages = append(pages, p)
		}
	}
	return writePDF(pages, Version)
}

// BoardPDF is the host-facing PDF export — River-parity site pack (0.9+).
func BoardPDF(board LPD, maxSeries int) ([]byte, error) {
	return SitePDF(board, maxSeries)
}

func nearPage(board LPD) pdfPage {
	floor := LPDKeepFloor
	rows := keepRows(board)
	sort.SliceStable(rows, func(i, j int) bool {
		if rows[i].RelAcc == rows[j].RelAcc {
			return rows[i].RAMKiB < rows[j].RAMKiB
		}
		return rows[i].RelAcc > rows[j].RelAcc
	})
	thr := board.PeakAcc * floor
	lines := []string{
		fmt.Sprintf("Acc keep band — %.0f%% to 100%% of champ", floor*100),
		fmt.Sprintf("Champ Acc %.1f%% (%s). Threshold %.1f%%. %d / %d cells in band.",
			board.PeakAcc, clipID(board.AccChamp.ID, 40), thr, len(rows), board.N),
		"",
		fmt.Sprintf("%-4s %-16s %7s %7s %8s %8s %6s", "#", "id", "Acc%", "keep%", "RAM", "Thru", "band"),
		strings.Repeat("-", 72),
	}
	max := 28
	for i, r := range rows {
		if i >= max {
			lines = append(lines, fmt.Sprintf("… %d more", len(rows)-max))
			break
		}
		lines = append(lines, fmt.Sprintf("%-4d %-16s %6.1f%% %6.1f%% %7.0f %7.0f %6s",
			i+1, clipID(r.ID, 16), r.Acc, r.RelAcc*100, r.RAMKiB, r.Thru, r.Band))
	}
	if len(rows) == 0 {
		lines = append(lines, "(no cells above keep floor)")
	}
	return pdfPage{title: "Acc keep", lines: lines}
}

func thruPage(board LPD) pdfPage {
	rows := append([]LPDRow(nil), board.Top...)
	if len(board.Pool) > len(rows) {
		rows = append([]LPDRow(nil), board.Pool...)
	}
	sort.SliceStable(rows, func(i, j int) bool {
		return rows[i].Thru > rows[j].Thru
	})
	bestID, bestThru := "", 0.0
	if len(rows) > 0 {
		bestID, bestThru = rows[0].ID, rows[0].Thru
	}
	lines := []string{
		"Throughput ranking",
		fmt.Sprintf("Best %.0f/s (%s). %d cells.", bestThru, clipID(bestID, 40), len(rows)),
		"",
		fmt.Sprintf("%-4s %-16s %8s %7s %7s %8s %6s", "#", "id", "Thru", "Acc%", "Avail%", "LPD", "band"),
		strings.Repeat("-", 72),
	}
	max := 28
	for i, r := range rows {
		if i >= max {
			lines = append(lines, fmt.Sprintf("… %d more", len(rows)-max))
			break
		}
		lines = append(lines, fmt.Sprintf("%-4d %-16s %7.0f %6.1f%% %6.1f%% %8.3g %6s",
			i+1, clipID(r.ID, 16), r.Thru, r.Acc, r.Avail, r.LPD, r.Band))
	}
	return pdfPage{title: "Throughput", lines: lines}
}

func bandsPage(board LPD) pdfPage {
	lines := []string{
		"Gold / lean / trap bands",
		fmt.Sprintf("gold=%d near=%d lean=%d trap=%d", len(board.Gold), len(board.Near), len(board.Lean), len(board.Trap)),
		"",
	}
	addBand := func(title string, rows []LPDRow) {
		lines = append(lines, title)
		if len(rows) == 0 {
			lines = append(lines, "  (none)", "")
			return
		}
		for i, r := range rows {
			if i >= 8 {
				lines = append(lines, fmt.Sprintf("  … %d more", len(rows)-8))
				break
			}
			lines = append(lines, fmt.Sprintf("  %-16s LPD=%.3g Q=%.1f%% RAM=%.0f Acc=%.1f",
				clipID(r.ID, 16), r.LPD, r.Q*100, r.RAMKiB, r.Acc))
		}
		lines = append(lines, "")
	}
	addBand("Gold", board.Gold)
	addBand("Near", board.Near)
	addBand("Lean", board.Lean)
	addBand("Trap", board.Trap)
	return pdfPage{title: "Bands", lines: lines}
}

func keepRows(board LPD) []LPDRow {
	src := board.Top
	if len(board.Pool) > 0 {
		src = board.Pool
	}
	out := make([]LPDRow, 0, len(src))
	for _, r := range src {
		if r.RelAcc >= LPDKeepFloor {
			out = append(out, r)
		}
	}
	return out
}

func clipID(s string, n int) string {
	if n <= 0 || len(s) <= n {
		return s
	}
	if n <= 1 {
		return s[:n]
	}
	return s[:n-1] + "…"
}
