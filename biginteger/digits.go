package biginteger

import "math"

func digitsAbs(i BigInteger) uint64 {
	i = i.Abs()

	if i.IsEqualTo(zero) {
		return 1
	}

	bitLength := i.BitLength().Uint()
	estimatedDigits := uint64(float64(bitLength)/math.Log2(10)) + 1

	divisor := ten.Power(OfUint64(estimatedDigits - 1))
	if i.IsLessThan(divisor) {
		return estimatedDigits - 1
	}

	return estimatedDigits
}
