package functions

import "github.com/borisskert/go-biginteger/digits"

func DivByDoubleDigit(dividend digits.Digits, divisor digits.DoubleDigit) (digits.Digits, digits.DoubleDigit) {
	quotient := digits.Empty()
	remainder := digits.Zero().AsDigits()

	for i := int64(dividend.Length()) - 1; i >= 0; i-- {
		di := dividend.DigitAt(uint(i))
		remainder = remainder.LeftShiftBits(64).AddDigit(di)

		r1 := remainder.DigitAt(2)
		r2 := remainder.DigitAt(1)
		r3 := remainder.DigitAt(0)

		qHat, rHat := divThreeByTwo(r1, r2, r3, divisor)

		quotient = quotient.LeftShiftBits(64).Add(qHat.AsDigits())
		remainder = rHat.AsDigits()
	}

	return quotient.Trim(), remainder.AsDoubleDigit()
}

func divThreeByTwo(a1, a2, a3 digits.Digit, b digits.DoubleDigit) (digits.DoubleDigit, digits.Digit) {
	a := digits.MakeDoubleDigitOfDigits(a1, a2)

	q, _ := a.DivideByDigit(b.High())
	qMulB, _ := q.MultiplyDigit(b.High())
	c, _ := a.Subtract(qMulB)

	d, _ := q.MultiplyDigit(b.Low())
	r, _ := c.LeftShift(64).AddDigit(a3)
	r, borrow := r.Subtract(d)

	if borrow > 0 {
		q, _ = q.Decrement()
		r, borrow = r.Add(b)

		if borrow > 0 {
			q, _ = q.Decrement()
			r, borrow = r.Add(b)

			if borrow > 0 {
				panic("This should never happen: r was negative twice")
			}
		}
	}

	return q, r.Low()
}

func DivByDigit(dividend digits.Digits, divisor digits.Digit) (digits.Digits, digits.Digit) {
	quotient := digits.Empty()
	remainder := digits.Zero().AsDoubleDigit()

	for i := int64(dividend.Length()) - 1; i >= 0; i-- {
		di := dividend.DigitAt(uint(i))
		remainder, _ = remainder.LeftShift(64).AddDigit(di)

		qHat, rHat := remainder.DivideByDigit(divisor)

		quotient = quotient.LeftShiftBits(64).AddDoubleDigit(qHat)
		remainder = rHat.AsDoubleDigit()
	}

	return quotient.Trim(), remainder.Low()
}
