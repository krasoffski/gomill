package node

import "golang.org/x/net/html"

// Attr ...
func Attr(n *html.Node, key string) (string, bool) {
	for _, a := range n.Attr {
		if a.Key == key {
			return a.Val, true
		}
	}
	return "", false
}

// Children ...
func Children(n *html.Node, check func(*html.Node) bool) []*html.Node {
	nodes := make([]*html.Node, 0)

	if check == nil {
		check = func(*html.Node) bool { return true }
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if !check(c) {
			continue
		}
		nodes = append(nodes, c)

	}
	return nodes
}
