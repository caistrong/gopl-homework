package popcount

import (
	"fmt"
	"time"
)

var pc [256]byte

// init
// example
// when i = 1 (0000 0001) pc[1] = 1
// when i = 2 (0000 0010) pc[1] = pc[0] (0000 0001) + (0000 0010 & 0000 0001) = 1
// when i = 3 (0000 0011) pc[3] = pc[1] (0000 0001) + (0000 0011 & 0000 0001) = 2
func init() {
	for i := range pc {
		pc[i] = pc[i>>1] + byte(i&1) // byte(i&1) only check the last bit is 1
	}
}

// PopCount returns the population count (number of set bits) of x. Example: PopCount(3) = 3 , 3 is 0000 0011 (binary) the count of 1 is 2
func PopCount(x uint64) int {
	start := time.Now()
	count := int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])

	fmt.Printf("PopCount cost %s\n", time.Since(start).String())
	return count
}

// LPopCount is loop version of PopCount
func LPopCount(x uint64) int {
	start := time.Now()
	count := 0
	for i := 0; i < 8; i++ {
		count += int(pc[byte(x>>(i*8))])
	}
	fmt.Printf("LPopCount cost %s\n", time.Since(start).String())
	return count
}

// SPopCount is shift version of PopCount
func SPopCount(x uint64) int {
	start := time.Now()
	count := 0

	for i := x; i > 0; i = i >> 1 {
		count += int(i & 1)
	}

	fmt.Printf("SPopCount cost %s\n", time.Since(start).String())
	return count
}

// CPopCount is clear last non zero version of PopCount
func CPopCount(x uint64) int {
	start := time.Now()
	count := 0

	for i := x; i > 0; i = i & (i - 1) {
		count++
	}

	fmt.Printf("CPopCount cost %s\n", time.Since(start).String())
	return count
}
