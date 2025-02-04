package multiply

import "github.com/borisskert/go-biginteger/digits"

func karatsubaMultiply(a, b digits.Digits) digits.Digits {
	m := min(a.Length(), b.Length())
	if m < 2 {
		return schoolbookMultiply(a, b)
	}

	k := m/2 + m%2

	a1, a0 := a.Split(k)
	b1, b0 := b.Split(k)

	signA := a0.Compare(a1) < 0
	signB := b0.Compare(b1) < 0

	c0 := karatsubaMultiply(a0, b0)
	c1 := karatsubaMultiply(a1, b1)
	c2 := karatsubaMultiply(
		a0.SubtractAbs(a1).Trim(),
		b0.SubtractAbs(b1).Trim(),
	)

	signC2 := signA != signB
	mid := c0.Add(c1).
		SubtractNoBorrow(c2.Sign(signC2)).
		LeftShiftDigits(k)

	return c0.Add(mid).Add(c1.LeftShiftDigits(k * 2)).Trim()
}
