package biginteger

import (
	"github.com/borisskert/go-biginteger/logarithm"
	"github.com/borisskert/go-biginteger/stringify"
)

type BigInteger struct {
	sign  bool
	value []uint64
}

func (i BigInteger) String() string {
	sign := ""
	if i.sign {
		sign = "-"
	}

	return sign + stringify.Stringify(i.value)
}

func (i BigInteger) Add(summand BigInteger) BigInteger {
	return add(i, summand)
}

func (i BigInteger) Subtract(subtrahend BigInteger) BigInteger {
	return subtract(i, subtrahend)
}

func (i BigInteger) Multiply(multiplier BigInteger) BigInteger {
	return multiply(i, multiplier)
}

func (i BigInteger) Divide(divisor BigInteger) BigInteger {
	quotient, _ := divMod(i, divisor)
	return quotient
}

func (i BigInteger) Modulo(divisor BigInteger) BigInteger {
	_, remainder := divMod(i, divisor)
	return remainder
}

func (i BigInteger) DivMod(divisor BigInteger) (BigInteger, BigInteger) {
	return divMod(i, divisor)
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

func (i BigInteger) Negate() BigInteger {
	if i.IsEqualTo(zero) {
		return zero
	}

	return BigInteger{sign: !i.sign, value: i.value}
}

func (i BigInteger) Power(exponent BigInteger) BigInteger {
	return power(i, exponent)
}

func (i BigInteger) ShiftLeft(count uint64) BigInteger {
	return shiftLeft(i, count)
}

func (i BigInteger) ShiftRight(count uint64) BigInteger {
	return shiftRight(i, count)
}

func (i BigInteger) BitLength() uint64 {
	return uint64(bitLength(i))
}

func (i BigInteger) Digits() uint64 {
	if i.IsEqualTo(zero) {
		return 1
	}

	return uint64(i.Abs().Log10()) + 1
}

func (i BigInteger) CompareTo(other BigInteger) int {
	return compareTo(i, other)
}

func (i BigInteger) IsLessThan(other BigInteger) bool {
	return compareTo(i, other) < 0
}

func (i BigInteger) IsGreaterThan(other BigInteger) bool {
	return compareTo(i, other) > 0
}

func (i BigInteger) IsEqualTo(other BigInteger) bool {
	return compareTo(i, other) == 0
}

func (i BigInteger) Log2() float64 {
	return logarithm.Log2(i.value)
}

func (i BigInteger) Log10() float64 {
	return logarithm.Log10(i.value)
}

func (i BigInteger) Log(base BigInteger) float64 {
	return logarithm.Log(i.value, base.value)
}

func (i BigInteger) LogE() float64 {
	return logarithm.LogE(i.value)
}

func (i BigInteger) LogF(base float64) float64 {
	return logarithm.LogF(i.value, base)
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

func OfUint64Array(array []uint64) BigInteger {
	values := make([]uint64, len(array))
	copy(values, array)

	return BigInteger{value: values}
}

func Of(s string) (*BigInteger, error) {
	return parse(s)
}
