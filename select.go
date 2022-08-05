package glinq

// Select return slice of T based on supplied func
func Select[T any, T2 any](s []T, f func(T) T2) []T2 {
	var ret []T2

	for _, x := range s {
		val := f(x)
		ret = append(ret, val)
	}

	return ret
}
