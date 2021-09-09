package gdual

type GDual struct {
	mat      *Matrix
	variable bool
}

func NewGDual(order int, variable bool) *GDual {
	mat := NewMatrix(order)
	gdual := &GDual{
		mat:      mat,
		variable: variable,
	}

	return gdual
}
