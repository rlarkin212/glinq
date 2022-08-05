package glinq

// Except retusn slice of items in first slice but not in second
func Except[T comparable](s []T, s1 []T) []T {
	var res []T

	for _, x := range s {
		if !Contains(x, s1) {
			res = append(res, x)
		}
	}

	return res
}
