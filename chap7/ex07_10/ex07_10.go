package main

import (
	"fmt"
	"sort"
)

func IsPalindrome(s sort.Interface) bool {
	for i, j := 0, s.Len() - 1; i < j; i, j = i + 1, j - 1 {
		if s.Less(i, j) || s.Less(j, i) {
			return false
		}
	}
	return true
}

func main() {
	stringSlice := sort.StringSlice{"abc", "def", "ghi", "ghi", "def", "abc"}
	fmt.Println(IsPalindrome(stringSlice))
	stringSlice = sort.StringSlice{"abc", "def", "ghi"}
	fmt.Println(IsPalindrome(stringSlice))
}
