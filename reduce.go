package glinq

//Reduce executes a user-supplied "reducer" function on each element of the slice
func Reduce[T any](s []T, f func(x T, y T) T, initial T) T {
	result := initial
	for _, x := range s {
		result = f(result, x)
	}

	return result
}
