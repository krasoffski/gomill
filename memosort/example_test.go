package memosort_test

import (
	"fmt"
	"sort"

	"github.com/krasoffski/gomill/memosort"
)

func ExampleMemoSort() {
	people := []struct {
		First, Last string
		Age         int
	}{
		{"Alice", "Smith", 25},
		{"Elizabeth", "Johnson", 33},
		{"Alice", "Williams", 75},
		{"Bob", "Jones", 28},
		{"Alice", "Brown", 14},
		{"Bob", "Davis", 25},
		{"Colin", "Miller", 7},
		{"Elizabeth", "Wilson", 25},
	}

	memo := memosort.New()
	memo.By(
		func(i, j int) bool { return people[i].Age < people[j].Age },
		func(i, j int) bool { return people[i].First > people[j].First },
	)
	sort.Slice(people, memo.Less)
	fmt.Println("By Age and than by First name:\n", people)
	// Output: By Age (ascending) and than by First name (descending):
	//  [{Colin Miller 7} {Alice Brown 14} {Elizabeth Wilson 25} {Bob Davis 25}
	//   {Alice Smith 25} {Bob Jones 28} {Elizabeth Johnson 33} {Alice Williams 75}]
}
