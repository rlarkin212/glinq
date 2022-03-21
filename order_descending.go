package glinq

import (
	"sort"

	"golang.org/x/exp/constraints"
)

func OrderByDescending[T any, T1 constraints.Ordered](s []T, f func(x T) T1) []T {
	sort.Slice(s, func(i, j int) bool {
		item1 := f(s[i])
		item2 := f(s[j])

		return item1 > item2
	})

	return s
}
