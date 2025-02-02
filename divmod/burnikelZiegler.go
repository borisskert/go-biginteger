package divmod

import (
	"github.com/borisskert/go-biginteger/digits"
)

func divThreeLongHalvesByTwo(a1, a2, a3, b1, b2 digits.Digits) (digits.Digits, digits.Digits) {
	b := b1.Concat(b2)
	a := a1.Concat(a2)

	q, _ := divModSelect(a, b1)
	c, _ := a.Subtract(q.Multiply(b1))

	d := q.Multiply(b2)
	r := c.LeftShiftDigits(a3.Length()).Add(a3)
	r.SubtractInPlace(d)

	if r.IsNegative() {
		one := digits.One().AsDigits()

		q.SubtractInPlace(one) // TODO decrement
		r.AddInPlace(b)

		if r.IsNegative() {
			q.SubtractInPlace(one) // TODO decrement
			r.AddInPlace(b)

			if r.IsNegative() {
				panic("This should never happen: r was negative twice")
			}
		}
	}

	return q, r
}

func DivMod(numerator, denominator []uint64) ([]uint64, []uint64) {
	quotient, remainder := divModSelect(digits.Wrap(numerator), digits.Wrap(denominator))
	return quotient.Trim().AsArray(), remainder.Trim().AsArray()
}

func divModSelect(numerator, denominator digits.Digits) (digits.Digits, digits.Digits) {
	if denominator.Compare(digits.ZeroAsDigits()) == 0 {
		panic("Division by zero")
	}

	if numerator.Compare(digits.ZeroAsDigits()) == 0 {
		return digits.ZeroAsDigits(), digits.ZeroAsDigits()
	}

	one := digits.One().AsDigits()

	if denominator.Compare(one) == 0 {
		return numerator, digits.ZeroAsDigits()
	}

	cmp := numerator.Compare(denominator)
	if cmp < 0 {
		return digits.ZeroAsDigits(), numerator
	}

	if cmp == 0 {
		return one, digits.ZeroAsDigits()
	}

	n := denominator.Length()
	m := numerator.Length()

	if n == 1 && m == 1 {
		return numerator.DigitAt(0).DivideToDigits(denominator.DigitAt(0))
	}

	if n == 1 && m == 2 {
		q, r := numerator.AsDoubleDigit().DivideByDigit(denominator.DigitAt(0))
		return q.AsDigits(), r.AsDigits()
	}

	if n == 1 {
		return divideByDigit(numerator, denominator.DigitAt(0))
	}

	if m < 40 {
		return divModByDonaldKnuthsTAOCPv2(numerator.Trim(), denominator.Trim())
	}

	if m == 2*n {
		return divModBurnikelZiegler(numerator.Trim(), denominator.Trim())
	}

	if m > 2*n {
		return divModChunked(numerator.Trim(), denominator.Trim())
	}

	return divModByDonaldKnuthsTAOCPv2(numerator.Trim(), denominator.Trim())
}

func divModBurnikelZiegler(a digits.Digits, b digits.Digits) (digits.Digits, digits.Digits) {
	n := b.Length()

	if a.Length() != 2*n {
		panic("Burnikel-Ziegler's precondition not met: " +
			"a's length must be 2n, b's length must be n")
	}

	a1, a2, a3, a4 := a.Quarter()
	b1, b2 := b.Halve()

	q1, r := divThreeLongHalvesByTwo(a1, a2, a3, b1, b2)

	r1, r2 := r.Halve()

	q2, s := divThreeLongHalvesByTwo(r1, r2, a4, b1, b2)

	q := q1.LeftShiftDigits(n / 2).Add(q2)

	return q.Trim(), s.Trim()
}
