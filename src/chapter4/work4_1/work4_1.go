package main

import (
	"fmt"
)

// MessageDigest is hash code generate by sha256
type MessageDigest [32]byte

// run command
// go run ./src/chapter4/work4_1/work4_1.go
func main() {
	md1, md2 := MessageDigest{14: 12, 31: 1}, MessageDigest{31: 1}
	bitDiff(&md1, &md2)
}

// XOR 按位异或
// a	b	a^b
// 0	0	0
// 0	1	1
// 1	0	1
// 1	1	0
func bitDiff(md1, md2 *MessageDigest) int {
	diffCnt := 0
	fmt.Println("  sha1  \t  sha2 ")
	for i, b1 := range md1 {
		b2 := md2[i]
		for i := 0; i < 8; i++ {
			lb1, lb2 := (b1>>i)&1, (b2>>i)&1 // get last bit of a byte
			if (lb1 ^ lb2) == 1 {
				diffCnt++
			}
		}
		fmt.Printf("%08b\t%08b\n", b1, b2)
	}
	fmt.Printf("bit diff count: %d\n", diffCnt)
	return diffCnt
}
