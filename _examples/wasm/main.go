//go:build js && wasm
// +build js,wasm

package main

import (
	"syscall/js"
)

func igOK(s string, _ bool) string { _ = "STUB: not implemented"; return "" }

func tokenize(_ js.Value, args []js.Value) interface{} { _ = "STUB: not implemented"; return nil }

//fmt.Printf("%s\t%+v%v\n", v.Surface, v.POS(), strings.Join(v.Features(), ","))

func registerCallbacks() { _ = "STUB: not implemented"; return }

func main() {
	c := make(chan struct{}, 0)
	registerCallbacks()
	println("Kagome Web Assembly Ready")
	<-c
}
