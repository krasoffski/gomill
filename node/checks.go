package node

import "golang.org/x/net/html"

// AllFn gets a slice of funcs and returns new func. This function calls each
// func from the slice with given *html.Node and returns okRes if all calls of
// func returns true otherwise this function returns !okRes.
func AllFn(okRes bool, funcs ...func(*html.Node) bool) func(*html.Node) bool {
	return func(n *html.Node) bool {
		for _, fn := range funcs {
			if !fn(n) {
				return !okRes
			}
		}
		return okRes
	}
}

// AnyFn gets a slice of funcs and returns new func. This function calls each
// func from the slice with given *html.Node and returns okRes if any call of
// func returns true. If all calls returns false otherwise this function returns
// !okRes.
func AnyFn(okRes bool, funcs ...func(*html.Node) bool) func(*html.Node) bool {
	return func(n *html.Node) bool {
		for _, fn := range funcs {
			if fn(n) {
				return okRes
			}
		}
		return !okRes
	}
}
