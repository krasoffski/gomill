package node

import "golang.org/x/net/html"

// Find walks through tree of nodes and call filter function for each one and
// returns node if filter returns true.
func Find(n *html.Node, filter func(*html.Node) bool) *html.Node {
	if filter(n) {
		return n
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		node := Find(c, filter)
		if node != nil {
			return node
		}
	}
	return nil
}
