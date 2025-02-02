package divmod

func ShiftRightUint64Array(a []uint64, n uint64) []uint64 {
	shifted := Wrap(a).RightShiftBits(n)
	shifted.NormalizeInPlace()

	return shifted.Trim().Array()
}
