package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)


func findLinks2(url string) {
	resp, fetchErr := http.Get(url)
	if fetchErr != nil {
		fmt.Fprintf(os.Stderr, "Error getting url: %s", url)
		os.Exit(1)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing body of url: %v", err)
		os.Exit(1)
	}
	tagCount := visit2(make(map[string]int), doc)
	for tag, cnt := range tagCount {
		fmt.Printf("%-5s %d\n", tag, cnt)
	}
}

func visit2(count map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		count[n.Data] += 1
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		count = visit2(count, c)
	}
	return count
}

func main() {
	findLinks2("https://golang.org")
}
