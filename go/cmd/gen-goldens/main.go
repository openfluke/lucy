package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/openfluke/lucy/lucy"
)

func main() {
	pts := []lucy.Sample{
		{ID: "f32", Mode: "sgd", DType: "float32", Arch: "single", Score: 100, Soft: 80, Acc: 90, Thru: 200, Avail: 40, RAMKiB: 1000},
		{ID: "int8", Mode: "sgd", DType: "int8", Arch: "single", Score: 85, Soft: 72, Acc: 82, Thru: 180, Avail: 38, RAMKiB: 180},
		{ID: "bin", Mode: "sgd", DType: "binary", Arch: "single", Score: 40, Soft: 20, Acc: 12, Thru: 400, Avail: 50, RAMKiB: 40},
		{ID: "fat", Mode: "tween", DType: "float32", Arch: "tricameral", Score: 95, Soft: 78, Acc: 88, Thru: 190, Avail: 35, RAMKiB: 3000},
	}
	l := lucy.BuildLPD(pts)
	out := map[string]any{
		"version":    lucy.Version,
		"keep_floor": lucy.LPDKeepFloor,
		"formula":    lucy.DensityFormula(),
		"samples":    pts,
		"expect": map[string]any{
			"champ_id":     l.Champ.ID,
			"acc_champ_id": l.AccChamp.ID,
			"gold_std_id":  l.GoldStd.ID,
			"trap_ids":     ids(l.Trap),
			"top":          rows(l.Top),
		},
	}
	b, _ := json.MarshalIndent(out, "", "  ")
	fmt.Println(string(b))
	_ = os.WriteFile("../testdata/goldens_lpd_v0.5.json", b, 0644)
}

func ids(rs []lucy.LPDRow) []string {
	out := make([]string, len(rs))
	for i, r := range rs {
		out[i] = r.ID
	}
	return out
}

func rows(rs []lucy.LPDRow) []map[string]any {
	out := make([]map[string]any, 0, len(rs))
	for _, r := range rs {
		out = append(out, map[string]any{
			"id": r.ID, "band": r.Band, "q": r.Q, "lpd": r.LPD,
			"rel_acc": r.RelAcc, "shrink": r.Shrink, "ram_kib": r.RAMKiB,
		})
	}
	return out
}
