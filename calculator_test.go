package calculator_test

import (
	"calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	type testCase struct {
		name string
		a, b float64
		want float64
	}

	var nameInput string

	testCases := []testCase{
		{nameInput, 2, 2, 4},
		{nameInput, 3, 4, 7},
		{nameInput, 0, 4, 4},
		{nameInput, -2, 3, 1},
		{nameInput, 0.3, 0.7, 1.0},
	}

	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)

		if tc.want != got {
			t.Errorf("Adding two plus two. %s, want %f, got %f", "Addition", tc.want, got)
			t.Errorf("Add(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}

	testCases := []testCase{
		{5, 2, 3},
		{2, 3, -1},
		{6.0, 2.1, 3.9},
	}

	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("want %f, got %f", tc.want, got)
			t.Errorf("Sub(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}

	testCases := []testCase{
		{2, 3, 6},
		{-1, 4, -4},
		{0.5, 2, 1},
	}

	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("want %f, got %f", tc.want, got)
			t.Errorf("Mul(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()

	type testCase struct {
		a, b        float64
		want        float64
		expectedErr bool
	}

	testCases := []testCase{
		{6,3,2, false},
		{5, 0,0,true},
		{12,4,3,false},
	}

	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)

		errReceived := err != nil
		if !tc.expectedErr && tc.want != got{
			t.Fatalf("Divide (%f, %f): unexpected error status: %v", tc.a, tc.b,
			errReceived)
		}

		if tc.want != got {
			t.Errorf("want %f, got %f", tc.want, got)
			t.Errorf("Div(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}
