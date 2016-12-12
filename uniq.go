package gomill

func UniqInt(s []int) []int {
	lst := make([]int, 0, 0)
	set := make(map[int]Empty)
	for _, i := range s {
		_, ok := set[i]
		if ok {
			continue
		}
		set[i] = Empty{}
		lst = append(lst, i)
	}
	return lst
}
