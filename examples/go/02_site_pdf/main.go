package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/openfluke/lucy/lucy"
)

func main() {
	_, file, _, _ := runtime.Caller(0)
	base := filepath.Join(filepath.Dir(file), "..", "..")
	raw, err := os.ReadFile(filepath.Join(base, "shared", "samples.json"))
	if err != nil {
		panic(err)
	}
	var req struct {
		Samples []lucy.Sample `json:"samples"`
	}
	_ = json.Unmarshal(raw, &req)
	board := lucy.BuildLPD(req.Samples)
	outDir := filepath.Join(base, "out", "go")
	_ = os.MkdirAll(outDir, 0o755)
	pdf, err := lucy.SitePDF(board, 8)
	if err != nil {
		panic(err)
	}
	_ = os.WriteFile(filepath.Join(outDir, "site.pdf"), pdf, 0o644)
	_ = os.WriteFile(filepath.Join(outDir, "board.csv"), []byte(lucy.BoardCSV(board)), 0o644)
	fmt.Printf("lucy %s floors keep=%.2f\n", lucy.Version, lucy.DefaultDensityOptions().KeepFloor)
	fmt.Printf("wrote %s/site.pdf (%d bytes)\n", outDir, len(pdf))
}
