package biginteger

func power(base BigInteger, exponent BigInteger) BigInteger {
	if exponent.IsEqualTo(zero) {
		return one
	}

	if exponent.IsEqualTo(one) {
		return base
	}

	if base.IsEqualTo(zero) {
		return zero
	}

	if base.IsEqualTo(one) {
		return one
	}

	sign := base.sign
	result := powerAbs(base, exponent)

	return BigInteger{
		sign:  sign,
		value: result.value,
	}
}

func powerAbs(a BigInteger, b BigInteger) BigInteger {
	a = a.Abs()
	b = b.Abs()

	if b.IsEqualTo(zero) {
		return one
	}

	if a.IsEqualTo(zero) {
		return zero
	}

	if a.IsEqualTo(one) || b.IsEqualTo(one) {
		return a
	}

	result := one

	for b.IsGreaterThan(zero) {
		if b.IsOdd() {
			result = result.Multiply(a)
		}

		a = a.Multiply(a)
		b = b.ShiftRight(1) // Proper halving of exponent
	}

	return result
}
