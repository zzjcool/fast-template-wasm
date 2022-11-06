package main

import (
	"syscall/js"

	"github.com/zzjcool/fast-template/render"
)

func renderf(this js.Value, args []js.Value) any {
	result, err := render.Str(args[0].String(), args[1].String())
	if err != nil {
		panic(err)
	}
	return result
}
func main() {
	c := make(chan struct{}, 0)
	// register functions
	js.Global().Set("render", js.FuncOf(renderf))
	<-c
}
