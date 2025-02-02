package uintArray

import "math/bits"

func BitLength(a []uint64) uint {
	lenA := len(a)

	if lenA == 0 {
		return 0
	}

	return uint((lenA-1)*64 + bits.Len64(a[lenA-1]))
}
