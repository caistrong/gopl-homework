package main

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

// run command
// go run ./src/chapter3/work3_10/work3_10.go
func main() {
	// testString := "A2C蔡x聪1二C45六"
	testString := "15225"
	fmt.Println(comma(testString))
}

func comma(s string) string {
	var buf bytes.Buffer
	runeCnt := utf8.RuneCountInString(s) // 先看有几个字符
	sep := runeCnt % 3
	cnt := 0
	for _, r := range s {
		// r is rune
		if cnt%3 == sep && cnt != 0 { // 每个逗号的位置为 sep + 3n
			buf.WriteString(",")
		}
		buf.WriteRune(r)
		cnt++
	}
	return buf.String()
}
