package glinq

import "testing"

type averageTest[T Number] struct {
	input    []T
	expected float64
}

var averageTestDataInt = []averageTest[int]{
	{
		input:    []int{1, 2, 3, 4, 5},
		expected: 3.0,
	},
	{
		input:    []int{1, 2, 3, 4, -5},
		expected: 1.0,
	},
	{
		input:    []int{1, 0, 0, 0, 0},
		expected: 0.2,
	},
	{
		input:    []int{0, 0, 0, 0, 0},
		expected: 0.0,
	},
}

var averageTestDataUnit = []averageTest[uint]{
	{
		input:    []uint{1, 2, 3, 4, 5},
		expected: 3.0,
	},
	{
		input:    []uint{1, 0, 0, 0, 0},
		expected: 0.2,
	},
	{
		input:    []uint{0, 0, 0, 0, 0},
		expected: 0,
	},
}

var averageDataFloat32 = []averageTest[float32]{
	{
		input:    []float32{1, 2.2, 3.5, 5, 5.5},
		expected: 3.4400000095367433,
	},
	{
		input:    []float32{1, 2.2, 3.5, 4, -5},
		expected: 1.140000009536743,
	},
}

var averageDataFloat64 = []averageTest[float64]{
	{
		input:    []float64{1, 2.2, 3.5, 5, 5.5},
		expected: 3.44,
	},
	{
		input:    []float64{1, 2.2, 3.5, 4, -5},
		expected: 1.14,
	},
}

func TestAverage(t *testing.T) {
	for _, test := range averageTestDataInt {
		actual := Average(test.input)

		if actual != test.expected {
			t.Errorf("expected %v; actual %v", test.expected, actual)
			t.Fail()
		}
	}

	for _, test := range averageTestDataUnit {
		actual := Average(test.input)

		if actual != test.expected {
			t.Errorf("expected %v; actual %v", test.expected, actual)
			t.Fail()
		}
	}

	for _, test := range averageDataFloat32 {
		actual := Average(test.input)

		if actual != test.expected {
			t.Errorf("expected %v; actual %v", test.expected, actual)
			t.Fail()
		}
	}

	for _, test := range averageDataFloat64 {
		actual := Average(test.input)

		if actual != test.expected {
			t.Errorf("expected %v; actual %v", test.expected, actual)
			t.Fail()
		}
	}
}
