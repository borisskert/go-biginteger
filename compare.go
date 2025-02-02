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
