package schoolbook

import (
	"github.com/borisskert/go-biginteger/digits"
	"github.com/borisskert/go-biginteger/multiply/api"
	"math/bits"
)

type schoolbookMultiplyAlgorithm struct {
}

func (s schoolbookMultiplyAlgorithm) Multiply(
	multiplicand digits.Digits,
	multiplier digits.Digits,
) (product digits.Digits) {
	result := multiply(multiplicand, multiplier)

	if multiplicand.IsNegative() != multiplier.IsNegative() && !result.IsZero() {
		return result.Trim().Negative()
	}

	return result.Trim()
}

func multiply(a digits.Digits, b digits.Digits) digits.Digits {
	result := make([]uint64, a.Length()+b.Length())

	for i := uint(0); i < a.Length(); i++ {
		ai := a.DigitAt(i)
		var carry uint64

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
				result[i+j+2], _ = bits.Add64(result[i+j+2], carry, 0)
			}
		}
	}

	return digits.Wrap(result).Trim()
}

// NewSchoolbookMultiplyAlgorithm creates a new instance of the schoolbook multiply algorithm.
func NewSchoolbookMultiplyAlgorithm() api.MultiplyAlgorithm {
	return &schoolbookMultiplyAlgorithm{}
}
