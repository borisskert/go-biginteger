package earlyExit

import (
	"github.com/borisskert/go-biginteger/digits"
	"github.com/borisskert/go-biginteger/multiply/api"
)

type earlyExitAlgorithm struct {
	multiply api.MultiplyAlgorithm
}

func (e earlyExitAlgorithm) Multiply(
	multiplicand digits.Digits, multiplier digits.Digits,
) (product digits.Digits) {
	if multiplicand.IsZero() || multiplier.IsZero() {
		return digits.Zero().AsDigits()
	}

	if multiplicand.IsOne() {
		return multiplier
	}

	if multiplier.IsOne() {
		return multiplicand
	}

	if multiplier.Abs().IsOne() {
		return multiplicand.Negate()
	}

	if multiplicand.Abs().IsOne() {
		return multiplier.Negate()
	}

	return e.multiply.Multiply(multiplicand, multiplier)
}

func DecorateWithEarlyExit(algorithm api.MultiplyAlgorithm) api.MultiplyAlgorithm {
	return &earlyExitAlgorithm{
		multiply: algorithm,
	}
}
