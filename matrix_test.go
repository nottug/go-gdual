package gdual

import (
	"testing"
)

func TestElementAdd(t *testing.T) {
	tests := []struct {
		order    int
		element  float64
		input    []float64
		expected []float64
	}{
		{
			order:    2,
			element:  5,
			input:    []float64{1, 2},
			expected: []float64{6, 7},
		},
		{
			order:    5,
			element:  2,
			input:    []float64{1, 2, 3, 4, 5},
			expected: []float64{3, 4, 5, 6, 7},
		},
	}

	// run tests for the standard matrix
	for i, tt := range tests {
		mat := importMatrix(tt.input)
		if mat.order != tt.order {
			t.Errorf("order mismatch on test %d: have %d want %d",
				i, mat.order, tt.order)
		}

		mat.ElementAdd(tt.element)

		for n := 0; n < mat.order; n++ {
			if mat.get(0, n) != tt.expected[n] {
				t.Errorf("value mismatch on standard test %d (col %d): have %f want %f",
					i, n, mat.get(0, n), tt.expected[n])
			}
		}
	}

	// run tests for the upper triangular Toeplitz matrix
	for i, tt := range tests {
		mat := importUpperTriToeplitz(tt.input)
		if mat.order != tt.order {
			t.Errorf("order mismatch on UTT test %d: have %d want %d",
				i, mat.order, tt.order)
		}

		mat.ElementAdd(tt.element)

		for n := 0; n < mat.order; n++ {
			if mat.get(n) != tt.expected[n] {
				t.Errorf("value mismatch on UTT test %d (col %d): have %f want %f",
					i, n, mat.get(n), tt.expected[n])
			}
		}
	}
}

func TestElementSub(t *testing.T) {
	tests := []struct {
		order    int
		element  float64
		input    []float64
		expected []float64
	}{
		{
			order:    2,
			element:  5,
			input:    []float64{1, 2},
			expected: []float64{-4, -3},
		},
		{
			order:    5,
			element:  2,
			input:    []float64{1, 2, 3, 4, 5},
			expected: []float64{-1, 0, 1, 2, 3},
		},
	}

	// run tests for the standard matrix
	for i, tt := range tests {
		mat := importMatrix(tt.input)
		if mat.order != tt.order {
			t.Errorf("order mismatch on test %d: have %d want %d",
				i, mat.order, tt.order)
		}

		mat.ElementSub(tt.element)

		for n := 0; n < mat.order; n++ {
			if mat.get(0, n) != tt.expected[n] {
				t.Errorf("value mismatch on standard test %d (col %d): have %f want %f",
					i, n, mat.get(0, n), tt.expected[n])
			}
		}
	}

	// run tests for the upper triangular Toeplitz matrix
	for i, tt := range tests {
		mat := importUpperTriToeplitz(tt.input)
		if mat.order != tt.order {
			t.Errorf("order mismatch on UTT test %d: have %d want %d",
				i, mat.order, tt.order)
		}

		mat.ElementSub(tt.element)

		for n := 0; n < mat.order; n++ {
			if mat.get(n) != tt.expected[n] {
				t.Errorf("value mismatch on UTT test %d (col %d): have %f want %f",
					i, n, mat.get(n), tt.expected[n])
			}
		}
	}
}

func TestElementMul(t *testing.T) {
	tests := []struct {
		order    int
		element  float64
		input    []float64
		expected []float64
	}{
		{
			order:    2,
			element:  5,
			input:    []float64{1, 2},
			expected: []float64{5, 10},
		},
		{
			order:    5,
			element:  2,
			input:    []float64{1, 2, 3, 4, 5},
			expected: []float64{2, 4, 6, 8, 10},
		},
	}

	// run tests for the standard matrix
	for i, tt := range tests {
		mat := importMatrix(tt.input)
		if mat.order != tt.order {
			t.Errorf("order mismatch on test %d: have %d want %d",
				i, mat.order, tt.order)
		}

		mat.ElementMul(tt.element)

		for n := 0; n < mat.order; n++ {
			if mat.get(0, n) != tt.expected[n] {
				t.Errorf("value mismatch on standard test %d (col %d): have %f want %f",
					i, n, mat.get(0, n), tt.expected[n])
			}
		}
	}

	// run tests for the upper triangular Toeplitz matrix
	for i, tt := range tests {
		mat := importUpperTriToeplitz(tt.input)
		if mat.order != tt.order {
			t.Errorf("order mismatch on UTT test %d: have %d want %d",
				i, mat.order, tt.order)
		}

		mat.ElementMul(tt.element)

		for n := 0; n < mat.order; n++ {
			if mat.get(n) != tt.expected[n] {
				t.Errorf("value mismatch on UTT test %d (col %d): have %f want %f",
					i, n, mat.get(n), tt.expected[n])
			}
		}
	}
}

func TestElementDiv(t *testing.T) {
	tests := []struct {
		order    int
		element  float64
		input    []float64
		expected []float64
	}{
		{
			order:    2,
			element:  5,
			input:    []float64{1, 2},
			expected: []float64{0.2, 0.4},
		},
		{
			order:    5,
			element:  2,
			input:    []float64{1, 2, 3, 4, 5},
			expected: []float64{0.5, 1.0, 1.5, 2.0, 2.5},
		},
	}

	// run tests for the standard matrix
	for i, tt := range tests {
		mat := importMatrix(tt.input)
		if mat.order != tt.order {
			t.Errorf("order mismatch on test %d: have %d want %d",
				i, mat.order, tt.order)
		}

		mat.ElementDiv(tt.element)

		for n := 0; n < mat.order; n++ {
			if mat.get(0, n) != tt.expected[n] {
				t.Errorf("value mismatch on standard test %d (col %d): have %f want %f",
					i, n, mat.get(0, n), tt.expected[n])
			}
		}
	}

	// run tests for the upper triangular Toeplitz matrix
	for i, tt := range tests {
		mat := importUpperTriToeplitz(tt.input)
		if mat.order != tt.order {
			t.Errorf("order mismatch on UTT test %d: have %d want %d",
				i, mat.order, tt.order)
		}

		mat.ElementDiv(tt.element)

		for n := 0; n < mat.order; n++ {
			if mat.get(n) != tt.expected[n] {
				t.Errorf("value mismatch on UTT test %d (col %d): have %f want %f",
					i, n, mat.get(n), tt.expected[n])
			}
		}
	}
}

func TestAdd(t *testing.T) {
	tests := []struct {
		order    int
		input1   []float64
		input2   []float64
		expected []float64
	}{
		{
			order:    2,
			input1:   []float64{1, 2},
			input2:   []float64{2, 4},
			expected: []float64{3.0, 6.0},
		},
		{
			order:    5,
			input1:   []float64{1, 2, 3, 4, 5},
			input2:   []float64{2, 4, 6, 8, 10},
			expected: []float64{3.0, 6.0, 9.0, 12.0, 15.0},
		},
	}

	// run tests for the standard matrix
	for i, tt := range tests {
		inp1 := importMatrix(tt.input1)
		if inp1.order != tt.order {
			t.Errorf("order mismatch on standard test %d: have %d want %d",
				i, inp1.order, tt.order)
		}

		inp2 := importMatrix(tt.input2)
		if inp2.order != tt.order {
			t.Errorf("order mismatch on standard test %d: have %d want %d",
				i, inp2.order, tt.order)
		}

		mat := inp1.Add(inp2)

		for n := 0; n < mat.order; n++ {
			if mat.get(0, n) != tt.expected[n] {
				t.Errorf("value mismatch on standard test %d (col %d): have %f want %f",
					i, n, mat.get(0, n), tt.expected[n])
			}
		}
	}

	// run tests for the upper triangular Toeplitz matrix
	for i, tt := range tests {
		inp1 := importUpperTriToeplitz(tt.input1)
		if inp1.order != tt.order {
			t.Errorf("order mismatch on UTT test %d: have %d want %d",
				i, inp1.order, tt.order)
		}

		inp2 := importUpperTriToeplitz(tt.input2)
		if inp2.order != tt.order {
			t.Errorf("order mismatch on UTT test %d: have %d want %d",
				i, inp2.order, tt.order)
		}

		mat := inp1.Add(inp2)

		for n := 0; n < mat.order; n++ {
			if mat.get(n) != tt.expected[n] {
				t.Errorf("value mismatch on UTT test %d (col %d): have %f want %f",
					i, n, mat.get(n), tt.expected[n])
			}
		}
	}
}

func TestSub(t *testing.T) {
	tests := []struct {
		order    int
		input1   []float64
		input2   []float64
		expected []float64
	}{
		{
			order:    2,
			input1:   []float64{1, 2},
			input2:   []float64{2, 4},
			expected: []float64{-1.0, -2.0},
		},
		{
			order:    5,
			input1:   []float64{1, 2, 3, 4, 5},
			input2:   []float64{2, 4, 6, 8, 10},
			expected: []float64{-1.0, -2.0, -3.0, -4.0, -5.0},
		},
	}

	// run tests for the standard matrix
	for i, tt := range tests {
		inp1 := importMatrix(tt.input1)
		if inp1.order != tt.order {
			t.Errorf("order mismatch on standard test %d: have %d want %d",
				i, inp1.order, tt.order)
		}

		inp2 := importMatrix(tt.input2)
		if inp2.order != tt.order {
			t.Errorf("order mismatch on standard test %d: have %d want %d",
				i, inp2.order, tt.order)
		}

		mat := inp1.Sub(inp2)

		for n := 0; n < mat.order; n++ {
			if mat.get(0, n) != tt.expected[n] {
				t.Errorf("value mismatch on standard test %d (col %d): have %f want %f",
					i, n, mat.get(0, n), tt.expected[n])
			}
		}
	}

	// run tests for the upper triangular Toeplitz matrix
	for i, tt := range tests {
		inp1 := importUpperTriToeplitz(tt.input1)
		if inp1.order != tt.order {
			t.Errorf("order mismatch on UTT test %d: have %d want %d",
				i, inp1.order, tt.order)
		}

		inp2 := importUpperTriToeplitz(tt.input2)
		if inp2.order != tt.order {
			t.Errorf("order mismatch on UTT test %d: have %d want %d",
				i, inp2.order, tt.order)
		}

		mat := inp1.Sub(inp2)

		for n := 0; n < mat.order; n++ {
			if mat.get(n) != tt.expected[n] {
				t.Errorf("value mismatch on UTT test %d (col %d): have %f want %f",
					i, n, mat.get(n), tt.expected[n])
			}
		}
	}
}

func TestMul(t *testing.T) {
	tests := []struct {
		order    int
		input1   []float64
		input2   []float64
		expected []float64
	}{
		{
			order:    2,
			input1:   []float64{1, 2},
			input2:   []float64{2, 4},
			expected: []float64{2.0, 8.0},
		},
		{
			order:    5,
			input1:   []float64{1, 2, 3, 4, 5},
			input2:   []float64{2, 4, 6, 8, 10},
			expected: []float64{2.0, 8.0, 20.0, 40.0, 70.0},
		},
	}

	// run tests for the standard matrix
	for i, tt := range tests {
		inp1 := importMatrix(tt.input1)
		if inp1.order != tt.order {
			t.Errorf("order mismatch on standard test %d: have %d want %d",
				i, inp1.order, tt.order)
		}

		inp2 := importMatrix(tt.input2)
		if inp2.order != tt.order {
			t.Errorf("order mismatch on standard test %d: have %d want %d",
				i, inp2.order, tt.order)
		}

		mat := inp1.Mul(inp2)

		for n := 0; n < mat.order; n++ {
			if mat.get(0, n) != tt.expected[n] {
				t.Errorf("value mismatch on standard test %d (col %d): have %f want %f",
					i, n, mat.get(0, n), tt.expected[n])
			}
		}
	}

	// run tests for the upper triangular Toeplitz matrix
	for i, tt := range tests {
		inp1 := importUpperTriToeplitz(tt.input1)
		if inp1.order != tt.order {
			t.Errorf("order mismatch on UTT test %d: have %d want %d",
				i, inp1.order, tt.order)
		}

		inp2 := importUpperTriToeplitz(tt.input2)
		if inp2.order != tt.order {
			t.Errorf("order mismatch on UTT test %d: have %d want %d",
				i, inp2.order, tt.order)
		}

		mat := inp1.Mul(inp2)

		for n := 0; n < mat.order; n++ {
			if mat.get(n) != tt.expected[n] {
				t.Errorf("value mismatch on UTT test %d (col %d): have %f want %f",
					i, n, mat.get(n), tt.expected[n])
			}
		}
	}
}

func TestDiv(t *testing.T) {
	tests := []struct {
		order    int
		input1   []float64
		input2   []float64
		expected []float64
	}{
		{
			order:    2,
			input1:   []float64{2, 4},
			input2:   []float64{1, 5},
			expected: []float64{2.0, -6.0},
		},
		{
			order:    5,
			input1:   []float64{1, 2, 3, 4, 5},
			input2:   []float64{2, 3, 10, 4, 10},
			expected: []float64{0.5, 0.25, -1.375, 1.8125, 3.65625},
		},
	}

	// run tests for the standard matrix
	for i, tt := range tests {
		inp1 := importMatrix(tt.input1)
		if inp1.order != tt.order {
			t.Errorf("order mismatch on standard test %d: have %d want %d",
				i, inp1.order, tt.order)
		}

		inp2 := importMatrix(tt.input2)
		if inp2.order != tt.order {
			t.Errorf("order mismatch on standard test %d: have %d want %d",
				i, inp2.order, tt.order)
		}

		mat := inp1.Div(inp2)

		for n := 0; n < mat.order; n++ {
			if mat.get(0, n) != tt.expected[n] {
				t.Errorf("value mismatch on standard test %d (col %d): have %f want %f",
					i, n, mat.get(0, n), tt.expected[n])
			}
		}
	}

	// run tests for the upper triangular Toeplitz matrix
	for i, tt := range tests {
		inp1 := importUpperTriToeplitz(tt.input1)
		if inp1.order != tt.order {
			t.Errorf("order mismatch on UTT test %d: have %d want %d",
				i, inp1.order, tt.order)
		}

		inp2 := importUpperTriToeplitz(tt.input2)
		if inp2.order != tt.order {
			t.Errorf("order mismatch on UTT test %d: have %d want %d",
				i, inp2.order, tt.order)
		}

		mat := inp1.Div(inp2)

		for n := 0; n < mat.order; n++ {
			if mat.get(n) != tt.expected[n] {
				t.Errorf("value mismatch on UTT test %d (col %d): have %f want %f",
					i, n, mat.get(n), tt.expected[n])
			}
		}
	}
}
