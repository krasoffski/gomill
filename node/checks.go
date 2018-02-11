package node

import "golang.org/x/net/html"

// AllFn gets a slice of filters and returns new filter func. This function
// calls each filter from the slice with given *html.Node and returns okRes if
// all filters return true otherwise this function returns !okRes.
func AllFn(okRes bool, filters ...func(*html.Node) bool) func(*html.Node) bool {
	return func(n *html.Node) bool {
		for _, filter := range filters {
			if !filter(n) {
				return !okRes
			}
		}
		return okRes
	}
}

// AnyFn gets a slice of filters and returns new func. This function calls each
// filter from the slice with given *html.Node and returns okRes if any call of
// filter returns true. If all filters returns false this function returns
// !okRes.
func AnyFn(okRes bool, filters ...func(*html.Node) bool) func(*html.Node) bool {
	return func(n *html.Node) bool {
		for _, filter := range filters {
			if filter(n) {
				return okRes
			}
		}
		return !okRes
	}
}
