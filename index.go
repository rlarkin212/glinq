package main

func Index[T comparable](term T, s []T) int {
	for i, x := range s {
		if x == term {
			return i
		}
	}

	return -1
}
