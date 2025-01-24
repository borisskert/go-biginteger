package biginteger

func power(a BigInteger, b BigInteger) BigInteger {
	if b.IsEqualTo(zero) {
		return one
	}

	if b.IsEqualTo(one) {
		return a
	}

	if a.IsEqualTo(zero) {
		return zero
	}

	if a.IsEqualTo(one) {
		return one
	}

	sign := a.sign
	result := powerAbs(a, b)

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

	if b.IsEqualTo(one) {
		return a
	}

	if a.IsEqualTo(zero) {
		return zero
	}

	if a.IsEqualTo(one) {
		return one
	}

	if b.IsEqualTo(one) {
		return a
	}

	result := one

	for b.IsGreaterThan(zero) {
		if b.IsOdd() {
			result = result.Multiply(a)
			b = b.Subtract(one)
		} else {
			a = a.Multiply(a)
			b = b.Divide(two)
		}
	}

	return result
}
