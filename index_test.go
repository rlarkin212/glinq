package glinq

import "testing"

type indexTest[T comparable] struct {
	input    []T
	term     T
	expected int
}

var indexTestDataInt = []indexTest[int]{
	{
		input:    []int{1, 2, 3, 4, 5},
		term:     1,
		expected: 0,
	},
	{
		input:    []int{1, 2, 3, 4, 5},
		term:     5,
		expected: 4,
	},
	{
		input:    []int{1, 2, 3, 4, 5},
		term:     20,
		expected: -1,
	},
}

var indexTestDataString = []indexTest[string]{
	{
		input:    []string{"a", "b", "c", "d", "e"},
		term:     "a",
		expected: 0,
	},
	{
		input:    []string{"a", "b", "c", "d", "e"},
		term:     "e",
		expected: 4,
	},
	{
		input:    []string{"a", "b", "c", "d", "e"},
		term:     "z",
		expected: -1,
	},
}

func TestIndexInt(t *testing.T) {
	for _, test := range indexTestDataInt {
		actual := Index(test.term, test.input)

		if actual != test.expected {
			t.Errorf("expected %v; actual %v", test.expected, actual)
			t.Fail()
		}
	}
}

func TestIndexString(t *testing.T) {
	for _, test := range indexTestDataString {
		actual := Index(test.term, test.input)

		if actual != test.expected {
			t.Errorf("expected %v; actual %v", test.expected, actual)
			t.Fail()
		}
	}
}
