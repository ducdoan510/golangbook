package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error fetching %s: %v", url, err)
			os.Exit(1)
		}
		fmt.Println("Status code: ", resp.StatusCode)
		body := resp.Body
		_, err = io.Copy(os.Stdout, body)
		body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error copying response body to Stdout for: %s", url)
		}
	}
}
