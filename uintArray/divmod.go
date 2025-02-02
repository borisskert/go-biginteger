package uintArray

func DivModUint64Arrays(a []uint64, b []uint64) ([]uint64, []uint64) {
	if Compare(b, []uint64{0}) == 0 {
		panic("Division by zero")
	}

	if Compare(a, []uint64{0}) == 0 {
		return []uint64{0}, []uint64{0}
	}

	if Compare(b, []uint64{1}) == 0 {
		return a, []uint64{0}
	}

	if Compare(a, b) < 0 {
		return []uint64{0}, a
	}

	result := []uint64{0}
	remaining := a

	divisor := b
	quotient := []uint64{1}

	for Compare(divisor, remaining) <= 0 {
		divisor = ShiftLeftBits(divisor, 1)
		quotient = ShiftLeftBits(quotient, 1)
	}

	for Compare(divisor, b) >= 0 {
		if Compare(remaining, divisor) >= 0 {
			remaining = Subtract(remaining, divisor)
			result = Add(result, quotient)
		}

		divisor = ShiftRightBits(divisor, 1)
		quotient = ShiftRightBits(quotient, 1)
	}

	return result, remaining
}
