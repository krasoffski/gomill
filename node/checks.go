package node

import "golang.org/x/net/html"

// AllFn ...
func AllFn(funcs ...func(*html.Node) bool) func(*html.Node) bool {
	return all(true, funcs...)
}

// AllFnN ...
func AllFnN(funcs ...func(*html.Node) bool) func(*html.Node) bool {
	return all(false, funcs...)
}

// AnyFn ...
func AnyFn(funcs ...func(*html.Node) bool) func(*html.Node) bool {
	return any(true, funcs...)
}

// AnyFnN ...
func AnyFnN(funcs ...func(*html.Node) bool) func(*html.Node) bool {
	return any(false, funcs...)
}

// AllFn ...
func all(okRes bool, funcs ...func(*html.Node) bool) func(*html.Node) bool {
	return func(n *html.Node) bool {
		for _, fn := range funcs {
			if !fn(n) {
				return !okRes
			}
		}
		return okRes
	}
}

func any(okRes bool, funcs ...func(*html.Node) bool) func(*html.Node) bool {
	return func(n *html.Node) bool {
		for _, fn := range funcs {
			if fn(n) {
				return okRes
			}
		}
		return !okRes
	}
}
