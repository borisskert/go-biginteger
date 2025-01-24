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

	if a.Abs().IsLessThan(b) {
		return zero
	}

	if a.Abs().IsEqualTo(b) {
		return one
	}

	result := zero
	remaining := a
	divisor := b
	quotient := one

	for {
		if remaining.IsGreaterThan(divisor) {
			remaining = remaining.Subtract(divisor)
			result = result.Add(quotient)
			divisor = divisor.Add(divisor)
			quotient = quotient.Add(quotient)
		} else if remaining.IsEqualTo(divisor) {
			result = result.Add(quotient)
			break
		} else if remaining.IsLessThan(b) {
			break
		} else if remaining.IsLessThan(divisor) {
			divisor = b
			quotient = one
		} else {
			break
		}
	}

	return result
}
