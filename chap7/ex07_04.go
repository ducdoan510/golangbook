package main

import (
	"fmt"
)

type StringReader struct {
	str string
}

func (sr *StringReader) Read(p []byte) (n int, err error) {
	inputByte := []byte(sr.str)
	if len(inputByte) < len(p) {
		n = len(inputByte)
	} else {
		n = len(p)
	}
	for i := range p {
		if i < len(inputByte) {
			p[i] = inputByte[i]
		}
	}
	return
}

func main() {
	p := make([]byte, 2)
	r := StringReader{"abc"}
	n, err := r.Read(p)
	fmt.Println(n, err, p, string(p))
}
