package biginteger

func multiply(i BigInteger, j BigInteger) BigInteger {
	if i.IsEqualTo(zero) || j.IsEqualTo(zero) {
		return zero
	}

	if i.IsEqualTo(one) {
		return j
	}

	if j.IsEqualTo(one) {
		return i
	}

	sign := i.sign != j.sign
	result := multiplyAbs(i, j)

	return BigInteger{
		sign:  sign,
		value: result.value,
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

		multiplier = multiplier.Add(multiplier)
		remaining = remaining.ShiftRight(one)
	}

	return result
}
