package uint64_test

import (
	"fmt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("uint64 Multi", func() {

	multiUint64 := func(a, b uint64) uint64 {
		if a == 0 || b == 0 {
			return 0
		}

		if a == 1 {
			return b
		}

		if b == 1 {
			return a
		}

		result := uint64(0)
		factor := uint64(1)
		multiplier := a
		remaining := b

		for {
			if remaining > factor {
				fmt.Print("(remaining > factor) | remaining: ", remaining, " factor: ", factor, " multiplier: ", multiplier, " result: ", result, "\n")

				result += multiplier
				remaining -= factor
				multiplier += multiplier
				factor += factor
			} else if remaining == factor {
				fmt.Print("(remaining == factor) | remaining: ", remaining, " factor: ", factor, " multiplier: ", multiplier, " result: ", result, "\n")

				result += multiplier
				break
			} else if remaining < factor {
				fmt.Print("(remaining < factor) | remaining: ", remaining, " factor: ", factor, " multiplier: ", multiplier, " result: ", result, "\n")

				factor = uint64(1)
				multiplier = a
			} else {
				fmt.Print("(else) | remaining: ", remaining, " factor: ", factor, " multiplier: ", multiplier, " result: ", result, "\n")

				break
			}
		}

		return result
	}

	It("4 * 2 should be 8", func() {
		Expect(multiUint64(4, 2)).To(Equal(uint64(8)))
	})

	It("8 * 4 should be 32", func() {
		Expect(multiUint64(8, 4)).To(Equal(uint64(32)))
	})

	It("16 * 4 should be 64", func() {
		Expect(multiUint64(16, 4)).To(Equal(uint64(64)))
	})

	It("16 * 5 should be 80", func() {
		Expect(multiUint64(16, 5)).To(Equal(uint64(80)))
	})

	It("80 * 10 should be 800", func() {
		Expect(multiUint64(80, 10)).To(Equal(uint64(800)))
	})

	It("Should multiply 10000000000 by 10000000000", func() {
		Expect(multiUint64(1000000000, 1000000000)).
			To(Equal(uint64(1000000000000000000)))
	})
})
