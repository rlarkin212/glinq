package glinq

// Average return the average of items in slice
func Average[T Number](s []T) float64 {
	n := len(s)
	var sum float64

	for _, x := range s {
		sum += float64(x)
	}

	avg := (sum / float64(n))

	return avg
}
