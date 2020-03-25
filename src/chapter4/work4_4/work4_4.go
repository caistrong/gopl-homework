package main

import (
	"fmt"
)

// run command
// go run ./src/chapter4/work4_4/work4_4.go
func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(rotate(s, 2))
}

// input: [0, 1, 2, 3, 4, 5], 2
// outpu: [2, 3, 4, 5, 0, 1]

func rotate(s []int, rotateTimes int) []int {
	var result []int
	for i := 0; i < rotateTimes; i++ {
		result = s[1:len(s)]
		result = append(result, s[0])
		s = result
	}
	return result
}
