package gdual

import (
	"testing"
)

func TestSimple(t *testing.T) {
	expected := []float64{16, 32, 24, 8}

	order := 4

	inp := NewMatrix(order)
	inp.Fill(0, 2.0)
	inp.Fill(1, 1.0)

	// f(x) = x^4
	out := inp.Pow(4)

	for i := 0; i < len(expected); i++ {
		if out.get(0, i) != expected[i] {
			t.Errorf("failed on simple test (iter %d): have %.2f want %.2f",
				i, out.get(0, i), expected[i])
		}
	}
}

func TestComplex(t *testing.T) {
	expected := []float64{
		-4.5,
		3.75,
		-2.75,
		1.875,
		-1.21875,
		0.765625,
		-0.46875,
		0.28125,
		-0.166015625,
	}

	order := 10

	inp := NewMatrix(order)
	inp.Fill(0, 3.0)
	inp.Fill(1, 1.0)

	one := NewMatrix(order)
	one.Fill(0, 1.0)

	four := NewMatrix(order)
	four.Fill(0, 4.0)

	numerator := inp.Pow(2).Mul(four)
	denominator := one.Sub(inp).Pow(3)

	// f(x) = 4x^2 / (1 - x)^3
	out := numerator.Div(denominator)

	for i := 0; i < len(expected); i++ {
		if out.get(0, i) != expected[i] {
			t.Errorf("failed on simple test (iter %d): have %.2f want %.2f",
				i, out.get(0, i), expected[i])
		}
	}
}
