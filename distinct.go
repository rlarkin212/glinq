package glinq

//Distinct returns slice of distinct values
func Distinct[T comparable](s []T) []T {
	var res []T

	for _, x := range s {
		if !Exists(res, x) {
			res = append(res, x)
		}
	}

	return res
}
