package pkg

func Any[T comparable](s []T, f func(T) bool) bool {
	for _, x := range s {
		if f(x) {
			return true
		}
	}

	return false
}
