package glinq

import (
	"testing"

	te "github.com/rlarkin212/glinq/test"
)

type sfun[T any, T2 any] func(x T) T2

type selectTest[T any, T2 any] struct {
	input    []T
	fun      sfun[T, T2]
	expected []T2
}

var selectTestDataUser = []selectTest[te.User, any]{
	{
		input: []te.User{
			{
				Name: "name1",
				Age:  1,
			},
			{
				Name: "name2",
				Age:  2,
			},
			{
				Name: "name3",
				Age:  3,
			},
			{
				Name: "name4",
				Age:  4,
			},
		},
		fun: func(x te.User) any {
			return x.Name
		},
		expected: []any{"name1", "name2", "name3", "name4"},
	},
	{
		input: []te.User{
			{
				Name: "name1",
				Age:  1,
			},
			{
				Name: "name2",
				Age:  2,
			},
			{
				Name: "name3",
				Age:  3,
			},
			{
				Name: "name4",
				Age:  4,
			},
		},
		fun: func(x te.User) any {
			return x.Age
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
