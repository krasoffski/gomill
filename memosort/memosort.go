// Package memosort implements a function which keeps order from previous sort
// for equal elements. It behaves as sort.Stable but only requires less function
// as argument and can be used with sort.Slice. Here is short example:
// m := memosort.New()
// sort.Slice(tracks, m.By(func(i, j int) bool {
// 	return tracks[i].Title < tracks[j].Title
// }))
// sort.Slice(tracks, m.By(func(i, j int) bool {
// 	return tracks[i].Year < tracks[j].Year
// }))
// sort.Slice(tracks, m.By(func(i, j int) bool {
// 	return tracks[i].Length < tracks[j].Length
// }))

package memosort

// MemoSort holds sort function call order.
type MemoSort []func(i, j int) bool

// By appends less function to internal storage and returns new one which keeps
// order from previous sort for equal elements.
func (m *MemoSort) By(f func(i, j int) bool) func(i, j int) bool {
	*m = append(*m, f)
	return m.Less
}

// Less implemented interface order from previous sort for equal elements.
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
