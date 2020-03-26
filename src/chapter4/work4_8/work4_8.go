// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

// run command:
// go build -o ./pkg/chapter4/work4_8/work4_8 ./src/chapter4/work4_8/work4_8.go
// less ./src/chapter4/work4_8/test | ./pkg/chapter4/work4_8/work4_8
func main() {
	counts := make(map[string]int)
	invalid := 0 // count of invalid UTF-8 characters
	total := 0
	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		total++
		if unicode.IsControl(r) {
			counts["Control"]++
		}
		if unicode.IsDigit(r) {
			counts["Digit"]++
		}
		if unicode.IsGraphic(r) {
			counts["Graphic"]++
		}
		if unicode.IsLetter(r) {
			counts["Letter"]++
		}
		if unicode.IsLower(r) {
			counts["Lower"]++
		}
		if unicode.IsMark(r) {
			counts["Mark"]++
		}
		if unicode.IsNumber(r) {
			counts["Number"]++
		}
		if unicode.IsPrint(r) {
			counts["Print"]++
		}
		if unicode.IsPunct(r) {
			counts["Punct"]++
		}
		if unicode.IsSpace(r) {
			counts["Space"]++
		}
		if unicode.IsSymbol(r) {
			counts["Symbol"]++
		}
		if unicode.IsTitle(r) {
			counts["Title"]++
		}
		if unicode.IsUpper(r) {
			counts["Upper"]++
		}
	}
	fmt.Printf("Total Rune Count: %d\n", total)

	fmt.Printf("Rune Type\tCount\n")
	for c, n := range counts {
		fmt.Printf("%s\t          %d\n", c, n)
	}

	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
