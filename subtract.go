package biginteger

import (
	"github.com/borisskert/go-biginteger/divmod"
	"math/bits"
)

func subtract(minuend BigInteger, subtrahend BigInteger) BigInteger {
	if minuend.IsEqualTo(subtrahend) {
		return zero
	}

	if !minuend.sign && !subtrahend.sign {
		if minuend.IsLessThan(subtrahend) {
			result := divmod.Subtract(subtrahend.value, minuend.value)

			return BigInteger{
				true,
				result,
			}
		}

		result := divmod.Subtract(minuend.value, subtrahend.value)

		return BigInteger{
			false,
			result,
		}
	}

	if minuend.sign && subtrahend.sign {
		result := divmod.Add(minuend.value, subtrahend.value)

		return BigInteger{
			true,
			result,
		}
	}

	if minuend.sign {
		return subtrahend.Add(minuend.Abs())
	}

	return minuend.Add(subtrahend.Abs())
}

func subtractUint64Arrays(a, b []uint64) []uint64 {
	if len(b) > len(a) {
		a, b = b, a
	}

	result := make([]uint64, len(a))
	carry := uint64(0)

	for i := 0; i < len(a); i++ {
		ai := a[i]
		bi := uint64(0)
		if i < len(b) {
			bi = b[i]
		}

		diff, borrow := bits.Sub64(ai, bi+carry, 0)
		result[i] = diff
		carry = borrow
	}

	return trimLeadingZeros(result)
}
