package divmod

import "math/bits"

func DivTwoDigitsByTwo(d, v DoubleDigit) (DoubleDigit, DoubleDigit) {
	if v.IsZero() {
		panic("Division by zero")
	}

	if d.IsLessThan(v) {
		return Digit(0).AsDoubleDigit(), d
	}

	if d.IsEqual(v) {
		return Digit(1).AsDoubleDigit(), Digit(0).AsDoubleDigit()
	}

	if v.High() == 0 {
		quotient, remainder := DivTwoDigitsByOne2(d, v.Low())
		return quotient, remainder.AsDoubleDigit()
	}

	shift := v.LeadingZeros()
	v = v.LeftShift(shift)
	d = d.LeftShift(shift)

	quot, _ := d.High().DivMod(v.High())

	prod := quot.Multiply2(v.Low())
	prod, _ = prod.Add128(quot.Multiply2(v.High()).Low().ShiftLeftToDoubleDigit(64))

	if prod.IsGreaterThanOrEqual(d) {
		quot = quot - 1
		prod = quot.Multiply2(v.Low())
		prod, _ = prod.Add128(quot.Multiply2(v.High()).Low().ShiftLeftToDoubleDigit(64))
	}

	rem, _ := d.Subtract(prod)
	rem = rem.RightShift(shift)

	return quot.AsDoubleDigit(), rem
}

// DivTwoDigitsByOne divides two Digits by One Digit.
// Returns quotient and remainder.
func DivTwoDigitsByOne(a DoubleDigit, b Digit) (Digit, Digit) {
	a1 := a.High().High()
	a2 := a.High().Low()
	a3 := a.Low().High()
	a4 := a.Low().Low()

	b1 := b.High()
	b2 := b.Low()

	q1, r := divThreeHalvesByTwo(a1, a2, a3, b1, b2)

	r1, r2 := r.Split()
	q2, s := divThreeHalvesByTwo(r1, r2, a4, b1, b2)

	return makeDigitOfHalfdigits(q1, q2), s
}

func DivTwoDigitsByOne2(D DoubleDigit, V Digit) (DoubleDigit, Digit) {
	if V == 0 {
		panic("division by zero")
	}

	// Step 1: Divide the high 64-bit part
	qHi, r := bits.Div64(0, uint64(D.High()), uint64(V))

	// Step 2: Divide the lower 64-bit part using the remainder from step 1
	qLo, r := bits.Div64(r, uint64(D.Low()), uint64(V))

	return DoubleDigitOf(Digit(qHi), Digit(qLo)), Digit(r)
}

func DivTwoDigitsByHalf(numerator DoubleDigit, denominator HalfDigit) (DoubleDigit, Digit) {
	if denominator == 0 {
		panic("Division by zero")
	}

	v := denominator.AsDigit()

	qHi, r := bits.Div64(0, uint64(numerator.High()), uint64(v))
	qLo, r := bits.Div64(r, uint64(numerator.Low()), uint64(v))

	return DoubleDigitOf(Digit(qHi), Digit(qLo)), Digit(r)
}

func makeDigitOfHalfdigits(hi HalfDigit, lo HalfDigit) Digit {
	return Digit(hi)<<32 | Digit(lo)
}

func divThreeHalvesByTwo(a1, a2, a3, b1, b2 HalfDigit) (HalfDigit, Digit) {
	B := makeDigitOfHalfdigits(b1, b2)
	A := makeDigitOfHalfdigits(a1, a2)

	q, _ := A.Divide32(b1)
	C, _ := A.Subtract(q.Multiply(b1))
	c := C.Low()
	d := q.Multiply(b2)
	r, borrow := makeDigitOfHalfdigits(c, a3).Subtract(d)

	if borrow > 0 {
		q = q - 1
		r, borrow = r.Subtract(B)

		if borrow > 0 {
			q = q - 1
			r, _ = r.Subtract(B)
		}
	}

	return q, r
}

func divThreeLongHalvesByTwo(a1, a2, a3, b1, b2 Digits) (Digits, Digits) {
	b := b1.Concat(b2)
	a := a1.Concat(a2)

	q, _ := divModSelect(a, b1)
	c, _ := a.Subtract(q.Multiply(b1))

	d := q.Multiply(b2)
	r := c.LeftShiftDigits(a3.Length()).Add(a3)
	r.SubtractInPlace(d)

	if r.IsNegative() {
		q.SubtractInPlace(One())
		r.AddInPlace(b)

		if r.IsNegative() {
			q.SubtractInPlace(One())
			r.AddInPlace(b)

			if r.IsNegative() {
				panic("This should never happen: r was negative twice")
			}
		}
	}

	return q, r
}

func DivMod(numerator, denominator []uint64) ([]uint64, []uint64) {
	quotient, remainder := divModSelect(Wrap(numerator), Wrap(denominator))
	return quotient.Trim().Array(), remainder.Trim().Array()
}

func divModSelect(numerator, denominator Digits) (Digits, Digits) {
	if denominator.Compare(ZeroAsDigits()) == 0 {
		panic("Division by zero")
	}

	if numerator.Compare(ZeroAsDigits()) == 0 {
		return ZeroAsDigits(), ZeroAsDigits()
	}

	if denominator.Compare(One()) == 0 {
		return numerator, ZeroAsDigits()
	}

	cmp := numerator.Compare(denominator)
	if cmp < 0 {
		return ZeroAsDigits(), numerator
	}

	if cmp == 0 {
		return One(), ZeroAsDigits()
	}

	n := denominator.Length()
	m := numerator.Length()

	if n == 1 && m == 1 {
		return numerator.DigitAt(0).DivModToDigits(denominator.DigitAt(0))
	}

	if n == 1 && m == 2 {
		q, r := numerator.AsDoubleDigit().DivideByDigit(denominator.DigitAt(0))
		return DigitsOfDoubleDigit(q), MakeDigitsOfDigit(r)
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

func divModBurnikelZiegler(a Digits, b Digits) (Digits, Digits) {
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
