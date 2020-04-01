package main

import (
	"fmt"
	"os"
	"strings"
)

// run command
// go run ./src/chapter5/work5_9/work5_9.go hello,foo
func main() {
	for _, input := range os.Args[1:] {
		fmt.Println("expand: ", expand(input, allAdd1))
	}
}

func expand(s string, f func(string) string) string {
	ret := strings.Replace(s, "foo", f("foo"), -1)
	return ret
}

func add1(r rune) rune {
	return r + 1
}

func allAdd1(s string) string {
	return strings.Map(add1, s)
}
