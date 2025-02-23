package karatsuba

import (
	"github.com/borisskert/go-biginteger/digits"
	"github.com/borisskert/go-biginteger/multiply/api"
)

type karatsubaMultiplyAlgorithm struct {
	algorithm api.MultiplyAlgorithm
}

func (k karatsubaMultiplyAlgorithm) Multiply(
	multiplicand digits.Digits, multiplier digits.Digits,
) (product digits.Digits) {
	return karatsubaMultiply(
		multiplicand,
		multiplier,
		k.algorithm.Multiply,
	)
}

func karatsubaMultiply(
	a, b digits.Digits,
	fn func(digits.Digits, digits.Digits) digits.Digits,
) digits.Digits {
	n := max(a.Length(), b.Length())

	k := max((n+1)/2, 1)

	a1, a0 := a.Split2(k)
	b1, b0 := b.Split2(k)

	signA := a0.Compare(a1) < 0
	signB := b0.Compare(b1) < 0

	c0 := fn(a0, b0)
	c1 := fn(a1, b1)
	c2 := fn(
		a0.Difference(a1),
		b0.Difference(b1),
	)

	signC2 := signA != signB
	mid := c0.Add(c1).
		Subtract(c2.Sign(signC2)).
		LeftShiftDigits(k)

	result := c0.Add(mid).Add(c1.LeftShiftDigits(k * 2))

	if a.IsNegative() != b.IsNegative() && !result.IsZero() {
		result = result.Negate()
	}

	return result.Trim()
}

func DecorateWithKaratsuba(algorithm api.MultiplyAlgorithm) api.MultiplyAlgorithm {
	return &karatsubaMultiplyAlgorithm{
		algorithm: algorithm,
	}
}
