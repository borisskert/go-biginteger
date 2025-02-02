package biginteger

import "github.com/borisskert/go-biginteger/divmod"

func shiftRight(value BigInteger, count uint64) BigInteger {
	if value.IsEqualTo(zero) {
		return zero
	}

	if count == 0 {
		return value
	}

	return BigInteger{
		sign:  value.sign,
		value: divmod.ShiftRightUint64Array(value.value, uint(count)),
	}
}

func shiftRightUint64Array(a []uint64, n uint64) []uint64 {
	if n == 0 || len(a) == 0 {
		if len(a) == 0 || (len(a) == 1 && a[0] == 0) {
			return []uint64{0}
		}

		return a
	}

	result := make([]uint64, len(a))
	bitsPerElement := uint64(64)
	shiftMask := uint64((1 << n) - 1)

	carry := uint64(0)
	for i := len(a) - 1; i >= 0; i-- {
		current := a[i]
		result[i] = (current >> n) | (carry << (bitsPerElement - n))
		carry = current & shiftMask
	}

	for len(result) > 1 && result[len(result)-1] == 0 {
		result = result[:len(result)-1]
	}

	if len(result) == 1 && result[0] == 0 {
		return []uint64{0}
	}

	return result
}
