package biginteger

func isEqualToUint64Array(a, b []uint64) bool {
	if len(a) != len(b) {
		return false
	}

	for k := 0; k < len(a); k++ {
		if a[k] != b[k] {
			return false
		}
	}

	return true
}
