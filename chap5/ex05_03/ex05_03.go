package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
)


func findLinks3(url string) {
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
	text := visit3(nil, doc)
	for _, t := range text {
		fmt.Println(t)
	}
}

func visit3(textList []string, n *html.Node) []string {
	if n.Type == html.TextNode {
		text := strings.TrimSpace(n.Data)
		if len(text) > 0 {
			textList = append(textList, n.Data)
		}
	}
	for c := n.FirstChild; c != nil && c.Data != "script" && c.Data != "style"; c = c.NextSibling {
		textList = visit3(textList, c)
	}
	return textList
}

func main() {
	url := "https://golang.org"
	findLinks3(url)
}
