// Package memosort implements a multi-tier ordering function Less with ability
// to specify primary, secondary and so forth comparing functions.
// Adding orders plays key role for comparison. Like first added function is
// primary comparing function, secondary added functin is secondary key.
// This functionality should be used with sort.Slice. Please, take a look on
// testing file for usage examples.
package memosort

// MemoSort holds sort functions call order.
type MemoSort []func(i, j int) bool

// By appends less function or functions to slice and returns new less function
// which keeps order from previous sort function.
// Don't define new type for less function for better documentation.
func (m *MemoSort) By(f ...func(i, j int) bool) func(i, j int) bool {
	*m = append(*m, f...)
	return m.Less
}

// Less compares two elements using all previous less functions calls.
func (m *MemoSort) Less(i, j int) bool {
	for _, fn := range *m {
		switch {
		case fn(i, j):
			return true
		case fn(j, i):
			return false
		}
	}
	return false
}

// New returns pointer to new MemoSort object.
func New() *MemoSort {
	return new(MemoSort)
}
