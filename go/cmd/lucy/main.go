// Command lucy is the native measuring CLI (stdin JSON → stdout JSON board).
//
//	echo '{"samples":[...]}' | lucy build-lpd
//	lucy version
package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/openfluke/lucy/lucy"
)

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(2)
	}
	switch os.Args[1] {
	case "version", "-version", "--version":
		fmt.Println(lucy.Version)
	case "chart-radar", "chart-scatter", "chart-bars",
		"chart-radar-png", "chart-scatter-png", "chart-bars-png", "chart-pack":
		raw, err := io.ReadAll(os.Stdin)
		if err != nil {
			fail(err)
		}
		resp, err := lucy.BuildFromJSON(raw)
		if err != nil {
			fail(err)
		}
		switch os.Args[1] {
		case "chart-radar":
			fmt.Print(lucy.RadarSVG("Consciousness radar", lucy.ConsciousnessSeries(resp.Board, 8)))
		case "chart-scatter":
			fmt.Print(lucy.ScatterSVG("Q% vs RAM", "RAM KiB", "Q %", lucy.LPDScatterPoints(resp.Board)))
		case "chart-bars":
			fmt.Print(lucy.BarsSVG("Top LPD", resp.Board, 12))
		case "chart-radar-png":
			b, err := lucy.RadarPNG("Consciousness radar", lucy.ConsciousnessSeries(resp.Board, 8))
			if err != nil {
				fail(err)
			}
			os.Stdout.Write(b)
		case "chart-scatter-png":
			b, err := lucy.ScatterPNG("Q% vs RAM", "RAM KiB", "Q %", lucy.LPDScatterPoints(resp.Board))
			if err != nil {
				fail(err)
			}
			os.Stdout.Write(b)
		case "chart-bars-png":
			b, err := lucy.BarsPNG("Top LPD", resp.Board, 12)
			if err != nil {
				fail(err)
			}
			os.Stdout.Write(b)
		case "chart-pack":
			pack := lucy.BuildBoardCharts(resp.Board, 8)
			enc := json.NewEncoder(os.Stdout)
			enc.SetIndent("", "  ")
			_ = enc.Encode(map[string]any{
				"version":            lucy.Version,
				"consciousness_svg":  pack.ConsciousnessSVG,
				"density_svg":        pack.DensitySVG,
				"scatter_svg":        pack.ScatterSVG,
				"bars_svg":           pack.BarsSVG,
				"consciousness_png_b64": encodeB64(pack.ConsciousnessPNG),
				"density_png_b64":       encodeB64(pack.DensityPNG),
				"scatter_png_b64":       encodeB64(pack.ScatterPNG),
				"bars_png_b64":          encodeB64(pack.BarsPNG),
			})
		}
	case "build-lpd", "lpd":
		raw, err := io.ReadAll(os.Stdin)
		if err != nil {
			fail(err)
		}
		resp, err := lucy.BuildFromJSON(raw)
		if err != nil {
			fail(err)
		}
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		if err := enc.Encode(resp); err != nil {
			fail(err)
		}
	case "help", "-h", "--help":
		usage()
	default:
		fmt.Fprintf(os.Stderr, "unknown command %q\n", os.Args[1])
		usage()
		os.Exit(2)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, `lucy %s — live-fit measuring CLI

Usage:
  lucy version
  lucy build-lpd < request.json   # stdin BuildRequest → stdout BuildResponse
  lucy chart-radar|chart-scatter|chart-bars < request.json       # SVG
  lucy chart-radar-png|chart-scatter-png|chart-bars-png < req.json # PNG
  lucy chart-pack < request.json  # JSON with svg + png_b64 fields

BuildRequest:
  {"samples":[{"id":"...","acc":90,"thru":200,"avail":40,"score":100,"ram_kib":1000}],
   "options":{"keep_floor":0.7}}
`, lucy.Version)
}

func fail(err error) {
	fmt.Fprintf(os.Stderr, "lucy: %v\n", err)
	os.Exit(1)
}

func encodeB64(b []byte) string {
	if len(b) == 0 {
		return ""
	}
	return base64.StdEncoding.EncodeToString(b)
}
