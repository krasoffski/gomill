package node

import (
	"bytes"

	"golang.org/x/net/html"
)

// Attr returns attribute of node by given key if any.
func Attr(n *html.Node, key string) (string, bool) {
	for _, a := range n.Attr {
		if a.Key == key {
			return a.Val, true
		}
	}
	return "", false
}

// Children returns slice of nodes where each node is a child of given node and
// filter function returns true for corresponding child node.
func Children(n *html.Node, filter func(*html.Node) bool) []*html.Node {
	nodes := make([]*html.Node, 0)

	if filter == nil {
		filter = func(*html.Node) bool { return true }
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if !filter(c) {
			continue
		}
		nodes = append(nodes, c)

	}
	return nodes
}

// JoinData returns attribute of node by given key if any.
func JoinData(n ...*html.Node) string {
	var buf bytes.Buffer
	for _, m := range n {
		buf.WriteString(m.Data)
	}
	return buf.String()
}
