package biginteger

func shiftRight(value BigInteger, count BigInteger) BigInteger {
	if value.IsEqualTo(zero) {
		return zero
	}

	if count.IsEqualTo(zero) {
		return value
	}

	if count.IsLessThan(zero) {
		return value.ShiftLeft(count.Abs())
	}

	return BigInteger{
		sign:  value.sign,
		value: shiftRightUint64Array(value.value, count.Uint()),
	}
}

func shiftRightUint64Array(a []uint64, n uint) []uint64 {
	if n == 0 || len(a) == 0 {
		if len(a) == 0 || (len(a) == 1 && a[0] == 0) {
			return []uint64{0}
		}

		return a
	}

	result := make([]uint64, len(a))
	bitsPerElement := uint(64)
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
