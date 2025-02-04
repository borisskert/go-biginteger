package multiply

import (
	"github.com/borisskert/go-biginteger/digits"
	"math/bits"
)

func schoolbookMultiply(a, b digits.Digits) digits.Digits {
	array := multiplyUint64Array(a.AsArray(), b.AsArray())
	result := digits.Wrap(array)

	if a.IsNegative() != b.IsNegative() {
		return result.Negate()
	}

	return result
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

	return result
}
