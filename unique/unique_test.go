package unique

import (
	"reflect"
	"testing"
)

func TestUniqueStrings(t *testing.T) {
	in := []string{"ab", "a", "abc", "A", "a", "t", "ab", "qwe", "p", "qwe"}
	exp := []string{"ab", "a", "abc", "A", "t", "qwe", "p"}
	res := Strings(in)
	if !reflect.DeepEqual(exp, res) {
		t.Errorf("expected slice %v is not equal to %v", exp, res)
	}
}
