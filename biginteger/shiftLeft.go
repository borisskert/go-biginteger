package biginteger

func shiftLeft(i BigInteger, j BigInteger) BigInteger {
	return BigInteger{
		sign:  i.sign,
		value: shiftLeftUint64Array(i.value, j.Uint()),
	}
}

func shiftLeftUint64Array(a []uint64, n uint) []uint64 {
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
