package glinq

import "testing"

type maxTest[T Number] struct {
	input    []T
	expected T
}

var maxTestDataInt = []maxTest[int]{
	{
		input:    []int{1, 2, 3, 4, 5},
		expected: 5,
	},
	{
		input:    []int{1, 2, 3, 4, -5},
		expected: 4,
	},
	{
		input:    []int{1, 0, 0, 0, 0},
		expected: 1,
	},
	{
		input:    []int{0, 0, 0, 0, 0},
		expected: 0,
	},
}

var maxTestDataUnit = []maxTest[uint]{
	{
		input:    []uint{1, 2, 3, 4, 5},
		expected: 5,
	},
	{
		input:    []uint{1, 0, 0, 0, 0},
		expected: 1,
	},
	{
		input:    []uint{0, 0, 0, 0, 0},
		expected: 0,
	},
}

var maxTestDataFloat32 = []maxTest[float32]{
	{
		input:    []float32{1, 2.2, 3.5, 5, 5.5},
		expected: 5.5,
	},
	{
		input:    []float32{1, 2.2, 3.5, 4, -5},
		expected: 4,
	},
}

func TestMax(t *testing.T) {
	for _, test := range maxTestDataInt {
		actual := Max(test.input)

		if actual != test.expected {
			t.Errorf("expected %v; actual %v", test.expected, actual)
			t.Fail()
		}
	}

	for _, test := range maxTestDataUnit {
		actual := Max(test.input)

		if actual != test.expected {
			t.Errorf("expected %v; actual %v", test.expected, actual)
			t.Fail()
		}
	}

	for _, test := range maxTestDataFloat32 {
		actual := Max(test.input)

		if actual != test.expected {
			t.Errorf("expected %v; actual %v", test.expected, actual)
			t.Fail()
		}
	}
}
