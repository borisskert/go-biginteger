package donaldKnuth

import (
	"github.com/borisskert/go-biginteger/digits"
)

// DonaldKnuthsAlgorithmD Implements Algorithm D (Division of nonnegative integers) introduced by Donald Knuth
// in his book "The Art of Computer Programming, Volume 2: Seminumerical Algorithms", 3rd edition, 1997 (chapter 4.3.1, p. 272)
type DonaldKnuthsAlgorithmD struct {
}

func (a *DonaldKnuthsAlgorithmD) DivMod(numerator digits.Digits, denominator digits.Digits) (digits.Digits, digits.Digits) {
	return divModByDonaldKnuthsTAOCPv2(numerator, denominator)
}

func divModByDonaldKnuthsTAOCPv2(
	u digits.Digits,
	v digits.Digits,
) (digits.Digits, digits.Digits) {
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

		qHat, _ = d3BTestAndCorrectQHat(b, qHat, rHat, v, u, j, n)

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
	remainder, r := divideByDoubleDigit(u.Trim(), d)

	if r.IsGreaterThan(digits.Zero().AsDoubleDigit()) {
		panic("Remainder must be less than divisor")
	}

	// Step 5: **Return Q (quotient) and R (remainder)**
	return q.Trim(), remainder.Trim()
}

func d1Normalize(n uint64, b digits.DoubleDigit, u digits.Digits, v digits.Digits) (digits.Digits, digits.Digits, digits.DoubleDigit) {
	var d digits.DoubleDigit

	bMaxValue, _ := b.SubtractDigit(digits.One()) // b-1
	vLast := v.DigitAt(uint(n - 1))               // Most significant digit of V

	// Let d = ⌊b / (v_{n-1} + 1)⌋
	if vLast.IsEqualDoubleDigit(bMaxValue) {
		d = digits.One().AsDoubleDigit() // If the most significant digit is b-1, set d to 1
	} else {
		d, _ = b.Divide(vLast.Increment().AsDoubleDigit()) // Divide b by v_{n-1}+1 to get the scaling factor
	}

	// Multiply U and V by d to scale both
	u = u.MultiplyByDoubleDigit(d).Trim()
	v = v.MultiplyByDoubleDigit(d).Trim()

	return u, v, d
}

func d3ACalculateQHat(u digits.Digits, v digits.Digits, j int64, n int64) (digits.DoubleDigit, digits.Digit) {
	uMostSignificantDigit := u.DigitAt(uint(j + n))
	uNextDigit := u.DigitAt(uint(j + n - 1))

	uBothDigits, _ := uMostSignificantDigit.AsDoubleDigit().
		LeftShift(64).
		AddDigit(uNextDigit)

	vLast := v.DigitAt(uint(n - 1))

	// Estimate quotient digit q̂ = ⌊(u_j * b + u_{j+1}) / v_{n-1}⌋
	qHat, rHat := uBothDigits.DivideByDigit(vLast)

	return qHat, rHat
}

func d3BTestAndCorrectQHat(
	b, qHat digits.DoubleDigit, rHat digits.Digit, v digits.Digits, u digits.Digits, j, n int64,
) (digits.DoubleDigit, digits.Digit) {
	vLast := v.DigitAt(uint(n - 1))

	// “Now test if qHat ≥ b or qHat * v[n−2] > b*rHat + u[j+n−2]; i”
	if checkIfQHatTooLarge(b, qHat, rHat, u, v, n, j) {
		qHat, _ = qHat.Decrement()
		rHat, _ = rHat.Add(vLast)

		if rHat.IsLessThanDoubleDigit(b) {
			if checkIfQHatTooLarge(b, qHat, rHat, u, v, n, j) {
				var borrowed bool
				qHat, borrowed = qHat.Decrement()

				if borrowed {
					panic("Decrement must not underflow")
				}

				rHat, _ = rHat.Add(vLast)
			}
		}
	}

	return qHat, rHat
}

func checkIfQHatTooLarge(b, qHat digits.DoubleDigit, rHat digits.Digit, u, v digits.Digits, n, j int64) bool {
	vDigit := v.DigitAt(uint(n - 2))
	uDigit := u.DigitAt(uint(j + n - 2))

	qHatIsGreaterOrEqualB := qHat.IsGreaterThanOrEqual(b)

	return qHatIsGreaterOrEqualB ||
		qHat.MultiplyIgnoreOverflow(vDigit.AsDoubleDigit()).
			IsGreaterThan(
				b.MultiplyIgnoreOverflow(rHat.AsDoubleDigit()).
					AddDigitIgnoreOverflow(uDigit),
			)
}

func d4MultiplyAndSubtract(j int64, n int64, u *digits.Digits, v digits.Digits, qHat digits.DoubleDigit) bool {
	uDigits := u.ChunkInclusive(uint(j), uint(j)+uint(n))
	uDigitsDifference, borrowed := uDigits.Trim().SubtractUnderflow(v.MultiplyByDoubleDigit(qHat).Trim())

	u.Replace(uint(j), uint(j)+uint(n), uDigitsDifference)

	return borrowed
}

func d6AddBack(j int64, qHat digits.DoubleDigit, u *digits.Digits, v digits.Digits) digits.DoubleDigit {
	var borrow bool
	qHat, borrow = qHat.Decrement()

	uChunk := u.ChunkInclusive(uint(j), uint(j)+v.Length())
	uChunk = uChunk.Add(v) // Add back V to U_j:j+n

	u.Replace(uint(j), uint(j)+uint(v.Length()), uChunk)

	if borrow {
		panic("Decrement must not underflow")
	}

	return qHat
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
		quotient, remainder := dividend.DivideByDigit(divisor.Low())
		return quotient, remainder.AsDoubleDigit()
	}

	return dividend.DivideByDoubleDigit(divisor)
}
