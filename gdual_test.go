package gdual

import (
	"testing"
)

func TestSimple(t *testing.T) {
	expected := []float64{
		16.0, 32.0, 24.0, 8.0,
	}

	order := 4
	inp := 2.0

	// f(2.0) = x^4
	x := NewGDual(order, inp, true)
	y := x.Pow(4)

	for i := 0; i < len(expected); i++ {
		if y.mat.get(i) != expected[i] {
			t.Errorf("failed on simple test (iter %d): have %.2f want %.2f",
				i, y.mat.get(i), expected[i])
		}
	}
}

func TestComplex(t *testing.T) {
	expected := []float64{
		-4.5, 3.75, -2.75, 1.875,
		-1.21875, 0.765625, -0.46875,
		0.28125, -0.166015625,
	}

	order := 10
	inp := 3.0

	x := NewGDual(order, inp, true)
	one := NewGDual(order, 1.0, false)
	four := NewGDual(order, 4.0, false)

	// f(3.0) = 4x^2 / (1 - x)^3
	y := x.Pow(2).Mul(four).Div(one.Sub(x).Pow(3))

	for i := 0; i < len(expected); i++ {
		if y.mat.get(i) != expected[i] {
			t.Errorf("failed on simple test (iter %d): have %.2f want %.2f",
				i, y.mat.get(i), expected[i])
		}
	}
}
