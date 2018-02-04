package node

import "golang.org/x/net/html"

// AllFn ...
func AllFn(funcs ...func(*html.Node) bool) func(*html.Node) bool {
	return func(n *html.Node) bool {
		for _, fn := range funcs {
			if !fn(n) {
				return false
			}
		}
		return true
	}
}

// AnyFn ...
func AnyFn(funcs ...func(*html.Node) bool) func(*html.Node) bool {
	return func(n *html.Node) bool {
		for _, fn := range funcs {
			if fn(n) {
				return true
			}
		}
		return false
	}
}
