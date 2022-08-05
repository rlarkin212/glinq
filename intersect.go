package glinq

// Intersect returns slice of items that are in both slices
func Intersct[T comparable](s []T, s1 []T) []T {
	var res []T

	for _, x := range s {
		if Contains(x, s1) {
			res = append(res, x)
		}
	}

	return res
}
