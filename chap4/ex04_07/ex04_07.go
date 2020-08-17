package main

import (
	"fmt"
	"unicode/utf8"
)

func shiftAndSwap(sl []byte, prefix int, suffix int) {
	n := len(sl)
	first := make([]byte, prefix) // byte(s) for first char
	last := make([]byte, suffix)  // byte(s) for last char
	copy(first, sl[:prefix])
	copy(last, sl[n - suffix:])

	if prefix > suffix {
		for i := suffix; i < n - prefix; i++ { sl[i] = sl[i + prefix - suffix] }
	} else {
		for i := n - prefix - 1; i >= suffix; i-- { sl[i] = sl[i - (suffix - prefix)] }
 	}
	copy(sl[:suffix], last)
	copy(sl[n - prefix:], first)
}

func reverseInPlace(bsl []byte) {
	for i, j := 0, len(bsl); i < j; {
		_, firstSize := utf8.DecodeRune(bsl[i:])
		_, lastSize := utf8.DecodeLastRune(bsl[:j])
		if i + firstSize == j { break }
		shiftAndSwap(bsl[i:j], firstSize, lastSize)
		i += lastSize
		j -= firstSize
	}
}

func main() {
	// Test 1:
	s := []byte{72, 101, 108, 108, 111, 44, 32, 228, 184, 150, 231, 149, 140} // "Hello, 世界"
	fmt.Println(string(s))
	reverseInPlace(s)
	fmt.Println(string(s))

	// Test 2:
	s1 := []byte{231, 149, 140, 72, 101, 108, 108, 111, 44, 32, 228, 184, 150} // "界Hello, 世"
	fmt.Println(string(s1))
	reverseInPlace(s1)
	fmt.Println(string(s1))
}
