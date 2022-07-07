package main

import (
	"fmt"
	"sort"
)

// map[key]value
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus": {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures": {"discrete math"},
	"databases": {"data structures"},
	"discrete math": {"intro to programming"},
	"formal languages": {"discrete math"},
	"networks": {"operating systems"},
	"operating systems": {
		"data structures",
		"computer organization",
	},
	"programming languages": {
		"data structures",
		"computer organization",
	},
}

func topSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	// 函数内递归定义匿名函数的名字
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}
	var keys []string
	for key := range m{
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}

// 5.10 []string --> map[int]string
func topSortMap(m map[string][]string) map[int]string {
	var cnt int = 0
	// 复合类型需要用make创建并分配空间?
	order := make(map[int]string)
	seen := make(map[string]bool)
	// 函数内递归定义匿名函数的名字
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				cnt++
				order[cnt] = item
			}
		}
	}
	var keys []string
	for key := range m{
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}

//func main() {
//	for i, course := range topSort(prereqs) {
//		fmt.Printf("%d:\t%s\n", i, course)
//	}
//}

func main() {
	for key, value := range topSortMap(prereqs) {
		fmt.Printf("%d:\t%s\n", key, value)
	}
}