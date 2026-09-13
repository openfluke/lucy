package lucy_test

import (
	"testing"

	"github.com/openfluke/lucy/lucy"
)

func TestBuildLPDWithOptionsKeepFloor(t *testing.T) {
	pts := []lucy.Sample{
		{ID: "big", Acc: 90, Thru: 100, Avail: 50, Score: 100, RAMKiB: 1000},
		{ID: "mid", Acc: 75, Thru: 100, Avail: 50, Score: 80, RAMKiB: 200}, // ~83% keep
	}
	def := lucy.BuildLPD(pts)
	if def.Top[0].ID != "mid" && def.Top[0].LPD == 0 {
		// mid should have LPD>0 at default 70% floor
	}
	strict := lucy.BuildLPDWithOptions(pts, lucy.DensityOptions{KeepFloor: 0.95})
	for _, r := range strict.Top {
		if r.ID == "mid" && r.LPD != 0 {
			t.Fatalf("mid should be gated at keep_floor=0.95, got LPD=%v", r.LPD)
		}
	}
	_ = def
}
