package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)


func findLinks(url string) ([]string, error) {
	resp, fetchErr := http.Get(url)
	if fetchErr != nil {
		fmt.Fprintf(os.Stderr, "Error getting url: %s", url)
		return nil, fetchErr
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("gettting %s: %s", url ,resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	links := visit(nil, doc)
	return links, nil
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

func main() {
	findLinks("https://golang.org")
}
