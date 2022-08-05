package glinq

import (
	"testing"

	"github.com/rlarkin212/glinq/internal"
)

type IntersectTest[T comparable] struct {
	input    [][]T
	expected []T
}

var intersectTestDataInt = []IntersectTest[int]{
	{
		input: [][]int{
			{1, 2, 3, 4, 5},
			{2, 3},
		},
		expected: []int{2, 3},
	},
	{
		input: [][]int{

			{6, 7},

			{10},
		},
		expected: []int{},
	},
}

func TestIntersect(t *testing.T) {
	for _, test := range intersectTestDataInt {
		actual := Intersct(test.input[0], test.input[1])

		if ok := internal.SliceCompare(test.expected, actual); !ok {
			t.Errorf("expected %v; actual %v", test.expected, actual)
			t.Fail()
		}
	}
}
