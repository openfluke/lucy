package lucy

import (
	"fmt"
	"strings"
)

// BoardCSV serializes board.Top as CSV (header + rows). Stable column order for hosts.
func BoardCSV(board LPD) string {
	var b strings.Builder
	b.WriteString("id,band,lpd,q,rel_acc,rel_thru,rel_avail,ram_kib,ram_frac,shrink,score,acc,thru,avail,mode,dtype,arch\n")
	for _, r := range board.Top {
		fmt.Fprintf(&b, "%s,%s,%.9g,%.9g,%.9g,%.9g,%.9g,%.9g,%.9g,%.9g,%.9g,%.9g,%.9g,%.9g,%s,%s,%s\n",
			csvCell(r.ID), csvCell(r.Band),
			r.LPD, r.Q, r.RelAcc, r.RelThru, r.RelAvail,
			r.RAMKiB, r.RAMFrac, r.Shrink, r.Score, r.Acc, r.Thru, r.Avail,
			csvCell(r.Mode), csvCell(r.DType), csvCell(r.Arch),
		)
	}
	return b.String()
}

func csvCell(s string) string {
	if strings.ContainsAny(s, ",\"\n\r") {
		return `"` + strings.ReplaceAll(s, `"`, `""`) + `"`
	}
	return s
}
