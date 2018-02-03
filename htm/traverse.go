package htm

import "golang.org/x/net/html"

// TraverseNode walks through tree of nodes and adds visited node to the slice
// if check function returns true for this node.
func TraverseNode(n *html.Node, check func(*html.Node) bool, nodes []*html.Node) []*html.Node {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if check(c) {
			nodes = append(nodes, c)
		}
		nodes = TraverseNode(c, check, nodes)
	}
	return nodes
}
