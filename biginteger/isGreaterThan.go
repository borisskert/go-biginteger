package biginteger

func isGreaterThenUint64Array(a, b []uint64) bool {
	if len(a) > len(b) {
		return true
	}

	if len(a) < len(b) {
		return false
	}

	for k := len(a) - 1; k > 0; k-- {
		if a[k] > b[k] {
			return true
		}

		if a[k] < b[k] {
			return false
		}
	}

	return a[0] > b[0]
}
