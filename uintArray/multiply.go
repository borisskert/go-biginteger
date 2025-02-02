package uintArray

import "math/bits"

func MultiplyUint64Array(a, b []uint64) []uint64 {
	result := make([]uint64, len(a)+len(b))

	for i := 0; i < len(a); i++ {
		var carry uint64 = 0

		for j := 0; j < len(b); j++ {
			high, low := bits.Mul64(a[i], b[j])
			low += result[i+j]

			if low < result[i+j] {
				high++
			}

			result[i+j] = low

			high += carry
			result[i+j+1] += high

			if result[i+j+1] < high {
				carry = 1
			} else {
				carry = 0
			}
		}

		result[i+len(b)] += carry
	}

	return result
}
