package biginteger

func multiply(multiplicand BigInteger, multiplier BigInteger) BigInteger {
	if multiplicand.IsEqualTo(zero) || multiplier.IsEqualTo(zero) {
		return zero
	}

	if multiplicand.IsEqualTo(one) {
		return multiplier
	}

	if multiplier.IsEqualTo(one) {
		return multiplicand
	}

	sign := multiplicand.sign != multiplier.sign
	product := multiplyAbs(multiplicand, multiplier)

	return BigInteger{
		sign:  sign,
		value: product.value,
	}
}

func multiplyAbs(a BigInteger, b BigInteger) BigInteger {
	a = a.Abs()
	b = b.Abs()

	result := zero
	multiplier := a
	remaining := b

	for remaining.IsGreaterThan(zero) {
		if remaining.IsOdd() {
			result = result.Add(multiplier)
		}

		multiplier = multiplier.ShiftLeft(1)
		remaining = remaining.ShiftRight(1)
	}

	return result
}
