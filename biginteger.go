package biginteger

import (
	"github.com/borisskert/go-biginteger/stringify"
	"math"
	"math/bits"
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
	if i.IsEqualTo(zero) {
		panic("Logarithm of zero is undefined")
	}
	if i.IsLessThan(zero) {
		panic("Logarithm of negative number is undefined")
	}

	return log2Uint64Array(i.value)
}

func (i BigInteger) Log10() float64 {
	return i.Log(ten)
}

func (i BigInteger) Log(base BigInteger) float64 {
	if base.IsEqualTo(zero) || base.IsEqualTo(one) {
		panic("Logarithm base must be greater than one")
	}

	log2i := i.Log2()
	log2base := base.Log2()

	return log2i / log2base
}

func log2Uint64Array(arr []uint64) float64 {
	// Find the highest set bit (integer part of log2)
	var integerPart int
	var msbIndex int
	var msbValue uint64

	// Iterate through the array to find the most significant set bit
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] != 0 {
			msbValue = arr[i]
			msbIndex = i
			break
		}
	}

	// If all elements are zero, return NaN (log2 of zero is undefined)
	if msbValue == 0 {
		return math.NaN()
	}

	// Find the position of the highest set bit in the most significant uint64
	integerPart = msbIndex*64 + 63 - bits.LeadingZeros64(msbValue)

	// Calculate the fractional part
	fractionalPart := 0.0
	if integerPart > 0 {
		// Shift the most significant uint64 to the left to isolate the bits below the MSB
		shift := uint(64 - bits.LeadingZeros64(msbValue))
		shiftedValue := msbValue << shift

		// Use the shifted value to approximate the fractional part
		fractionalPart = math.Log2(float64(shiftedValue)) - float64(shift)
	}

	return float64(integerPart) + fractionalPart
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
