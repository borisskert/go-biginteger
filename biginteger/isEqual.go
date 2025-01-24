package biginteger

func isEqualTo(a BigInteger, b BigInteger) bool {
	if a.sign != b.sign {
		return false
	}

	if a.sign && b.sign {
		return isEqualToUint64Array(b.value, a.value)
	}

	return isEqualToUint64Array(a.value, b.value)
}

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
