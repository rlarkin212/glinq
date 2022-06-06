package glinq

import "errors"

// Last returns last item in slice or error if slice contains no items
func Last[T any](s []T) (T, error) {
	var res T

	if len(s) == 0 {
		return res, errors.New("slice has 0 items")
	}

	res = s[len(s)-1]

	return res, nil
}

// LastWhere returns last item in slice in slice that satisfies func or error if slice contains no items
func LastWhere[T any](s []T, f func(x T) bool) (T, error) {
	var res T

	if len(s) == 0 {
		return res, errors.New("slice has 0 items")
	}

	for i := len(s) - 1; i >= 0; i-- {
		if f(s[i]) {
			res = s[i]
			break
		}
	}

	return res, nil
}
