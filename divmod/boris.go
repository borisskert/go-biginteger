package divmod

func DivModBoris(numerator []uint64, denominator []uint64) ([]uint64, []uint64) {
	a := MakeDigitsByCopy(numerator)
	b := MakeDigitsByCopy(denominator)

	q, r := divModBoris(a, b)

	return q.Array(), r.Array()
}

func divModBoris(numerator Digits, denominator Digits) (Digits, Digits) {
	if denominator.Compare(ZeroAsDigits()) == 0 {
		panic("Division by zero")
	}

	if numerator.Compare(ZeroAsDigits()) == 0 {
		return ZeroAsDigits(), ZeroAsDigits()
	}

	if denominator.Compare(One()) == 0 {
		return numerator, ZeroAsDigits()
	}

	if numerator.Compare(denominator) < 0 {
		return ZeroAsDigits(), numerator
	}

	result := ZeroAsDigits()
	remaining := numerator

	divisor := denominator
	quotient := One()

	for divisor.Compare(remaining) <= 0 {
		divisor = divisor.LeftShiftBits(1)
		quotient = quotient.LeftShiftBits(1)
	}

	for divisor.Compare(denominator) >= 0 {
		if remaining.Compare(divisor) >= 0 {
			remaining, _ = remaining.Subtract(divisor)
			result = result.Add(quotient)
		}

		divisor = divisor.RightShiftBits(1)
		quotient = quotient.RightShiftBits(1)
	}

	result = result.Trim()
	remaining = remaining.Trim()

	return result, remaining
}
