package memosort

type MemoSort []func(i, j int) bool

func (m *MemoSort) By(f func(i, j int) bool) func(i, j int) bool {
	*m = append(*m, f)
	return m.Less
}

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

func New() *MemoSort {
	return new(MemoSort)
}
