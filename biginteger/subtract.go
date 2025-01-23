package biginteger

func subtractUint64Arrays(a, b []uint64, borrow bool) ([]uint64, bool) {
	// Ensure a is the longer array (swap if necessary)
	if len(b) > len(a) {
		a, b = b, a
	}

	result := make([]uint64, len(a))
	carry := uint64(0)
	if borrow {
		carry = 1
	}

	for i := 0; i < len(a); i++ {
		ai := a[i]
		bi := uint64(0)
		if i < len(b) {
			bi = b[i]
		}

		// Subtract with carry
		diff, newBorrow := subtract(ai, bi+carry)
		result[i] = diff
		carry = 0
		if newBorrow {
			carry = 1
		}
	}

	// Handle any remaining carry
	finalBorrow := (carry > 0)

	// Remove leading zeros from the result (if needed)
	for len(result) > 1 && result[len(result)-1] == 0 {
		result = result[:len(result)-1]
	}

	return result, finalBorrow
}

func subtract(a, b uint64) (uint64, bool) {
	return a - b, a < b
}
