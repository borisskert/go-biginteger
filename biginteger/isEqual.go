package biginteger

func isEqualTo(left BigInteger, right BigInteger) bool {
	if left.sign != right.sign {
		return false
	}

	if left.sign && right.sign {
		return isEqualToUint64Array(right.value, left.value)
	}

	return isEqualToUint64Array(left.value, right.value)
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
