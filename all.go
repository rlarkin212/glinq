package glinq

// All return true if all items in slice satisfy func
func All[T any](s []T, f func(x T) bool) bool {
	for _, x := range s {
		if !f(x) {
			return false
		}
	}

	return true
}
