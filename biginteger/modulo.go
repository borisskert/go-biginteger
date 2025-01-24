package biginteger

func modulo(dividend BigInteger, divisor BigInteger) BigInteger {
	if divisor.IsEqualTo(zero) {
		panic("Division by zero")
	}

	if dividend.IsEqualTo(zero) {
		return zero
	}

	if divisor.IsEqualTo(one) {
		return zero
	}

	if dividend.IsEqualTo(divisor) {
		return zero
	}

	sign := dividend.sign
	remainder := moduloAbs(dividend, divisor)

	return BigInteger{
		sign:  sign,
		value: remainder.value,
	}
}

func moduloAbs(a BigInteger, b BigInteger) BigInteger {
	a = a.Abs()
	b = b.Abs()

	if b.IsEqualTo(zero) {
		panic("Division by zero")
	}

	if a.IsLessThan(b) {
		return a
	}

	division := a.Divide(b)
	product := division.Multiply(b)

	return a.Subtract(product)
}
