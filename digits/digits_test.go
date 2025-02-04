package digits_test

import (
	. "github.com/borisskert/go-biginteger/digits"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Digits", func() {
	Context("SubtractUnderflow", func() {
		It("1 - 2", func() {
			a := Digit(1).AsDigits()
			b := Digit(2).AsDigits()

			result, borrow := a.SubtractUnderflow(b)

			Expect(borrow).To(BeTrue())
			Expect(result.Length()).To(Equal(uint(1)))
			Expect(result.DigitAt(0)).To(Equal(Digit(18446744073709551615)))
		})

		It("2 - 1", func() {
			a := Digit(2).AsDigits()
			b := Digit(1).AsDigits()

			result, borrow := a.SubtractUnderflow(b)

			Expect(borrow).To(BeFalse())
			Expect(result.Length()).To(Equal(uint(1)))
			Expect(result.DigitAt(0)).To(Equal(Digit(1)))
		})

		It("1 - (2 pow 64 - 1)", func() {
			a := Digit(1).AsDigits()
			b, _ := Digit(1).AsDigits().LeftShiftBits(64).Decrement()

			result, borrow := a.SubtractUnderflow(b.Trim())

			Expect(borrow).To(BeTrue())
			Expect(result.Length()).To(Equal(uint(1)))
			Expect(result.DigitAt(0)).To(Equal(Digit(2)))
		})

		It("1 - 2 pow 64", func() {
			a := Digit(1).AsDigits()
			b := Digit(1).AsDigits().LeftShiftBits(64)

			result, borrow := a.SubtractUnderflow(b)

			Expect(borrow).To(BeTrue())
			Expect(result.Length()).To(Equal(uint(2)))
			Expect(result.DigitAt(0)).To(Equal(Digit(1)))
			Expect(result.DigitAt(1)).To(Equal(Digit(18446744073709551615)))
		})

		It("2 - 2 pow 64", func() {
			a := Digit(2).AsDigits()
			b := Digit(1).AsDigits().LeftShiftBits(64)

			result, borrow := a.SubtractUnderflow(b)

			Expect(borrow).To(BeTrue())
			Expect(result.Length()).To(Equal(uint(2)))
			Expect(result.DigitAt(0)).To(Equal(Digit(2)))
			Expect(result.DigitAt(1)).To(Equal(Digit(18446744073709551615)))
		})
	})

	Context("MultiplyByDoubleDigit", func() {
		It("1 * 1", func() {
			a := Digit(1).AsDigits()
			b := Digit(1).AsDoubleDigit()

			result := a.MultiplyByDoubleDigit(b)

			Expect(result.Length()).To(Equal(uint(1)))
			Expect(result.DigitAt(0)).To(Equal(Digit(1)))
		})

		It("2 * 2", func() {
			a := Digit(2).AsDigits()
			b := Digit(2).AsDoubleDigit()

			result := a.MultiplyByDoubleDigit(b)

			Expect(result.Length()).To(Equal(uint(1)))
			Expect(result.DigitAt(0)).To(Equal(Digit(4)))
			Expect(result.DigitAt(1)).To(Equal(Digit(0)))
		})

		It("2 pow 64 * 2 pow 64", func() {
			a := Digit(1).AsDigits().LeftShiftBits(64)
			b := Digit(1).AsDigits().LeftShiftBits(64).AsDoubleDigit()

			result := a.MultiplyByDoubleDigit(b)

			Expect(result.Length()).To(Equal(uint(3)))
			Expect(result.DigitAt(0)).To(Equal(Digit(0)))
			Expect(result.DigitAt(1)).To(Equal(Digit(0)))
			Expect(result.DigitAt(2)).To(Equal(Digit(1)))
		})

		It("2 pow 64 * 2 pow 64 - 1", func() {
			a := Digit(1).AsDigits().LeftShiftBits(64)
			b, _ := Digit(1).AsDigits().LeftShiftBits(64).Decrement()

			result := a.MultiplyByDoubleDigit(b.AsDoubleDigit())

			Expect(result.Length()).To(Equal(uint(2)))
			Expect(result.DigitAt(0)).To(Equal(Digit(0)))
			Expect(result.DigitAt(1)).To(Equal(Digit(18446744073709551615)))
		})

		It("(2 pow 64) * (2 pow 64)", func() {
			a, _ := Digit(1).AsDigits().LeftShiftBits(64).Decrement()
			b, _ := Digit(1).AsDigits().LeftShiftBits(64).Decrement()

			result := a.MultiplyByDoubleDigit(b.AsDoubleDigit())

			Expect(result.Length()).To(Equal(uint(2)))
			Expect(result.DigitAt(0)).To(Equal(Digit(1)))
			Expect(result.DigitAt(1)).To(Equal(Digit(18446744073709551614)))
		})
	})
})
