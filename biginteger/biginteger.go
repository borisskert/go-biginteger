package biginteger

type BigInteger struct {
	sign  bool
	value []uint64
}

func (i BigInteger) String() string {
	return toString(i)
}

func (i BigInteger) Add(j BigInteger) BigInteger {
	return add(i, j)
}

func (i BigInteger) Subtract(j BigInteger) BigInteger {
	return subtract(i, j)
}

func (i BigInteger) Multiply(j BigInteger) BigInteger {
	return multiply(i, j)
}

func (i BigInteger) Divide(j BigInteger) BigInteger {
	return divide(i, j)
}

func (i BigInteger) Modulo(j BigInteger) BigInteger {
	return modulo(i, j)
}

func (i BigInteger) IsEven() bool {
	if len(i.value) == 0 {
		return true
	}

	return i.value[0]%2 == 0
}

func (i BigInteger) IsOdd() bool {
	if len(i.value) == 0 {
		return false
	}

	return i.value[0]%2 == 1
}

func (i BigInteger) Abs() BigInteger {
	return BigInteger{value: i.value}
}

func (i BigInteger) Power(j BigInteger) BigInteger {
	return power(i, j)
}

func (i BigInteger) ShiftLeft(j BigInteger) BigInteger {
	return shiftLeft(i, j)
}

func (i BigInteger) ShiftRight(j BigInteger) BigInteger {
	return shiftRight(i, j)
}

func (i BigInteger) BitLength() BigInteger {
	if i.Abs().IsLessThan(two) {
		return one
	}

	length := bitLengthUint64Array(i.value)

	return OfUint64(length)
}

func (i BigInteger) IsLessThan(j BigInteger) bool {
	return isLessThan(i, j)
}

func (i BigInteger) IsGreaterThan(j BigInteger) bool {
	return isGreaterThan(i, j)
}

func (i BigInteger) IsEqualTo(j BigInteger) bool {
	return isEqualTo(i, j)
}

func (i BigInteger) Uint() uint {
	if len(i.value) == 0 {
		return 0
	}

	return uint(i.value[0])
}

func OfUint64(i uint64) BigInteger {
	return BigInteger{value: []uint64{i}}
}

func Of(s string) (*BigInteger, error) {
	return parse(s)
}
