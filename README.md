Gomega Matcher for gonum's mat64
==============================

This package aims to provide customized [Gomega](http://onsi.github.io/gomega) matchers for gonum's [mat64](https://godoc.org/github.com/gonum/matrix/mat64).

## Usage

### AllCloseTo Matcher

This matcher compare two `*mat64.Dense`, with absolute/relative error. The matcher's signature:

```go
AllCloseTo(expected *mat64.Dense, tol float64, relative bool)
```

Used with [Ginkgo](http://onsi.github.io/ginkgo) and Gomega:

```go
pt3 := mat64.NewDense(3, 1, []float64{1, 3, -2})

Ω(pt3).Should(gomegamat64.AllCloseTo(
        mat64.NewDense(3, 1, []float64{1, 3, -3}),
        1e-6,
        false,
))
```

And if the matrix doesn't match, it will raise the error (row, column):

```
Expected -2 to close to -3: pos=(2, 0), tol=1e-06, relative=false
```
