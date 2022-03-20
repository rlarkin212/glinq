package glinq

// Contains returns true if term is contained in slice, else false
func Contains[T comparable](term T, s []T) bool {
	for _, x := range s {
		if x == term {
			return true
		}
	}

	return false
}
