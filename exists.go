package glinq

// Exists return true if item exists in slice
func Exists[T comparable](s []T, item T) bool {
	for _, x := range s {
		if x == item {
			return true
		}
	}

	return false
}
