package biginteger

func shiftLeft(value BigInteger, count uint64) BigInteger {
	if value.IsEqualTo(zero) {
		return zero
	}

	if count == 0 {
		return value
	}

	return BigInteger{
		sign:  value.sign,
		value: shiftLeftUint64Array(value.value, count),
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

	return removeLeadingZeros(result)
}

func removeLeadingZeros(array []uint64) []uint64 {
	for len(array) > 1 && array[len(array)-1] == 0 {
		array = array[:len(array)-1]
	}

	return array
}
