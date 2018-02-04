package htm

import "golang.org/x/net/html"

// AttrOfNode ...
func AttrOfNode(n *html.Node, key string) (string, bool) {
	for _, a := range n.Attr {
		if a.Key == key {
			return a.Val, true
		}
	}
	return "", false
}
