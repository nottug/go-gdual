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
([Audi](github.com/darioizzo/audi/)), which takes it a step further
implementing the algebra of truncated polynomials. 

This library is meant to be a starting point for generalized dual
numbers, mostly out of curiousity. It will probably end up moving
towards Audi at some point to implement [DCGP](https://arxiv.org/pdf/1611.04766.pdf).

The matrix stuff could probably be optimized a lot more, but since
we can use so many shortcuts from the assurances dual numbers give us,
I'll avoid using something like [gonum](https://github.com/gonum/gonum) and
maybe one day try to implement some BLAS support.


# References

 - [Wikipedia: Dual Number](https://en.wikipedia.org/wiki/Dual_number)
 - [jliszka: Exact numeric nth derivatives](http://blog.jliszka.org/2013/10/24/exact-numeric-nth-derivatives.html)
 - [Demofox: Dual Numbers and Automatic Differentiation](https://blog.demofox.org/2014/12/30/dual-numbers-automatic-differentiation/)
 - [Demofox: Multivariate Dual Numbers and Automatic Differentiation](https://blog.demofox.org/2017/02/20/multivariable-dual-numbers-automatic-differentiation/)
 - [Dario Izzo: Audi - Truncated Polynomial Algebra Implementation](https://darioizzo.github.io/audi/notebooks/example00)
