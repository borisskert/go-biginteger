package biginteger

func divide(dividend BigInteger, divisor BigInteger) BigInteger {
	if divisor.IsEqualTo(zero) {
		panic("Division by zero")
	}

	sign := dividend.sign != divisor.sign
	quotient := divideAbs(dividend, divisor)

	return BigInteger{
		sign:  sign,
		value: quotient.value,
	}
}

func divideAbs(a BigInteger, b BigInteger) BigInteger {
	a = a.Abs()
	b = b.Abs()

	if b.IsEqualTo(zero) {
		panic("Division by zero")
	}

	if a.IsEqualTo(zero) {
		return zero
	}

	if b.IsEqualTo(one) {
		return a
	}

	if a.IsLessThan(b) {
		return zero
	}

	result := zero
	remaining := a

	divisor := b
	quotient := one

	for divisor.CompareTo(remaining) <= 0 {
		divisor = divisor.ShiftLeft(1)
		quotient = quotient.ShiftLeft(1)
	}

	for divisor.CompareTo(b) >= 0 {
		if remaining.CompareTo(divisor) >= 0 {
			remaining = remaining.Subtract(divisor)
			result = result.Add(quotient)
		}

		divisor = divisor.ShiftRight(1)
		quotient = quotient.ShiftRight(1)
	}

	return result
}
