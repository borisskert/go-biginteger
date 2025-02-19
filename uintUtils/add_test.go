package uintUtils

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"testing"
)

var _ = Describe("AddFour64", func() {
	It("Add 0, 0, 0, 0", func() {
		sum, carry := AddFour64(0, 0, 0, 0)
		Expect(sum).To(Equal(uint64(0)))
		Expect(carry).To(Equal(uint64(0)))
	})

	It("Add 1, 0, 0, 0", func() {
		sum, carry := AddFour64(1, 0, 0, 0)
		Expect(sum).To(Equal(uint64(1)))
		Expect(carry).To(Equal(uint64(0)))
	})

	It("Add 2, 1, 0, 0", func() {
		sum, carry := AddFour64(2, 1, 0, 0)
		Expect(sum).To(Equal(uint64(3)))
		Expect(carry).To(Equal(uint64(0)))
	})

	It("Add 3, 2, 1, 0", func() {
		sum, carry := AddFour64(3, 2, 1, 0)
		Expect(sum).To(Equal(uint64(6)))
		Expect(carry).To(Equal(uint64(0)))
	})

	It("Add 4, 3, 2, 1", func() {
		sum, carry := AddFour64(4, 3, 2, 1)
		Expect(sum).To(Equal(uint64(10)))
		Expect(carry).To(Equal(uint64(0)))
	})

	It("Add 5, 4, 3, 2", func() {
		sum, carry := AddFour64(5, 4, 3, 2)
		Expect(sum).To(Equal(uint64(14)))
		Expect(carry).To(Equal(uint64(0)))
	})

	It("Add 18446744073709551615, 0, 0, 0", func() {
		sum, carry := AddFour64(18446744073709551615, 0, 0, 0)
		Expect(sum).To(Equal(uint64(18446744073709551615)))
		Expect(carry).To(Equal(uint64(0)))
	})

	It("Add 18446744073709551615, 18446744073709551615, 0, 0", func() {
		sum, carry := AddFour64(18446744073709551615, 18446744073709551615, 0, 0)
		Expect(sum).To(Equal(uint64(18446744073709551614)))
		Expect(carry).To(Equal(uint64(1)))
	})

	It("Add 18446744073709551615, 18446744073709551615, 18446744073709551615, 0", func() {
		sum, carry := AddFour64(18446744073709551615, 18446744073709551615, 18446744073709551615, 0)
		Expect(sum).To(Equal(uint64(18446744073709551613)))
		Expect(carry).To(Equal(uint64(2)))
	})

	It("Add 18446744073709551615, 18446744073709551615, 18446744073709551615, 3", func() {
		sum, carry := AddFour64(18446744073709551615, 18446744073709551615, 18446744073709551615, 3)
		Expect(sum).To(Equal(uint64(0)))
		Expect(carry).To(Equal(uint64(3)))
	})

	It("Add 18446744073709551615, 18446744073709551615, 18446744073709551615, 18446744073709551615", func() {
		sum, carry := AddFour64(18446744073709551615, 18446744073709551615, 18446744073709551615, 18446744073709551615)
		Expect(sum).To(Equal(uint64(18446744073709551612)))
		Expect(carry).To(Equal(uint64(3)))
	})
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "multiply Test Suite")
}
