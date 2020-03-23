package main

import (
	"fmt"

	"github.com/caistrong/gopl-homework/src/chapter2/work2_3/popcount"
)

// go command:
// go run ./src/chapter2/work2_5/work2_5.go
func main() {
	x := uint64(100232132223123130)
	fmt.Printf("count: %d\n", popcount.PopCount(x))
	fmt.Printf("count: %d\n", popcount.CPopCount(x))
}
