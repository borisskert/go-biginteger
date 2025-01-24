package biginteger

func modulo(a BigInteger, b BigInteger) BigInteger {
	if b.IsEqualTo(zero) {
		panic("Division by zero")
	}

	if a.IsEqualTo(zero) {
		return zero
	}

	if b.IsEqualTo(one) {
		return zero
	}

	if a.IsEqualTo(b) {
		return zero
	}

	sign := a.sign
	result := moduloAbs(a, b)

	return BigInteger{
		sign:  sign,
		value: result.value,
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
