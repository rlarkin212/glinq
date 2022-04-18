package glinq

import (
	"testing"

	"github.com/rlarkin212/glinq/internal"
	te "github.com/rlarkin212/glinq/test"
)

type takeTest[T any] struct {
	input     []T
	n         int
	expected  []T
	expectErr bool
}

var takeTestDataInt = []takeTest[int]{
	{
		input:     []int{1, 2, 3, 4, 5},
		n:         2,
		expected:  []int{1, 2},
		expectErr: false,
	},
	{
		input:     []int{1, 2},
		n:         2,
		expected:  []int{1, 2},
		expectErr: false,
	},
	{
		input:     []int{1, 2},
		n:         3,
		expected:  []int{},
		expectErr: true,
	},
	{
		input:     []int{},
		n:         2,
		expected:  []int{},
		expectErr: true,
	},
}

var takeTestDataUser = []takeTest[te.User]{
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
		n: 2,
		expected: []te.User{
			{
				Name: "z",
				Age:  12,
			},
			{
				Name: "a",
				Age:  100,
			},
		},
		expectErr: false,
	},
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
		},
		n: 2,
		expected: []te.User{
			{
				Name: "z",
				Age:  12,
			},
			{
				Name: "a",
				Age:  100,
			},
		},
		expectErr: false,
	},
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
		},
		n:         3,
		expected:  []te.User{},
		expectErr: true,
	},
	{
		input:     []te.User{},
		n:         2,
		expected:  []te.User{},
		expectErr: true,
	},
}

func TestTake(t *testing.T) {
	for _, test := range takeTestDataInt {
		actual, err := Take(test.input, test.n)

		if test.expectErr {
			if err == nil {
				t.Errorf("expected %v; actual %v", test.expected, actual)
				t.Fail()
			}
		} else {
			if ok := internal.SliceCompare(test.expected, actual); !ok {
				t.Errorf("expected %v; actual %v", test.expected, actual)
				t.Fail()
			}
		}

	}

	for _, test := range takeTestDataUser {
		actual, err := Take(test.input, test.n)

		if test.expectErr {
			if err == nil {
				t.Errorf("expected %v; actual %v", test.expected, actual)
				t.Fail()
			}
		} else {
			if ok := internal.SliceCompare(test.expected, actual); !ok {
				t.Errorf("expected %v; actual %v", test.expected, actual)
				t.Fail()
			}
		}
	}
}
