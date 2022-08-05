package glinq

import (
	"testing"

	"github.com/rlarkin212/glinq/internal"
)

type ExceptTest[T comparable] struct {
	input    [][]T
	expected []T
}

var exceptTestDataInt = []ExceptTest[int]{
	{
		input: [][]int{
			{1, 2, 3, 4, 5},
			{2, 3},
		},
		expected: []int{1, 4, 5},
	},
	{
		input: [][]int{
			{1, 2, 3, 4, 5},
			{7, 8},
		},
		expected: []int{1, 2, 3, 4, 5},
	},
	{
		input: [][]int{

			{6, 7},

			{},
		},
		expected: []int{6, 7},
	},
	{
		input: [][]int{

			{},

			{},
		},
		expected: []int{},
	},
}

func TestExcept(t *testing.T) {
	for _, test := range exceptTestDataInt {
		actual := Except(test.input[0], test.input[1])

		if ok := internal.SliceCompare(test.expected, actual); !ok {
			t.Errorf("expected %v; actual %v", test.expected, actual)
			t.Fail()
		}
	}
}
