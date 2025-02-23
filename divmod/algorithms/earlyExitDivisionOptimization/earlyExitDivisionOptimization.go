package earlyExitDivisionOptimization

import (
	"github.com/borisskert/go-biginteger/digits"
	"github.com/borisskert/go-biginteger/divmod/api"
)

type earlyExitDivisionOptimization struct {
	algorithm api.DivisionAlgorithm
}

func (d *earlyExitDivisionOptimization) DivMod(
	numerator digits.Digits,
	denominator digits.Digits,
) (digits.Digits, digits.Digits) {
	return earlyExitDivision(numerator, denominator, d.algorithm.DivMod)
}

func earlyExitDivision(
	numerator digits.Digits,
	denominator digits.Digits,
	fn func(digits.Digits, digits.Digits) (digits.Digits, digits.Digits),
) (digits.Digits, digits.Digits) {
	if denominator.IsZero() {
		panic("Division by zero")
	}

	if numerator.IsZero() {
		return digits.Zero().AsDigits(), digits.Zero().AsDigits()
	}

	if denominator.IsOne() {
		return numerator, digits.Zero().AsDigits()
	}

	cmp := numerator.Compare(denominator)

	if cmp < 0 {
		return digits.Zero().AsDigits(), numerator
	}

	if cmp == 0 {
		return digits.One().AsDigits(), digits.Zero().AsDigits()
	}

	return fn(numerator, denominator)
}

func DecorateWithEarlyExitDivisionOptimization(algorithm api.DivisionAlgorithm) api.DivisionAlgorithm {
	return &earlyExitDivisionOptimization{
		algorithm: algorithm,
	}
}
