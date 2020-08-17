package main

import (
	"crypto/sha256"
	"fmt"
)

func countDiff(b1 byte, b2 byte) int {
	ans := 0
	for i := 0; i < 8; i++ {
		ans += int(((b1 >> i) & 1) ^ ((b2 >> i) & 1))
	}
	return ans
}

func main() {
	s1 := []string{"x", "abc"}
	s2 := []string{"X", "abc"}
	for i := range s1 {
		diff := 0
		hash1 := sha256.Sum256([]byte(s1[i]))
		hash2 := sha256.Sum256([]byte(s2[i]))
		for i := range hash1 {
			diff += countDiff(hash1[i], hash2[i])
		}
		fmt.Printf("Number of different bits between hashes of %4s and %4s is %d\n", s1[i], s2[i], diff)
	}
}
