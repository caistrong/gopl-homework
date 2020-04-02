// prereqs记录了每个课程的前置课程
package main

import (
	"fmt"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

// run command
// go run ./src/chapter5/work5_10/work5_10.go
func main() {
	for key, val := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", key, val)
	}
}

func topoSort(m map[string][]string) map[int]string {
	var order = make(map[int]string)
	seen := make(map[string]bool)
	var visitAll func(items []string)
	index := 1
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order[index] = item
				index++
			}
		}
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	visitAll(keys)
	return order
}
