package biginteger

func divMod(a BigInteger, b BigInteger) (BigInteger, BigInteger) {
	if b.IsEqualTo(zero) {
		panic("Division by zero")
	}

	if a.IsEqualTo(zero) {
		return zero, zero
	}

	if b.IsEqualTo(one) {
		return a, zero
	}

	if a.IsEqualTo(b) {
		return one, zero
	}

	if a.IsGreaterThan(b) {
		return divModAbs(a, b)
	}

	return zero, a
}

func divModAbs(a BigInteger, b BigInteger) (BigInteger, BigInteger) {
	a = a.Abs()
	b = b.Abs()

	if b.IsEqualTo(zero) {
		panic("Division by zero")
	}

	if a.IsEqualTo(zero) {
		return zero, zero
	}

	if b.IsEqualTo(one) {
		return a, zero
	}

	if a.IsLessThan(b) {
		return zero, a
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

	return result, remaining
}
