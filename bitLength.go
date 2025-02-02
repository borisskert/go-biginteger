package biginteger

import "github.com/borisskert/go-biginteger/uintArray"

func bitLength(i BigInteger) uint64 {
	if i.Abs().IsLessThan(two) {
		return 1
	}

	return uintArray.BitLength(i.value)
}
