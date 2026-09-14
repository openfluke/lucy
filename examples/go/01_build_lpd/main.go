// Example: BuildLPD — why int8 leads LPD and bin is a trap.
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/openfluke/lucy/lucy"
)

func samplesPath() string {
	_, file, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(file), "..", "..", "shared", "samples.json")
}

func main() {
	raw, err := os.ReadFile(samplesPath())
	if err != nil {
		panic(err)
	}
	var req struct {
		Samples []lucy.Sample       `json:"samples"`
		Options lucy.DensityOptions `json:"options"`
	}
	if err := json.Unmarshal(raw, &req); err != nil {
		panic(err)
	}
	fmt.Println("lucy", lucy.Version)
	board := lucy.BuildLPDWithOptions(req.Samples, req.Options)
	fmt.Printf("Acc champ=%s  Score champ=%s  Gold-std=%s\n", board.AccChamp.ID, board.Champ.ID, board.GoldStd.ID)
	fmt.Println("Top LPD:")
	for i, r := range board.Top {
		if i >= 4 {
			break
		}
		fmt.Printf("  #%d %-6s band=%-4s LPD=%.4g Q=%.1f%% RAM=%.0f\n", i+1, r.ID, r.Band, r.LPD, r.Q*100, r.RAMKiB)
	}
	if board.Top[0].ID != "int8" {
		panic("expected int8 to lead LPD")
	}
	fmt.Println("OK — int8 leads (gold denser than Acc champ)")
}
