package biginteger

func divideAbs(i BigInteger, j BigInteger) BigInteger {
	i = i.Abs()
	j = j.Abs()

	if j.IsEqualTo(Zero) {
		panic("Division by zero")
	}

	if i.IsEqualTo(Zero) {
		return Zero
	}

	if j.IsEqualTo(One) {
		return i
	}

	if i.Abs().IsLessThan(j) {
		return Zero
	}

	if i.Abs().IsEqualTo(j) {
		return One
	}

	result := Zero
	remaining := i
	divisor := j
	quotient := One

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
			quotient = One
		} else {
			break
		}
	}

	return result
}
