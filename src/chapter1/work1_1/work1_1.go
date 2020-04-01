package main

import (
	"fmt"
	"os"
	"strings"
)

// run command:
// go run ./src/chapter1/work1_1/work1_1.go hello world
func main() {
	fmt.Println(strings.Join(os.Args, " "))
}
