package main

import (
	"fmt"
	"unicode"
)

// run command
// go run ./src/chapter4/work4_6/work4_6.go
func main() {
	rs := []rune{'H', 'e', 'l', 'l', 'o', ' ', ' ', ' ', '世', '界'}
	fmt.Println("input string:\t", string(rs))
	bs := []byte(string(rs))
	fmt.Println("output string:\t", string(uniqueSpaceSlice(bs)))
}

func uniqueSpaceSlice(bs []byte) []byte {
	fmt.Println("input []bytes:\t", bs)
	flag := 0 // 连续出现space次数的flag
	for i := 0; i < len(bs); i++ {
		if unicode.IsSpace(rune(bs[i])) {
			flag++
			if flag > 1 {
				bs = append(bs[:i], bs[i+1:]...)
				i-- // we delete a element in slice, so we should make i minus 1
			}
		} else {
			flag = 0
		}
	}
	fmt.Println("output []bytes:\t", bs)
	return bs
}
