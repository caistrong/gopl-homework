package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
)

// run command
// go run ./src/chapter3/work3_11/work3_11.go
func main() {
	// testString := "A2C蔡x聪1二C45六"
	testString := "-152215.232"
	fmt.Println(comma(testString))
}

func comma(s string) string {
	symbol := s[0]
	s = s[1:]
	dotIndex := strings.LastIndex(s, ".")
	decimals := s[dotIndex:] // 小数部分
	s = s[:dotIndex]
	var buf bytes.Buffer
	buf.WriteByte(symbol)
	runeCnt := utf8.RuneCountInString(s) // 先看有几个字符
	sep := runeCnt % 3                   // 距字符数计算余数应为几才是逗号位置
	cnt := 0
	for _, r := range s {
		// r is rune
		if cnt%3 == sep && cnt != 0 { // 每个逗号的位置为 sep + 3n
			buf.WriteString(",")
		}
		buf.WriteRune(r)
		cnt++
	}
	buf.WriteString(decimals)
	return buf.String()
}
