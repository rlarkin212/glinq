package glinq

import (
	"testing"

	"github.com/rlarkin212/glinq/internal"
)

type rangeTest[T any] struct {
	input    []T
	start    int
	end      int
	expected []T
}

var rangeTestDataInt = []rangeTest[int]{
	{
		input:    []int{1, 2, 3, 4, 5},
		start:    1,
		end:      3,
		expected: []int{2, 3, 4},
	},
	{
		input:    []int{1, 2, 3, 4, 5},
		start:    0,
		end:      0,
		expected: []int{1},
	},
	{
		input:    []int{},
		start:    0,
		end:      0,
		expected: []int{},
	},
}

func TestRange(t *testing.T) {
	for _, test := range rangeTestDataInt {
		actual := Range(test.input, test.start, test.end)

		if ok := internal.SliceCompare(test.expected, actual); !ok {
			t.Errorf("expected %v; actual %v", test.expected, actual)
			t.Fail()
		}
	}
}
