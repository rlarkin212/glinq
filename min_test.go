package glinq

import "testing"

type minTest[T Number] struct {
	input    []T
	expected T
}

var minTestDataInt = []minTest[int]{
	{
		input:    []int{1, 2, 3, 4, 5},
		expected: 1,
	},
	{
		input:    []int{1, 2, 3, 4, -5},
		expected: -5,
	},
	{
		input:    []int{1, 0, 0, 0, 0},
		expected: 0,
	},
	{
		input:    []int{0, 0, 0, 0, 0},
		expected: 0,
	},
}

var minTestUnit = []minTest[uint]{
	{
		input:    []uint{1, 2, 3, 4, 5},
		expected: 1,
	},
	{
		input:    []uint{1, 0, 0, 0, 0},
		expected: 0,
	},
	{
		input:    []uint{0, 0, 0, 0, 0},
		expected: 0,
	},
}

var minTestDataFloat32 = []minTest[float32]{
	{
		input:    []float32{1.4, 2.2, 3.5, 5, 5.5},
		expected: 1.4,
	},
	{
		input:    []float32{1, 2.2, 3.5, 4, -5},
		expected: -5,
	},
}

func TestMin(t *testing.T) {
	for _, test := range minTestDataInt {
		actual := Min(test.input)

		if actual != test.expected {
			t.Errorf("expected %v; actual %v", test.expected, actual)
			t.Fail()
		}
	}

	for _, test := range minTestUnit {
		actual := Min(test.input)

		if actual != test.expected {
			t.Errorf("expected %v; actual %v", test.expected, actual)
			t.Fail()
		}
	}

	for _, test := range minTestDataFloat32 {
		actual := Min(test.input)

		if actual != test.expected {
			t.Errorf("expected %v; actual %v", test.expected, actual)
			t.Fail()
		}
	}
}
