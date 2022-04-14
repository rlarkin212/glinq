package glinq

import "testing"

type reducefunc[T any] func(x T, y T) T

type reduceTest[T any] struct {
	input    []T
	inital   T
	fun      reducefunc[T]
	expected T
}

var reduceTestDataInt = []reduceTest[int]{
	{
		input:  []int{1, 2, 3, 4, 5},
		inital: 0,
		fun: func(x, y int) int {
			return x + y
		},
		expected: 15,
	},
	{
		input:  []int{1, 2, 3, 4, 5},
		inital: 5,
		fun: func(x, y int) int {
			return x + y
		},
		expected: 20,
	},
	{
		input:  []int{1, 2, 3, 4, 5},
		inital: 1,
		fun: func(x, y int) int {
			return x * y
		},
		expected: 120,
	},
}

func TestReduce(t *testing.T) {
	for _, test := range reduceTestDataInt {
		actual := Reduce(test.input, test.fun, test.inital)

		if actual != test.expected {
			t.Errorf("expected %v; actual %v", test.expected, actual)
			t.Fail()
		}
	}
}
