package main

import (
	"fmt"
	"reflect"
)

// run command
// go run ./src/chapter3/work3_12/work3_12.go
func main() {
	fmt.Println(containCharsEqual("haha", "ahaq"))
}

func containCharsEqual(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	m1 := make(map[rune]int)
	m2 := make(map[rune]int)
	for _, r := range s1 {
		m1[r]++
	}
	for _, r := range s2 {
		m2[r]++
	}
	return reflect.DeepEqual(m1, m2)
}
