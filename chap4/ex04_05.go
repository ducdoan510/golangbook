package main

import "fmt"

func removeAdjDuplicate(sl []string) []string {
	i := 0
	for slIdx := 0; slIdx < len(sl); {
		idx := slIdx
		sl[i] = sl[slIdx]
		for idx < len(sl) && sl[idx] == sl[slIdx] { idx++ }
		slIdx = idx
		i++
	}
	return sl[:i]
}

func main() {
	sl := []string{"abc", "abc", "d", "abc", "d", "d", "e", "f", "f"}
	fmt.Println(removeAdjDuplicate(sl))
}