package glinq

type Number interface {
	uint | uint8 | uint16 | uint32 | uint64 | int | int8 | int16 | int32 | int64 | float32 | float64
}

// Sum returns sum of items in slice
func Sum[T Number](s []T) T {
	var sum T

	for _, x := range s {
		sum += x
	}

	return sum
}
