package main

import (
	"bytes"
	"fmt"
	"strings"
)

func commaBackward(s string) string {
	// comma inserts commas in a non-negative decimal integer string\
	// same as comma2 in exercise 3.10
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

func commaForward(s string) string {
	var buf bytes.Buffer
	n := len(s)
	for idx, char := range s {
		buf.WriteRune(char)
		if idx != n - 1 && idx % 3 == 2 {
			buf.WriteByte(',')
		}
	}
	return buf.String()
}

func comma3(s string) string {
	// comma inserts commas in a float number string
	n := len(s)
	decimalPointIdx := strings.LastIndex(s, ".")
	if decimalPointIdx == -1 {
		decimalPointIdx = n
	}
	hasSign := s[0] == '-'
	var signPart, integerPart, decimalPart string
	if hasSign {
		signPart = "-"
		integerPart = s[1: decimalPointIdx]
	} else {
		integerPart = s[: decimalPointIdx]
	}

	ans := signPart + commaBackward(integerPart)
	if decimalPointIdx != n {
		decimalPart = s[decimalPointIdx+1:]
		ans += "." + commaForward(decimalPart)
	}
	return ans
}

func main() {
	strs := []string{"123", "123456", "1234", "12345", "-1234.3235", "-12343", "-12.21"}
	for _, s := range strs {
		fmt.Printf("%8s\t%s\n", s, comma3(s))
	}
}
