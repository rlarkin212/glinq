package glinq

// Where return slice of items where item in slice satisfies func
func Where[T comparable](s []T, f func(T) bool) []T {
	var ret []T

	for _, x := range s {
		if f(x) {
			ret = append(ret, x)
		}
	}

	return ret
}
