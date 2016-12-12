package gomill

import "testing"

func TestUniqInt(t *testing.T) {
	in := []int{1, 2, 3, 4, 5, 2, 1, 3, -2, 4, 1, 99, 12, 2, 12, 3, 3}
	exp := []int{1, 2, 3, 4, 5, -2, 99, 12}
	res := UniqInt(in)
	if len(exp) != len(res) {
		t.Fail()
	}
	for i := range exp {
		if exp[i] != res[i] {
			t.Fail()
		}
	}
}
