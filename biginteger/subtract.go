package biginteger

func subtractUint64Arrays(a, b []uint64, borrow bool) ([]uint64, bool) {
	if len(a) == 0 && len(b) == 0 {
		if borrow {
			return []uint64{1}, false
		}

		return []uint64{}, false
	}

	carry := uint64(0)
	if borrow {
		carry = uint64(1)
	}

	if len(a) == 0 {
		diff, borrow := subtract(b[0], carry)
		return []uint64{diff}, borrow
	}

	if len(b) == 0 {
		diff, borrow := subtract(a[0], carry)
		return []uint64{diff}, borrow
	}

	diff, borrow := subtract(a[0], b[0])
	result := []uint64{diff - carry}
	rest, borrow := subtractUint64Arrays(a[1:], b[1:], borrow)

	if len(rest) == 1 && rest[0] == 0 {
		return result, borrow
	}

	return append(result, rest...), borrow
}

func subtract(a, b uint64) (uint64, bool) {
	return a - b, a < b
}
