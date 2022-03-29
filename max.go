package glinq

//Max returns max value in slice
func Max[T Number](s []T) T {
	var max T

	for i, x := range s {
		if i == 0 || x > max {
			max = x
		}
	}

	return max
}
