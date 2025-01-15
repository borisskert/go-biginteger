package uint64_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("uint64", func() {
	Context("Divide", func() {
		divUint64 := func(a, b uint64) uint64 {
			if b == 0 {
				panic("division by zero")
			}

			if b == 1 {
				return a
			}

			if a == 0 {
				return 0
			}

			if a < b {
				return 0
			}

			result := 0
			remaining := a
			divisor := b
			quotient := 1

			for {
				if remaining > divisor {
					remaining -= divisor
					result += quotient
					divisor += divisor
					quotient += quotient
				} else if remaining == divisor {
					result += quotient
					break
				} else if remaining < b {
					break
				} else if remaining < divisor {
					divisor = b
					quotient = 1
				} else {
					break
				}
			}

			return uint64(result)
		}

		It("4 / 2 should be 2", func() {
			Expect(divUint64(4, 2)).To(Equal(uint64(2)))
		})

		It("8 / 4 should be 2", func() {
			Expect(divUint64(8, 4)).To(Equal(uint64(2)))
		})

		It("16 / 4 should be 4", func() {
			Expect(divUint64(16, 4)).To(Equal(uint64(4)))
		})

		It("16 / 5 should be 3", func() {
			Expect(divUint64(16, 5)).To(Equal(uint64(3)))
		})

		It("80 / 10 should be 8", func() {
			Expect(divUint64(80, 10)).To(Equal(uint64(8)))
		})

		It("4294967296 / 10 should be 429496729", func() {
			Expect(divUint64(4294967296, 10)).To(Equal(uint64(429496729)))
		})
	})
})
