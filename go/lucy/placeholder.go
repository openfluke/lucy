package lucy

// KeepFloor is the default Acc-keep gate for LPD (RelAcc ≥ this, else LPD = 0).
// Tunable — hosts may override for their own duty-cycle experiments.
const KeepFloor = 0.70

// Sample is one finished cell for density ranking. Hosts map their own IDs.
type Sample struct {
	ID     string
	Mode   string
	DType  string
	Arch   string
	Score  float64
	Soft   float64
	Acc    float64 // hard argmax
	Thru   float64
	Avail  float64
	RAMKiB float64
}

// Board is a ranked Lucy snapshot (consciousness then Pareto density).
// Placeholder — see BuildLPD once formulas land here from welvet/lucy.
type Board struct {
	Formula string
	Top     []Row
}

// Row is one cell vs learner peaks and Acc-champ RAM.
type Row struct {
	ID  string
	Q   float64
	LPD float64 // Q × shrink; 0 unless Acc keep ≥ KeepFloor
}

// BuildLPD ranks samples for consciousness then Lucy Pareto density.
// TODO: port from github.com/openfluke/welvet/lucy and keep testdata goldens in sync.
func BuildLPD(samples []Sample) Board {
	_ = samples
	return Board{
		Formula: "placeholder — Score = T×Avail×Acc/10_000; LPD = Q×shrink if RelAcc≥KeepFloor else 0",
	}
}
