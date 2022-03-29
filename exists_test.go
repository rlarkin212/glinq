package glinq

import (
	"testing"

	te "github.com/rlarkin212/glinq/test"
)

type existsTest[T comparable] struct {
	input    []T
	item     T
	expected bool
}

var existsTestDataString = []existsTest[string]{
	{
		input:    []string{"a", "b", "c"},
		item:     "a",
		expected: true,
	},
	{
		input:    []string{"a", "b", "c"},
		item:     "z",
		expected: false,
	},
	{
		input:    []string{},
		item:     "a",
		expected: false,
	},
}

var existsTestDataInt = []existsTest[int]{
	{
		input:    []int{1, 2, 3},
		item:     1,
		expected: true,
	},
	{
		input:    []int{1, 2, 3},
		item:     10,
		expected: false,
	},
	{
		input:    []int{},
		item:     1,
		expected: false,
	},
}

var existsTestDataUser = []existsTest[te.User]{
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
		item: te.User{
			Name: "a",
			Age:  1,
		},
		expected: true,
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
		item: te.User{
			Name: "z",
			Age:  10,
		},
		expected: false,
	},
	{
		input: []te.User{},
		item: te.User{
			Name: "z",
			Age:  10,
		},
		expected: false,
	},
}

func TestExists(t *testing.T) {
	for _, test := range existsTestDataString {
		actual := Exists(test.input, test.item)

		if actual != test.expected {
			t.Errorf("expected %v; actual %v", test.expected, actual)
			t.Fail()
		}
	}

	for _, test := range existsTestDataInt {
		actual := Exists(test.input, test.item)

		if actual != test.expected {
			t.Errorf("expected %v; actual %v", test.expected, actual)
			t.Fail()
		}
	}

	for _, test := range existsTestDataUser {
		actual := Exists(test.input, test.item)

		if actual != test.expected {
			t.Errorf("expected %v; actual %v", test.expected, actual)
			t.Fail()
		}
	}
}
