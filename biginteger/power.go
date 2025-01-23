package biginteger

func powerAbs(i BigInteger, j BigInteger) BigInteger {
	i = i.Abs()
	j = j.Abs()

	if j.IsEqualTo(Zero) {
		return One
	}

	if j.IsEqualTo(One) {
		return i
	}

	if i.IsEqualTo(Zero) {
		return Zero
	}

	if i.IsEqualTo(One) {
		return One
	}

	if j.IsEqualTo(One) {
		return i
	}

	result := One

	for j.IsGreaterThan(Zero) {
		if j.IsOdd() {
			result = result.Multiply(i)
			j = j.Subtract(One)
		} else {
			i = i.Multiply(i)
			j = j.Divide(Two)
		}
	}

	return result
}
