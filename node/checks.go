package htm

import "golang.org/x/net/html"

// All ...
func All(funcs ...func(*html.Node) bool) func(*html.Node) bool {
	return func(n *html.Node) bool {
		for _, fn := range funcs {
			if !fn(n) {
				return false
			}
		}
		return true
	}
}

// Any ...
func Any(funcs ...func(*html.Node) bool) func(*html.Node) bool {
	return func(n *html.Node) bool {
		for _, fn := range funcs {
			if fn(n) {
				return true
			}
		}
		return false
	}
}
