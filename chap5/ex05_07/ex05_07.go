package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
)

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil && n.Data != "script" && n.Data != "style"; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {
	switch n.Type {
	case html.ElementNode:
		formattedAttrs := make([]string, len(n.Attr))
		for i, attr := range n.Attr {
			formattedAttrs[i] = fmt.Sprintf("%s=%s", attr.Key, attr.Val)
		}
		joinedAttrs := strings.Join(formattedAttrs, " ")
		if n.FirstChild != nil {
			fmt.Printf("%*s<%s %s>\n", depth * 2, "", n.Data, joinedAttrs)
			depth++
		} else {
			fmt.Printf("%*s<%s %s/>\n", depth * 2, "", n.Data, joinedAttrs)
		}
	case html.TextNode, html.CommentNode:
		data := strings.TrimSpace(n.Data)
		if len(data) != 0 {
			fmt.Printf("%*s%s\n", depth*2, "", n.Data)
		}
	}
}

func endElement(n *html.Node) {
	switch n.Type {
	case html.ElementNode:
		if n.FirstChild != nil {
			depth--
			fmt.Printf("%*s</%s>\n", depth * 2, "", n.Data)
		}
	}
}

func pprint(url string) (err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return
	}

	forEachNode(doc, startElement, endElement)
	err = nil
	return
}

func main() {
	url := "https://golang.org"
	err := pprint(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v", err)
	}
}
