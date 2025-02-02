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

	shiftedBits := uintArray.ShiftLeftBits(value.value, uint(count))

	return BigInteger{
		sign:  value.sign,
		value: trimLeadingZeros(shiftedBits),
	}
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
