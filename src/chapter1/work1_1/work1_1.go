package main

import (
	"fmt"
	"os"
	"strings"
)

// run command:
// go run work1_1.go hello world
func main() {
	fmt.Println(strings.Join(os.Args, " "))
}
