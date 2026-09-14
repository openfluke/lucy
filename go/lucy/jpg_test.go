package lucy_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/openfluke/lucy/lucy"
)

func TestJPGAndReport(t *testing.T) {
	board := lucy.BuildLPD([]lucy.Sample{
		{ID: "f32", Acc: 90, Thru: 200, Avail: 40, Score: 100, RAMKiB: 1000},
		{ID: "int8", Acc: 82, Thru: 180, Avail: 38, Score: 85, RAMKiB: 180},
	})
	pack := lucy.BuildBoardCharts(board, 8)
	if len(pack.ConsciousnessJPG) < 50 || pack.ConsciousnessJPG[0] != 0xff {
		t.Fatalf("jpg magic missing len=%d", len(pack.ConsciousnessJPG))
	}
	dir := t.TempDir()
	if err := lucy.WriteReportDir(dir, board, 8); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(dir, "index.html")); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(dir, "bars.jpg")); err != nil {
		t.Fatal(err)
	}
}

func TestHandlerLPD(t *testing.T) {
	h := lucy.Handler()
	reqBody, _ := json.Marshal(lucy.BuildRequest{Samples: []lucy.Sample{
		{ID: "a", Acc: 90, Thru: 100, Avail: 50, Score: 80, RAMKiB: 500},
	}})
	req := httptest.NewRequest(http.MethodPost, "/api/lpd", bytes.NewReader(reqBody))
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)
	if rr.Code != 200 {
		t.Fatalf("status %d %s", rr.Code, rr.Body.String())
	}
	var resp lucy.BuildResponse
	if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
		t.Fatal(err)
	}
	if resp.Version == "" {
		t.Fatal("empty version")
	}
}
