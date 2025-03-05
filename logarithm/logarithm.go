package logarithm

import (
	"github.com/borisskert/go-biginteger/digits"
	"math"
)

func Log2(x []uint64) float64 {
	return log2(digits.Wrap(x))
}

func log2(x digits.Digits) float64 {
	if x.IsZero() {
		panic("Logarithm of zero is undefined")
	}
	if x.IsNegative() {
		panic("Logarithm of negative number is undefined")
	}

	n := x.BitLength()

	if n <= 64 {
		return math.Log2(float64(x.DigitAt(0)))
	}

	// Calculate the position of the most significant bit (MSB)
	msbPos := n - 1

	// Normalize the value to the range [1, 2)
	numerator := x.RightShiftBits(n - 64).DigitAt(0)
	denominator := digits.One().AsDigits().LeftShiftBits(63).DigitAt(0)
	normValue := float64(numerator) / float64(denominator)

	// Return the log2 value
	return float64(msbPos) + math.Log2(normValue)
}

func Log10(i []uint64) float64 {
	return log(digits.Wrap(i), digits.Ten().AsDigits())
}

func Log(i []uint64, base []uint64) float64 {
	return log(digits.Wrap(i), digits.Wrap(base))
}

func log(i digits.Digits, base digits.Digits) float64 {
	if base.IsZero() || base.IsOne() {
		panic("Logarithm base must be greater than one")
	}

	log2i := log2(i)
	log2base := log2(base)

	return log2i / log2base
}

func LogE(i []uint64) float64 {
	return logE(digits.Wrap(i))
}

const e = 2.71828182845904523536028747135266249775724709369995957496696763

func logE(i digits.Digits) float64 {
	log2i := log2(i)
	log2e := math.Log2(e)

	return log2i / log2e
}
