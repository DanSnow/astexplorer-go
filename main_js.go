package main

import "syscall/js"

func main() {
	block := make(chan struct{}, 0)
	js.Global().Set("__GO_PARSE_FILE__", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		code := args[0].String()
		res := parseFile(code)
		return string(res)
	}))
	<-block
}
