package glinq

// Join merges the provided slices into one
func Join[T any](s ...[]T) []T {
	var res []T

	for _, x := range s {
		res = append(res, x...)
	}

	return res
}

// JoinWhere merges items in provided slices where item satisfies func
func JoinWhere[T any](f func(x T) bool, s ...[]T) []T {
	var res []T

	for _, x := range s {
		for _, item := range x {
			if f(item) {
				res = append(res, item)
			}
		}
	}

	return res
}
