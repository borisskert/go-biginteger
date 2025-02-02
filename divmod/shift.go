package divmod

import "github.com/borisskert/go-biginteger/digits"

func ShiftRightUint64Array(a []uint64, n uint64) []uint64 {
	shifted := digits.Wrap(a).RightShiftBits(n)
	shifted.NormalizeInPlace()

	return shifted.Trim().AsArray()
}
