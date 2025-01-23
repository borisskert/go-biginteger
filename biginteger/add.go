package biginteger

func addUint64Arrays(a, b []uint64) []uint64 {
	if len(a) == 0 {
		return b
	}

	if len(b) == 0 {
		return a
	}

	if len(a) < len(b) {
		a = append(a, make([]uint64, len(b)-len(a))...)
	} else {
		b = append(b, make([]uint64, len(a)-len(b))...)
	}

	result := make([]uint64, len(a))

	sum := uint64(0)
	carry := false
	for i := 0; i < len(a); i++ {
		sum, carry = add(a[i], b[i], carry)
		result[i] = sum
	}

	if carry {
		result = append(result, 1)
	}

	return result
}

func add(a, b uint64, carry bool) (uint64, bool) {
	c := uint64(0)
	if carry {
		c = 1
	}

	sum := a + b + c
	carry = sum < a || sum < b

	return sum, carry
}
