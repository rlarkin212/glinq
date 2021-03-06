package glinq

import (
	"testing"

	"github.com/rlarkin212/glinq/internal"
	te "github.com/rlarkin212/glinq/test"
	"golang.org/x/exp/constraints"
)

type obdfun[T any, T2 any] func(x T) T2

type orderByDescendingTest[T any, T2 constraints.Ordered] struct {
	input    []T
	fun      obdfun[T, T2]
	expected []T
}

var orderByDescendingTestDataUserInt = []orderByDescendingTest[te.User, int]{
	{
		input: []te.User{
			{
				Name: "z",
				Age:  12,
			},
			{
				Name: "a",
				Age:  100,
			},
			{
				Name: "d",
				Age:  43,
			},
			{
				Name: "y",
				Age:  1,
			},
		},
		fun: func(x te.User) int {
			return x.Age
		},
		expected: []te.User{
			{
				Name: "a",
				Age:  100,
			},
			{
				Name: "d",
				Age:  43,
			},
			{
				Name: "z",
				Age:  12,
			},
			{
				Name: "y",
				Age:  1,
			},
		},
	},
}

var orderByDescendingTestDataUserString = []orderByDescendingTest[te.User, string]{
	{
		input: []te.User{
			{
				Name: "z",
				Age:  12,
			},
			{
				Name: "a",
				Age:  100,
			},
			{
				Name: "d",
				Age:  43,
			},
			{
				Name: "y",
				Age:  1,
			},
		},
		fun: func(x te.User) string {
			return x.Name
		},
		expected: []te.User{
			{
				Name: "z",
				Age:  12,
			},
			{
				Name: "y",
				Age:  1,
			},
			{
				Name: "d",
				Age:  43,
			},
			{
				Name: "a",
				Age:  100,
			},
		},
	},
}
var orderByDescendingTestDataString = []orderByTest[string, string]{
	{
		input: []string{"z", "b", "v", "c"},
		fun: func(x string) string {
			return x
		},
		expected: []string{"z", "v", "c", "b"},
	},
}

func TestOrderByDescending(t *testing.T) {
	for _, test := range orderByDescendingTestDataUserInt {
		actual := OrderByDescending(test.input, test.fun)

		if ok := internal.SliceCompare(test.expected, actual); !ok {
			t.Errorf("expected %v; actual %v", test.expected, actual)
			t.Fail()
		}
	}

	for _, test := range orderByDescendingTestDataUserString {
		actual := OrderByDescending(test.input, test.fun)

		if ok := internal.SliceCompare(test.expected, actual); !ok {
			t.Errorf("expected %v; actual %v", test.expected, actual)
			t.Fail()
		}
	}

	for _, test := range orderByDescendingTestDataString {
		actual := OrderByDescending(test.input, test.fun)

		if ok := internal.SliceCompare(test.expected, actual); !ok {
			t.Errorf("expected %v; actual %v", test.expected, actual)
			t.Fail()
		}
	}
}
