package main

import (
	"fmt"
	"io"
	"strings"
)

type LimitedReader struct {
	innerReader io.Reader
	byteLimit int64
}

func (lr *LimitedReader) Read(p []byte) (int, error) {
	r, limit := lr.innerReader, lr.byteLimit
	n, err := r.Read(p[:limit])
	if int64(n) < limit || err != nil {
		return n, err
	}
	return int(limit), io.EOF
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &LimitedReader{r, n}
}

func main() {
	limit := int64(6)
	reader := strings.NewReader("golang")
	limitedReader := LimitReader(reader, limit)
	p := make([]byte, limit)
	n, err := limitedReader.Read(p)
	fmt.Println(n, err, string(p))
}
