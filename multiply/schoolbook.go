package multiply

import (
	"github.com/borisskert/go-biginteger/digits"
	"math/bits"
)

func SchoolbookMultiply(a, b digits.Digits) digits.Digits {
	result := multiplyByDigits(a, b)

	if a.IsNegative() != b.IsNegative() && !result.IsZero() {
		return result.Trim().Negative()
	}

	return result.Trim()
}

func multiplyByDigits(a digits.Digits, b digits.Digits) digits.Digits {
	if a.IsZero() || b.IsZero() {
		return digits.Zero().AsDigits()
	}

	if a.IsOne() {
		return b
	}

	if b.IsOne() {
		return a
	}

	result := make([]uint64, a.Length()+b.Length())

	for i := uint(0); i < a.Length(); i++ {
		ai := a.DigitAt(i)
		var carry uint64 = 0

		for j := uint(0); j < b.Length(); j++ {
			bi := b.DigitAt(j)

			hi, lo := bits.Mul64(uint64(ai), uint64(bi))

			sumLo, carryLo := bits.Add64(result[i+j], lo, 0)
			result[i+j] = sumLo

			sumMid, carryMid := bits.Add64(result[i+j+1], hi, carryLo)
			result[i+j+1] = sumMid

			// Carry propagation
			carry = carryMid

			if carry > 0 && i+j+2 < uint(len(result)) {
				sumCarry, carryNext := bits.Add64(result[i+j+2], carry, 0)
				result[i+j+2] = sumCarry
				carry = carryNext
			}
		}
	}

	return digits.Wrap(result).Trim()
}
