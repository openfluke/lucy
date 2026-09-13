package lucy_test

import (
	"strings"
	"testing"

	"github.com/openfluke/lucy/lucy"
)

func TestRadarSVG(t *testing.T) {
	board := lucy.BuildLPD([]lucy.Sample{
		{ID: "f32", Acc: 90, Thru: 200, Avail: 40, Score: 100, RAMKiB: 1000},
		{ID: "int8", Acc: 82, Thru: 180, Avail: 38, Score: 85, RAMKiB: 180},
	})
	svg := lucy.RadarSVG("Consciousness radar", lucy.ConsciousnessSeries(board, 8))
	if !strings.Contains(svg, "<svg") || !strings.Contains(svg, "Acc") {
		t.Fatalf("bad svg")
	}
	sc := lucy.ScatterSVG("Q", "RAM", "Q%", lucy.LPDScatterPoints(board))
	if !strings.Contains(sc, "<svg") {
		t.Fatal("scatter")
	}
	bars := lucy.BarsSVG("Top", board, 8)
	if !strings.Contains(bars, "int8") && !strings.Contains(bars, "f32") {
		t.Fatal("bars")
	}
}
