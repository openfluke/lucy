//go:build js && wasm

package main

import (
	"encoding/json"
	"syscall/js"

	"github.com/openfluke/lucy/lucy"
)

func main() {
	js.Global().Set("lucyVersion", js.FuncOf(func(this js.Value, args []js.Value) any {
		return lucy.Version
	}))
	js.Global().Set("lucyBuildLPD", js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) < 1 {
			return errObj("lucyBuildLPD requires a JSON string")
		}
		raw := args[0].String()
		resp, err := lucy.BuildFromJSON([]byte(raw))
		if err != nil {
			return errObj(err.Error())
		}
		b, err := json.Marshal(resp)
		if err != nil {
			return errObj(err.Error())
		}
		return string(b)
	}))
	select {}
}

func errObj(msg string) string {
	b, _ := json.Marshal(map[string]string{"error": msg})
	return string(b)
}
