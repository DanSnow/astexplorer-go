//go:build !js

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var code []byte
	var err error
	if len(os.Args) < 2 || os.Args[1] == "-" {
		code, err = io.ReadAll(os.Stdin)
	} else {
		code, err = os.ReadFile(os.Args[1])
	}
	if err == nil {
		res := parseFile(string(code))
		fmt.Printf("%s\n", res)
	}
}
