package glinq

import (
	"sort"

	"golang.org/x/exp/constraints"
)

func OrderByDescending[T1 any, T2 constraints.Ordered](s []T1, f func(x T1) T2) []T1 {
	sort.Slice(s, func(i, j int) bool {
		item1 := f(s[i])
		item2 := f(s[j])

		return item1 > item2
	})

	return s
}
