package main

import (
	"fmt"
	"strings"
)

func join(sep string, strs ...string) string {
	return strings.Join(strs, sep)
}

func main() {
	fmt.Println(join(",", "abc", "def", "ghi"))
}
