package biginteger

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
