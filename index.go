package main

// Index returns index of item if found in slice else returns -1
func Index[T comparable](term T, s []T) int {
	for i, x := range s {
		if x == term {
			return i
		}
	}

	return -1
}
