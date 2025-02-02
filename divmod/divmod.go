package divmod

import (
	"github.com/borisskert/go-biginteger/digits"
	"github.com/borisskert/go-biginteger/divmod/algorithms/recursive"
)

var recursiveDivisionAlgorithm = recursive.NewRecursiveDivisionAlgorithm()

func DivMod(numerator, denominator []uint64) ([]uint64, []uint64) {
	wrappedNumerator := digits.Wrap(numerator)
	wrappedDenominator := digits.Wrap(denominator)

	quotient, remainder := recursiveDivisionAlgorithm.DivMod(
		wrappedNumerator,
		wrappedDenominator,
	)

	return quotient.Trim().AsArray(), remainder.Trim().AsArray()
}
