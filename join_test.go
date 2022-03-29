package glinq

import (
	"fmt"
	"testing"

	"github.com/rlarkin212/glinq/internal"
)

type jwFun[T any] func(x T) bool

type joinTest[T any] struct {
	input    [][]T
	expected []T
}

type joinWhereTest[T any] struct {
	input    [][]T
	fun      jwFun[T]
	expected []T
}

var joinTestDataInt = []joinTest[int]{
	{
		input: [][]int{
			{1, 2, 3, 4, 5},
			{6, 7},
			{8, 9},
			{10},
		},
		expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	},
	{
		input: [][]int{

			{6, 7},

			{10},
		},
		expected: []int{6, 7, 10},
	},
}

var joinWhereTestDataInt = []joinWhereTest[int]{
	{
		input: [][]int{
			{1, 2, 3, 4, 5},
			{6, 7},
			{8, 9},
			{10},
		},
		fun: func(x int) bool {
			return x%2 == 0
		},
		expected: []int{2, 4, 6, 8, 10},
	},
	{
		input: [][]int{
			{6, 7},
			{10},
		},
		fun: func(x int) bool {
			return x/2 == 3 || x/2 == 5
		},
		expected: []int{6, 7, 10},
	},
}

func TestJoin(t *testing.T) {
	for _, test := range joinTestDataInt {
		actual := Join(test.input...)

		fmt.Println(actual)
		if ok := internal.SliceCompare(test.expected, actual); !ok {
			t.Errorf("expected %v; actual %v", test.expected, actual)
			t.Fail()
		}
	}
}

func TestJoinWhere(t *testing.T) {
	for _, test := range joinWhereTestDataInt {
		actual := JoinWhere(test.fun, test.input...)

		fmt.Println(actual)
		if ok := internal.SliceCompare(test.expected, actual); !ok {
			t.Errorf("expected %v; actual %v", test.expected, actual)
			t.Fail()
		}
	}
}
