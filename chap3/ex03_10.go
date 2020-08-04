package main

import (
	"bytes"
	"fmt"
)

func comma2(s string) string {
	// comma inserts commas in a non-negative decimal integer string
	var buf bytes.Buffer
	n := len(s)  // can use length because it is certain that s only contains digits
	for idx, char := range s {
		if idx != 0 && idx % 3 == n % 3 {
			buf.WriteByte(',')
		}
		buf.WriteRune(char)
	}
	return buf.String()
}

func main() {
	strs := []string{"123", "123456", "1234", "12345"}
	for _, s := range strs {
		fmt.Println(s, comma2(s))
	}
}
