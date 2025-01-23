package biginteger

func multiplyAbs(a BigInteger, b BigInteger) BigInteger {
	a = a.Abs()
	b = b.Abs()

	result := Zero
	factor := One
	multiplier := a
	remaining := b

	for {
		if remaining.IsGreaterThan(factor) {
			result = result.Add(multiplier)
			remaining = remaining.Subtract(factor)
			multiplier = multiplier.Add(multiplier)
			factor = factor.Add(factor)
		} else if remaining.IsEqualTo(factor) {
			result = result.Add(multiplier)
			break
		} else if remaining.IsLessThan(factor) {
			factor = One
			multiplier = a
		} else {
			break
		}
	}

	return result
}
