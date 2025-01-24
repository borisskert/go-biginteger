package biginteger

import (
	"fmt"
	"strconv"
)

var e19 = BigInteger{value: []uint64{1000000000000000000}}

func stringAbs(i BigInteger) string {
	i = i.Abs()

	if i.IsLessThan(e19) {
		return strconv.FormatUint(i.value[0], 10)
	}

	result := ""
	for i.IsGreaterThan(zero) {
		remainder := i.Modulo(e19)
		i = i.Divide(e19)

		if i.IsGreaterThan(zero) {
			result = fmt.Sprintf("%018d", remainder.value[0]) + result
		} else {
			result = strconv.FormatUint(remainder.value[0], 10) + result
		}
	}

	return result
}
