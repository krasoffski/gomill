package htm

import "golang.org/x/net/html"

// TraverseNode walks through tree of nodes and add visited node to the slice if
// all skip functions returns false for this node.
func TraverseNode(n *html.Node, nodes []*html.Node, skip ...func(*html.Node) bool) []*html.Node {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		skipped := false
		for _, fn := range skip {
			if fn(c) {
				skipped = true
				break
			}
		}
		if !skipped {
			nodes = append(nodes, c)
		}
		nodes = TraverseNode(c, nodes, skip...)
	}
	return nodes
}
