package lucy_test

import (
	"bytes"
	"testing"

	"github.com/openfluke/lucy/lucy"
)

func TestBoardPDF(t *testing.T) {
	board := lucy.BuildLPD([]lucy.Sample{
		{ID: "f32", Acc: 90, Thru: 200, Avail: 40, Score: 100, RAMKiB: 1000},
		{ID: "int8", Acc: 82, Thru: 180, Avail: 38, Score: 85, RAMKiB: 180},
	})
	pdf, err := lucy.BoardPDF(board, 8)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.HasPrefix(pdf, []byte("%PDF-1.4")) {
		t.Fatalf("bad header %q", pdf[:min(16, len(pdf))])
	}
	if !bytes.Contains(pdf, []byte("%%EOF")) {
		t.Fatal("missing EOF")
	}
	// cover + table are text streams
	if !bytes.Contains(pdf, []byte("Lucy LPD board")) {
		t.Fatal("missing cover title")
	}
	if !bytes.Contains(pdf, []byte("Top LPD ranking")) {
		t.Fatal("missing table page")
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
