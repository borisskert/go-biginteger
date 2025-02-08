package multiply

import "github.com/borisskert/go-biginteger/digits"

func KaratsubaMultiply(a, b digits.Digits) digits.Digits {
	if min(a.Length(), b.Length()) < 2 {
		return SchoolbookMultiply(a, b)
	}

	n := max(a.Length(), b.Length())

	k := max((n+1)/2, 1)

	a1, a0 := a.Split2(k)
	b1, b0 := b.Split2(k)

	signA := a0.Compare(a1) < 0
	signB := b0.Compare(b1) < 0

	c0 := KaratsubaMultiply(a0, b0)
	c1 := KaratsubaMultiply(a1, b1)
	c2 := KaratsubaMultiply(
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
