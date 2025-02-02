package divmod

func divModOptimized(numerator Digits, denominator Digits) (Digits, Digits) {
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

	length := numerator.Length()

	if length == 1 {
		return divideByDigit(numerator, denominator.AsDigit())
	}

	if length >= 6 {
		trailingZeroBits := min(numerator.TrailingZeros(), denominator.TrailingZeros())
		numerator = numerator.RightShiftBits(trailingZeroBits)
		denominator = denominator.RightShiftBits(trailingZeroBits)

		q, r := divModOptimized(numerator, denominator)

		return q.LeftShiftBits(trailingZeroBits), r.LeftShiftBits(trailingZeroBits)
	}

	return divModByDonaldKnuthsTAOCPv2(numerator.Trim(), denominator.Trim())
}

func d1Normalize(n uint64, b DoubleDigit, u Digits, v Digits) (Digits, Digits, DoubleDigit) {
	var d DoubleDigit

	bMaxValue, _ := b.SubtractDigit(One().AsDigit()) // b-1
	vLast := v.DigitAt(uint(n - 1))                  // Most significant digit of V

	// Let d = ⌊b / (v_{n-1} + 1)⌋

	if vLast.IsEqualDoubleDigit(bMaxValue) {
		// If the most significant digit is b-1, set d to 1
		d = One().AsDoubleDigit()
	} else {
		// Otherwise, calculate d
		vLastPlus1, _ := vLast.AsDoubleDigit().AddDigit(OneAsDigit())
		d, _ = b.Divide128(vLastPlus1) // Divide b by v_{n-1}+1 to get the scaling factor
	}

	// Multiply U and V by d to scale both
	r := u.MultiplyDoubleDigit(d).Trim()
	v = v.MultiplyDoubleDigit(d).Trim()

	return r, v, d
}

func d3ACalculateQHat(u Digits, v Digits, j int64, n int64) (DoubleDigit, Digit) {
	ujPlusN := u.DigitAt(uint(j + n))
	uJPlusNMinus1 := u.DigitAt(uint(j + n - 1))

	uJPlusNAndUJPlusNMinus1, _ := ujPlusN.AsDoubleDigit().
		LeftShift(64).
		AddDigit(uJPlusNMinus1)

	vLast := v.DigitAt(uint(n - 1))

	// Estimate quotient digit q̂ = ⌊(u_j * b + u_{j+1}) / v_{n-1}⌋
	qHat, rHat := uJPlusNAndUJPlusNMinus1.DivideByDigit(vLast)

	return qHat, rHat
}

func d3BTestAndCorrectQHat(b, qHat DoubleDigit, rHat Digit, v Digits, r Digits, j int64, n int64) (DoubleDigit, Digit) {
	vLast := v.DigitAt(uint(n - 1))

	vNMinus2 := v.DigitAt(uint(n - 2))
	qHatMulVMinus2, _ := qHat.Multiply(vNMinus2.AsDoubleDigit())
	bMulRHat, _ := b.Multiply(rHat.AsDoubleDigit())
	uJPlusNMinus2 := r.DigitAt(uint(j + n - 2))
	bMulRHatPlusUJPlusNMinus2, _ := bMulRHat.AddDigit(uJPlusNMinus2)

	// “Now test if qHat ≥ b or qHat * v[n−2] > b*rHat + u[j+n−2]; i”
	if qHat.IsGreaterThanOrEqual(b) || qHatMulVMinus2.IsGreaterThan(bMulRHatPlusUJPlusNMinus2) {
		qHat, _ = qHat.Subtract(OneAsDigit().AsDoubleDigit())
		rHat, _ = rHat.Add(vLast)

		if rHat.IsLessThan128(b) {
			vNMinus2 = v.DigitAt(uint(n - 2))
			qHatMulVMinus2, _ = qHat.Multiply(vNMinus2.AsDoubleDigit())
			bMulRHat, _ = b.Multiply(rHat.AsDoubleDigit())
			uJPlusNMinus2 = r.DigitAt(uint(j + n - 2))
			bMulRHatPlusUJPlusNMinus2, _ = bMulRHat.AddDigit(uJPlusNMinus2)

			if qHat.IsGreaterThanOrEqual(b) || qHatMulVMinus2.IsGreaterThan(bMulRHatPlusUJPlusNMinus2) {
				qHat, _ = qHat.Subtract(OneAsDigit().AsDoubleDigit())
				rHat, _ = rHat.Add(vLast)
			}
		}
	}

	return qHat, rHat
}

func d4MultiplyAndSubtract(j int64, n int64, u *Digits, v Digits, qHat DoubleDigit) bool {
	ujToJPlusN := u.TakeDigits(uint(j), uint(j)+uint(n))
	vMulQHat := v.MultiplyDoubleDigit(qHat)
	ujToJPlusNMinusVMulQHat, borrowed := ujToJPlusN.Trim().SubtractUnderflow(vMulQHat.Trim())

	u.Replace(uint(j), uint(j)+uint(n), ujToJPlusNMinusVMulQHat)

	return borrowed
}

func d6AddBack(j int64, qHat DoubleDigit, u *Digits, v Digits) {
	borrow := Zero()
	qHat, borrow = qHat.Subtract(OneAsDigit().AsDoubleDigit()) // Decrement q̂

	ujToJPlusNMinusVMulQHat := u.TakeDigits(uint(j), uint(j)+uint(v.Length())) // TODO rename
	ujToJPlusNMinusVMulQHat = ujToJPlusNMinusVMulQHat.Add(v)                   // Add back V to U_j:j+n

	u.Replace(uint(j), uint(j)+uint(v.Length()), ujToJPlusNMinusVMulQHat)

	if borrow != 0 {
		panic("Borrow should be false")
	}
}

func divModByDonaldKnuthsTAOCPv2(u Digits, v Digits) (Digits, Digits) {
	// Input:
	// - U = (u_0, u_1, ..., u_{m+n-1})  ⟶ Dividend (m+n digits)
	// - V = (v_0, v_1, ..., v_{n-1})     ⟶ Divisor (n digits, v_{n-1} ≠ 0)

	// Output:
	// - Q = (q_0, q_1, ..., q_{m-1}) ⟶ Quotient
	// - R = (r_0, r_1, ..., r_{n-1}) ⟶ Remainder (same size as V)

	n := int64(v.Length())
	m := int64(u.Length()) - n

	// Base b (e.g., 10 or 2^64)
	b := DoubleDigitOf(1, 0)

	// D1. [Normalize.]
	u, v, d := d1Normalize(uint64(n), b, u, v)

	q := Zero().AsDigits()

	// D2. [Initialize j.]
	for j := m; j >= 0; j-- {
		qHat, rHat := d3ACalculateQHat(u, v, j, n)

		qHat, rHat = d3BTestAndCorrectQHat(b, qHat, rHat, v, u, j, n)

		// Multiply and subtract: U_j:j+n = U_j:j+n - q̂ * V
		wasNegative := d4MultiplyAndSubtract(j, n, &u, v, qHat)

		// D5. [Test remainder.] Store q̂ in Q[j]
		q.SetDigitAt(uint(j), qHat.Low())
		if j > 0 {
			q.SetDigitAt(uint(j-1), qHat.High())
		}

		if wasNegative {
			d6AddBack(j, qHat, &u, v)
		}
	}

	// D8. [Unnormalize.]
	remainder, _ := divideByDoubleWord(u.Trim(), d)

	// Step 5: **Return Q (quotient) and R (remainder)**
	return q.Trim(), remainder.Trim()
}

func divideByDigit(a Digits, b Digit) (Digits, Digits) {
	q, r := divideByWord(a, b)
	return q, MakeDigitsOfDigit(r)
}

func divideByWord(dividend Digits, divisor Digit) (Digits, Digit) {
	quotient := Empty()
	remainder := Zero().AsDoubleDigit()

	for i := int64(dividend.Length()) - 1; i >= 0; i-- {
		di := dividend.DigitAt(uint(i))
		remainder, _ = remainder.LeftShift(64).AddDigit(di)

		qHat, rHat := remainder.DivideByDigit(divisor)

		quotient = quotient.LeftShiftBits(64).AddDoubleDigit(qHat)
		remainder = rHat.AsDoubleDigit()
	}

	return quotient.Trim(), remainder.Low()
}

func divideByDoubleDigit(dividend Digits, divisor DoubleDigit) (Digits, DoubleDigit) {
	if divisor.IsZero() {
		panic("Division by zero")
	}

	if dividend.IsZero() {
		return Zero().AsDigits(), Zero().AsDoubleDigit()
	}

	if divisor.IsOne() {
		return dividend, Zero().AsDoubleDigit()
	}

	if dividend.Compare(divisor.AsDigits()) < 0 {
		return Zero().AsDigits(), dividend.AsDoubleDigit()
	}

	if divisor.High() == 0 {
		quotient, remainder := divideByWord(dividend, divisor.Low())
		return quotient, remainder.AsDoubleDigit()
	}

	quotient, remainder := divideByDoubleWord(dividend, divisor)

	return quotient.Trim(), remainder
}

func divideByDoubleWord(dividend Digits, divisor DoubleDigit) (Digits, DoubleDigit) {
	quotient := Empty()
	remainder := Zero().AsDigits()

	for i := int64(dividend.Length()) - 2; i >= 0; i-- {
		di := dividend.DoubleDigitAt(uint(i))
		remainder = remainder.LeftShiftBits(64).AddDoubleDigit(di)

		qHat, rHat := divideByDoubleDigit(remainder, divisor)

		quotient = quotient.LeftShiftBits(64).Add(qHat)
		remainder = rHat.AsDigits()
	}

	return quotient.Trim(), remainder.AsDoubleDigit()
}
