package lucy_test

import (
	"testing"

	"github.com/openfluke/lucy/lucy"
)

func TestBoardChartsPNG(t *testing.T) {
	board := lucy.BuildLPD([]lucy.Sample{
		{ID: "f32", Acc: 90, Thru: 200, Avail: 40, Score: 100, RAMKiB: 1000},
		{ID: "int8", Acc: 82, Thru: 180, Avail: 38, Score: 85, RAMKiB: 180},
	})
	pack := lucy.BuildBoardCharts(board, 8)
	if len(pack.ConsciousnessPNG) < 100 || pack.ConsciousnessPNG[0] != 0x89 {
		t.Fatalf("expected PNG magic, got len=%d", len(pack.ConsciousnessPNG))
	}
	if pack.ConsciousnessSVG == "" || pack.BarsSVG == "" {
		t.Fatal("missing svg")
	}
}
