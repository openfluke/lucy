package lucy_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/openfluke/lucy/lucy"
)

func TestGoldensLPDV01(t *testing.T) {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("caller")
	}
	path := filepath.Join(filepath.Dir(file), "..", "..", "testdata", "goldens_lpd_v1.0.json")
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	var g struct {
		Version string        `json:"version"`
		Samples []lucy.Sample `json:"samples"`
		Expect  struct {
			ChampID    string `json:"champ_id"`
			AccChampID string `json:"acc_champ_id"`
			GoldStdID  string `json:"gold_std_id"`
			TrapIDs    []string `json:"trap_ids"`
			Top        []struct {
				ID  string  `json:"id"`
				LPD float64 `json:"lpd"`
				Band string `json:"band"`
			} `json:"top"`
		} `json:"expect"`
	}
	if err := json.Unmarshal(raw, &g); err != nil {
		t.Fatal(err)
	}
	if g.Version != lucy.Version {
		t.Fatalf("golden version %s != lucy.Version %s", g.Version, lucy.Version)
	}
	l := lucy.BuildLPD(g.Samples)
	if l.Champ.ID != g.Expect.ChampID || l.AccChamp.ID != g.Expect.AccChampID || l.GoldStd.ID != g.Expect.GoldStdID {
		t.Fatalf("champs got champ=%s acc=%s gold=%s", l.Champ.ID, l.AccChamp.ID, l.GoldStd.ID)
	}
	if len(l.Top) != len(g.Expect.Top) {
		t.Fatalf("top len %d != %d", len(l.Top), len(g.Expect.Top))
	}
	for i, want := range g.Expect.Top {
		got := l.Top[i]
		if got.ID != want.ID || got.Band != want.Band {
			t.Fatalf("top[%d] got %s/%s want %s/%s", i, got.ID, got.Band, want.ID, want.Band)
		}
		if abs(got.LPD-want.LPD) > 1e-9 {
			t.Fatalf("top[%d].LPD got %v want %v", i, got.LPD, want.LPD)
		}
	}
	if len(l.Trap) != len(g.Expect.TrapIDs) || (len(l.Trap) > 0 && l.Trap[0].ID != g.Expect.TrapIDs[0]) {
		t.Fatalf("trap %#v want %#v", ids(l.Trap), g.Expect.TrapIDs)
	}
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func ids(rs []lucy.LPDRow) []string {
	out := make([]string, len(rs))
	for i, r := range rs {
		out[i] = r.ID
	}
	return out
}
