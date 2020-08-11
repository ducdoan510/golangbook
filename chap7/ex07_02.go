package main

import (
	"fmt"
	"io"
	"os"
)

type WriterWrapper struct {
	innerWriter io.Writer
	counterPtr *int64
}

func (ww *WriterWrapper) Write(p []byte) (int, error) {
	n, err := ww.innerWriter.Write(p)
	*ww.counterPtr += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var count int64
	newWriter := WriterWrapper{innerWriter: w, counterPtr: &count}
	return &newWriter, newWriter.counterPtr
}

func main() {
	writer, counter := CountingWriter(os.Stdout)
	input := []string{"abc", "def", "1234"}
	for _, s := range input {
		_, _ = fmt.Fprintf(writer, "%s\n", s)
		fmt.Println(*counter)
	}
}
