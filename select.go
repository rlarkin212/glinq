package glinq

// Select return slice of V based on supplied func
func Select[T any, V any](s []T, f func(T) V) []V {
	var ret []V

	for _, x := range s {
		v := f(x)
		ret = append(ret, v)
	}

	return ret
}
