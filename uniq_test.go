package gomill

import (
	"reflect"
	"testing"
)

func TestUniqInt(t *testing.T) {
	in := []int{1, 2, 3, 4, 5, 2, 1, 3, -2, 4, 1, 99, 12, 2, 12, 3, 3}
	exp := []int{1, 2, 3, 4, 5, -2, 99, 12}
	res := UniqInt(in)
	if !reflect.DeepEqual(exp, res) {
		t.Errorf("expected slice %v is not equal to %v", exp, res)
	}
}
