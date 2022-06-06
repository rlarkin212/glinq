package glinq

import (
	"sort"

	"golang.org/x/exp/constraints"
)

// OrderBy returns slice with items ordered by returned item in func in ascending order
func OrderBy[T any, T2 constraints.Ordered](s []T, f func(x T) T2) []T {
	sort.Slice(s, func(i, j int) bool {
		item1 := f(s[i])
		item2 := f(s[j])

		return item1 < item2
	})

	return s
}
