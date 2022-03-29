package glinq

import (
	"testing"

	"github.com/rlarkin212/glinq/internal"
	te "github.com/rlarkin212/glinq/test"
)

type distinctTest[T comparable] struct {
	input    []T
	expected []T
}

var distinctTestDataString = []distinctTest[string]{
	{
		input:    []string{"a", "a", "c"},
		expected: []string{"a", "c"},
	},
	{
		input:    []string{"a", "b", "c"},
		expected: []string{"a", "b", "c"},
	},
	{
		input:    []string{},
		expected: []string{},
	},
}

var distinctTestDataInt = []distinctTest[int]{
	{
		input:    []int{1, 1, 3},
		expected: []int{1, 3},
	},
	{
		input:    []int{1, 2, 3},
		expected: []int{1, 2, 3},
	},
	{
		input:    []int{},
		expected: []int{},
	},
}

var distinctTestDataUser = []distinctTest[te.User]{
	{
		input: []te.User{
			{
				Name: "a",
				Age:  1,
			},
			{
				Name: "a",
				Age:  1,
			},
			{
				Name: "c",
				Age:  3,
			},
		},
		expected: []te.User{
			{
				Name: "a",
				Age:  1,
			},
			{
				Name: "c",
				Age:  3,
			},
		},
	},
	{
		input: []te.User{
			{
				Name: "a",
				Age:  1,
			},
			{
				Name: "b",
				Age:  2,
			},
			{
				Name: "c",
				Age:  3,
			},
		},
		expected: []te.User{
			{
				Name: "a",
				Age:  1,
			},
			{
				Name: "b",
				Age:  2,
			},
			{
				Name: "c",
				Age:  3,
			},
		},
	},
	{
		input:    []te.User{},
		expected: []te.User{},
	},
}

func TestDistinct(t *testing.T) {
	for _, test := range distinctTestDataInt {
		actual := Distinct(test.input)

		if ok := internal.SliceCompare(test.expected, actual); !ok {
			t.Errorf("expected %v; actual %v", test.expected, actual)
			t.Fail()
		}
	}

	for _, test := range distinctTestDataString {
		actual := Distinct(test.input)

		if ok := internal.SliceCompare(test.expected, actual); !ok {
			t.Errorf("expected %v; actual %v", test.expected, actual)
			t.Fail()
		}
	}

	for _, test := range distinctTestDataUser {
		actual := Distinct(test.input)

		if ok := internal.SliceCompare(test.expected, actual); !ok {
			t.Errorf("expected %v; actual %v", test.expected, actual)
			t.Fail()
		}
	}
}
