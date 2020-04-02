// prereqs记录了每个课程的前置课程
package main

import (
	"fmt"
	"sort"
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
	// "linear algebra":        {"calculus"},
}

// run command
// go run ./src/chapter5/work5_11/work5_11.go
func main() {
	for key, val := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", key, val)
	}
}

func topoSort(m map[string][]string) map[int]string {
	// 有向图找环
	var findCycle func(m map[string][]string, v string, seen map[string]bool, deep int, trace map[int]string) bool
	findCycle = func(m map[string][]string, v string, seen map[string]bool, deep int, trace map[int]string) bool {
		seen[v] = true            //标记该课程已经访问
		deep++                    //deep是递归的深度，也记录整个课程的顺序
		trace[deep] = v           //trace记录整个课程顺序与课程名
		for _, it := range m[v] { //对该课程（v）的每一个邻居（前置课程）进行遍历
			if !seen[it] {
				if findCycle(m, it, seen, deep, trace) { //深度优先递归遍历
					return true
				}
			} else {
				if seen[it] { //已经访问的课程说明存在环
					return true
				}
			}
		}
		//如果没有环，那么以上循环会以递归树执行，而且不会return，所以会返回下面的终点
		seen[v] = false
		return false
	}
	for k := range m {
		seen := make(map[string]bool)
		trace := make(map[int]string)
		i := 0
		fmt.Printf("Course: %s\n", k)
		fmt.Printf("Contains Cycle: %t\n", findCycle(m, k, seen, i, trace))
		var sortedkeys []int
		for k := range trace {
			sortedkeys = append(sortedkeys, k)
		}
		sort.Ints(sortedkeys)
		for _, k := range sortedkeys {
			if k == len(sortedkeys) {
				fmt.Printf("%d/%d %s", k, len(sortedkeys), trace[k])

			} else {
				fmt.Printf("%d/%d %s -> ", k, len(sortedkeys), trace[k])
			}
		}
		fmt.Printf("\n************************************\n")
	}

	// 完整路径
	var order = make(map[int]string)
	marked := make(map[string]bool)
	var visitAll func(items []string)
	index := 1
	visitAll = func(items []string) {
		for _, item := range items {
			if !marked[item] {
				marked[item] = true
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

// refer: https://blog.csdn.net/mmmm433/article/details/103496429
