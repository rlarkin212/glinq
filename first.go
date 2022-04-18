package glinq

import "errors"

// First returns first item in slice or error if slice contains no items
func First[T any](s []T) (T, error) {
	var res T

	if len(s) == 0 {
		return res, errors.New("slice has 0 items")
	}

	res = s[0]

	return res, nil
}

// FirstWhere returns first item in slice in slice that satisfies func or error if slice contains no items
func FirstWhere[T any](s []T, f func(x T) bool) (T, error) {
	var res T

	if len(s) == 0 {
		return res, errors.New("slice has 0 items")
	}

	for _, x := range s {
		if f(x) {
			res = x
			break
		}
	}

	return res, nil
}
