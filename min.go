package glinq

// Min returns min value of slice
func Min[T Number](s []T) T {
	var min T

	for i, x := range s {
		if i == 0 || x < min {
			min = x
		}
	}

	return min
}
