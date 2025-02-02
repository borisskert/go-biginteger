package trailingZeroReduction

import (
	"github.com/borisskert/go-biginteger/digits"
	"github.com/borisskert/go-biginteger/divmod/common"
)

type trailingZeroReduction struct {
	algorithm common.DivisionAlgorithm
}

func (t trailingZeroReduction) DivMod(numerator digits.Digits, denominator digits.Digits) (digits.Digits, digits.Digits) {
	return t.reduceTrailingZeroBits(numerator, denominator)
}

func (t trailingZeroReduction) reduceTrailingZeroBits(
	numerator digits.Digits,
	denominator digits.Digits,
) (digits.Digits, digits.Digits) {
	length := numerator.Length()

	if length >= 6 {
		trailingZeroBits := min(numerator.TrailingZeros(), denominator.TrailingZeros())

		if trailingZeroBits > 0 {
			// Right shift both numbers to remove common trailing zeros
			nReduced := numerator.RightShiftBits(trailingZeroBits)
			dReduced := denominator.RightShiftBits(trailingZeroBits)

			// Perform division on the reduced numbers
			q, r := t.algorithm.DivMod(nReduced, dReduced)

			// Ensure remainder is still valid
			r = r.LeftShiftBits(trailingZeroBits)

			return q, r
		}
	}

	return t.algorithm.DivMod(numerator, denominator)
}

func DecorateWithTrailingZeroReduction(algorithm common.DivisionAlgorithm) common.DivisionAlgorithm {
	return &trailingZeroReduction{
		algorithm: algorithm,
	}
}
