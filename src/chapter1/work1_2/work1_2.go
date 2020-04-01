package main

import (
	"fmt"
	"os"
	"strconv"
)

// run command:
// go run ./src/chapter1/work1_2/work1_2.go hello world
func main() {
	var s string
	for idx, arg := range os.Args[1:] {
		s += strconv.Itoa(idx) + "\t" + arg + "\n"
	}
	fmt.Println(s)
}
