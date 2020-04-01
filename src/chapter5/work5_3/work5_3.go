package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

// run command:
// go build -o ./pkg/chapter5/work5_3/work5_3 ./src/chapter5/work5_3/work5_3.go
// ./pkg/chapter1/work1_7/work1_7 https://golang.org/ | ./pkg/chapter5/work5_3/work5_3

// 【or】less ./src/chapter5/work5_1/golang.org.html | ./pkg/chapter5/work5_3/work5_3
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	getTextContent(doc)
}

// 循环版本（书上原版）
func getTextContent(n *html.Node) {
	if n.Type == html.ElementNode {
		if n.Data == "script" || n.Data == "style" {
			return
		}
	}
	if n.Type == html.TextNode {
		fmt.Println(n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		getTextContent(c)
	}
}
