package divmod

import (
	"github.com/borisskert/go-biginteger/digits"
	"github.com/borisskert/go-biginteger/divmod/algorithms/recursive"
)

var divisionAlgorithmFactory = recursive.NewRecursiveDivisionAlgorithm()

func DivMod(numerator, denominator []uint64) ([]uint64, []uint64) {
	wrappedNumerator := digits.Wrap(numerator)
	wrappedDenominator := digits.Wrap(denominator)

	quotient, remainder := divisionAlgorithmFactory.DivMod(
		wrappedNumerator,
		wrappedDenominator,
	)

	return quotient.Trim().AsArray(), remainder.Trim().AsArray()
}

type divmodFn func(digits.Digits, digits.Digits) (digits.Digits, digits.Digits)
