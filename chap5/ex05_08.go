package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)


func startElement2(n *html.Node, id string) bool {
	if n.Type == html.ElementNode {
		for _, attr := range n.Attr {
			if attr.Key == "id" && attr.Val == id {
				return false
			}
		}
	}
	return true
}

func forEachNode2(n *html.Node, id string, pre func(n *html.Node, id string) bool, post func(n *html.Node) bool) (node *html.Node){
	if pre != nil {
		if cont := pre(n, id); !cont {
			node = n
			return
		}
	}

	for c := n.FirstChild; c != nil && n.Data != "script" && n.Data != "style"; c = c.NextSibling {
		if node = forEachNode2(c, id, pre, post); node != nil {
			return
		}
	}

	if post != nil {
		post(n)
	}
	return
}

func ElementByID(doc *html.Node, id string) *html.Node {
	return forEachNode2(doc, id, startElement2, nil)
}

func ElementByIDFromURL(url, id string) (node *html.Node, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return
	}

	node = ElementByID(doc, id)
	err = nil
	return
}

func main() {
	url := "https://golang.org"
	node, err := ElementByIDFromURL(url, "nav")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v", err)
	}
	fmt.Printf("found node %v: <%s>\n", node, node.Data)
}
