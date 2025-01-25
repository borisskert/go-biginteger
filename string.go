package biginteger

import (
	"fmt"
	"strconv"
)

func toString(i BigInteger) string {
	sign := ""
	if i.sign {
		sign = "-"
	}

	return sign + stringAbs(i)
}

func stringAbs(i BigInteger) string {
	i = i.Abs()

	if i.IsLessThan(e19) {
		return strconv.FormatUint(i.value[0], 10)
	}

	result := ""
	for i.IsGreaterThan(zero) {
		quotient, remainder := i.DivMod(e19)
		i = quotient

		if i.IsGreaterThan(zero) {
			result = fmt.Sprintf("%018d", remainder.value[0]) + result
		} else {
			result = strconv.FormatUint(remainder.value[0], 10) + result
		}
	}

	return result
}
