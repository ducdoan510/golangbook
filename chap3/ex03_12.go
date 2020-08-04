package main

import "fmt"

func countApp(s string) map[rune]int {
	ans := make(map[rune]int)
	characters := []rune(s)
	for _, ch := range characters {
		ans[ch] += 1
	}
	return ans
}

func IsAnagram(s1 string, s2 string) bool {
	app1 := countApp(s1)
	app2 := countApp(s2)
	for ch, cnt := range app1 {
		if app2[ch] != cnt {
			return false
		}
	}
	for ch, cnt := range app2 {
		if app1[ch] != cnt {
			return false
		}
	}
	return true
}

func main() {
	testcases := map[string]string {
		"abc": "cba",
		"abcsd": "acewb",
		"abcd": "abc",
		"aaa": "aaa",
	}
	for s1, s2 := range testcases {
		fmt.Printf("%6s %6s\t%v\n", s1, s2, IsAnagram(s1, s2))
	}
}
