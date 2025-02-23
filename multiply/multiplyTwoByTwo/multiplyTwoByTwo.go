package multiplyTwoByTwo

import (
	"github.com/borisskert/go-biginteger/digits"
	"github.com/borisskert/go-biginteger/multiply/api"
)

type multiplyTwoByTwoAlgorithm struct {
}

func (m multiplyTwoByTwoAlgorithm) Multiply(
	multiplicand digits.Digits, multiplier digits.Digits,
) (product digits.Digits) {
	return multiplyTwoByTwoAbs(
		multiplicand.AsDoubleDigit(),
		multiplier.AsDoubleDigit(),
	).
		Sign(multiplicand.IsNegative() != multiplier.IsNegative()).
		Trim()
}

func multiplyTwoByTwoAbs(a, b digits.DoubleDigit) digits.Digits {
	hi, lo := a.Multiply(b)

	result := digits.OfUint64Array([]uint64{
		uint64(lo.Low()),
		uint64(lo.High()),
		uint64(hi.Low()),
		uint64(hi.High()),
	})

	return result
}

func NewMultiplyTwoByTwo() api.MultiplyAlgorithm {
	return &multiplyTwoByTwoAlgorithm{}
}
