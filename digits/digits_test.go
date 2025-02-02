package digits_test

import (
	. "github.com/borisskert/go-biginteger/digits"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Digits", func() {
	Context("SubtractUnderflow", func() {
		It("1 - 2", func() {
			a := One()
			b := Digit(2).AsDigits()

			result, borrow := a.SubtractUnderflow(b)

			Expect(borrow).To(BeTrue())
			Expect(result.Length()).To(Equal(uint64(1)))
			Expect(result.DigitAt(0)).To(Equal(Digit(18446744073709551615)))
		})

		It("2 - 1", func() {
			a := Digit(2).AsDigits()
			b := One()

			result, borrow := a.SubtractUnderflow(b)

			Expect(borrow).To(BeFalse())
			Expect(result.Length()).To(Equal(uint64(1)))
			Expect(result.DigitAt(0)).To(Equal(Digit(1)))
		})

		It("1 - (2 pow 64 - 1)", func() {
			a := One()
			b, _ := Digit(1).AsDigits().LeftShiftBits(64).Subtract(One())

			result, borrow := a.SubtractUnderflow(b.Trim())

			Expect(borrow).To(BeTrue())
			Expect(result.Length()).To(Equal(uint64(1)))
			Expect(result.DigitAt(0)).To(Equal(Digit(2)))
		})

		It("1 - 2 pow 64", func() {
			a := One()
			b := Digit(1).AsDigits().LeftShiftBits(64)

			result, borrow := a.SubtractUnderflow(b)

			Expect(borrow).To(BeTrue())
			Expect(result.Length()).To(Equal(uint64(2)))
			Expect(result.DigitAt(0)).To(Equal(Digit(1)))
			Expect(result.DigitAt(1)).To(Equal(Digit(18446744073709551615)))
		})

		It("2 - 2 pow 64", func() {
			a := One().Add(One())
			b := Digit(1).AsDigits().LeftShiftBits(64)

			result, borrow := a.SubtractUnderflow(b)

			Expect(borrow).To(BeTrue())
			Expect(result.Length()).To(Equal(uint64(2)))
			Expect(result.DigitAt(0)).To(Equal(Digit(2)))
			Expect(result.DigitAt(1)).To(Equal(Digit(18446744073709551615)))
		})
	})
})
