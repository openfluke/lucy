package lucy

// DensityOptions tunes LPD / gold / lean floors. Zero values mean defaults.
type DensityOptions struct {
	KeepFloor float64 `json:"keep_floor,omitempty"` // Acc keep for LPD > 0 (default 0.70)
	GoldKeep  float64 `json:"gold_keep,omitempty"`  // pillar keep for gold/near (default 0.80)
	LeanKeep  float64 `json:"lean_keep,omitempty"`  // Acc keep for lean band (default 0.95)
	GoldRAM   float64 `json:"gold_ram,omitempty"`   // max RAM frac of Acc champ for gold/trap (default 0.20)
	NearRAM   float64 `json:"near_ram,omitempty"`   // max RAM frac for near band (default 0.50)
	ShrinkCap float64 `json:"shrink_cap,omitempty"` // max AccChampRAM/thisRAM (default 32)
}

func (o DensityOptions) normalized() DensityOptions {
	if o.KeepFloor <= 0 {
		o.KeepFloor = LPDKeepFloor
	}
	if o.GoldKeep <= 0 {
		o.GoldKeep = LPDGoldKeep
	}
	if o.LeanKeep <= 0 {
		o.LeanKeep = LPDLeanKeep
	}
	if o.GoldRAM <= 0 {
		o.GoldRAM = LPDGoldRAM
	}
	if o.NearRAM <= 0 {
		o.NearRAM = LPDNearRAM
	}
	if o.ShrinkCap <= 0 {
		o.ShrinkCap = LPDShrinkCap
	}
	return o
}
