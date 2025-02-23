package biginteger

import (
	mul "github.com/borisskert/go-biginteger/multiply"
)

func multiply(multiplicand BigInteger, multiplier BigInteger) BigInteger {
	if multiplicand.IsEqualTo(zero) || multiplier.IsEqualTo(zero) {
		return zero
	}

	if multiplicand.IsEqualTo(one) {
		return multiplier
	}

	if multiplier.IsEqualTo(one) {
		return multiplicand
	}

	sign := multiplicand.sign != multiplier.sign
	product := mul.Multiply(multiplicand.value, multiplier.value)

	return BigInteger{
		sign:  sign,
		value: product,
	}
}
