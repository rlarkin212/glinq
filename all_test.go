package glinq

import (
	"strings"
	"testing"
)

type allfun[T comparable] func(x T) bool

type allTest[T comparable] struct {
	input    []T
	fun      allfun[T]
	expected bool
}

var allTestDataInt = []allTest[int]{
	{
		input: []int{4},
		fun: func(x int) bool {
			return x-2 == 2
		},
		expected: true,
	},
	{
		input: []int{1, 2, 3, 4, 5},
		fun: func(x int) bool {
			return x+1 == 2
		},
		expected: false,
	},
	{
		input: []int{5, 5, 5, 5, 5},
		fun: func(x int) bool {
			return x == 5
		},
		expected: true,
	},
}

var allTestDataString = []allTest[string]{
	{
		input: []string{"a", "a"},
		fun: func(x string) bool {
			return x == "a"
		},
		expected: true,
	},
	{
		input: []string{"a", "b", "c", "d", "e"},
		fun: func(x string) bool {
			return x == "z"
		},
		expected: false,
	},
	{
		input: []string{"ab", "ab", "abc", "abd", "abe", "abcde"},
		fun: func(x string) bool {
			return strings.HasPrefix(x, "ab")
		},
		expected: true,
	},
}

func TestAll(t *testing.T) {
	for _, test := range allTestDataInt {
		actual := All(test.input, test.fun)

		if actual != test.expected {
			t.Errorf("expected %v; actual %v", test.expected, actual)
			t.Fail()
		}
	}

	for _, test := range allTestDataString {
		actual := All(test.input, test.fun)

		if actual != test.expected {
			t.Errorf("expected %v; actual %v", test.expected, actual)
			t.Fail()
		}
	}
}
