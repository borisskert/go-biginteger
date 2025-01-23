package biginteger

func bitLengthUint64Array(a []uint64) uint64 {
	if len(a) == 0 {
		return uint64(1)
	}

	lastPart := a[len(a)-1]
	if lastPart == 0 {
		return uint64(1)
	}

	result := uint64(64 * (len(a) - 1))

	for lastPart > 0 {
		result = result + 1
		lastPart = lastPart >> 1
	}

	return result
}
