package biginteger

import (
	mul "github.com/borisskert/go-biginteger/multiply"
	"math/bits"
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

func multiplyUint64Array(a, b []uint64) []uint64 {
	result := make([]uint64, len(a)+len(b))

	for i := 0; i < len(a); i++ {
		var carry uint64 = 0

		for j := 0; j < len(b); j++ {
			high, low := bits.Mul64(a[i], b[j])
			low += result[i+j]

			if low < result[i+j] {
				high++
			}

			result[i+j] = low

			high += carry
			result[i+j+1] += high

			if result[i+j+1] < high {
				carry = 1
			} else {
				carry = 0
			}
		}

		result[i+len(b)] += carry
	}

	return trimLeadingZeros(result)
}
