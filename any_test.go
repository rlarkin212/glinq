package glinq

import (
	"strings"
	"testing"
)

type afun[T comparable] func(x T) bool

type anyTest[T comparable] struct {
	input    []T
	fun      afun[T]
	expected bool
}

var anyTestDataInt = []anyTest[int]{
	{
		input: []int{1, 2, 3, 4, 5},
		fun: func(x int) bool {
			return x+1 == 2
		},
		expected: true,
	},
	{
		input: []int{1, 2, 3, 4, 5},
		fun: func(x int) bool {
			return x*2 == 2
		},
		expected: true,
	},
	{
		input: []int{1, 2, 3, 4, 5},
		fun: func(x int) bool {
			return x+1000 == 2
		},
		expected: false,
	},
}

var anyTestDataString = []anyTest[string]{
	{
		input: []string{"a", "b", "c", "d", "e"},
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
		input: []string{"a", "b", "c", "d", "e", "abcde"},
		fun: func(x string) bool {
			return strings.HasPrefix(x, "ab")
		},
		expected: true,
	},
}

func TestAny(t *testing.T) {
	for _, test := range anyTestDataInt {
		actual := Any(test.input, test.fun)

		if actual != test.expected {
			t.Errorf("expected %v; actual %v", test.expected, actual)
			t.Fail()
		}
	}

	for _, test := range anyTestDataString {
		actual := Any(test.input, test.fun)

		if actual != test.expected {
			t.Errorf("expected %v; actual %v", test.expected, actual)
			t.Fail()
		}
	}
}
