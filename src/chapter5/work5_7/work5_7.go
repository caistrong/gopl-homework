package main

import (
	"fmt"
	"net/http"
	"os"
	"regexp"

	"golang.org/x/net/html"
)

// run command
// go run ./src/chapter5/work5_7/work5_7.go https://golang.org/
func main() {
	for _, url := range os.Args[1:] {
		doc, err := getDoc(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "err: %v\n", err)
			os.Exit(1)
		}
		forEachNode(doc, startElement, endElement)
	}
}
func getDoc(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing HTML: %s", err)
	}
	return doc, nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		attrs := ""
		for _, a := range n.Attr {
			attrs += fmt.Sprintf("%s=\"%s\" ", a.Key, a.Val)
		}
		if n.FirstChild == nil {
			fmt.Printf("%*s<%s %s\\>\n", depth*2, "", n.Data, attrs)
		} else {
			fmt.Printf("%*s<%s %s>\n", depth*2, "", n.Data, attrs)
		}
		depth++
	}
	if n.Type == html.CommentNode {
		fmt.Printf("%*s//%s\n", depth*2, "", n.Data)
	}
	if nData := regexp.MustCompile(`\s+`).ReplaceAllString(n.Data, ""); n.Type == html.TextNode && nData != "" {
		fmt.Printf("%*s%s", depth*2, "", n.Data)
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		if n.FirstChild != nil {
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}
}
