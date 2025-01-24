package biginteger

func isGreaterThan(left BigInteger, right BigInteger) bool {
	if left.sign && !right.sign {
		return false
	}

	if !left.sign && right.sign {
		return true
	}

	if left.sign && right.sign {
		return isGreaterThenUint64Array(right.value, left.value)
	}

	return isGreaterThenUint64Array(left.value, right.value)
}

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
