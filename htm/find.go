package htm

import "golang.org/x/net/html"

// FindNode walks through tree of nodes and call check function for each one and
// returns node if check function returns true.
func FindNode(n *html.Node, check func(*html.Node) bool) *html.Node {
	if check(n) {
		return n
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		node := FindNode(c, check)
		if node != nil {
			return node
		}
	}
	return nil
}
