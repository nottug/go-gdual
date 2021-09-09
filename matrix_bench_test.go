package gdual

import (
	"math/rand"
	"testing"
)

const (
	minBound = -10000
	maxBound = 10000
)

var standardMat *Matrix
var toeplitzMat *UpperTriToeplitz

/* utils */

func randFloats(min, max float64, n int) []float64 {
	res := make([]float64, n)
	for i := range res {
		res[i] = min + rand.Float64()*(max-min)
	}
	return res
}

/* base functions */

func benchmarkStandardAdd(order int, b *testing.B) {
	var mat *Matrix
	input1 := randFloats(minBound, maxBound, order)
	input2 := randFloats(minBound, maxBound, order)

	for i := 0; i < b.N; i++ {
		inp1 := importMatrix(input1)
		inp2 := importMatrix(input2)
		mat = inp1.Add(inp2)
	}

	standardMat = mat
}

func benchmarkStandardSub(order int, b *testing.B) {
	var mat *Matrix
	input1 := randFloats(minBound, maxBound, order)
	input2 := randFloats(minBound, maxBound, order)

	for i := 0; i < b.N; i++ {
		inp1 := importMatrix(input1)
		inp2 := importMatrix(input2)
		mat = inp1.Sub(inp2)
	}

	standardMat = mat
}

func benchmarkStandardMul(order int, b *testing.B) {
	var mat *Matrix
	input1 := randFloats(minBound, maxBound, order)
	input2 := randFloats(minBound, maxBound, order)

	for i := 0; i < b.N; i++ {
		inp1 := importMatrix(input1)
		inp2 := importMatrix(input2)
		mat = inp1.Mul(inp2)
	}

	standardMat = mat
}

func benchmarkStandardDiv(order int, b *testing.B) {
	var mat *Matrix
	input1 := randFloats(minBound, maxBound, order)
	input2 := randFloats(minBound, maxBound, order)

	for i := 0; i < b.N; i++ {
		inp1 := importMatrix(input1)
		inp2 := importMatrix(input2)
		mat = inp1.Div(inp2)
	}

	standardMat = mat
}

func benchmarkToeplitzAdd(order int, b *testing.B) {
	var mat *UpperTriToeplitz
	input1 := randFloats(minBound, maxBound, order)
	input2 := randFloats(minBound, maxBound, order)

	for i := 0; i < b.N; i++ {
		inp1 := importUpperTriToeplitz(input1)
		inp2 := importUpperTriToeplitz(input2)
		mat = inp1.Add(inp2)
	}

	toeplitzMat = mat
}

func benchmarkToeplitzSub(order int, b *testing.B) {
	var mat *UpperTriToeplitz
	input1 := randFloats(minBound, maxBound, order)
	input2 := randFloats(minBound, maxBound, order)

	for i := 0; i < b.N; i++ {
		inp1 := importUpperTriToeplitz(input1)
		inp2 := importUpperTriToeplitz(input2)
		mat = inp1.Sub(inp2)
	}

	toeplitzMat = mat
}

func benchmarkToeplitzMul(order int, b *testing.B) {
	var mat *UpperTriToeplitz
	input1 := randFloats(minBound, maxBound, order)
	input2 := randFloats(minBound, maxBound, order)

	for i := 0; i < b.N; i++ {
		inp1 := importUpperTriToeplitz(input1)
		inp2 := importUpperTriToeplitz(input2)
		mat = inp1.Mul(inp2)
	}

	toeplitzMat = mat
}

func benchmarkToeplitzDiv(order int, b *testing.B) {
	var mat *UpperTriToeplitz
	input1 := randFloats(minBound, maxBound, order)
	input2 := randFloats(minBound, maxBound, order)

	for i := 0; i < b.N; i++ {
		inp1 := importUpperTriToeplitz(input1)
		inp2 := importUpperTriToeplitz(input2)
		mat = inp1.Div(inp2)
	}

	toeplitzMat = mat
}

/* standard benchmarks */

func BenchmarkTestStandardAdd10(b *testing.B) {
	benchmarkStandardAdd(10, b)
}

func BenchmarkTestStandardAdd100(b *testing.B) {
	benchmarkStandardAdd(100, b)
}

func BenchmarkTestStandardSub10(b *testing.B) {
	benchmarkStandardSub(10, b)
}

func BenchmarkTestStandardSub100(b *testing.B) {
	benchmarkStandardSub(100, b)
}

func BenchmarkTestStandardMul10(b *testing.B) {
	benchmarkStandardMul(10, b)
}

func BenchmarkTestStandardMul100(b *testing.B) {
	benchmarkStandardMul(100, b)
}

func BenchmarkTestStandardDiv10(b *testing.B) {
	benchmarkStandardDiv(10, b)
}

func BenchmarkTestStandardDiv100(b *testing.B) {
	benchmarkStandardDiv(100, b)
}

/* Toeplitz benchmarks */

func BenchmarkTestToeplitzAdd10(b *testing.B) {
	benchmarkToeplitzAdd(10, b)
}

func BenchmarkTestToeplitzAdd100(b *testing.B) {
	benchmarkToeplitzAdd(100, b)
}

func BenchmarkTestToeplitzSub10(b *testing.B) {
	benchmarkToeplitzSub(10, b)
}

func BenchmarkTestToeplitzSub100(b *testing.B) {
	benchmarkToeplitzSub(100, b)
}

func BenchmarkTestToeplitzMul10(b *testing.B) {
	benchmarkToeplitzMul(10, b)
}

func BenchmarkTestToeplitzMul100(b *testing.B) {
	benchmarkToeplitzMul(100, b)
}

func BenchmarkTestToeplitzDiv10(b *testing.B) {
	benchmarkToeplitzDiv(10, b)
}

func BenchmarkTestToeplitzDiv100(b *testing.B) {
	benchmarkToeplitzDiv(100, b)
}
