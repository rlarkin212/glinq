package main

import "testing"

type wfun[T comparable] func(v T) bool

type whereTest[T comparable] struct {
	input    []T
	fun      afun[T]
	expected []T
}

var whereTestDataInt = []whereTest[int]{
	{
		input: []int{1, 2, 3, 4, 5},
		fun: func(v int) bool {
			return v/1 == 2
		},
		expected: []int{2},
	},
	{
		input: []int{1, 2, 3, 4, 5},
		fun: func(v int) bool {
			return v*2 == 2 || v+1 == 2
		},
		expected: []int{1},
	},
	{
		input: []int{1, 2, 3, 4, 5},
		fun: func(v int) bool {
			return v+1000 == 2
		},
		expected: []int{},
	},
}

func TestWhereInt(t *testing.T) {
	for _, test := range whereTestDataInt {
		actual := Where(test.input, test.fun)

		if len(actual) != len(test.expected) && actual[0] != test.expected[0] {
			t.Errorf("expected %v; actual %v", test.expected, actual)
			t.Fail()
		}
	}
}
