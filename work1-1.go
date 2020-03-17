package main

import (
	"fmt"
	"os"
	"strings"
)

// run command:
// go run work1-1.go hello world
func main() {
	fmt.Println(strings.Join(os.Args, " "))
}
