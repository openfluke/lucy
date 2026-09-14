package lucy

import (
	"bytes"
	"fmt"
	"strings"
)

// BoardPDF builds a multi-page PDF (one chart page each) from a board.
// Embeds JPEG charts (stdlib). Full Tide/River PDF parity still later.
func BoardPDF(board LPD, maxSeries int) ([]byte, error) {
	if maxSeries <= 0 {
		maxSeries = 8
	}
	pack := BuildBoardCharts(board, maxSeries)
	barH := 40 + minInt(maxSeries+4, maxInt(1, len(board.Top)))*28
	if barH < 120 {
		barH = 120
	}
	pages := []pdfPage{
		{title: "Consciousness radar", jpg: pack.ConsciousnessJPG, w: 960, h: 480},
		{title: "Memory density radar", jpg: pack.DensityJPG, w: 960, h: 480},
		{title: "Q% vs RAM", jpg: pack.ScatterJPG, w: 960, h: 440},
		{title: "Top LPD", jpg: pack.BarsJPG, w: 960, h: barH},
	}
	var filtered []pdfPage
	for _, p := range pages {
		if len(p.jpg) > 0 {
			filtered = append(filtered, p)
		}
	}
	if len(filtered) == 0 {
		return nil, fmt.Errorf("no chart images to embed")
	}
	return writePDF(filtered, Version)
}

type pdfPage struct {
	title string
	jpg   []byte
	w, h  int
}

func writePDF(pages []pdfPage, version string) ([]byte, error) {
	n := len(pages)
	fontID := 3
	first := 4
	kids := make([]string, n)
	for i := 0; i < n; i++ {
		kids[i] = fmt.Sprintf("%d 0 R", first+i*3)
	}

	type obj struct {
		id   int
		body string
		dict string
		raw  []byte
	}
	objs := []obj{
		{id: 1, body: "<< /Type /Catalog /Pages 2 0 R >>"},
		{id: 2, body: fmt.Sprintf("<< /Type /Pages /Kids [%s] /Count %d >>", strings.Join(kids, " "), n)},
		{id: 3, body: "<< /Type /Font /Subtype /Type1 /BaseFont /Helvetica >>"},
	}

	const pageW, pageH = 612.0, 792.0
	for i, p := range pages {
		pageID := first + i*3
		contentID := pageID + 1
		imageID := pageID + 2
		margin := 36.0
		maxW := pageW - 2*margin
		maxH := pageH - 100
		scale := maxW / float64(p.w)
		if float64(p.h)*scale > maxH {
			scale = maxH / float64(p.h)
		}
		imgW := float64(p.w) * scale
		imgH := float64(p.h) * scale
		x := (pageW - imgW) / 2
		y := pageH - 72 - imgH
		title := pdfEscape(fmt.Sprintf("%s (lucy %s)", p.title, version))
		content := fmt.Sprintf(
			"BT /F1 12 Tf 36 760 Td (%s) Tj ET\nq %.2f 0 0 %.2f %.2f %.2f cm /Im0 Do Q\n",
			title, imgW, imgH, x, y,
		)
		objs = append(objs,
			obj{id: pageID, body: fmt.Sprintf(
				"<< /Type /Page /Parent 2 0 R /MediaBox [0 0 %.0f %.0f] /Contents %d 0 R /Resources << /Font << /F1 %d 0 R >> /XObject << /Im0 %d 0 R >> >> >>",
				pageW, pageH, contentID, fontID, imageID,
			)},
			obj{id: contentID, raw: []byte(content)},
			obj{id: imageID, dict: fmt.Sprintf(
				"/Type /XObject /Subtype /Image /Width %d /Height %d /ColorSpace /DeviceRGB /BitsPerComponent 8 /Filter /DCTDecode",
				p.w, p.h,
			), raw: p.jpg},
		)
	}

	var buf bytes.Buffer
	buf.WriteString("%PDF-1.4\n%\xff\xff\xff\xff\n")
	offsets := make([]int, len(objs)+1)
	for _, o := range objs {
		offsets[o.id] = buf.Len()
		if o.raw != nil {
			dict := o.dict
			if dict == "" {
				dict = fmt.Sprintf("/Length %d", len(o.raw))
			} else {
				dict += fmt.Sprintf(" /Length %d", len(o.raw))
			}
			fmt.Fprintf(&buf, "%d 0 obj\n<< %s >>\nstream\n", o.id, dict)
			buf.Write(o.raw)
			buf.WriteString("\nendstream\nendobj\n")
		} else {
			fmt.Fprintf(&buf, "%d 0 obj\n%s\nendobj\n", o.id, o.body)
		}
	}
	xref := buf.Len()
	fmt.Fprintf(&buf, "xref\n0 %d\n", len(objs)+1)
	buf.WriteString("0000000000 65535 f \n")
	for id := 1; id <= len(objs); id++ {
		fmt.Fprintf(&buf, "%010d 00000 n \n", offsets[id])
	}
	fmt.Fprintf(&buf, "trailer\n<< /Size %d /Root 1 0 R >>\nstartxref\n%d\n%%%%EOF\n", len(objs)+1, xref)
	return buf.Bytes(), nil
}

func pdfEscape(s string) string {
	r := strings.NewReplacer(`\`, `\\`, `(`, `\(`, `)`, `\)`)
	return r.Replace(s)
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
