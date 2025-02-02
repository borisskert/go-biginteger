package digits_test

import (
	. "github.com/borisskert/go-biginteger/divmod"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Digit", func() {
	Context("Hexadecimal", func() {
		It("1", func() {
			a := One()
			Expect(a.Hexadecimal()).To(Equal("0x1"))
		})

		It("2^64-1", func() {
			a := Digit(18446744073709551615)
			Expect(a.Hexadecimal()).To(Equal("0xFFFFFFFFFFFFFFFF"))
		})

		It("9223372036854775808", func() {
			a := Digit(9223372036854775808)
			Expect(a.Hexadecimal()).To(Equal("0x8000000000000000"))
		})

		It("15 * 15", func() {
			a := Digit(15 * 15)
			Expect(a.Hexadecimal()).To(Equal("0xE1")) // 1110 0001
		})

		It("31 * 31", func() {
			a := Digit(31 * 31)
			Expect(a.Hexadecimal()).To(Equal("0x3C1")) // 0011 1100 0001
		})

	})
})
