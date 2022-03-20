package glinq

// Sum returns sum of items in slice
func Sum[T Number](s []T) T {
	var sum T

	for _, x := range s {
		sum += x
	}

	return sum
}
