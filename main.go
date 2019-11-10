// +build !js

package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	code, err := ioutil.ReadFile(os.Args[1])
	if err == nil {
		res := parseFile(string(code))
		fmt.Printf("%s\n", res)
	}
}
