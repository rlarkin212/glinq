package glinq

import (
	"strings"
	"testing"

	"github.com/rlarkin212/glinq/internal"
)

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

var whereTestDataString = []whereTest[string]{
	{
		input: []string{"a", "b", "c", "d", "abcde", "abxyz", "vwxyz"},
		fun: func(v string) bool {
			return v == "a"
		},
		expected: []string{"a"},
	},
	{
		input: []string{"a", "b", "c", "d", "abcde", "abxyz", "vwxyz"},
		fun: func(v string) bool {
			return strings.HasPrefix(v, "ab")
		},
		expected: []string{"abcde", "abxyz"},
	},
	{
		input: []string{"a", "b", "c", "d", "abcde", "abxyz", "vwxyz"},
		fun: func(v string) bool {
			return strings.HasSuffix(v, "yz")
		},
		expected: []string{"abxyz", "vwxyz"},
	},
	{
		input: []string{"a", "b", "c", "d", "abcde", "abxyz", "vwxyz"},
		fun: func(v string) bool {
			return strings.Contains(v, "f")
		},
		expected: []string{},
	},
}

func TestWhere(t *testing.T) {
	for _, test := range whereTestDataInt {
		actual := Where(test.input, test.fun)

		if ok := internal.SliceCompare(test.expected, actual); !ok {
			t.Errorf("expected %v; actual %v", test.expected, actual)
			t.Fail()
		}
	}
}
