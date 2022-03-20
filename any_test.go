package glinq

import (
	"strings"
	"testing"
)

type afun[T comparable] func(v T) bool

type anyTest[T comparable] struct {
	input    []T
	fun      afun[T]
	expected bool
}

var anyTestDataInt = []anyTest[int]{
	{
		input: []int{1, 2, 3, 4, 5},
		fun: func(v int) bool {
			return v+1 == 2
		},
		expected: true,
	},
	{
		input: []int{1, 2, 3, 4, 5},
		fun: func(v int) bool {
			return v*2 == 2
		},
		expected: true,
	},
	{
		input: []int{1, 2, 3, 4, 5},
		fun: func(v int) bool {
			return v+1000 == 2
		},
		expected: false,
	},
}

var anyTestDataString = []anyTest[string]{
	{
		input: []string{"a", "b", "c", "d", "e"},
		fun: func(v string) bool {
			return v == "a"
		},
		expected: true,
	},
	{
		input: []string{"a", "b", "c", "d", "e"},
		fun: func(v string) bool {
			return v == "z"
		},
		expected: false,
	},
	{
		input: []string{"a", "b", "c", "d", "e", "abcde"},
		fun: func(v string) bool {
			return strings.HasPrefix(v, "ab")
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
