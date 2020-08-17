package main

import (
	"fmt"
	"golangbook/chap2/popcount"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		num, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Fprintln(os.Stderr, arg, " is not a number")
			continue
		}
		num64 := uint64(num)
		fmt.Println(num64)
		fmt.Println("Bit count using count table: ", popcount.PopCount(num64))
		fmt.Println("Bit count using shifting 64 bit: ", popcount.Count2(num64))
		fmt.Println("Bit count using clearing bit: ", popcount.Count3(num64))
	}
}