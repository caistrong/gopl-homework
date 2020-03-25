package main

import (
	"fmt"
)

// run command
// go run ./src/chapter4/work4_5/work4_5.go
func main() {
	s := []string{"c", "a", "a", "a", "i"}
	fmt.Println(uniqueSlice(s))
}

// input: ["c", "a", "a", "a","i"]
// outpu: ["c", "a", "i"]

func uniqueSlice(strSlice []string) []string {
	tempStr := ""
	for i := 0; i < len(strSlice); i++ {
		if tempStr == strSlice[i] {
			strSlice = append(strSlice[:i], strSlice[i+1:]...)
			i-- // we delete a element in slice, so we should make i minus 1
		}
		tempStr = strSlice[i]
	}
	return strSlice
}
