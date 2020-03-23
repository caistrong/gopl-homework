package main

import (
	"fmt"

	"github.com/caistrong/gopl-homework/src/chapter2/work2_3/popcount"
)

// go command:
// go run ./src/chapter2/work2_4/work2_4.go
func main() {
	x := uint64(1002321322130)
	fmt.Printf("count: %d\n", popcount.PopCount(x))
	fmt.Printf("count: %d\n", popcount.SPopCount(x))
}
