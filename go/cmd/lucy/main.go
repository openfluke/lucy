// Command lucy is the native measuring CLI (stdin JSON → stdout JSON board).
//
//	echo '{"samples":[...]}' | lucy build-lpd
//	lucy version
package main

import (
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
	case "chart-radar", "chart-scatter", "chart-bars":
		raw, err := io.ReadAll(os.Stdin)
		if err != nil {
			fail(err)
		}
		resp, err := lucy.BuildFromJSON(raw)
		if err != nil {
			fail(err)
		}
		var svg string
		switch os.Args[1] {
		case "chart-radar":
			svg = lucy.RadarSVG("Consciousness radar", lucy.ConsciousnessSeries(resp.Board, 8))
		case "chart-scatter":
			svg = lucy.ScatterSVG("Q% vs RAM", "RAM KiB", "Q %", lucy.LPDScatterPoints(resp.Board))
		case "chart-bars":
			svg = lucy.BarsSVG("Top LPD", resp.Board, 12)
		}
		fmt.Print(svg)
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
  lucy chart-radar|chart-scatter|chart-bars < request.json  # SVG on stdout

BuildRequest:
  {"samples":[{"id":"...","acc":90,"thru":200,"avail":40,"score":100,"ram_kib":1000}],
   "options":{"keep_floor":0.7}}
`, lucy.Version)
}

func fail(err error) {
	fmt.Fprintf(os.Stderr, "lucy: %v\n", err)
	os.Exit(1)
}
