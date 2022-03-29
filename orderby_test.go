package glinq

import (
	"testing"

	"github.com/rlarkin212/glinq/internal"
	"github.com/rlarkin212/glinq/test"
	te "github.com/rlarkin212/glinq/test"
	"golang.org/x/exp/constraints"
)

type obfun[T any, T2 any] func(x T) T2

type orderByTest[T any, T2 constraints.Ordered] struct {
	input    []T
	fun      obfun[T, T2]
	expected []T
}

var orderByTestDataUserInt = []orderByTest[te.User, int]{
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
		fun: func(x test.User) int {
			return x.Age
		},
		expected: []test.User{
			{
				Name: "y",
				Age:  1,
			},
			{
				Name: "z",
				Age:  12,
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

var orderByTestDataUserString = []orderByTest[te.User, string]{
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
		fun: func(x test.User) string {
			return x.Name
		},
		expected: []test.User{
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
			{
				Name: "z",
				Age:  12,
			},
		},
	},
}

var orderByTestDataString = []orderByTest[string, string]{
	{
		input: []string{"z", "b", "v", "c"},
		fun: func(x string) string {
			return x
		},
		expected: []string{"b", "c", "v", "z"},
	},
}

func TestOrderBy(t *testing.T) {
	for _, test := range orderByTestDataUserInt {
		actual := OrderBy(test.input, test.fun)

		if ok := internal.SliceCompare(test.expected, actual); !ok {
			t.Errorf("expected %v; actual %v", test.expected, actual)
			t.Fail()
		}
	}

	for _, test := range orderByTestDataUserString {
		actual := OrderBy(test.input, test.fun)

		if ok := internal.SliceCompare(test.expected, actual); !ok {
			t.Errorf("expected %v; actual %v", test.expected, actual)
			t.Fail()
		}
	}

	for _, test := range orderByTestDataString {
		actual := OrderBy(test.input, test.fun)

		if ok := internal.SliceCompare(test.expected, actual); !ok {
			t.Errorf("expected %v; actual %v", test.expected, actual)
			t.Fail()
		}
	}
}
