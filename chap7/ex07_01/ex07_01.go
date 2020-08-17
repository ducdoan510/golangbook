package main

import (
	"bufio"
	"fmt"
	"strings"
)

type WordCounter int
type LineCounter int

func scan(p []byte, splitFunc bufio.SplitFunc) (n int) {
	sc := bufio.NewScanner(strings.NewReader(string(p)))
	sc.Split(splitFunc)
	for sc.Scan() {
		n++
	}
	return
}

func (wc *WordCounter) Write(p []byte) (int, error) {
	ans := scan(p, bufio.ScanWords)
	*wc = WordCounter(ans)
	return ans, nil
}

func (lc *LineCounter) Write(p []byte) (int, error) {
	ans := scan(p, bufio.ScanLines)
	*lc = LineCounter(ans)
	return ans, nil
}

func main() {
	input1 := " word1 word2 word3 word4 "
	var wc WordCounter
	_, _ = wc.Write([]byte(input1))
	fmt.Println(wc)

	input2 := "word1\n word2 \nword3"
	var lc LineCounter
	_, _ = lc.Write([]byte(input2))
	fmt.Println(lc)
}
