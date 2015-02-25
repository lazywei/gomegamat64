package gomegamat64

import (
	"fmt"
	"math"

	"github.com/gonum/matrix/mat64"
	"github.com/onsi/gomega/types"
)

type mat64Matcher struct {
	expected *mat64.Dense
	errRow   int
	errCol   int

	expectedVal float64
	actualVal   float64

	tol      float64
	relative bool
}

func AllCloseTo(expected *mat64.Dense, tol float64, relative bool) types.GomegaMatcher {
	return &mat64Matcher{
		expected: expected,
		tol:      tol,
		relative: relative,
	}
}

func (matcher *mat64Matcher) Match(actual interface{}) (success bool, err error) {
	mat, ok := actual.(*mat64.Dense)
	if !ok {
		return false, fmt.Errorf("AllCloseTo matcher expects an *mat64.Dense")
	}

	eRows, eCols := matcher.expected.Dims()
	aRows, aCols := mat.Dims()

	if eRows != aRows || eCols != aCols {
		return false, fmt.Errorf(
			"The dimensions of expected (%s, %s) and actual (%s, %s) don't match",
			eRows, eCols, aRows, aCols)
	}

	for i := 0; i < eRows; i++ {
		for j := 0; j < eCols; j++ {

			residual := math.Abs(matcher.expected.At(i, j) - mat.At(i, j))

			if matcher.relative {
				// Use relative error
				residual = math.Abs(residual / mat.At(i, j))
			}

			if residual >= matcher.tol {
				matcher.errRow = i
				matcher.errCol = j
				matcher.expectedVal = matcher.expected.At(i, j)
				matcher.actualVal = mat.At(i, j)
				return false, nil
			}

		}
	}

	return true, nil
}

func (matcher *mat64Matcher) FailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected %v to close to %v: pos=(%v, %v), tol=%v, relative=%v",
		matcher.actualVal, matcher.expectedVal,
		matcher.errRow, matcher.errCol,
		matcher.tol, matcher.relative)
}

func (matcher *mat64Matcher) NegatedFailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected %v to close to %v: pos=(%v, %v), tol=%v, relative=%v",
		matcher.actualVal, matcher.expectedVal,
		matcher.errRow, matcher.errCol,
		matcher.tol, matcher.relative)
}
