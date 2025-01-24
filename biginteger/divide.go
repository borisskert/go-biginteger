package biginteger

func divideAbs(i BigInteger, j BigInteger) BigInteger {
	i = i.Abs()
	j = j.Abs()

	if j.IsEqualTo(zero) {
		panic("Division by zero")
	}

	if i.IsEqualTo(zero) {
		return zero
	}

	if j.IsEqualTo(one) {
		return i
	}

	if i.Abs().IsLessThan(j) {
		return zero
	}

	if i.Abs().IsEqualTo(j) {
		return one
	}

	result := zero
	remaining := i
	divisor := j
	quotient := one

	for {
		if remaining.IsGreaterThan(divisor) {
			remaining = remaining.Subtract(divisor)
			result = result.Add(quotient)
			divisor = divisor.Add(divisor)
			quotient = quotient.Add(quotient)
		} else if remaining.IsEqualTo(divisor) {
			result = result.Add(quotient)
			break
		} else if remaining.IsLessThan(j) {
			break
		} else if remaining.IsLessThan(divisor) {
			divisor = j
			quotient = one
		} else {
			break
		}
	}

	return result
}
