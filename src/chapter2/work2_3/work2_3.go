package main

import (
	"fmt"

	"github.com/caistrong/gopl-homework/src/chapter2/work2_3/popcount"
)

// go command:
// go run ./src/chapter2/work2_3/work2_3.go

func main() {
	x := uint64(189264826)
	fmt.Printf("count: %d\n", popcount.PopCount(x))
	fmt.Printf("count: %d\n", popcount.LPopCount(x))
}
