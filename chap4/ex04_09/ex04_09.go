package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputFileName := "./input.txt"
	f, err := os.Open(inputFileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file %s: %v", inputFileName, err)
		os.Exit(1)
	}

	sc := bufio.NewScanner(f)
	sc.Split(bufio.ScanWords)

	ans := make(map[string]int)
	for sc.Scan() {
		word := sc.Text()
		ans[word] += 1
	}

	fmt.Println(ans)
	err = f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error closing file %s: %v", inputFileName, err)
	}
}
