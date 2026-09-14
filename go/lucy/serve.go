package lucy

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
)

// Handler returns an http.Handler with:
//
//	GET  /api/version
//	GET  /api/floors
//	POST /api/lpd          BuildRequest → BuildResponse
//	POST /api/chart-pack   BuildRequest → chart-pack JSON (svg + png/jpg b64)
//	POST /api/pdf          BuildRequest → application/pdf (River-parity site pack)
//	POST /api/site-pdf     same as /api/pdf
//	POST /api/csv          BuildRequest → text/csv
func Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/version", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]string{"version": Version})
	})
	mux.HandleFunc("/api/floors", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(FloorsMap())
	})
	mux.HandleFunc("/api/lpd", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "POST only", http.StatusMethodNotAllowed)
			return
		}
		raw, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		resp, err := BuildFromJSON(raw)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resp)
	})
	sitePDF := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "POST only", http.StatusMethodNotAllowed)
			return
		}
		raw, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		resp, err := BuildFromJSON(raw)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		pdf, err := SitePDF(resp.Board, 8)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "attachment; filename=lucy-site.pdf")
		_, _ = w.Write(pdf)
	}
	mux.HandleFunc("/api/pdf", sitePDF)
	mux.HandleFunc("/api/site-pdf", sitePDF)
	mux.HandleFunc("/api/csv", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "POST only", http.StatusMethodNotAllowed)
			return
		}
		raw, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		resp, err := BuildFromJSON(raw)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		w.Header().Set("Content-Type", "text/csv; charset=utf-8")
		w.Header().Set("Content-Disposition", "attachment; filename=lucy-board.csv")
		_, _ = io.WriteString(w, BoardCSV(resp.Board))
	})
	mux.HandleFunc("/api/chart-pack", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "POST only", http.StatusMethodNotAllowed)
			return
		}
		raw, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		resp, err := BuildFromJSON(raw)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		pack := BuildBoardCharts(resp.Board, 8)
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(ChartPackJSON(pack))
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		_, _ = io.WriteString(w, "lucy serve "+Version+"\nGET /api/version\nGET /api/floors\nPOST /api/lpd\nPOST /api/chart-pack\nPOST /api/pdf\nPOST /api/site-pdf\nPOST /api/csv\n")
	})
	return mux
}

// ChartPackJSON is the stable JSON shape for chart-pack (CLI + HTTP).
func ChartPackJSON(pack BoardCharts) map[string]any {
	return map[string]any{
		"version":               Version,
		"consciousness_svg":     pack.ConsciousnessSVG,
		"density_svg":           pack.DensitySVG,
		"scatter_svg":           pack.ScatterSVG,
		"bars_svg":              pack.BarsSVG,
		"consciousness_png_b64": b64(pack.ConsciousnessPNG),
		"density_png_b64":       b64(pack.DensityPNG),
		"scatter_png_b64":       b64(pack.ScatterPNG),
		"bars_png_b64":          b64(pack.BarsPNG),
		"consciousness_jpg_b64": b64(pack.ConsciousnessJPG),
		"density_jpg_b64":       b64(pack.DensityJPG),
		"scatter_jpg_b64":       b64(pack.ScatterJPG),
		"bars_jpg_b64":          b64(pack.BarsJPG),
	}
}

func b64(b []byte) string {
	if len(b) == 0 {
		return ""
	}
	return base64.StdEncoding.EncodeToString(b)
}
