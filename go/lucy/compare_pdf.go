package lucy

import (
	"fmt"
	"sort"
	"strings"
)

type gridKey struct {
	Mode, DType, Arch string
}

type gridCell struct {
	Mode, DType, Arch string
	N                 int
	SumAcc, BestAcc   float64
	SumThru           float64
}

// compareGridPage summarizes mode × dtype × arch from the board (River compare subset).
func compareGridPage(board LPD) pdfPage {
	src := board.Top
	if len(board.Pool) > 0 {
		src = board.Pool
	}
	m := map[gridKey]*gridCell{}
	for _, r := range src {
		k := gridKey{Mode: r.Mode, DType: r.DType, Arch: r.Arch}
		c, ok := m[k]
		if !ok {
			c = &gridCell{Mode: r.Mode, DType: r.DType, Arch: r.Arch}
			m[k] = c
		}
		c.N++
		c.SumAcc += r.Acc
		c.SumThru += r.Thru
		if r.Acc > c.BestAcc {
			c.BestAcc = r.Acc
		}
	}
	cells := make([]*gridCell, 0, len(m))
	for _, c := range m {
		cells = append(cells, c)
	}
	sort.SliceStable(cells, func(i, j int) bool {
		if cells[i].BestAcc == cells[j].BestAcc {
			return cells[i].Mode+cells[i].DType < cells[j].Mode+cells[j].DType
		}
		return cells[i].BestAcc > cells[j].BestAcc
	})
	lines := []string{
		"Compare — mode × dtype × arch",
		fmt.Sprintf("%d groups from %d cells (board-derived; Tide Store grids stay optional).", len(cells), len(src)),
		"",
		fmt.Sprintf("%-14s %-10s %-12s %4s %8s %8s %8s", "mode", "dtype", "arch", "n", "meanAcc", "bestAcc", "meanThru"),
		strings.Repeat("-", 72),
	}
	max := 32
	for i, c := range cells {
		if i >= max {
			lines = append(lines, fmt.Sprintf("… %d more groups", len(cells)-max))
			break
		}
		meanAcc := 0.0
		meanThru := 0.0
		if c.N > 0 {
			meanAcc = c.SumAcc / float64(c.N)
			meanThru = c.SumThru / float64(c.N)
		}
		lines = append(lines, fmt.Sprintf("%-14s %-10s %-12s %4d %7.1f%% %7.1f%% %8.0f",
			clipID(nz(c.Mode, "—"), 14), clipID(nz(c.DType, "—"), 10), clipID(nz(c.Arch, "—"), 12),
			c.N, meanAcc, c.BestAcc, meanThru))
	}
	if len(cells) == 0 {
		lines = append(lines, "(no groups — samples need mode/dtype/arch)")
	}
	return pdfPage{title: "Compare", lines: lines}
}

func nz(s, fallback string) string {
	if s == "" {
		return fallback
	}
	return s
}
