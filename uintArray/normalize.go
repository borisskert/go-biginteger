package uintArray

func normalize(value []uint64) []uint64 {
	if len(value) == 0 {
		return []uint64{0}
	}

	return trimLeadingZeros(value)
}

func trimLeadingZeros(a []uint64) []uint64 {
	for len(a) > 1 && a[len(a)-1] == 0 {
		a = a[:len(a)-1]
	}

	return a
}
