package multiplyTwoByOne

import (
	"github.com/borisskert/go-biginteger/digits"
	"github.com/borisskert/go-biginteger/multiply/api"
)

type multiplyTwoByOne struct {
}

func (m multiplyTwoByOne) Multiply(
	multiplicand digits.Digits, multiplier digits.Digits,
) (product digits.Digits) {
	lo, hi := multiplicand.AsDoubleDigit().MultiplyDigit(multiplier.AsDigit())

	result := digits.OfUint64Array([]uint64{
		uint64(lo.Low()),
		uint64(lo.High()),
		uint64(hi),
	})

	return result.
		Sign(multiplicand.IsNegative() != multiplier.IsNegative()).
		Trim()
}

func NewMultiplyTwoByOne() api.MultiplyAlgorithm {
	return multiplyTwoByOne{}
}
