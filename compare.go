package biginteger

import (
	"github.com/borisskert/go-biginteger/uintArray"
)

func compareTo(left BigInteger, right BigInteger) int {
	if left.sign && !right.sign {
		return -1
	}

	if !left.sign && right.sign {
		return 1
	}

	if left.sign && right.sign {
		return uintArray.Compare(right.value, left.value)
	}

	return uintArray.Compare(left.value, right.value)
}

func compareUint64Arrays(a, b []uint64) int {
	if len(a) < len(b) {
		return -1
	}

	if len(a) > len(b) {
		return 1
	}

	for k := len(a) - 1; k > 0; k-- {
		if a[k] > b[k] {
			return 1
		}

		if a[k] < b[k] {
			return -1
		}
	}

	if a[0] > b[0] {
		return 1
	}

	if a[0] < b[0] {
		return -1
	}

	return 0
}
