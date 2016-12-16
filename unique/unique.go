// Package unique provides a simple function for removing
// string duplicates from a slice of string.
package unique

// Strings removes duplicated strings from a slice of strings.
// It returns a new slice of strings without duplicates.
func Strings(s []string) []string {
	lst := make([]string, 0, 0)
	set := make(map[string]struct{})
	for _, i := range s {
		_, ok := set[i]
		if ok {
			continue
		}
		set[i] = struct{}{}
		lst = append(lst, i)
	}
	return lst
}
