// Command lucy is the native measuring CLI (stdin JSON → stdout JSON board).
//
//	echo '{"samples":[...]}' | lucy build-lpd
//	lucy version
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
		"chart-radar-png", "chart-scatter-png", "chart-bars-png",
		"chart-radar-jpg", "chart-scatter-jpg", "chart-bars-jpg", "chart-pack":
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
		case "chart-radar-jpg":
			b, err := lucy.RadarJPG("Consciousness radar", lucy.ConsciousnessSeries(resp.Board, 8), 85)
			if err != nil {
				fail(err)
			}
			os.Stdout.Write(b)
		case "chart-scatter-jpg":
			b, err := lucy.ScatterJPG("Q% vs RAM", "RAM KiB", "Q %", lucy.LPDScatterPoints(resp.Board), 85)
			if err != nil {
				fail(err)
			}
			os.Stdout.Write(b)
		case "chart-bars-jpg":
			b, err := lucy.BarsJPG("Top LPD", resp.Board, 12, 85)
			if err != nil {
				fail(err)
			}
			os.Stdout.Write(b)
		case "chart-pack":
			pack := lucy.BuildBoardCharts(resp.Board, 8)
			enc := json.NewEncoder(os.Stdout)
			enc.SetIndent("", "  ")
			_ = enc.Encode(lucy.ChartPackJSON(pack))
		}
	case "report":
		if len(os.Args) < 3 {
			fail(fmt.Errorf("usage: lucy report <outdir> < request.json"))
		}
		raw, err := io.ReadAll(os.Stdin)
		if err != nil {
			fail(err)
		}
		resp, err := lucy.BuildFromJSON(raw)
		if err != nil {
			fail(err)
		}
		if err := lucy.WriteReportDir(os.Args[2], resp.Board, 8); err != nil {
			fail(err)
		}
		fmt.Println(os.Args[2])
	case "serve":
		addr := ":7474"
		if len(os.Args) >= 3 {
			addr = os.Args[2]
		}
		fmt.Fprintf(os.Stderr, "lucy serve %s on %s\n", lucy.Version, addr)
		if err := http.ListenAndServe(addr, lucy.Handler()); err != nil {
			fail(err)
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
  lucy build-lpd < request.json
  lucy chart-radar|chart-scatter|chart-bars < request.json
  lucy chart-*-png|chart-*-jpg < request.json
  lucy chart-pack < request.json
  lucy report <outdir> < request.json
  lucy serve [addr]                 # default :7474

BuildRequest:
  {"samples":[{"id":"...","avg_accuracy":90,"throughput":200,"availability":40,"score":100,"ram_kib":1000}],
   "options":{"keep_floor":0.7}}
`, lucy.Version)
}

func fail(err error) {
	fmt.Fprintf(os.Stderr, "lucy: %v\n", err)
	os.Exit(1)
}

