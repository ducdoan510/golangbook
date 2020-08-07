package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"strings"
)

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(node *html.Node) (words, images int) {
	if node.Type == html.TextNode {
		data := strings.TrimSpace(node.Data)
		words = len(strings.Fields(data))
	} else if node.Type == html.ElementNode && node.Data == "img" {
		fmt.Println(node.Attr)
		images = 1
	}
	for c := node.FirstChild; c != nil && c.Data != "script" && c.Data != "style"; c = c.NextSibling {
		cWords, cImages := countWordsAndImages(c)
		words += cWords
		images += cImages
	}
	return
}

func main() {
	url := "https://golang.org"
	words, images, err := CountWordsAndImages(url)
	if err == nil {
		fmt.Println(words, images)
	}
}
