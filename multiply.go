package biginteger

func multiply(multiplicand BigInteger, multiplier BigInteger) BigInteger {
	if multiplicand.IsEqualTo(zero) || multiplier.IsEqualTo(zero) {
		return zero
	}

	if multiplicand.IsEqualTo(one) {
		return multiplier
	}

	if multiplier.IsEqualTo(one) {
		return multiplicand
	}

	sign := multiplicand.sign != multiplier.sign

	return BigInteger{
		sign:  sign,
		value: multiplyUint64Array(multiplicand.value, multiplier.value),
	}
}

func multiplyUint64Array(a, b []uint64) []uint64 {
	result := make([]uint64, len(a)+len(b))

	for i := 0; i < len(a); i++ {
		var carry uint64 = 0
		for j := 0; j < len(b); j++ {
			high, low := multiplyUint64(a[i], b[j])
			low += result[i+j]

			if low < result[i+j] {
				high++
			}

			result[i+j] = low

			high += carry
			result[i+j+1] += high

			if result[i+j+1] < high {
				carry = 1
			} else {
				carry = 0
			}
		}

		result[i+len(b)] += carry
	}

	return trimLeadingZeros(result)
}

func multiplyUint64(a, b uint64) (uint64, uint64) {
	const mask uint64 = 0xFFFFFFFF // Mask for lower 32 bits

	// Split inputs into lower and upper 32-bit halves
	aLow := a & mask
	aHigh := a >> 32
	bLow := b & mask
	bHigh := b >> 32

	// Calculate partial products
	lowLow := aLow * bLow
	lowHigh := aLow * bHigh
	highLow := aHigh * bLow
	highHigh := aHigh * bHigh

	// Combine results
	carry := (lowLow >> 32) + (lowHigh & mask) + (highLow & mask)
	low := (lowLow & mask) + (carry << 32)
	high := highHigh + (lowHigh >> 32) + (highLow >> 32) + (carry >> 32)

	return high, low
}
