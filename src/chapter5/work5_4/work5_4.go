package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

var linkMap = map[string]string{
	"a":      "href",
	"img":    "src",
	"script": "src",
	"link":   "href",
}

// run command:
// go build -o ./pkg/chapter5/work5_4/work5_4 ./src/chapter5/work5_4/work5_4.go
// ./pkg/chapter1/work1_7/work1_7 https://golang.org/ | ./pkg/chapter5/work5_4/work5_4

// 【or】less ./src/chapter5/work5_1/golang.org.html | ./pkg/chapter5/work5_1/work5_1
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && linkMap[n.Data] != "" {
		for _, a := range n.Attr {
			if a.Key == linkMap[n.Data] {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
