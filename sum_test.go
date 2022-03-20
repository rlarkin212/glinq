package glinq

import "testing"

type sumTest[T Number] struct {
	input    []T
	expected T
}

var sumTestDataInt = []sumTest[int]{
	{
		input:    []int{1, 2, 3, 4, 5},
		expected: 15,
	},
	{
		input:    []int{1, 2, 3, 4, -5},
		expected: 5,
	},
}

var sumTestDataUInt = []sumTest[uint]{
	{
		input:    []uint{1, 2, 3, 4, 5},
		expected: 15,
	},
}

var sumTestDataFloat32 = []sumTest[float32]{
	{
		input:    []float32{1, 2.2, 3.5, 4, 5},
		expected: 15.7,
	},
	{
		input:    []float32{1, 2.2, 3.5, 4, -5},
		expected: 5.7,
	},
}

func TestSumInt(t *testing.T) {
	for _, test := range sumTestDataInt {
		actual := Sum(test.input)

		if actual != test.expected {
			t.Errorf("expected %v; actual %v", test.expected, actual)
			t.Fail()
		}
	}
}

func TestSumUInt(t *testing.T) {
	for _, test := range sumTestDataUInt {
		actual := Sum(test.input)

		if actual != test.expected {
			t.Errorf("expected %v; actual %v", test.expected, actual)
			t.Fail()
		}
	}
}

func TestSumFloat32(t *testing.T) {
	for _, test := range sumTestDataFloat32 {
		actual := Sum(test.input)

		if actual != test.expected {
			t.Errorf("expected %v; actual %v", test.expected, actual)
			t.Fail()
		}
	}
}
