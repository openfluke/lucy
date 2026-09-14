package lucy

import (
	"bytes"
	"fmt"
	"strings"
)

// PDF helpers (pages + writePDF). BoardPDF / SitePDF live in site_pdf.go.

func coverPage(board LPD) pdfPage {
	lines := []string{
		"Lucy LPD board",
		fmt.Sprintf("version %s · n=%d", Version, board.N),
		"",
		"Champs",
		fmt.Sprintf("  Score champ   %s", champLine(board.Champ)),
		fmt.Sprintf("  Acc champ     %s", champLine(board.AccChamp)),
		fmt.Sprintf("  Live Q champ  %s", champLine(board.LiveChamp)),
		fmt.Sprintf("  Gold-std      %s  LPD=%.3g", board.GoldStd.ID, board.GoldStd.LPD),
		fmt.Sprintf("  Lean champ    %s  RAM=%.0f KiB", board.LeanChamp.ID, board.LeanChamp.RAMKiB),
		"",
		"Formula",
	}
	lines = append(lines, wrapWords(board.Formula, 86)...)
	return pdfPage{title: "Cover", lines: lines}
}

func champLine(c LPDChamp) string {
	if c.ID == "" {
		return "—"
	}
	return fmt.Sprintf("%s  Acc=%.1f Thru=%.0f Avail=%.0f RAM=%.0f", c.ID, c.Acc, c.Thru, c.Avail, c.RAMKiB)
}

func tablePage(board LPD) pdfPage {
	lines := []string{
		"Top LPD ranking",
		fmt.Sprintf("%-4s %-16s %-6s %8s %7s %7s %8s %6s", "#", "id", "band", "LPD", "Q", "Acc%", "RAM", "shrink"),
		strings.Repeat("-", 72),
	}
	max := 28
	for i, r := range board.Top {
		if i >= max {
			lines = append(lines, fmt.Sprintf("… %d more (see board.csv)", len(board.Top)-max))
			break
		}
		id := r.ID
		if len(id) > 16 {
			id = id[:15] + "…"
		}
		lines = append(lines, fmt.Sprintf("%-4d %-16s %-6s %8.3g %6.1f%% %6.1f%% %7.0f %6.2f",
			i+1, id, r.Band, r.LPD, r.Q*100, r.RelAcc*100, r.RAMKiB, r.Shrink))
	}
	return pdfPage{title: "LPD table", lines: lines}
}

func wrapWords(s string, width int) []string {
	words := strings.Fields(s)
	if len(words) == 0 {
		return nil
	}
	var out []string
	var cur strings.Builder
	for _, w := range words {
		if cur.Len() == 0 {
			cur.WriteString(w)
			continue
		}
		if cur.Len()+1+len(w) > width {
			out = append(out, cur.String())
			cur.Reset()
			cur.WriteString(w)
			continue
		}
		cur.WriteByte(' ')
		cur.WriteString(w)
	}
	if cur.Len() > 0 {
		out = append(out, cur.String())
	}
	return out
}

type pdfPage struct {
	title string
	jpg   []byte
	w, h  int
	lines []string
}

func writePDF(pages []pdfPage, version string) ([]byte, error) {
	if len(pages) == 0 {
		return nil, fmt.Errorf("no pages")
	}
	type obj struct {
		id   int
		body string
		dict string
		raw  []byte
	}
	objs := []obj{
		{id: 1, body: "<< /Type /Catalog /Pages 2 0 R >>"},
		{id: 3, body: "<< /Type /Font /Subtype /Type1 /BaseFont /Helvetica >>"},
		{id: 4, body: "<< /Type /Font /Subtype /Type1 /BaseFont /Courier >>"},
	}
	fontID, monoID := 3, 4
	nextID := 5
	var kids []string
	const pageW, pageH = 612.0, 792.0

	for _, p := range pages {
		pageID := nextID
		contentID := nextID + 1
		kids = append(kids, fmt.Sprintf("%d 0 R", pageID))
		if len(p.jpg) > 0 {
			imageID := nextID + 2
			nextID += 3
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
			continue
		}
		nextID += 2
		var sb strings.Builder
		sb.WriteString("BT\n")
		y := 760.0
		for i, line := range p.lines {
			font := "/F1"
			size := 10.0
			if i == 0 {
				size = 16
			} else if strings.HasPrefix(line, "  ") || (len(line) > 0 && line[0] >= '0' && line[0] <= '9') || strings.HasPrefix(line, "#") || strings.HasPrefix(line, "-") || strings.Contains(line, "  LPD") {
				font = "/F2"
				size = 8
			}
			if i == 1 && strings.Contains(line, "id") && strings.Contains(line, "band") {
				font = "/F2"
				size = 8
			}
			if i > 0 && (strings.HasPrefix(p.lines[0], "Top LPD") || p.title == "LPD table") {
				font = "/F2"
				size = 8
			}
			esc := pdfEscape(line)
			if i == 0 {
				fmt.Fprintf(&sb, "%s %.0f Tf 36 %.0f Td (%s) Tj\n", font, size, y, esc)
			} else {
				fmt.Fprintf(&sb, "0 -%.0f Td %s %.0f Tf (%s) Tj\n", size+4, font, size, esc)
			}
		}
		sb.WriteString("ET\n")
		objs = append(objs,
			obj{id: pageID, body: fmt.Sprintf(
				"<< /Type /Page /Parent 2 0 R /MediaBox [0 0 %.0f %.0f] /Contents %d 0 R /Resources << /Font << /F1 %d 0 R /F2 %d 0 R >> >> >>",
				pageW, pageH, contentID, fontID, monoID,
			)},
			obj{id: contentID, raw: []byte(sb.String())},
		)
	}

	objs = append(objs, obj{id: 2, body: fmt.Sprintf("<< /Type /Pages /Kids [%s] /Count %d >>", strings.Join(kids, " "), len(kids))})

	// sort objs by id for xref
	byID := map[int]obj{}
	maxID := 0
	for _, o := range objs {
		byID[o.id] = o
		if o.id > maxID {
			maxID = o.id
		}
	}

	var buf bytes.Buffer
	buf.WriteString("%PDF-1.4\n%\xff\xff\xff\xff\n")
	offsets := make([]int, maxID+1)
	for id := 1; id <= maxID; id++ {
		o, ok := byID[id]
		if !ok {
			continue
		}
		offsets[id] = buf.Len()
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
	fmt.Fprintf(&buf, "xref\n0 %d\n", maxID+1)
	buf.WriteString("0000000000 65535 f \n")
	for id := 1; id <= maxID; id++ {
		if _, ok := byID[id]; !ok {
			buf.WriteString("0000000000 65535 f \n")
			continue
		}
		fmt.Fprintf(&buf, "%010d 00000 n \n", offsets[id])
	}
	fmt.Fprintf(&buf, "trailer\n<< /Size %d /Root 1 0 R >>\nstartxref\n%d\n%%%%EOF\n", maxID+1, xref)
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
