package biginteger

type BigInteger struct {
	sign  bool
	value []uint64
}

var Zero = BigInteger{value: []uint64{0}}
var One = BigInteger{value: []uint64{1}}
var Two = BigInteger{value: []uint64{2}}
var Ten = BigInteger{value: []uint64{10}}

func (i BigInteger) String() string {
	sign := ""
	if i.sign {
		sign = "-"
	}

	return sign + stringAbs(i)
}

func (i BigInteger) Add(j BigInteger) BigInteger {
	if i.sign == j.sign {
		result := addUint64Arrays(i.value, j.value)

		return BigInteger{
			i.sign,
			result,
		}
	}

	if i.sign {
		return j.Subtract(i.Abs())
	}

	return i.Subtract(j.Abs())
}

func (i BigInteger) Subtract(j BigInteger) BigInteger {
	if i.IsEqualTo(j) {
		return Zero
	}

	if !i.sign && !j.sign {
		if i.IsLessThan(j) {
			result, _ := subtractUint64Arrays(j.value, i.value, false)

			return BigInteger{
				true,
				result,
			}
		}

		result, _ := subtractUint64Arrays(i.value, j.value, false)

		return BigInteger{
			false,
			result,
		}
	}

	if i.sign && j.sign {
		result := addUint64Arrays(i.value, j.value)

		return BigInteger{
			true,
			result,
		}
	}

	if i.sign {
		return j.Add(i.Abs())
	}

	return i.Add(j.Abs())
}

func (i BigInteger) Multiply(j BigInteger) BigInteger {
	if i.IsEqualTo(Zero) || j.IsEqualTo(Zero) {
		return Zero
	}

	if i.IsEqualTo(One) {
		return j
	}

	if j.IsEqualTo(One) {
		return i
	}

	sign := i.sign != j.sign
	result := multiplyAbs(i, j)

	return BigInteger{
		sign:  sign,
		value: result.value,
	}
}

func (i BigInteger) Divide(j BigInteger) BigInteger {
	if j.IsEqualTo(Zero) {
		panic("Division by zero")
	}

	sign := i.sign != j.sign
	result := divideAbs(i, j)

	return BigInteger{
		sign:  sign,
		value: result.value,
	}
}

func (i BigInteger) Modulo(j BigInteger) BigInteger {
	if j.IsEqualTo(Zero) {
		panic("Division by zero")
	}

	if i.IsEqualTo(Zero) {
		return Zero
	}

	if j.IsEqualTo(One) {
		return Zero
	}

	if i.IsEqualTo(j) {
		return Zero
	}

	sign := i.sign
	result := moduloAbs(i, j)

	return BigInteger{
		sign:  sign,
		value: result.value,
	}
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

	sign := i.sign
	result := powerAbs(i, j)

	return BigInteger{
		sign:  sign,
		value: result.value,
	}
}

func (i BigInteger) ShiftLeft(j BigInteger) BigInteger {
	if i.IsEqualTo(Zero) {
		return Zero
	}

	if j.IsEqualTo(Zero) {
		return i
	}

	return shiftLeft(i, j)
}

func (i BigInteger) ShiftRight(j BigInteger) BigInteger {
	if i.IsEqualTo(Zero) {
		return Zero
	}

	if j.IsEqualTo(Zero) {
		return i
	}

	if j.IsLessThan(Zero) {
		return i.ShiftLeft(j.Abs())
	}

	return shiftRight(i, j)
}

func (i BigInteger) BitLength() BigInteger {
	if i.Abs().IsLessThan(Two) {
		return One
	}

	return OfUint64(bitLengthUint64Array(i.value))
}

func (i BigInteger) IsLessThan(j BigInteger) bool {
	if i.sign && !j.sign {
		return true
	}

	if !i.sign && j.sign {
		return false
	}

	if i.sign && j.sign {
		return isLessThenUint64Array(j.value, i.value)
	}

	return isLessThenUint64Array(i.value, j.value)
}

func (i BigInteger) IsGreaterThan(j BigInteger) bool {
	if i.sign && !j.sign {
		return false
	}

	if !i.sign && j.sign {
		return true
	}

	if i.sign && j.sign {
		return isGreaterThenUint64Array(j.value, i.value)
	}

	return isGreaterThenUint64Array(i.value, j.value)
}

func (i BigInteger) IsEqualTo(j BigInteger) bool {
	if i.sign != j.sign {
		return false
	}

	if i.sign && j.sign {
		return isEqualToUint64Array(j.value, i.value)
	}

	return isEqualToUint64Array(i.value, j.value)
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
	sign := false
	if s[0] == '-' {
		sign = true
		s = s[1:]
	}

	i, err := parseToBigInteger(s)

	if err != nil {
		return nil, err
	}

	return &BigInteger{
		sign:  sign,
		value: i.value,
	}, nil
}
