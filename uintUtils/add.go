package uintUtils

import "math/bits"

func AddFour64(a, b, c, d uint64) (uint64, uint64) {
	lo1, carry1 := bits.Add64(a, b, 0)
	lo2, carry2 := bits.Add64(c, d, 0)
	lo, carry3 := bits.Add64(lo1, lo2, 0)

	hi := carry1 + carry2 + carry3

	return lo, hi
}

// AddThree64 adds three uint64 numbers and returns the sum and the carry.
// Because this test for bits.Add64 fails:
//
//	It("Why is this a thing?", func() {
//		sum, out := bits.Add64(0, 0, 18446744073709551615)
//		Expect(sum).To(Equal(uint64(18446744073709551615)))
//		Expect(out).To(Equal(uint64(0)))
//	})
func AddThree64(a, b, c uint64) (uint64, uint64) {
	lo1, carry1 := bits.Add64(a, b, 0)
	lo, carry2 := bits.Add64(lo1, c, 0)

	hi := carry1 + carry2

	return lo, hi
}
