package divmod_test

import (
	. "github.com/borisskert/go-biginteger/divmod"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("DoubleDigits", func() {
	Context("Subtract", func() {
		It("[1,0] - [0,1] == [0,2^64-1],0", func() {
			a := DoubleDigitOf(1, 0)
			b := One().AsDoubleDigit()

			result, borrow := a.Subtract(b)

			Expect(borrow).To(Equal(Digit(0)))
			Expect(result.High()).To(Equal(Digit(0)))
			Expect(result.Low()).To(Equal(Digit(18446744073709551615)))
		})

		It("[0,1] - [1,0] == [2^64-1,1],1", func() {
			a := One().AsDoubleDigit()
			b := DoubleDigitOf(1, 0)

			result, borrow := a.Subtract(b)

			Expect(borrow).To(Equal(Digit(1)))
			Expect(result.High()).To(Equal(Digit(18446744073709551615)))
			Expect(result.Low()).To(Equal(Digit(1)))
		})
	})
})
