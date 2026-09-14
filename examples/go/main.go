// Example: feed samples → LPD board → CSV + PDF.
package main

import (
	"fmt"
	"os"

	"github.com/openfluke/lucy/lucy"
)

func main() {
	samples := []lucy.Sample{
		{ID: "f32", Acc: 90, Thru: 200, Avail: 40, Score: 100, RAMKiB: 1000},
		{ID: "int8", Acc: 82, Thru: 180, Avail: 38, Score: 85, RAMKiB: 180},
	}
	board := lucy.BuildLPD(samples)
	fmt.Printf("top=%s LPD=%.3g\n", board.Top[0].ID, board.Top[0].LPD)
	_ = os.WriteFile("board.csv", []byte(lucy.BoardCSV(board)), 0o644)
	pdf, err := lucy.BoardPDF(board, 8)
	if err != nil {
		panic(err)
	}
	_ = os.WriteFile("board.pdf", pdf, 0o644)
	fmt.Println("wrote board.csv board.pdf")
}
