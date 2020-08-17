package main

import (
	"fmt"
	"os"
)

var prereqs = map[string][]string {
	"algorithms": {"data structures"},
	"calculus": {"linear algebra"},
	"compilers": {"data structures", "formal languages", "computer organization"},
	"data structures": {"discrete math"},
	"databases": {"data structures"},
	"discrete math": {"intro to programming"},
	"formal languages": {"discrete math"},
	"networks": {"operating systems"},
	"operating systems": {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
	"linear algebra": {"calculus"},
}

func topoSort2(m map[string][]string) (order []string, err error) {
	group := make(map[string]int)
	var visit func(string, int)
	visit = func (item string, gr int) {
		for _, prereq := range m[item] {
			if _, ok := group[prereq]; !ok {
				group[prereq] = gr
				visit(prereq, gr)
			} else if gr == group[prereq] {
				err = fmt.Errorf("cycle is detected in map")
			}
		}
		order = append(order, item)
	}

	groupIndex := 0
	for item := range m {
		if _, ok := group[item]; !ok {
			group[item] = groupIndex
			visit(item, groupIndex)
			groupIndex++
		}
	}
	return
}

func main() {
	subjects, err := topoSort2(prereqs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error topoSort: %v", err)
		os.Exit(1)
	}
	for idx, subject := range subjects {
		fmt.Printf("%3d\t%s\n", idx, subject)
	}
}
