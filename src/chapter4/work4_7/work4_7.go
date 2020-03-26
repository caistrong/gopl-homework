package main

import (
	"fmt"
)

// run command
// go run ./src/chapter4/work4_7/work4_7.go
func main() {
	s := []byte("Hello 世界")

	fmt.Println(string(reverse(s)))
}

func reverse(bs []byte) []byte {
	runes := []rune(string(bs))
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return []byte(string(runes))
}
