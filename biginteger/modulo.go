package biginteger

func moduloAbs(i BigInteger, j BigInteger) BigInteger {
	i = i.Abs()
	j = j.Abs()

	if j.IsEqualTo(Zero) {
		panic("Division by zero")
	}

	if i.IsLessThan(j) {
		return i
	}

	division := i.Divide(j)
	multiply := division.Multiply(j)

	return i.Subtract(multiply)
}
