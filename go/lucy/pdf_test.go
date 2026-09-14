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
		{ID: "bin", Acc: 12, Thru: 400, Avail: 50, Score: 40, RAMKiB: 40},
	})
	pdf, err := lucy.SitePDF(board, 8)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.HasPrefix(pdf, []byte("%PDF-1.4")) {
		t.Fatalf("bad header %q", pdf[:min(16, len(pdf))])
	}
	if !bytes.Contains(pdf, []byte("%%EOF")) {
		t.Fatal("missing EOF")
	}
	for _, want := range []string{
		"Lucy LPD board",
		"Top LPD ranking",
		"Acc keep band",
		"Throughput ranking",
		"Gold / lean / trap bands",
	} {
		if !bytes.Contains(pdf, []byte(want)) {
			t.Fatalf("missing %q", want)
		}
	}
	// BoardPDF aliases SitePDF
	pdf2, err := lucy.BoardPDF(board, 8)
	if err != nil || len(pdf2) < 100 {
		t.Fatal(err, len(pdf2))
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
