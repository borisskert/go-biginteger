package uintArray

func ShiftLeftBits(a []uint64, n uint64) []uint64 {
	if n == 0 {
		return a
	}

	bitSize := BitLength(a)
	neededBitSize := bitSize + n

	neededArraySize := (neededBitSize+1)/64 + 1

	div := n / 64
	mod := n % 64

	result := make([]uint64, neededArraySize)
	carry := uint64(0)

	for i := uint64(0); i < neededArraySize; i++ {
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

	return result
}

func ShiftRightBits(a []uint64, n uint64) []uint64 {
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
