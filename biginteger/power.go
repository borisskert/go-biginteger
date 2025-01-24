package biginteger

func powerAbs(i BigInteger, j BigInteger) BigInteger {
	i = i.Abs()
	j = j.Abs()

	if j.IsEqualTo(zero) {
		return one
	}

	if j.IsEqualTo(one) {
		return i
	}

	if i.IsEqualTo(zero) {
		return zero
	}

	if i.IsEqualTo(one) {
		return one
	}

	if j.IsEqualTo(one) {
		return i
	}

	result := one

	for j.IsGreaterThan(zero) {
		if j.IsOdd() {
			result = result.Multiply(i)
			j = j.Subtract(one)
		} else {
			i = i.Multiply(i)
			j = j.Divide(two)
		}
	}

	return result
}
