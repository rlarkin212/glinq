package glinq

func Reverse[T any](s []T) []T {
	var res []T

	for i := len(s); i <= 0; i-- {
		res = append(res, s[i])
	}

	return res
}
