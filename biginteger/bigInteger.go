package biginteger

import (
	"strconv"
)

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

	return sign + i.stringAbs()
}

func (i BigInteger) stringAbs() string {
	j := i.Abs()

	if j.IsLessThan(Ten) {
		return strconv.FormatUint(j.value[0], 10)
	}

	result := ""
	for j.IsGreaterThan(Zero) {
		remainder := j.Modulo(Ten)
		result = strconv.FormatUint(remainder.value[0], 10) + result
		j = j.Divide(Ten)
	}

	return result
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

	return i.shiftLeftAbs(j)
}

func (i BigInteger) shiftLeftAbs(j BigInteger) BigInteger {
	return BigInteger{
		sign:  i.sign,
		value: shiftLeft(i.value, j.Uint()),
	}
}

func shiftLeft(a []uint64, n uint) []uint64 {
	if n == 0 {
		return a
	}

	div := n / uint(64) // Number of 64-bit word shifts
	mod := n % uint(64) // Remaining bit shift within a word

	size := uint(len(a)) + div
	if mod > 0 {
		size++ // Extra space for carry if mod > 0
	}

	result := make([]uint64, size)
	carry := uint64(0)

	for i := uint(0); i < size; i++ {
		if i < div {
			result[i] = 0
			continue
		}

		var value uint64
		if i-div < uint(len(a)) {
			value = a[i-div]
		} else {
			value = 0
		}

		newValue := (value << mod) | carry
		result[i] = newValue
		carry = value >> (64 - mod)
	}

	if carry > 0 {
		result = append(result, carry)
	}

	// Remove leading zeros
	for len(result) > 1 && result[len(result)-1] == 0 {
		result = result[:len(result)-1]
	}

	return result
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

	divisor := Two.Power(j)
	return i.Divide(divisor)
}

func (i BigInteger) BitLength() BigInteger {
	if i.Abs().IsLessThan(Two) {
		return One
	}

	return OfUint64(bitLength(i.value))
}

func bitLength(a []uint64) uint64 {
	if len(a) == 0 {
		return uint64(1)
	}

	lastPart := a[len(a)-1]
	if lastPart == 0 {
		return uint64(1)
	}

	//result := OfUint64(64).
	//	Multiply(OfUint64(uint64(len(a) - 1)))

	result := uint64(64 * (len(a) - 1))

	for lastPart > 0 {
		result = result + 1
		lastPart = lastPart >> 1
	}

	return result
}

func (i BigInteger) IsLessThan(j BigInteger) bool {
	if i.sign && !j.sign {
		return true
	}

	if !i.sign && j.sign {
		return false
	}

	if i.sign && j.sign {
		return isLessThenArrays(j.value, i.value)
	}

	return isLessThenArrays(i.value, j.value)
}

func isLessThenArrays(a, b []uint64) bool {
	if len(a) < len(b) {
		return true
	}

	if len(a) > len(b) {
		return false
	}

	for k := len(a) - 1; k > 0; k-- {
		if a[k] > b[k] {
			return false
		}

		if a[k] < b[k] {
			return true
		}
	}

	return a[0] < b[0]
}

func (i BigInteger) IsGreaterThan(j BigInteger) bool {
	if i.sign && !j.sign {
		return false
	}

	if !i.sign && j.sign {
		return true
	}

	if i.sign && j.sign {
		return isGreaterThenArrays(j.value, i.value)
	}

	return isGreaterThenArrays(i.value, j.value)
}

func isGreaterThenArrays(a, b []uint64) bool {
	if len(a) > len(b) {
		return true
	}

	if len(a) < len(b) {
		return false
	}

	for k := len(a) - 1; k > 0; k-- {
		if a[k] > b[k] {
			return true
		}

		if a[k] < b[k] {
			return false
		}
	}

	return a[0] > b[0]
}

func (i BigInteger) IsEqualTo(j BigInteger) bool {
	if i.sign != j.sign {
		return false
	}

	if i.sign && j.sign {
		return isEqualToArrays(j.value, i.value)
	}

	return isEqualToArrays(i.value, j.value)
}

func (i BigInteger) Uint() uint {
	if len(i.value) == 0 {
		return 0
	}

	return uint(i.value[0])
}

func isEqualToArrays(a, b []uint64) bool {
	if len(a) != len(b) {
		return false
	}

	for k := 0; k < len(a); k++ {
		if a[k] != b[k] {
			return false
		}
	}

	return true
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

	i, err := parseToUint64BLA(s)

	if err != nil {
		return nil, err
	}

	return &BigInteger{
		sign:  sign,
		value: i.value,
	}, nil
}

func parseToUint64BLA(s string) (*BigInteger, error) {
	if len(s) == 0 {
		return &Zero, nil
	}

	result := Zero

	for len(s) > 0 {
		first := firstRunes(s, 1)
		parsedDigit, err := strconv.ParseUint(first, 10, 64)
		if err != nil {
			return nil, err
		}

		digit := OfUint64(parsedDigit)
		result = result.Multiply(Ten).Add(digit)
		s = lastRunes(s, len(s)-1)
	}

	return &result, nil
}

func firstRunes(s string, n int) string {
	if n <= 0 {
		return ""
	}

	if n >= len(s) {
		return s
	}

	runes := []rune(s)
	if len(runes) < n {
		return s
	}

	return string(runes[:n])
}

func lastRunes(s string, n int) string {
	runes := []rune(s)
	if len(runes) < n {
		return s
	}

	return string(runes[len(runes)-n:])
}
