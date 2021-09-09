# Generalized Dual Numbers

Dual numbers are a strange part of math that allow you
to calculate derivatives at a point for a given function.
The references below give far better explanations than I 
could here.

While the single variable, single order implementation of
dual numbers is trivial, increasing to multivariate, nth order
dual numbers rises in complexity quickly. It isn't very hard
to find good examples of multivariate or nth order dual numbers,
but together there is really only a single implementation 
([Audi](https://github.com/darioizzo/audi/)), which takes it a step further
implementing the algebra of truncated polynomials. 

This library is meant to be a starting point for generalized dual
numbers, mostly out of curiousity. It will probably end up moving
towards Audi at some point to implement [DCGP](https://arxiv.org/pdf/1611.04766.pdf).

# Usage

The library is still in progress, but basic functionality works like this:

```go
x := NewGDual(5, 2.0, true)

// f(2.0) = x^2
y := x.Pow(2)
```

Constant terms and variables are handled differently so it's important to differentiate
between the two. An example with a constant would be:

```go
x := NewGDual(5, 2.0, true)
four := NewGDual(5, 4.0, false)

// f(2.0) = 4*x^2
y := x.Pow(2).Mul(four)
```

# Implementation

When using matrices for generalized dual numbers, we have the assurance that
each matrix will be of a specific form: a square, upper triangular
[Toeplitz](https://en.wikipedia.org/wiki/Toeplitz_matrix) matrix. This primarily
means that we only have to store a 1D array instead of a 2D matrix. For example,

```
1 2 3 4 5
0 1 2 3 4
0 0 1 2 3       =>      1 2 3 4 5
0 0 0 1 2
0 0 0 0 1
```

By doing this, we reduce our memory footprint by a factor of `n-1`, where `n`
is the order of the matrix. This also reduces all element-wise operations to
`O(n)` instead of `O(n^2)`, along with matrix addition and subtraction. Matrix
multiplication goes from `O(n^3)` to something like `O(nlogn)`. Matrix division 
requires the inverse and multiplication, but goes from something like `O(n^3 + n^5)` 
to something like `O(nlogn + n^3logn)` (maybe?).

# Performance

In terms of performance, our Toeplitz matrix performs somewhere around `2.5 * n` times 
faster than the standard, where `n` is the order of the matrix. We're only benchmarking
performance, not memory load, but the results for the four core matrix functions
are pretty clear:

```
goos: darwin
goarch: amd64
pkg: github.com/sencha-dev/go-gdual
cpu: Intel(R) Core(TM) i5-1038NG7 CPU @ 2.00GHz

BenchmarkTestStandardAdd10-8          602390          1847 ns/op
BenchmarkTestStandardAdd100-8          12966         92756 ns/op
BenchmarkTestStandardSub10-8          657568          1866 ns/op
BenchmarkTestStandardSub100-8          12321         96109 ns/op
BenchmarkTestStandardMul10-8          243154          4514 ns/op
BenchmarkTestStandardMul100-8            561       2103823 ns/op
BenchmarkTestStandardDiv10-8           25075         47598 ns/op
BenchmarkTestStandardDiv100-8              5     207308988 ns/op

BenchmarkTestToeplitzAdd10-8        14504365            81.88 ns/op
BenchmarkTestToeplitzAdd100-8        3113448           383.5 ns/op
BenchmarkTestToeplitzSub10-8        14486222            81.13 ns/op
BenchmarkTestToeplitzSub100-8        3134450           382.3 ns/op
BenchmarkTestToeplitzMul10-8         7004923           169.3 ns/op
BenchmarkTestToeplitzMul100-8         142678          8302 ns/op
BenchmarkTestToeplitzDiv10-8          421898          2759 ns/op
BenchmarkTestToeplitzDiv100-8           1358        872147 ns/op
PASS
ok      github.com/sencha-dev/go-gdual  23.948s
```

For further optimization, there is the possibility the multiplication and
inverting functions for Toeplitz could be improved. For multiplication, 
there are a few approaches using FFT (see 
[scipy: PR #11346](https://github.com/scipy/scipy/pull/11346)), but since
we know its also square and upper triangular, our approach may already have
equivalent performance to these methods (which are `O(nlogn)` time too).
For inversion, we already are using a few shortcuts, but since that is
the most expensive operation it may be nice to see if something like
[this](https://math.stackexchange.com/a/786200/798807) could help.

# TODO

 - [ ] Clean up matrix implementations, probably make an interface
 - [ ] Clean up interaction pattern between gdual and matrix
 - [ ] Add lazy evaluation (and possible simplification/optimization)
 - [ ] Implement partials and total derivative
 - [ ] Implement special functions like `exp, log, power, sin, cos, tan`
 - [ ] Make a better mechanism for defining variables, constants, and expressions

# References

 - [Wikipedia: Dual Number](https://en.wikipedia.org/wiki/Dual_number)
 - [jliszka: Exact numeric nth derivatives](http://blog.jliszka.org/2013/10/24/exact-numeric-nth-derivatives.html)
 - [Demofox: Dual Numbers and Automatic Differentiation](https://blog.demofox.org/2014/12/30/dual-numbers-automatic-differentiation/)
 - [Demofox: Multivariate Dual Numbers and Automatic Differentiation](https://blog.demofox.org/2017/02/20/multivariable-dual-numbers-automatic-differentiation/)
 - [Dario Izzo: Audi - Truncated Polynomial Algebra Implementation](https://darioizzo.github.io/audi/notebooks/example00)
