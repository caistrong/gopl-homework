package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var s string
	for idx, arg := range os.Args[1:] {
		s += strconv.Itoa(idx) + "\t" + arg + "\n"
	}
	fmt.Println(s)
}
