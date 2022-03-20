package glinq

// Range returns slice of items between index start & end
func Range[T any](s []T, start int, end int) []T {
	var ret []T

	for i, x := range s {
		if i >= start && i <= end {
			ret = append(ret, x)
		}
	}

	return ret
}
