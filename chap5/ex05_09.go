package main

import (
	"fmt"
	"strings"
)

func expand(s string, f func(string) string) string {
	replacedBy := f("foo")
	return strings.Replace(s, "$foo", replacedBy, -1)
}

func setEmpty(s string) string {
	return ""
}

func duplicate(s string) string {
	return s + s
}

func main() {
	testString := "$foo hello world $foo"
	fmt.Println(expand(testString, setEmpty))
	fmt.Println(expand(testString, duplicate))
	fmt.Println(expand(testString, strings.ToUpper))
}
