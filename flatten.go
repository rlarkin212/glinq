package glinq

// Flatten returns slice from slice of slice
func Flatten[T any](s [][]T) []T {
	res := []T{}

	for _, x := range s {
		res = append(res, x...)
	}

	return res
}
