package divmod

import "github.com/borisskert/go-biginteger/digits"

func divModOptimized(numerator digits.Digits, denominator digits.Digits) (digits.Digits, digits.Digits) {
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

func d1Normalize(n uint64, b digits.DoubleDigit, u digits.Digits, v digits.Digits) (digits.Digits, digits.Digits, digits.DoubleDigit) {
	var d digits.DoubleDigit

	bMaxValue, _ := b.SubtractDigit(digits.One()) // b-1
	vLast := v.DigitAt(uint(n - 1))               // Most significant digit of V

	// Let d = ⌊b / (v_{n-1} + 1)⌋

	if vLast.IsEqualDoubleDigit(bMaxValue) {
		// If the most significant digit is b-1, set d to 1
		d = digits.One().AsDoubleDigit()
	} else {
		// Otherwise, calculate d
		vLastPlus1, _ := vLast.AsDoubleDigit().AddDigit(digits.OneAsDigit())
		d, _ = b.Divide128(vLastPlus1) // Divide b by v_{n-1}+1 to get the scaling factor
	}

	// Multiply U and V by d to scale both
	r := u.MultiplyByDoubleDigit(d).Trim()
	v = v.MultiplyByDoubleDigit(d).Trim()

	return r, v, d
}

func d3ACalculateQHat(u digits.Digits, v digits.Digits, j int64, n int64) (digits.DoubleDigit, digits.Digit) {
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

func d3BTestAndCorrectQHat(b, qHat digits.DoubleDigit, rHat digits.Digit, v digits.Digits, r digits.Digits, j int64, n int64) (digits.DoubleDigit, digits.Digit) {
	vLast := v.DigitAt(uint(n - 1))

	vNMinus2 := v.DigitAt(uint(n - 2))
	qHatMulVMinus2, _ := qHat.Multiply(vNMinus2.AsDoubleDigit())
	bMulRHat, _ := b.Multiply(rHat.AsDoubleDigit())
	uJPlusNMinus2 := r.DigitAt(uint(j + n - 2))
	bMulRHatPlusUJPlusNMinus2, _ := bMulRHat.AddDigit(uJPlusNMinus2)

	// “Now test if qHat ≥ b or qHat * v[n−2] > b*rHat + u[j+n−2]; i”
	if qHat.IsGreaterThanOrEqual(b) || qHatMulVMinus2.IsGreaterThan(bMulRHatPlusUJPlusNMinus2) {
		qHat, _ = qHat.Subtract(digits.OneAsDigit().AsDoubleDigit())
		rHat, _ = rHat.Add(vLast)

		if rHat.IsLessThanDoubleDigit(b) {
			vNMinus2 = v.DigitAt(uint(n - 2))
			qHatMulVMinus2, _ = qHat.Multiply(vNMinus2.AsDoubleDigit())
			bMulRHat, _ = b.Multiply(rHat.AsDoubleDigit())
			uJPlusNMinus2 = r.DigitAt(uint(j + n - 2))
			bMulRHatPlusUJPlusNMinus2, _ = bMulRHat.AddDigit(uJPlusNMinus2)

			if qHat.IsGreaterThanOrEqual(b) || qHatMulVMinus2.IsGreaterThan(bMulRHatPlusUJPlusNMinus2) {
				qHat, _ = qHat.Subtract(digits.OneAsDigit().AsDoubleDigit())
				rHat, _ = rHat.Add(vLast)
			}
		}
	}

	return qHat, rHat
}

func d4MultiplyAndSubtract(j int64, n int64, u *digits.Digits, v digits.Digits, qHat digits.DoubleDigit) bool {
	ujToJPlusN := u.TakeDigits(uint(j), uint(j)+uint(n))
	vMulQHat := v.MultiplyByDoubleDigit(qHat)
	ujToJPlusNMinusVMulQHat, borrowed := ujToJPlusN.Trim().SubtractUnderflow(vMulQHat.Trim())

	u.Replace(uint(j), uint(j)+uint(n), ujToJPlusNMinusVMulQHat)

	return borrowed
}

func d6AddBack(j int64, qHat digits.DoubleDigit, u *digits.Digits, v digits.Digits) {
	borrow := digits.Zero()                                           // TODO empty
	qHat, borrow = qHat.Subtract(digits.OneAsDigit().AsDoubleDigit()) // Decrement q̂

	ujToJPlusNMinusVMulQHat := u.TakeDigits(uint(j), uint(j)+uint(v.Length())) // TODO rename
	ujToJPlusNMinusVMulQHat = ujToJPlusNMinusVMulQHat.Add(v)                   // Add back V to U_j:j+n

	u.Replace(uint(j), uint(j)+uint(v.Length()), ujToJPlusNMinusVMulQHat)

	if borrow != 0 {
		panic("Borrow should be false")
	}
}

func divModByDonaldKnuthsTAOCPv2(u digits.Digits, v digits.Digits) (digits.Digits, digits.Digits) {
	// Input:
	// - U = (u_0, u_1, ..., u_{m+n-1})  ⟶ Dividend (m+n digits)
	// - V = (v_0, v_1, ..., v_{n-1})     ⟶ Divisor (n digits, v_{n-1} ≠ 0)

	// Output:
	// - Q = (q_0, q_1, ..., q_{m-1}) ⟶ Quotient
	// - R = (r_0, r_1, ..., r_{n-1}) ⟶ Remainder (same size as V)

	n := int64(v.Length())
	m := int64(u.Length()) - n

	// Base b (e.g., 10 or 2^64)
	b := digits.DoubleDigitOf(1, 0)

	// D1. [Normalize.]
	u, v, d := d1Normalize(uint64(n), b, u, v)

	q := digits.Zero().AsDigits()

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

func divideByDigit(a digits.Digits, b digits.Digit) (digits.Digits, digits.Digits) {
	q, r := divideByWord(a, b)
	return q, r.AsDigits()
}

func divideByWord(dividend digits.Digits, divisor digits.Digit) (digits.Digits, digits.Digit) {
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

func divideByDoubleDigit(dividend digits.Digits, divisor digits.DoubleDigit) (digits.Digits, digits.DoubleDigit) {
	if divisor.IsZero() {
		panic("Division by zero")
	}

	if dividend.IsZero() {
		return digits.Zero().AsDigits(), digits.Zero().AsDoubleDigit()
	}

	if divisor.IsOne() {
		return dividend, digits.Zero().AsDoubleDigit()
	}

	if dividend.Compare(divisor.AsDigits()) < 0 {
		return digits.Zero().AsDigits(), dividend.AsDoubleDigit()
	}

	if divisor.High() == 0 {
		quotient, remainder := divideByWord(dividend, divisor.Low())
		return quotient, remainder.AsDoubleDigit()
	}

	quotient, remainder := divideByDoubleWord(dividend, divisor)

	return quotient.Trim(), remainder
}

func divideByDoubleWord(dividend digits.Digits, divisor digits.DoubleDigit) (digits.Digits, digits.DoubleDigit) {
	quotient := digits.Empty()
	remainder := digits.Zero().AsDigits()

	for i := int64(dividend.Length()) - 2; i >= 0; i-- {
		di := dividend.DoubleDigitAt(uint(i))
		remainder = remainder.LeftShiftBits(64).AddDoubleDigit(di)

		qHat, rHat := divideByDoubleDigit(remainder, divisor)

		quotient = quotient.LeftShiftBits(64).Add(qHat)
		remainder = rHat.AsDigits()
	}

	return quotient.Trim(), remainder.AsDoubleDigit()
}
