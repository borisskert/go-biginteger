package uintUtils

import "math/bits"

func AddFour64(a, b, c, d uint64) (uint64, uint64) {
	sum, carry := bits.Add64(a, b, 0)
	sum, carry = bits.Add64(sum, c, carry)
	sum, carry = bits.Add64(sum, d, carry)

	return sum, carry
}
