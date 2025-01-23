package biginteger

import "strconv"

func stringAbs(i BigInteger) string {
	i = i.Abs()

	if i.IsLessThan(Ten) {
		return strconv.FormatUint(i.value[0], 10)
	}

	result := ""
	for i.IsGreaterThan(Zero) {
		remainder := i.Modulo(Ten)
		result = strconv.FormatUint(remainder.value[0], 10) + result
		i = i.Divide(Ten)
	}

	return result
}
