package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

// run command:
// go build -o ./pkg/chapter5/work5_2/work5_2 ./src/chapter5/work5_2/work5_2.go
// ./pkg/chapter1/work1_7/work1_7 https://golang.org/ | ./pkg/chapter5/work5_2/work5_2

// 【or】less ./src/chapter5/work5_1/golang.org.html | ./pkg/chapter5/work5_2/work5_2
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	counts := map[string]int{}
	for tag, count := range count(counts, doc) {
		fmt.Printf("<%s>\t%d\n", tag, count)
	}
}

// 循环版本（书上原版）
func count(counts map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		counts[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		count(counts, c)
	}
	return counts
}
