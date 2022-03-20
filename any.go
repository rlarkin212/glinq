package glinq

// Any return true if item in slice satisfies func
func Any[T comparable](s []T, f func(T) bool) bool {
	for _, x := range s {
		if f(x) {
			return true
		}
	}

	return false
}
