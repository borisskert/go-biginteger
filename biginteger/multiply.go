package biginteger

func multiplyAbs(a BigInteger, b BigInteger) BigInteger {
	a = a.Abs()
	b = b.Abs()

	result := Zero
	multiplier := a
	remaining := b

	for remaining.IsGreaterThan(Zero) {
		if remaining.IsOdd() {
			result = result.Add(multiplier)
		}

		multiplier = multiplier.Add(multiplier)
		remaining = remaining.ShiftRight(One)
	}

	return result
}
