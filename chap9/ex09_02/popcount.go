package main

import (
	"fmt"
	"sync"
)

var computeCountOnce sync.Once
var byteCount [256]int

func computeCount() {
	for i := 0; i < 256; i++ {
		byteCount[i] = byteCount[i >> 1] + (i & 1)
	}
}

func popCount(x uint64) int {
	computeCountOnce.Do(computeCount)
	ans := 0
	for i := 0; i < 8; i++ {
		ans += byteCount[x >> (8 * i)]
	}
	return ans
}

func main() {
	go fmt.Println(popCount(16))
	fmt.Println(popCount(7))
}
