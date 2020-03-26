package main

import (
	"bufio"
	"fmt"
	"os"
)

// run command:
// go build -o ./pkg/chapter4/work4_9/work4_9 ./src/chapter4/work4_9/work4_9.go
// less ./src/chapter4/work4_9/test | ./pkg/chapter4/work4_9/work4_9
func main() {
	counts := make(map[string]int) // a set of strings
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords) // functional programming
	total := 0
	for input.Scan() {
		word := input.Text()
		counts[word]++
		total++
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "work4_9: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Total Work Count: ", total)
	fmt.Println("word\tcount\tfrequency")
	for word, count := range counts {
		fmt.Printf("%s\t%d\t%.2f\n", word, count, float64(count)/float64(total))
	}
}
