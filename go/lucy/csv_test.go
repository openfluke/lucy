package lucy_test

import (
	"strings"
	"testing"

	"github.com/openfluke/lucy/lucy"
)

func TestBoardCSV(t *testing.T) {
	board := lucy.BuildLPD([]lucy.Sample{
		{ID: "f32", Acc: 90, Thru: 200, Avail: 40, Score: 100, RAMKiB: 1000},
		{ID: "int8", Acc: 82, Thru: 180, Avail: 38, Score: 85, RAMKiB: 180},
	})
	csv := lucy.BoardCSV(board)
	if !strings.HasPrefix(csv, "id,band,lpd,") {
		t.Fatalf("bad header %q", csv[:min(40, len(csv))])
	}
	if !strings.Contains(csv, "int8") || !strings.Contains(csv, "f32") {
		t.Fatal("missing rows")
	}
	if lines := strings.Count(csv, "\n"); lines < 3 {
		t.Fatalf("want header+rows, got %d lines", lines)
	}
}

func TestFloorsMap(t *testing.T) {
	m := lucy.FloorsMap()
	if m["version"] != lucy.Version {
		t.Fatal(m["version"])
	}
	if m["keep_floor"].(float64) != lucy.LPDKeepFloor {
		t.Fatal(m["keep_floor"])
	}
}
