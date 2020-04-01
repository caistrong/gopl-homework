package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

// run command:
// go build -o ./pkg/chapter5/work5_1/work5_1 ./src/chapter5/work5_1/work5_1.go

// 【optional】go build -o ./pkg/chapter1/work1_7/work1_7 ./src/chapter1/work1_7/work1_7.go
// ./pkg/chapter1/work1_7/work1_7 https://golang.org/ | ./pkg/chapter5/work5_1/work5_1

// 【or】less ./src/chapter5/work5_1/golang.org.html | ./pkg/chapter5/work5_1/work5_1
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visitRecursive(nil, doc) {
		fmt.Println(link)
	}
}

// 循环版本（书上原版）
// func visit(links []string, n *html.Node) []string {
// 	if n.Type == html.ElementNode && n.Data == "a" {
// 		for _, a := range n.Attr {
// 			if a.Key == "href" {
// 				links = append(links, a.Val)
// 			}
// 		}
// 	}
// 	for c := n.FirstChild; c != nil; c = c.NextSibling {
// 		links = visit(links, c)
// 	}
// 	return links
// }

// 递归版本
// NOTE: 禁止套娃！（大多数情况下，代码的可读性才是第一目标，对于这里，很明显循环的版本要比下面这个递归版本要好理解得多
func visitRecursive(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	return visitRecursion(links, n.FirstChild)

}

func visitRecursion(links []string, c *html.Node) []string {
	if c == nil {
		return links
	}
	links = visitRecursive(links, c)
	c = c.NextSibling
	return visitRecursion(links, c)
}

// NOTE: https://softwareengineering.stackexchange.com/questions/279004/general-way-to-convert-a-loop-while-for-to-recursion-or-from-a-recursion-to-a
// 循环和递归的互转参考了上述答主的优秀指示，不然我根本写不出来这个套娃版本的func
