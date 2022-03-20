package glinq

import (
	"testing"
)

type containsTest[T comparable] struct {
	input    []T
	term     T
	expected bool
}

var containsTestDataInt = []containsTest[int]{
	{
		input:    []int{1, 2, 3, 4, 5},
		term:     3,
		expected: true,
	},
	{
		input:    []int{1, 2, 3, 4, 5},
		term:     10,
		expected: false,
	},
}

var containsTestDataString = []containsTest[string]{
	{
		input:    []string{"a", "b", "c", "d", "e"},
		term:     "a",
		expected: true,
	},
	{
		input:    []string{"a", "b", "c", "d", "e"},
		term:     "z",
		expected: false,
	},
}

func TestContains(t *testing.T) {
	for _, test := range containsTestDataInt {
		actual := Contains(test.term, test.input)

		if actual != test.expected {
			t.Errorf("expected %v; actual %v", test.expected, actual)
			t.Fail()
		}
	}

	for _, test := range containsTestDataString {
		actual := Contains(test.term, test.input)

		if actual != test.expected {
			t.Errorf("expected %v; actual %v", test.expected, actual)
			t.Fail()
		}
	}
}
