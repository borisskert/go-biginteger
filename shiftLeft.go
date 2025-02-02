package biginteger

import (
	"github.com/borisskert/go-biginteger/uintArray"
)

func shiftLeft(value BigInteger, count uint64) BigInteger {
	if value.IsEqualTo(zero) {
		return zero
	}

	if count == 0 {
		return value
	}

	shiftedBits := uintArray.ShiftLeftBits(value.value, count)

	return BigInteger{
		sign:  value.sign,
		value: trimLeadingZeros(shiftedBits),
	}
}

func shiftLeftUint64Array(a []uint64, n uint64) []uint64 {
	if n == 0 {
		return a
	}

	div := n / uint64(64)
	mod := n % uint64(64)

	size := uint64(len(a)) + div
	if mod > 0 {
		size++
	}

	result := make([]uint64, size)
	carry := uint64(0)

	for i := uint64(0); i < size; i++ {
		if i < div {
			result[i] = 0
			continue
		}

		var value uint64
		if i-div < uint64(len(a)) {
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

	return trimLeadingZeros(result)
}

func trimLeadingZeros(array []uint64) []uint64 {
	if len(array) == 0 {
		return []uint64{0}
	}

	size := len(array)

	for size > 1 && array[size-1] == 0 {
		size--
	}

	return array[:size]
}
