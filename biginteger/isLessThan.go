package biginteger

func isLessThan(left BigInteger, right BigInteger) bool {
	if left.sign && !right.sign {
		return true
	}

	if !left.sign && right.sign {
		return false
	}

	if left.sign && right.sign {
		return isLessThenUint64Array(right.value, left.value)
	}

	return isLessThenUint64Array(left.value, right.value)
}

func isLessThenUint64Array(a, b []uint64) bool {
	if len(a) < len(b) {
		return true
	}

	if len(a) > len(b) {
		return false
	}

	for k := len(a) - 1; k > 0; k-- {
		if a[k] > b[k] {
			return false
		}

		if a[k] < b[k] {
			return true
		}
	}

	return a[0] < b[0]
}
