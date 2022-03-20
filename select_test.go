package glinq

import (
	"testing"
)

type sfun[T any, V any] func(v T) V

type selectTest[T any, V any] struct {
	input    []T
	fun      sfun[T, V]
	expected []V
}

type user struct {
	name string
	age  int
}

var selectTestDataUser = []selectTest[user, any]{
	{
		input: []user{
			{
				name: "name1",
				age:  1,
			},
			{
				name: "name2",
				age:  2,
			},
			{
				name: "name3",
				age:  3,
			},
			{
				name: "name4",
				age:  4,
			},
		},
		fun: func(v user) any {
			return v.name
		},
		expected: []any{"name1", "name2", "name3", "name4"},
	},
	{
		input: []user{
			{
				name: "name1",
				age:  1,
			},
			{
				name: "name2",
				age:  2,
			},
			{
				name: "name3",
				age:  3,
			},
			{
				name: "name4",
				age:  4,
			},
		},
		fun: func(v user) any {
			return v.age
		},
		expected: []any{1, 2, 3, 4},
	},
}

func TestSelect(t *testing.T) {
	for _, test := range selectTestDataUser {
		actual := Select(test.input, test.fun)

		if len(actual) != len(test.expected) && actual[0] != test.expected[0] {
			t.Errorf("expected %v; actual %v", test.expected, actual)
			t.Fail()
		}
	}
}
