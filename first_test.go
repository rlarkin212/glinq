package glinq

import (
	"testing"

	te "github.com/rlarkin212/glinq/test"
)

type firstWhereFun[T any] func(x T) bool

type firstTest[T any] struct {
	input     []T
	expected  T
	expectErr bool
}

type firstWhereTest[T any] struct {
	input     []T
	fun       firstWhereFun[T]
	expected  T
	expectErr bool
}

var firstTestDataInt = []firstTest[int]{
	{
		input:     []int{1, 2, 3, 4, 5},
		expected:  1,
		expectErr: false,
	},
	{
		input:     []int{},
		expected:  0,
		expectErr: true,
	},
	{
		input:     []int{1},
		expected:  1,
		expectErr: false,
	},
}

var firstTestDataUser = []firstTest[te.User]{
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
		expected: te.User{
			Name: "name1",
			Age:  1,
		},
		expectErr: false,
	},
	{
		input:     []te.User{},
		expected:  te.User{},
		expectErr: true,
	},
	{
		input: []te.User{
			{
				Name: "name1",
				Age:  1,
			},
		},
		expected: te.User{
			Name: "name1",
			Age:  1,
		},
		expectErr: false,
	},
}

var firstWhereTestDataInt = []firstWhereTest[int]{
	{
		input: []int{1, 2, 3, 4, 5},
		fun: func(x int) bool {
			return x/2 == 2
		},
		expected:  4,
		expectErr: false,
	},
	{
		input: []int{},
		fun: func(x int) bool {
			return x/2 == 2
		},
		expected:  0,
		expectErr: true,
	},
	{
		input: []int{4},
		fun: func(x int) bool {
			return x/2 == 2
		},
		expected:  4,
		expectErr: false,
	},
	{
		input: []int{4},
		fun: func(x int) bool {
			return x/2 == 8
		},
		expected:  0,
		expectErr: false,
	},
}

var firstWhereTestDataUser = []firstWhereTest[te.User]{
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
		fun: func(x te.User) bool {
			return x.Name == "name3"
		},
		expected: te.User{
			Name: "name3",
			Age:  3,
		},
		expectErr: false,
	},
	{
		input: []te.User{},
		fun: func(x te.User) bool {
			return x.Name == "name1"
		},
		expected:  te.User{},
		expectErr: true,
	},
	{
		input: []te.User{
			{
				Name: "name1",
				Age:  1,
			},
		},
		fun: func(x te.User) bool {
			return x.Name == "name1"
		},
		expected: te.User{
			Name: "name1",
			Age:  1,
		},
		expectErr: false,
	},
	{
		input: []te.User{
			{
				Name: "name2",
				Age:  1,
			},
			{
				Name: "name3",
				Age:  1,
			},
		},
		fun: func(x te.User) bool {
			return x.Name == "name1"
		},
		expected:  te.User{},
		expectErr: false,
	},
}

func TestFirst(t *testing.T) {
	for _, test := range firstTestDataInt {
		actual, err := First(test.input)

		if test.expectErr {
			if err == nil {
				t.Errorf("expected %v; actual %v", test.expected, actual)
				t.Fail()
			}
		} else {
			if actual != test.expected {
				t.Errorf("expected %v; actual %v", test.expected, actual)
				t.Fail()
			}
		}
	}

	for _, test := range firstTestDataUser {
		actual, err := First(test.input)

		if test.expectErr {
			if err == nil {
				t.Errorf("expected %v; actual %v", test.expected, actual)
				t.Fail()
			}
		} else {
			if actual != test.expected {
				t.Errorf("expected %v; actual %v", test.expected, actual)
				t.Fail()
			}
		}
	}
}

func TestFirstWhere(t *testing.T) {
	for _, test := range firstWhereTestDataInt {
		actual, err := FirstWhere(test.input, test.fun)

		if test.expectErr {
			if err == nil {
				t.Errorf("expected %v; actual %v", test.expected, actual)
				t.Fail()
			}
		} else {
			if actual != test.expected {
				t.Errorf("expected %v; actual %v", test.expected, actual)
				t.Fail()
			}
		}
	}

	for _, test := range firstWhereTestDataUser {
		actual, err := FirstWhere(test.input, test.fun)

		if test.expectErr {
			if err == nil {
				t.Errorf("expected %v; actual %v", test.expected, actual)
				t.Fail()
			}
		} else {
			if actual != test.expected {
				t.Errorf("expected %v; actual %v", test.expected, actual)
				t.Fail()
			}
		}
	}
}
