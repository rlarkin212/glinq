package glinq

import (
	"testing"

	"github.com/rlarkin212/glinq/internal"
)

type flattenTest[T any] struct {
	input    [][]T
	expected []T
}

var flattenTestDataInt = []flattenTest[int]{
	{
		input: [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
		expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
	},
}

var flattenTestDataString = []flattenTest[string]{
	{
		input: [][]string{
			{"1", "2", "3"},
			{"4", "5", "6"},
			{"7", "8", "9"},
		},
		expected: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"},
	},
}

func TestFlatten(t *testing.T) {
	for _, test := range flattenTestDataInt {
		actual := Flatten(test.input)

		if ok := internal.SliceCompare(test.expected, actual); !ok {
			t.Errorf("expected %v; actual %v", test.expected, actual)
			t.Fail()
		}
	}

	for _, test := range flattenTestDataString {
		actual := Flatten(test.input)

		if ok := internal.SliceCompare(test.expected, actual); !ok {
			t.Errorf("expected %v; actual %v", test.expected, actual)
			t.Fail()
		}
	}
}
