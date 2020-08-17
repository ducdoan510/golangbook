package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func removeSpaceInPlace(bs []byte) []byte {
	idx := 0
	for i := 0; i < len(bs); {
		r, size := utf8.DecodeRune(bs[i:])
		if !unicode.IsSpace(r) {
			for offset := 0; offset < size; offset++ { bs[idx + offset] = bs[i + offset] }
			idx += size
		}
		i += size
	}
	return bs[:idx]
}

func main() {
	s := []byte{72, 101, 108, 108, 111, 44, 32, 228, 184, 150, 32, 231, 149, 140}
	fmt.Println(string(s))
	fmt.Println(string(removeSpaceInPlace(s)))
}
