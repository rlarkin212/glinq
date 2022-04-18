package glinq

import (
	"errors"
)

func Take[T any](s []T, n int) ([]T, error) {
	if len(s) == 0 {
		return []T{}, errors.New("slice has 0 items")
	}

	if n > len(s) {
		return []T{}, errors.New("slice range out of bounds")
	}

	return s[:n], nil
}
