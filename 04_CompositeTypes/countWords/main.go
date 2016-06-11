package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, _ := html.Parse(os.Stdin)

	count := wcount(nil, doc)
	fmt.Println(count)

}

func wcount(links []string, n *html.Node) int {
	var w int
	w = 0
	if n.Type == html.TextNode && n.Data == "p" {
		for _, b := range n.Attr {
			w += len(WordCount(b.Val))
		}
	}
	return w

}
func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, f := range strings.Fields(s) {
		m[f] += 1
	}

	return m
}
