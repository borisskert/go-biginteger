package multiplyOneByOne

import (
	"github.com/borisskert/go-biginteger/digits"
	"github.com/borisskert/go-biginteger/multiply/api"
)

type multiplyOneByOneAlgorithm struct {
}

func (m multiplyOneByOneAlgorithm) Multiply(
	multiplicand digits.Digits, multiplier digits.Digits,
) (product digits.Digits) {
	return multiplicand.AsDigit().
		Multiply(multiplier.AsDigit()).
		AsDigits().
		Sign(multiplicand.IsNegative() != multiplier.IsNegative())
}

func NewMultiplyOneByOneAlgorithm() api.MultiplyAlgorithm {
	return &multiplyOneByOneAlgorithm{}
}
