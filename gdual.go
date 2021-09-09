package gdual

type GDual struct {
	mat      *UpperTriToeplitz
	variable bool
}

func NewGDual(order int, seed float64, variable bool) *GDual {
	mat := NewUpperTriToeplitz(order)
	mat.Fill(0, seed)
	if variable {
		mat.Fill(1, 1.0)
	}

	gdual := &GDual{
		mat:      mat,
		variable: variable,
	}

	return gdual
}

func importGDual(mat *UpperTriToeplitz, variable bool) *GDual {
	gdual := &GDual{
		mat:      mat,
		variable: variable,
	}

	return gdual
}

func (g *GDual) Add(inp *GDual) *GDual {
	mat := g.mat.Add(inp.mat)
	gdual := importGDual(mat, g.variable || inp.variable)

	return gdual
}

func (g *GDual) Sub(inp *GDual) *GDual {
	mat := g.mat.Sub(inp.mat)
	gdual := importGDual(mat, g.variable || inp.variable)

	return gdual
}

func (g *GDual) Mul(inp *GDual) *GDual {
	mat := g.mat.Mul(inp.mat)
	gdual := importGDual(mat, g.variable || inp.variable)

	return gdual
}

func (g *GDual) Div(inp *GDual) *GDual {
	mat := g.mat.Div(inp.mat)
	gdual := importGDual(mat, g.variable || inp.variable)

	return gdual
}

func (g *GDual) Pow(n int) *GDual {
	mat := g.mat.Pow(n)
	gdual := importGDual(mat, g.variable)

	return gdual
}
