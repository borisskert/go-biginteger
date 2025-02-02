package uintArray

import "math/bits"

func Subtract(a, b []uint64) []uint64 {
	if len(b) > len(a) {
		a, b = b, a
	}

	result := make([]uint64, len(a))
	carry := uint64(0)

	for i := 0; i < len(a); i++ {
		ai := a[i]
		bi := uint64(0)
		if i < len(b) {
			bi = b[i]
		}

		diff, borrow := bits.Sub64(ai, bi+carry, 0)
		result[i] = diff
		carry = borrow
	}

	return result
}
