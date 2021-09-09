/*

this is a library for upper triangular matrices with the a non-zero, constant
term across the full main diagonal. filling is based off of upper diagonals,
the determinant assumes the matrix is an upper triangular matrix, and the
inverse function assumes the matrix is an upper triangular matrix with
a non-zero, constant term across the main diagonal. this library is specifically
designed for (generalized) dual numbers.

*/

package gdual

type Matrix struct {
	order int
	val   [][]float64
}

func NewMatrix(order int) *Matrix {
	val := make([][]float64, order)
	for i := 0; i < order; i++ {
		val[i] = make([]float64, order)
	}

	mat := &Matrix{
		order: order,
		val:   val,
	}

	return mat
}

/* utility functions */

func (m *Matrix) get(i, j int) float64 {
	if i > m.order || j > m.order {
		return 0.0
	}

	return m.val[i][j]
}

func (m *Matrix) set(i, j int, val float64) {
	if i > m.order || j > m.order {
		return
	}

	m.val[i][j] = val
}

func (m *Matrix) Fill(diagonal int, val float64) {
	// fill the given upper diagonal of the matrix
	for i := diagonal; i < m.order; i++ {
		m.set(i-diagonal, i, val)
	}
}

func (m *Matrix) Reset(val float64) {
	for i := 0; i < m.order; i++ {
		for j := 0; j < m.order; j++ {
			m.set(i, j, val)
		}
	}
}

func (m *Matrix) Copy() *Matrix {
	copy := NewMatrix(m.order)

	for i := 0; i < m.order; i++ {
		for j := 0; j < m.order; j++ {
			val := m.get(i, j)
			copy.set(i, j, val)
		}
	}

	return copy
}

/* general matrix operations */

func (m *Matrix) Determinant() float64 {
	det := 1.0
	for i := 0; i < m.order; i++ {
		det *= m.get(i, i)
	}

	return det
}

/* element-wise matrix operations */

func (m *Matrix) ElementAdd(val float64) {
	for i := 0; i < m.order; i++ {
		for j := 0; j < m.order; j++ {
			sum := m.get(i, j) + val
			m.set(i, j, sum)
		}
	}
}

func (m *Matrix) ElementSub(val float64) {
	for i := 0; i < m.order; i++ {
		for j := 0; j < m.order; j++ {
			difference := m.get(i, j) - val
			m.set(i, j, difference)
		}
	}
}

func (m *Matrix) ElementMul(val float64) {
	for i := 0; i < m.order; i++ {
		for j := 0; j < m.order; j++ {
			product := m.get(i, j) * val
			m.set(i, j, product)
		}
	}
}

func (m *Matrix) ElementDiv(val float64) {
	for i := 0; i < m.order; i++ {
		for j := 0; j < m.order; j++ {
			quotient := m.get(i, j) / val
			m.set(i, j, quotient)
		}
	}
}

/* matrix operations */

func (m *Matrix) Add(inp *Matrix) *Matrix {
	out := NewMatrix(m.order)

	for i := 0; i < m.order; i++ {
		for j := 0; j < m.order; j++ {
			sum := m.get(i, j) + inp.get(i, j)
			out.set(i, j, sum)
		}
	}

	return out
}

func (m *Matrix) Sub(inp *Matrix) *Matrix {
	out := NewMatrix(m.order)

	for i := 0; i < m.order; i++ {
		for j := 0; j < m.order; j++ {
			difference := m.get(i, j) - inp.get(i, j)
			out.set(i, j, difference)
		}
	}

	return out
}

func (m *Matrix) Mul(inp *Matrix) *Matrix {
	out := NewMatrix(m.order)

	for i := 0; i < m.order; i++ {
		for j := 0; j < m.order; j++ {
			product := 0.0
			for k := 0; k < m.order; k++ {
				product += m.get(i, k) * inp.get(k, j)
			}

			out.set(i, j, product)
		}
	}

	return out
}

/*
shortcut borrowed from
blog.jliszka.org/2013/10/24/exact-numeric-nth-derivatives.html

a nilpotent matrix follows the form:
(I - N)^-1 = Σ(I + N + N^2 + ... N^n-1)

since the square upper triangular matrices we're using
follow the form a*I + D, where a is some constant and D
is a nilpotent matrix, we can use the above equation
to simplify our inverse calculation.

we calculate D by subtracting our input matrix by a*I.

in order to find (a*I + D)^-1, the formula is:
(a*I + D)^-1 = (1 / a*I) * Σ(I + N + N^2 + ... N^n-1)
*/
func (m *Matrix) Inv() *Matrix {
	// derive a*I
	a := m.get(0, 0)
	A := NewMatrix(m.order)
	A.Fill(0, a)

	// derive the nilpotent matrix D
	D := m.Sub(A)

	// derive N from D and a
	N := D
	N.ElementDiv(-a)

	// initialize the identity matrix
	inv := NewMatrix(m.order)
	inv.Fill(0, 1.0)

	firstN := N.Copy()
	for i := 0; i < m.order; i++ {
		inv = inv.Add(N)
		if i != m.order-1 {
			N = firstN.Mul(N)
		}
	}

	// divide by a to find the true inverse
	inv.ElementDiv(a)

	return inv
}

func (m *Matrix) Div(inp *Matrix) *Matrix {
	inv := inp.Inv()
	out := m.Mul(inv)

	return out
}

func (m *Matrix) Pow(n int) *Matrix {
	out := m.Copy()
	for i := 0; i < n-1; i++ {
		out = out.Mul(m)
	}

	return out
}
