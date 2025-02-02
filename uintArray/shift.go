package uintArray

func ShiftLeftBits(a []uint64, n uint) []uint64 {
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

	for i := uint(0); i < neededArraySize; i++ {
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

	return result
}
