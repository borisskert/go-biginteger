package biginteger_test

import (
	"github.com/borisskert/go-biginteger"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"testing"
)

var _ = Describe("BigInteger", func() {
	Context("Of and String", func() {
		It("Should create 1", func() {
			bigint, err := biginteger.Of("1")

			Expect(err).To(BeNil())
			Expect(bigint.String()).To(Equal("1"))
		})

		It("Should create -1", func() {
			bigint, err := biginteger.Of("-1")

			Expect(err).To(BeNil())
			Expect(bigint.String()).To(Equal("-1"))
		})

		It("Should create 2", func() {
			bigint, err := biginteger.Of("2")

			Expect(err).To(BeNil())
			Expect(bigint.String()).To(Equal("2"))
		})

		It("Should create 4", func() {
			bigint, err := biginteger.Of("4")

			Expect(err).To(BeNil())
			Expect(bigint.String()).To(Equal("4"))
		})

		It("Should create 8", func() {
			bigint, err := biginteger.Of("8")

			Expect(err).To(BeNil())
			Expect(bigint.String()).To(Equal("8"))
		})

		It("Should create 4294967295", func() {
			bigint, err := biginteger.Of("4294967295")

			Expect(err).To(BeNil())
			Expect(bigint.String()).To(Equal("4294967295"))
		})

		It("Should create 4294967296", func() {
			bigint, err := biginteger.Of("4294967296")

			Expect(err).To(BeNil())
			Expect(bigint.String()).To(Equal("4294967296"))
		})

		It("Should create 18446744073709551615", func() {
			bigint, err := biginteger.Of("18446744073709551615")

			Expect(err).To(BeNil())
			Expect(bigint.String()).To(Equal("18446744073709551615"))
		})

		It("Should create 18446744073709551616", func() {
			bigint, err := biginteger.Of("18446744073709551616")

			Expect(err).To(BeNil())
			Expect(bigint.String()).To(Equal("18446744073709551616"))
		})

		It("Should create 36893488147419103232", func() {
			bigint, err := biginteger.Of("36893488147419103232")

			Expect(err).To(BeNil())
			Expect(bigint.String()).To(Equal("36893488147419103232"))
		})

		It("Should create 73786976294838206464", func() {
			bigint, err := biginteger.Of("73786976294838206464")

			Expect(err).To(BeNil())
			Expect(bigint.String()).To(Equal("73786976294838206464"))
		})

		It("Should create 340282366920938463444927863358058659840", func() {
			bigint, err := biginteger.Of("340282366920938463444927863358058659840")

			Expect(err).To(BeNil())
			Expect(bigint.String()).To(Equal("340282366920938463444927863358058659840"))
		})

		It("Should create 20000000000000000000", func() {
			bigint, err := biginteger.Of("20000000000000000000")

			Expect(err).To(BeNil())
			Expect(bigint.String()).To(Equal("20000000000000000000"))
		})
	})

	Context("OfUint64Array", func() {
		var array []uint64
		var a biginteger.BigInteger

		BeforeEach(func() {
			array = []uint64{1, 2, 3}
			a = biginteger.OfUint64Array(array)
		})

		It("Should create instance", func() {
			Expect(a.String()).To(Equal("1020847100762815390427017310442723737601"))
		})

		It("Should be immutable", func() {
			array[0] = 0
			Expect(a.String()).To(Equal("1020847100762815390427017310442723737601"))
		})
	})

	Context("Add", func() {
		It("Should add 1 and 2", func() {
			bigint1, _ := biginteger.Of("1")
			bigint2, _ := biginteger.Of("2")

			Expect(bigint1.Add(*bigint2).String()).To(Equal("3"))
		})

		It("Should add 1 and 0", func() {
			bigint1, _ := biginteger.Of("1")
			bigint2, _ := biginteger.Of("0")

			Expect(bigint1.Add(*bigint2).String()).To(Equal("1"))
		})

		It("Should add 0 and 2", func() {
			bigint1, _ := biginteger.Of("0")
			bigint2, _ := biginteger.Of("2")

			Expect(bigint1.Add(*bigint2).String()).To(Equal("2"))
		})

		It("Should add -1 and 2", func() {
			bigint1, _ := biginteger.Of("-1")
			bigint2, _ := biginteger.Of("2")

			Expect(bigint1.Add(*bigint2).String()).To(Equal("1"))
		})

		It("Should add 1 and -2", func() {
			bigint1, _ := biginteger.Of("1")
			bigint2, _ := biginteger.Of("-2")

			Expect(bigint1.Add(*bigint2).String()).To(Equal("-1"))
		})

		It("Should add -1 and -2", func() {
			bigint1, _ := biginteger.Of("-1")
			bigint2, _ := biginteger.Of("-2")

			Expect(bigint1.Add(*bigint2).String()).To(Equal("-3"))
		})

		It("Should add 2 and 4", func() {
			bigint1, _ := biginteger.Of("2")
			bigint2, _ := biginteger.Of("4")

			Expect(bigint1.Add(*bigint2).String()).To(Equal("6"))
		})

		It("Should add 4 and 8", func() {
			bigint1, _ := biginteger.Of("4")
			bigint2, _ := biginteger.Of("8")

			Expect(bigint1.Add(*bigint2).String()).To(Equal("12"))
		})

		It("Should add 4294967295 and 4294967296", func() {
			bigint1, _ := biginteger.Of("4294967295")
			bigint2, _ := biginteger.Of("4294967296")

			Expect(bigint1.Add(*bigint2).String()).To(Equal("8589934591"))
		})

		It("Should add 4294967296 and 4294967296", func() {
			bigint1, _ := biginteger.Of("4294967296")
			bigint2, _ := biginteger.Of("4294967296")

			Expect(bigint1.Add(*bigint2).String()).To(Equal("8589934592"))
		})

		It("Should add 10000000000000000000 and 10000000000000000000", func() {
			bigint1, _ := biginteger.Of("10000000000000000000")
			bigint2, _ := biginteger.Of("10000000000000000000")
			result := bigint1.Add(*bigint2)

			Expect(result.String()).To(Equal("20000000000000000000"))
		})

		It("Should add 18446744073709551615 and 1", func() {
			bigint1, _ := biginteger.Of("18446744073709551615")
			bigint2, _ := biginteger.Of("1")

			Expect(bigint1.Add(*bigint2).String()).To(Equal("18446744073709551616"))
		})

		It("Should add 18446744073709551615 and 18446744073709551615", func() {
			bigint1, _ := biginteger.Of("18446744073709551615")
			bigint2, _ := biginteger.Of("18446744073709551615")
			sum := bigint1.Add(*bigint2)

			Expect(sum.String()).To(Equal("36893488147419103230"))
		})

		It("Should add 18446744073709551615 and 18446744073709551616", func() {
			bigint1, _ := biginteger.Of("18446744073709551615")
			bigint2, _ := biginteger.Of("18446744073709551616")

			Expect(bigint1.Add(*bigint2).String()).To(Equal("36893488147419103231"))
		})
	})

	Context("Subtract", func() {
		It("Should subtract 2 from 1", func() {
			bigint1, _ := biginteger.Of("1")
			bigint2, _ := biginteger.Of("2")
			result := bigint1.Subtract(*bigint2)

			Expect(result.String()).To(Equal("-1"))
		})

		It("Should subtract 1 from 2", func() {
			bigint1, _ := biginteger.Of("2")
			bigint2, _ := biginteger.Of("1")
			result := bigint1.Subtract(*bigint2)

			Expect(result.String()).To(Equal("1"))
		})

		It("Should subtract -2 from 1", func() {
			bigint1, _ := biginteger.Of("1")
			bigint2, _ := biginteger.Of("-2")
			result := bigint1.Subtract(*bigint2)

			Expect(result.String()).To(Equal("3"))
		})

		It("Should subtract -1 from 2", func() {
			bigint1, _ := biginteger.Of("2")
			bigint2, _ := biginteger.Of("-1")
			result := bigint1.Subtract(*bigint2)

			Expect(result.String()).To(Equal("3"))
		})

		It("Should subtract 2 from -1", func() {
			bigint1, _ := biginteger.Of("-1")
			bigint2, _ := biginteger.Of("2")
			result := bigint1.Subtract(*bigint2)

			Expect(result.String()).To(Equal("-3"))
		})

		It("Should subtract 2 from 2", func() {
			bigint1, _ := biginteger.Of("2")
			bigint2, _ := biginteger.Of("2")
			result := bigint1.Subtract(*bigint2)

			Expect(result.String()).To(Equal("0"))
		})

		It("Should subtract -2 from -2", func() {
			bigint1, _ := biginteger.Of("-2")
			bigint2, _ := biginteger.Of("-2")
			result := bigint1.Subtract(*bigint2)

			Expect(result.String()).To(Equal("0"))
		})

		It("Should subtract -1 from -2", func() {
			bigint1, _ := biginteger.Of("-2")
			bigint2, _ := biginteger.Of("-1")
			result := bigint1.Subtract(*bigint2)

			Expect(result.String()).To(Equal("-3"))
		})

		It("Should subtract 2 from 4", func() {
			bigint1, _ := biginteger.Of("4")
			bigint2, _ := biginteger.Of("2")

			Expect(bigint1.Subtract(*bigint2).String()).To(Equal("2"))
		})

		It("Should subtract 4 from 8", func() {
			bigint1, _ := biginteger.Of("8")
			bigint2, _ := biginteger.Of("4")

			Expect(bigint1.Subtract(*bigint2).String()).To(Equal("4"))
		})

		It("Should subtract 4294967296 from 4294967296", func() {
			bigint1, _ := biginteger.Of("4294967296")
			bigint2, _ := biginteger.Of("4294967296")

			Expect(bigint1.Subtract(*bigint2).String()).To(Equal("0"))
		})

		It("Should subtract 1 from 18446744073709551616", func() {
			bigint1, _ := biginteger.Of("18446744073709551616")
			bigint2, _ := biginteger.Of("1")
			result := bigint1.Subtract(*bigint2)

			Expect(result.String()).To(Equal("18446744073709551615"))
		})

		It("Should subtract 18446744073709551615 from 18446744073709551616", func() {
			bigint1, _ := biginteger.Of("18446744073709551616")
			bigint2, _ := biginteger.Of("18446744073709551615")

			Expect(bigint1.Subtract(*bigint2).String()).To(Equal("1"))
		})

		It("Should subtract 18446744073709551616 from 18446744073709551616", func() {
			bigint1, _ := biginteger.Of("18446744073709551616")
			bigint2, _ := biginteger.Of("18446744073709551616")

			Expect(bigint1.Subtract(*bigint2).String()).To(Equal("0"))
		})

		It("Should subtract 18446744073709551616 from 18446744073709551615", func() {
			bigint1, _ := biginteger.Of("18446744073709551615")
			bigint2, _ := biginteger.Of("18446744073709551616")

			Expect(bigint1.Subtract(*bigint2).String()).To(Equal("-1"))
		})

		It("Should subtract 18446744073709551615 from 36893488147419103232", func() {
			bigint1, _ := biginteger.Of("36893488147419103232")
			bigint2, _ := biginteger.Of("18446744073709551615")

			Expect(bigint1.Subtract(*bigint2).String()).To(Equal("18446744073709551617"))
		})

		It("Should subtract 18446744073709551617 from 36893488147419103232", func() {
			bigint1, _ := biginteger.Of("36893488147419103232")
			bigint2, _ := biginteger.Of("18446744073709551617")

			Expect(bigint1.Subtract(*bigint2).String()).To(Equal("18446744073709551615"))
		})

		It("Should subtract 10 from 11579208923731619542357098500868790785326998466564056403945758400791312963993", func() {
			bigint1, _ := biginteger.Of("11579208923731619542357098500868790785326998466564056403945758400791312963993")
			bigint2, _ := biginteger.Of("10")

			Expect(bigint1.Subtract(*bigint2).String()).To(
				Equal("11579208923731619542357098500868790785326998466564056403945758400791312963983"))
		})

		It("[1, 0, 1] + -[0, 0, 1, 1]", func() {
			b := biginteger.OfUint64Array([]uint64{
				0, 0, 1, 1,
			}).Negate()
			a := biginteger.OfUint64Array([]uint64{
				1, 0, 1,
			})

			result := a.Add(b)

			expected := biginteger.OfUint64Array([]uint64{
				18446744073709551615,
				18446744073709551615,
				18446744073709551615,
			}).Negate()

			Expect(result.IsEqualTo(expected)).To(BeTrue())
		})

		It("[1, 0, 0, 1] + -[0, 0, 0, 1, 1]", func() {
			b := biginteger.OfUint64Array([]uint64{
				0, 0, 0, 1, 1,
			}).Negate()
			a := biginteger.OfUint64Array([]uint64{
				1, 0, 0, 1,
			})

			result := a.Add(b)

			expected := biginteger.OfUint64Array([]uint64{
				18446744073709551615,
				18446744073709551615,
				18446744073709551615,
				18446744073709551615,
			}).Negate()

			Expect(result.IsEqualTo(expected)).To(BeTrue())
		})

		It("[1, 0, 1] - [0, 0, 1, 1]", func() {
			b := biginteger.OfUint64Array([]uint64{
				0, 0, 1, 1,
			})
			a := biginteger.OfUint64Array([]uint64{
				1, 0, 1,
			})

			result := a.Subtract(b)

			expected := biginteger.OfUint64Array([]uint64{
				18446744073709551615,
				18446744073709551615,
				18446744073709551615,
			}).Negate()

			Expect(result.IsEqualTo(expected)).To(BeTrue())
		})
	})

	Context("Multiply", func() {
		It("Should multiply 2 by 1", func() {
			bigint1, _ := biginteger.Of("2")
			bigint2, _ := biginteger.Of("1")

			Expect(bigint1.Multiply(*bigint2).String()).To(Equal("2"))
		})

		It("Should multiply 1 by 2", func() {
			bigint1, _ := biginteger.Of("1")
			bigint2, _ := biginteger.Of("2")

			Expect(bigint1.Multiply(*bigint2).String()).To(Equal("2"))
		})

		It("Should multiply 2 by 0", func() {
			bigint1, _ := biginteger.Of("2")
			bigint2, _ := biginteger.Of("0")

			Expect(bigint1.Multiply(*bigint2).String()).To(Equal("0"))
		})

		It("Should multiply -2 by 0", func() {
			bigint1, _ := biginteger.Of("-2")
			bigint2, _ := biginteger.Of("0")

			Expect(bigint1.Multiply(*bigint2).String()).To(Equal("0"))
		})

		It("Should multiply 0 by 2", func() {
			bigint1, _ := biginteger.Of("0")
			bigint2, _ := biginteger.Of("2")

			Expect(bigint1.Multiply(*bigint2).String()).To(Equal("0"))
		})

		It("Should multiply 0 by -2", func() {
			bigint1, _ := biginteger.Of("0")
			bigint2, _ := biginteger.Of("-2")

			Expect(bigint1.Multiply(*bigint2).String()).To(Equal("0"))
		})

		It("Should multiply 2 by -1", func() {
			bigint1, _ := biginteger.Of("2")
			bigint2, _ := biginteger.Of("-1")

			Expect(bigint1.Multiply(*bigint2).String()).To(Equal("-2"))
		})

		It("Should multiply 2 by 2", func() {
			bigint1, _ := biginteger.Of("2")
			bigint2, _ := biginteger.Of("2")

			Expect(bigint1.Multiply(*bigint2).String()).To(Equal("4"))
		})

		It("Should multiply 2 by -2", func() {
			bigint1, _ := biginteger.Of("2")
			bigint2, _ := biginteger.Of("-2")

			Expect(bigint1.Multiply(*bigint2).String()).To(Equal("-4"))
		})

		It("Should multiply 2 by 4", func() {
			bigint1, _ := biginteger.Of("2")
			bigint2, _ := biginteger.Of("4")

			Expect(bigint1.Multiply(*bigint2).String()).To(Equal("8"))
		})

		It("Should multiply 4 by 8", func() {
			bigint1, _ := biginteger.Of("4")
			bigint2, _ := biginteger.Of("8")

			Expect(bigint1.Multiply(*bigint2).String()).To(Equal("32"))
		})

		It("Should multiply 10000000000 by 10000000000", func() {
			bigint1, _ := biginteger.Of("10000000000")
			bigint2, _ := biginteger.Of("10000000000")

			Expect(bigint1.Multiply(*bigint2).String()).To(Equal("100000000000000000000"))
		})

		It("Should multiply 4294967295 by 4294967296", func() {
			bigint1, _ := biginteger.Of("4294967295")
			bigint2, _ := biginteger.Of("4294967296")

			Expect(bigint1.Multiply(*bigint2).String()).To(Equal("18446744069414584320"))
		})

		It("Should multiply 4294967295 by -4294967296", func() {
			bigint1, _ := biginteger.Of("4294967295")
			bigint2, _ := biginteger.Of("-4294967296")

			Expect(bigint1.Multiply(*bigint2).String()).To(Equal("-18446744069414584320"))
		})

		It("Should multiply [0, 4294967295] by [0, 4294967295]", func() {
			bigint1 := biginteger.OfUint64Array([]uint64{0, 4294967295})
			bigint2 := biginteger.OfUint64Array([]uint64{0, 4294967295})

			result := bigint1.Multiply(bigint2)

			expected := biginteger.OfUint64Array([]uint64{0, 0, 18446744065119617025})

			Expect(result.IsEqualTo(expected)).To(BeTrue())
		})

		It("Should multiply 4294967296 by 4294967296", func() {
			bigint1, _ := biginteger.Of("4294967296")
			bigint2, _ := biginteger.Of("4294967296")
			result := bigint1.Multiply(*bigint2)

			Expect(result.String()).To(Equal("18446744073709551616"))
		})

		It("Should multiply 18446744073709551615 by 1", func() {
			bigint1, _ := biginteger.Of("18446744073709551615")
			bigint2, _ := biginteger.Of("1")

			Expect(bigint1.Multiply(*bigint2).String()).To(Equal("18446744073709551615"))
		})

		It("Should multiply 18446744073709551615 by 2", func() {
			bigint1, _ := biginteger.Of("18446744073709551615")
			bigint2, _ := biginteger.Of("2")

			Expect(bigint1.Multiply(*bigint2).String()).To(Equal("36893488147419103230"))
		})

		It("Should multiply 18446744073709551615 by 18446744073709551616", func() {
			bigint1, _ := biginteger.Of("18446744073709551615")
			bigint2, _ := biginteger.Of("18446744073709551616")
			result := bigint1.Multiply(*bigint2)

			Expect(result.String()).To(Equal("340282366920938463444927863358058659840"))
		})

		It("[2, 3] * [1]", func() {
			bigint1 := biginteger.Two().
				Add(biginteger.OfUint64(3).Multiply(biginteger.Two().Power(biginteger.OfUint64(64))))
			bigint2 := biginteger.OfUint64(1)

			result := bigint1.Multiply(bigint2)

			Expect(result.String()).To(Equal("55340232221128654850"))
		})

		It("[0, 0, 3] * [0, 1]", func() {
			bigint1 := biginteger.OfUint64Array([]uint64{0, 0, 3})
			bigint2 := biginteger.OfUint64Array([]uint64{0, 1})

			result := bigint1.Multiply(bigint2)
			expected := biginteger.OfUint64Array([]uint64{0, 0, 0, 3})

			Expect(result.String()).To(Equal(expected.String()))
		})

		It("[0, 0, 3] * [0, 2]", func() {
			bigint1 := biginteger.OfUint64Array([]uint64{0, 0, 3})
			bigint2 := biginteger.OfUint64Array([]uint64{0, 2})

			result := bigint1.Multiply(bigint2)
			expected := biginteger.OfUint64Array([]uint64{0, 0, 0, 6})

			Expect(result.String()).To(Equal(expected.String()))
		})

		It("[0, 0, 3] * [0, 0, 2]", func() {
			bigint1 := biginteger.OfUint64Array([]uint64{0, 0, 3})
			bigint2 := biginteger.OfUint64Array([]uint64{0, 0, 2})

			result := bigint1.Multiply(bigint2)
			expected := biginteger.OfUint64Array([]uint64{0, 0, 0, 0, 6})

			Expect(result.String()).To(Equal(expected.String()))
		})

		It("[1, 0, 3] * [0, 0, 2]", func() {
			bigint1 := biginteger.OfUint64Array([]uint64{1, 0, 3})
			bigint2 := biginteger.OfUint64Array([]uint64{0, 0, 2})

			result := bigint1.Multiply(bigint2)
			expected := biginteger.OfUint64Array([]uint64{0, 0, 2, 0, 6})

			Expect(result.String()).To(Equal(expected.String()))
		})

		It("[1, 1, 3] * [0, 0, 2]", func() {
			bigint1 := biginteger.OfUint64Array([]uint64{1, 1, 3})
			bigint2 := biginteger.OfUint64Array([]uint64{0, 0, 2})

			result := bigint1.Multiply(bigint2)
			expected := biginteger.OfUint64Array([]uint64{0, 0, 2, 2, 6})

			Expect(result.String()).To(Equal(expected.String()))
		})

		It("[0, 2, 0] * [0, 3, 0]", func() {
			bigint1 := biginteger.OfUint64Array([]uint64{0, 2, 0})
			bigint2 := biginteger.OfUint64Array([]uint64{0, 3, 0})

			result := bigint1.Multiply(bigint2)
			expected := biginteger.OfUint64Array([]uint64{0, 0, 6})

			Expect(result.String()).To(Equal(expected.String()))
		})

		It("[0, 2, 3] * [1, 5, 7]", func() {
			bigint1 := biginteger.OfUint64Array([]uint64{0, 2, 3})
			bigint2 := biginteger.OfUint64Array([]uint64{1, 5, 7})

			result := bigint1.Multiply(bigint2)
			expected := biginteger.OfUint64Array([]uint64{0, 2, 13, 29, 21})

			Expect(result.String()).To(Equal(expected.String()))
		})

		It("[1] * [2, 3]", func() {
			bigint2 := biginteger.Two().
				Add(biginteger.OfUint64(3).Multiply(biginteger.Two().Power(biginteger.OfUint64(64))))
			bigint1 := biginteger.OfUint64(1)

			result := bigint1.Multiply(bigint2)

			Expect(result.String()).To(Equal("55340232221128654850"))
		})

		It("[2, 3] * [5, 7]", func() {
			bigint1 := biginteger.Two().
				Add(biginteger.OfUint64(3).Multiply(biginteger.Two().Power(biginteger.OfUint64(64))))
			bigint2 := biginteger.OfUint64(5).
				Add(biginteger.OfUint64(7).Multiply(biginteger.Two().Power(biginteger.OfUint64(64))))

			result := bigint1.Multiply(bigint2)

			Expect(result.String()).To(Equal("7145929705339707733265822334204709437450"))
		})

		It("-[2, 3] * [5, 7]", func() {
			bigint1 := biginteger.Two().
				Add(biginteger.OfUint64(3).Multiply(biginteger.Two().Power(biginteger.OfUint64(64)))).
				Negate()
			bigint2 := biginteger.OfUint64(5).
				Add(biginteger.OfUint64(7).Multiply(biginteger.Two().Power(biginteger.OfUint64(64))))

			result := bigint1.Multiply(bigint2)

			Expect(result.String()).To(Equal("-7145929705339707733265822334204709437450"))
		})

		It("[2, 3] * -[5, 7]", func() {
			bigint1 := biginteger.Two().
				Add(biginteger.OfUint64(3).Multiply(biginteger.Two().Power(biginteger.OfUint64(64))))
			bigint2 := biginteger.OfUint64(5).
				Add(biginteger.OfUint64(7).Multiply(biginteger.Two().Power(biginteger.OfUint64(64)))).
				Negate()

			result := bigint1.Multiply(bigint2)

			Expect(result.String()).To(Equal("-7145929705339707733265822334204709437450"))
		})

		It("-[2, 3] * -[5, 7]", func() {
			bigint1 := biginteger.Two().
				Add(biginteger.OfUint64(3).Multiply(biginteger.Two().Power(biginteger.OfUint64(64)))).
				Negate()
			bigint2 := biginteger.OfUint64(5).
				Add(biginteger.OfUint64(7).Multiply(biginteger.Two().Power(biginteger.OfUint64(64)))).
				Negate()

			result := bigint1.Multiply(bigint2)

			Expect(result.String()).To(Equal("7145929705339707733265822334204709437450"))
		})

		It("[2, 3, 5] * [7, 11, 13]", func() {
			bigint1 := biginteger.Two().
				Add(biginteger.OfUint64(3).Multiply(biginteger.Two().Power(biginteger.OfUint64(64)))).
				Add(biginteger.OfUint64(5).Multiply(biginteger.Two().Power(biginteger.OfUint64(64 * 2))))
			bigint2 := biginteger.OfUint64(7).
				Add(biginteger.OfUint64(11).Multiply(biginteger.Two().Power(biginteger.OfUint64(64)))).
				Add(biginteger.OfUint64(13).Multiply(biginteger.Two().Power(biginteger.OfUint64(64 * 2))))

			result := bigint1.Multiply(bigint2)

			Expect(result.String()).To(Equal("7526485800425552703122161588691062002295099751538725521244714795387728393404430"))
		})

		It("[2, 3, 5] * -[7, 11, 13]", func() {
			bigint1 := biginteger.Two().
				Add(biginteger.OfUint64(3).Multiply(biginteger.Two().Power(biginteger.OfUint64(64)))).
				Add(biginteger.OfUint64(5).Multiply(biginteger.Two().Power(biginteger.OfUint64(64 * 2))))
			bigint2 := biginteger.OfUint64(7).
				Add(biginteger.OfUint64(11).Multiply(biginteger.Two().Power(biginteger.OfUint64(64)))).
				Add(biginteger.OfUint64(13).Multiply(biginteger.Two().Power(biginteger.OfUint64(64 * 2)))).
				Negate()

			result := bigint1.Multiply(bigint2)

			Expect(result.String()).To(Equal("-7526485800425552703122161588691062002295099751538725521244714795387728393404430"))
		})

		It("-[2, 3, 5] * [7, 11, 13]", func() {
			bigint1 := biginteger.Two().
				Add(biginteger.OfUint64(3).Multiply(biginteger.Two().Power(biginteger.OfUint64(64)))).
				Add(biginteger.OfUint64(5).Multiply(biginteger.Two().Power(biginteger.OfUint64(64 * 2)))).
				Negate()
			bigint2 := biginteger.OfUint64(7).
				Add(biginteger.OfUint64(11).Multiply(biginteger.Two().Power(biginteger.OfUint64(64)))).
				Add(biginteger.OfUint64(13).Multiply(biginteger.Two().Power(biginteger.OfUint64(64 * 2))))

			result := bigint1.Multiply(bigint2)

			Expect(result.String()).To(Equal("-7526485800425552703122161588691062002295099751538725521244714795387728393404430"))
		})

		It("-[2, 3, 5] * -[7, 11, 13]", func() {
			bigint1 := biginteger.Two().
				Add(biginteger.OfUint64(3).Multiply(biginteger.Two().Power(biginteger.OfUint64(64)))).
				Add(biginteger.OfUint64(5).Multiply(biginteger.Two().Power(biginteger.OfUint64(64 * 2)))).
				Negate()
			bigint2 := biginteger.OfUint64(7).
				Add(biginteger.OfUint64(11).Multiply(biginteger.Two().Power(biginteger.OfUint64(64)))).
				Add(biginteger.OfUint64(13).Multiply(biginteger.Two().Power(biginteger.OfUint64(64 * 2)))).
				Negate()

			result := bigint1.Multiply(bigint2)

			Expect(result.String()).To(Equal("7526485800425552703122161588691062002295099751538725521244714795387728393404430"))
		})

		It("[2, 3, 5, 7] * [11, 13, 17, 19]", func() {
			bigint1 := biginteger.Two().
				Add(biginteger.OfUint64(3).Multiply(biginteger.Two().Power(biginteger.OfUint64(64)))).
				Add(biginteger.OfUint64(5).Multiply(biginteger.Two().Power(biginteger.OfUint64(64 * 2)))).
				Add(biginteger.OfUint64(7).Multiply(biginteger.Two().Power(biginteger.OfUint64(64 * 3))))
			bigint2 := biginteger.OfUint64(11).
				Add(biginteger.OfUint64(13).Multiply(biginteger.Two().Power(biginteger.OfUint64(64)))).
				Add(biginteger.OfUint64(17).Multiply(biginteger.Two().Power(biginteger.OfUint64(64 * 2)))).
				Add(biginteger.OfUint64(19).Multiply(biginteger.Two().Power(biginteger.OfUint64(64 * 3))))

			result := bigint1.Multiply(bigint2)

			Expect(result.String()).To(Equal("5240466824120465735690213559006175393735119524884483232897243840556903294164097445318873535703388519421361975688429590"))
		})

		It("-[2, 3, 5, 7] * [11, 13, 17, 19]", func() {
			bigint1 := biginteger.Two().
				Add(biginteger.OfUint64(3).Multiply(biginteger.Two().Power(biginteger.OfUint64(64)))).
				Add(biginteger.OfUint64(5).Multiply(biginteger.Two().Power(biginteger.OfUint64(64 * 2)))).
				Add(biginteger.OfUint64(7).Multiply(biginteger.Two().Power(biginteger.OfUint64(64 * 3)))).
				Negate()
			bigint2 := biginteger.OfUint64(11).
				Add(biginteger.OfUint64(13).Multiply(biginteger.Two().Power(biginteger.OfUint64(64)))).
				Add(biginteger.OfUint64(17).Multiply(biginteger.Two().Power(biginteger.OfUint64(64 * 2)))).
				Add(biginteger.OfUint64(19).Multiply(biginteger.Two().Power(biginteger.OfUint64(64 * 3))))

			result := bigint1.Multiply(bigint2)

			Expect(result.String()).To(Equal("-5240466824120465735690213559006175393735119524884483232897243840556903294164097445318873535703388519421361975688429590"))
		})

		It("[2, 3, 5, 7] * -[11, 13, 17, 19]", func() {
			bigint1 := biginteger.Two().
				Add(biginteger.OfUint64(3).Multiply(biginteger.Two().Power(biginteger.OfUint64(64)))).
				Add(biginteger.OfUint64(5).Multiply(biginteger.Two().Power(biginteger.OfUint64(64 * 2)))).
				Add(biginteger.OfUint64(7).Multiply(biginteger.Two().Power(biginteger.OfUint64(64 * 3))))
			bigint2 := biginteger.OfUint64(11).
				Add(biginteger.OfUint64(13).Multiply(biginteger.Two().Power(biginteger.OfUint64(64)))).
				Add(biginteger.OfUint64(17).Multiply(biginteger.Two().Power(biginteger.OfUint64(64 * 2)))).
				Add(biginteger.OfUint64(19).Multiply(biginteger.Two().Power(biginteger.OfUint64(64 * 3)))).
				Negate()

			result := bigint1.Multiply(bigint2)

			Expect(result.String()).To(Equal("-5240466824120465735690213559006175393735119524884483232897243840556903294164097445318873535703388519421361975688429590"))
		})

		It("-[2, 3, 5, 7] * -[11, 13, 17, 19]", func() {
			bigint1 := biginteger.Two().
				Add(biginteger.OfUint64(3).Multiply(biginteger.Two().Power(biginteger.OfUint64(64)))).
				Add(biginteger.OfUint64(5).Multiply(biginteger.Two().Power(biginteger.OfUint64(64 * 2)))).
				Add(biginteger.OfUint64(7).Multiply(biginteger.Two().Power(biginteger.OfUint64(64 * 3)))).
				Negate()
			bigint2 := biginteger.OfUint64(11).
				Add(biginteger.OfUint64(13).Multiply(biginteger.Two().Power(biginteger.OfUint64(64)))).
				Add(biginteger.OfUint64(17).Multiply(biginteger.Two().Power(biginteger.OfUint64(64 * 2)))).
				Add(biginteger.OfUint64(19).Multiply(biginteger.Two().Power(biginteger.OfUint64(64 * 3)))).
				Negate()

			result := bigint1.Multiply(bigint2)

			Expect(result.String()).To(Equal("5240466824120465735690213559006175393735119524884483232897243840556903294164097445318873535703388519421361975688429590"))
		})

		It("[2, 3, 5, 7, 11] * [13, 17, 19, 23, 29]", func() {
			bigint1 := biginteger.Two().
				Add(biginteger.OfUint64(3).Multiply(biginteger.Two().Power(biginteger.OfUint64(64)))).
				Add(biginteger.OfUint64(5).Multiply(biginteger.Two().Power(biginteger.OfUint64(64 * 2)))).
				Add(biginteger.OfUint64(7).Multiply(biginteger.Two().Power(biginteger.OfUint64(64 * 3)))).
				Add(biginteger.OfUint64(11).Multiply(biginteger.Two().Power(biginteger.OfUint64(64 * 4))))
			bigint2 := biginteger.OfUint64(13).
				Add(biginteger.OfUint64(17).Multiply(biginteger.Two().Power(biginteger.OfUint64(64)))).
				Add(biginteger.OfUint64(19).Multiply(biginteger.Two().Power(biginteger.OfUint64(64 * 2)))).
				Add(biginteger.OfUint64(23).Multiply(biginteger.Two().Power(biginteger.OfUint64(64 * 3)))).
				Add(biginteger.OfUint64(29).Multiply(biginteger.Two().Power(biginteger.OfUint64(64 * 4))))

			result := bigint1.Multiply(bigint2)

			Expect(result.String()).To(Equal("4277090729651688475095552432706461656776701386357046711954515251904984437254503326852270488771354389001848379176532800563528495084040123606862209293454671898"))
		})

		It("-[2, 3, 5, 7, 11] * [13, 17, 19, 23, 29]", func() {
			bigint1 := biginteger.Two().
				Add(biginteger.OfUint64(3).Multiply(biginteger.Two().Power(biginteger.OfUint64(64)))).
				Add(biginteger.OfUint64(5).Multiply(biginteger.Two().Power(biginteger.OfUint64(64 * 2)))).
				Add(biginteger.OfUint64(7).Multiply(biginteger.Two().Power(biginteger.OfUint64(64 * 3)))).
				Add(biginteger.OfUint64(11).Multiply(biginteger.Two().Power(biginteger.OfUint64(64 * 4)))).
				Negate()
			bigint2 := biginteger.OfUint64(13).
				Add(biginteger.OfUint64(17).Multiply(biginteger.Two().Power(biginteger.OfUint64(64)))).
				Add(biginteger.OfUint64(19).Multiply(biginteger.Two().Power(biginteger.OfUint64(64 * 2)))).
				Add(biginteger.OfUint64(23).Multiply(biginteger.Two().Power(biginteger.OfUint64(64 * 3)))).
				Add(biginteger.OfUint64(29).Multiply(biginteger.Two().Power(biginteger.OfUint64(64 * 4))))

			result := bigint1.Multiply(bigint2)

			Expect(result.String()).To(Equal("-4277090729651688475095552432706461656776701386357046711954515251904984437254503326852270488771354389001848379176532800563528495084040123606862209293454671898"))
		})

		It("[2, 3, 5, 7, 11] * -[13, 17, 19, 23, 29]", func() {
			bigint1 := biginteger.Two().
				Add(biginteger.OfUint64(3).Multiply(biginteger.Two().Power(biginteger.OfUint64(64)))).
				Add(biginteger.OfUint64(5).Multiply(biginteger.Two().Power(biginteger.OfUint64(64 * 2)))).
				Add(biginteger.OfUint64(7).Multiply(biginteger.Two().Power(biginteger.OfUint64(64 * 3)))).
				Add(biginteger.OfUint64(11).Multiply(biginteger.Two().Power(biginteger.OfUint64(64 * 4))))
			bigint2 := biginteger.OfUint64(13).
				Add(biginteger.OfUint64(17).Multiply(biginteger.Two().Power(biginteger.OfUint64(64)))).
				Add(biginteger.OfUint64(19).Multiply(biginteger.Two().Power(biginteger.OfUint64(64 * 2)))).
				Add(biginteger.OfUint64(23).Multiply(biginteger.Two().Power(biginteger.OfUint64(64 * 3)))).
				Add(biginteger.OfUint64(29).Multiply(biginteger.Two().Power(biginteger.OfUint64(64 * 4)))).
				Negate()

			result := bigint1.Multiply(bigint2)

			Expect(result.String()).To(Equal("-4277090729651688475095552432706461656776701386357046711954515251904984437254503326852270488771354389001848379176532800563528495084040123606862209293454671898"))
		})

		It("-[2, 3, 5, 7, 11] * -[13, 17, 19, 23, 29]", func() {
			bigint1 := biginteger.Two().
				Add(biginteger.OfUint64(3).Multiply(biginteger.Two().Power(biginteger.OfUint64(64)))).
				Add(biginteger.OfUint64(5).Multiply(biginteger.Two().Power(biginteger.OfUint64(64 * 2)))).
				Add(biginteger.OfUint64(7).Multiply(biginteger.Two().Power(biginteger.OfUint64(64 * 3)))).
				Add(biginteger.OfUint64(11).Multiply(biginteger.Two().Power(biginteger.OfUint64(64 * 4)))).
				Negate()
			bigint2 := biginteger.OfUint64(13).
				Add(biginteger.OfUint64(17).Multiply(biginteger.Two().Power(biginteger.OfUint64(64)))).
				Add(biginteger.OfUint64(19).Multiply(biginteger.Two().Power(biginteger.OfUint64(64 * 2)))).
				Add(biginteger.OfUint64(23).Multiply(biginteger.Two().Power(biginteger.OfUint64(64 * 3)))).
				Add(biginteger.OfUint64(29).Multiply(biginteger.Two().Power(biginteger.OfUint64(64 * 4)))).
				Negate()

			result := bigint1.Multiply(bigint2)

			Expect(result.String()).To(Equal("4277090729651688475095552432706461656776701386357046711954515251904984437254503326852270488771354389001848379176532800563528495084040123606862209293454671898"))
		})

		It("[2, 3, 5, 7, 11, 13] * [17, 19, 23, 29, 31, 37]", func() {
			bigint1 := biginteger.OfUint64Array([]uint64{2, 3, 5, 7, 11, 13})
			bigint2 := biginteger.OfUint64Array([]uint64{17, 19, 23, 29, 31, 37})

			result := bigint1.Multiply(bigint2)

			Expect(result.String()).To(Equal("2194533937076275900366741167535054799875749382360025467904073191803561576452710729579359653785381716606050939386956344352048930805278246717119350772195259030481664056998925675893250901174357327906"))
		})

		It("-[2, 3, 5, 7, 11, 13] * [17, 19, 23, 29, 31, 37]", func() {
			bigint1 := biginteger.OfUint64Array([]uint64{2, 3, 5, 7, 11, 13}).Negate()
			bigint2 := biginteger.OfUint64Array([]uint64{17, 19, 23, 29, 31, 37})

			result := bigint1.Multiply(bigint2)

			Expect(result.String()).To(Equal("-2194533937076275900366741167535054799875749382360025467904073191803561576452710729579359653785381716606050939386956344352048930805278246717119350772195259030481664056998925675893250901174357327906"))
		})

		It("[2, 3, 5, 7, 11, 13] * -[17, 19, 23, 29, 31, 37]", func() {
			bigint1 := biginteger.OfUint64Array([]uint64{2, 3, 5, 7, 11, 13})
			bigint2 := biginteger.OfUint64Array([]uint64{17, 19, 23, 29, 31, 37}).Negate()

			result := bigint1.Multiply(bigint2)

			Expect(result.String()).To(Equal("-2194533937076275900366741167535054799875749382360025467904073191803561576452710729579359653785381716606050939386956344352048930805278246717119350772195259030481664056998925675893250901174357327906"))
		})

		It("-[2, 3, 5, 7, 11, 13] * -[17, 19, 23, 29, 31, 37]", func() {
			bigint1 := biginteger.OfUint64Array([]uint64{2, 3, 5, 7, 11, 13}).Negate()
			bigint2 := biginteger.OfUint64Array([]uint64{17, 19, 23, 29, 31, 37}).Negate()

			result := bigint1.Multiply(bigint2)

			Expect(result.String()).To(Equal("2194533937076275900366741167535054799875749382360025467904073191803561576452710729579359653785381716606050939386956344352048930805278246717119350772195259030481664056998925675893250901174357327906"))
		})

		It("-[2, 3, 5, 7, 11] * -[17, 19, 23, 29, 31, 37]", func() {
			bigint1 := biginteger.OfUint64Array([]uint64{2, 3, 5, 7, 11}).Negate()
			bigint2 := biginteger.OfUint64Array([]uint64{17, 19, 23, 29, 31, 37}).Negate()

			result := bigint1.Multiply(bigint2)

			Expect(result.String()).To(Equal("100663473399553545370669210812961964604566924803333859089625794036125654728254304435718191721276530529593616055713017888717257920343814613824652276020843748436229242230746906658"))
		})

		It("-[2, 3, 5, 7] * -[17, 19, 23, 29, 31, 37]", func() {
			bigint1 := biginteger.OfUint64Array([]uint64{2, 3, 5, 7}).Negate()
			bigint2 := biginteger.OfUint64Array([]uint64{17, 19, 23, 29, 31, 37}).Negate()

			result := bigint1.Multiply(bigint2)

			Expect(result.String()).To(Equal("3472622253855132649081861641702148117036463459210516718392135432595196689365808939506596144213874079967253209885808263144992544392451121881350518022663569442"))
		})

		It("-[2, 3, 5] * -[17, 19, 23, 29, 31, 37]", func() {
			bigint1 := biginteger.OfUint64Array([]uint64{2, 3, 5}).Negate()
			bigint2 := biginteger.OfUint64Array([]uint64{17, 19, 23, 29, 31, 37}).Negate()

			result := bigint1.Multiply(bigint2)

			Expect(result.String()).To(Equal("134465163994687274762105838107521770326556304348999366369964713378245628850002121928477531488290893427306456222684114611038808332300189730"))
		})

		It("-[2, 3] * -[17, 19, 23, 29, 31, 37]", func() {
			bigint1 := biginteger.OfUint64Array([]uint64{2, 3}).Negate()
			bigint2 := biginteger.OfUint64Array([]uint64{17, 19, 23, 29, 31, 37}).Negate()

			result := bigint1.Multiply(bigint2)

			Expect(result.String()).To(Equal("4373622687799787192919683286114733116141072705248339896399821491791524735944525748429014922836055156486784396417826850"))
		})

		It("[2, 5, 1] * [0, 3, 1]", func() {
			bigint1 := biginteger.OfUint64Array([]uint64{2, 5, 1})
			bigint2 := biginteger.OfUint64Array([]uint64{0, 3, 1})

			result := bigint1.Multiply(bigint2)

			Expect(result.String()).To(Equal("115792089237316195473787798891781353969741100288957849247264476354407722647552"))
		})

		It("Should multiply 340282366920938463463374607431768211456 by 340282366920938463463374607431768211456", func() {
			bigint1, _ := biginteger.Of("340282366920938463463374607431768211456")
			bigint2, _ := biginteger.Of("340282366920938463463374607431768211456")
			result := bigint1.Multiply(*bigint2)

			Expect(result.String()).To(
				Equal("115792089237316195423570985008687907853269984665640564039457584007913129639936"))
		})

		It("Should multiply 340282366920938463444927863358058659840 by 340282366920938463444927863358058659841", func() {
			bigint1, _ := biginteger.Of("340282366920938463444927863358058659840")
			bigint2, _ := biginteger.Of("340282366920938463444927863358058659841")
			result := bigint1.Multiply(*bigint2)

			Expect(result.String()).To(
				Equal("115792089237316195411016781537914546326278970553067108134161175589774887485440"))
		})

		It("Should multiply 115792089237316195423570985008687907853269984665640564039457584007913129639936 by 115792089237316195423570985008687907853269984665640564039457584007913129639936", func() {
			bigint1, _ := biginteger.Of("115792089237316195423570985008687907853269984665640564039457584007913129639936")
			bigint2, _ := biginteger.Of("115792089237316195423570985008687907853269984665640564039457584007913129639936")
			result := bigint1.Multiply(*bigint2)

			Expect(result.String()).To(
				Equal("13407807929942597099574024998205846127479365820592393377723561443721764030073546976801874298166903427690031858186486050853753882811946569946433649006084096"))
		})

		It("Should multiply 115792089237316195423570985008687907853269984665640564039457584007913129639936 by 115792089237316195423570985008687907853269984665640564039457584007913129639935", func() {
			bigint1, _ := biginteger.Of("115792089237316195423570985008687907853269984665640564039457584007913129639936")
			bigint2, _ := biginteger.Of("115792089237316195423570985008687907853269984665640564039457584007913129639935")
			result := bigint1.Multiply(*bigint2)

			Expect(result.String()).To(
				Equal("13407807929942597099574024998205846127479365820592393377723561443721764030073431184712636981971479856705023170278632780869088242247907112362425735876444160"))
		})

		It("Should multiply 1 * (2 pow (1 * 64)) + 2 * (2 pow (2 * 64)) + 3 * (2 pow (3 * 64)) + 4 * (2 pow (4 * 64)) by 5 * (2 pow (1 * 64)) + 6 * (2 pow (2 * 64)) + 7 * (2 pow (3 * 64)) + 8 * (2 pow (4 * 64))", func() {
			// 1 * (2 ^ (1 * 64)) + 2 * (2 ^ (2 * 64)) + 3 * (2 ^ (3 * 64)) + 4 * (2 ^ (4 * 64))
			bigint1, _ := biginteger.Of("463168356949264781713115245240911673705267871666027132333082598323981868072960")

			// 5 * (2 ^ (1 * 64)) + 6 * (2 ^ (2 * 64)) + 7 * (2 ^ (3 * 64)) + 8 * (2 ^ (4 * 64))
			bigint2, _ := biginteger.Of("926336713898529563432507592217210028175052097489103808009249641539512435736576")
			result := bigint1.Multiply(*bigint2)

			Expect(result.String()).To(
				Equal("429049853758163107224164413605958634390308066647112887105620497623101003144048495943955854647503111715161403869100645633766616505135266290336675715308584960"))
		})

		It("Should multiply 1 * (2 pow (1 * 64)) + 2 * (2 pow (2 * 64)) + 3 * (2 pow (3 * 64)) + 4 * (2 pow (4 * 64)) + 5 * (2 pow (5 * 64)) + 6 * (2 pow (6 * 64)) by 7 * (2 pow (1 * 64)) + 8 * (2 pow (2 * 64)) + 9 * (2 pow (3 * 64)) + 10 * (2 pow (4 * 64)) + 11 * (2 pow (5 * 64)) + 12 * (2 pow (6 * 64))", func() {
			// 1 * (2 ^ (1 * 64)) + 2 * (2 ^ (2 * 64)) + 3 * (2 ^ (3 * 64)) + 4 * (2 ^ (4 * 64)) + 5 * (2 ^ (5 * 64)) + 6 * (2 ^ (6 * 64))
			bigint1, _ := biginteger.Of("236412037178366875284354175780466233242916712510589705362416398282498506152423739579262723290545779193278560244596736")

			// 7 * (2 ^ (1 * 64)) + 8 * (2 ^ (2 * 64)) + 9 * (2 ^ (3 * 64)) + 10 * (2 ^ (4 * 64)) + 11 * (2 ^ (5 * 64)) + 12 * (2 ^ (6 * 64))
			bigint2, _ := biginteger.Of("472824074356733750570844338596853376568460030905823595230301474362577200932793929233416254188053366903079396559552512")
			result := bigint1.Multiply(*bigint2)

			Expect(result.String()).To(
				Equal("111781302645651043341330957146821704969837306656960913212445251545967385814394946310465628903825385247865954089346563586529231026931576869977620392152317636317939302391742552473990007097776303763130396778309561449760511058894055800832"))
		})

		It("Should multiply 2003... (19k digits) ...56736 by itself", func() {
			bigint1, _ := biginteger.Of("2003529930406846464979072351560255750447825475569751419265016973710894059556311453089506130880933348101038234342907263181822949382118812668869506364761547029165041871916351587966347219442930927982084309104855990570159318959639524863372367203002916969592156108764948889254090805911457037675208500206671563702366126359747144807111774815880914135742720967190151836282560618091458852699826141425030123391108273603843767876449043205960379124490905707560314035076162562476031863793126484703743782954975613770981604614413308692118102485959152380195331030292162800160568670105651646750568038741529463842244845292537361442533614373729088303794601274724958414864915930647252015155693922628180691650796381064132275307267143998158508811292628901134237782705567421080070065283963322155077831214288551675554073345107213112427399562982719769150054883905223804357045848197956393157853510018992000024141963706813559840464039472194016069517690156119726982337890017641517190051133466306898140219383481435426387306539552969691388024158161859561100640362119796101859534802787167200122604642492385111393400464351623867567078745259464670903886547743483217897012764455529409092021959585751622973333576159552394885297579954028471943529913543763705986928913757153740001986394332464890052543106629669165243419174691389632476560289415199775477703138064781342309596190960654591300890188887588084733625956065444888501447335706058817090162108499714529568344061979690565469813631162053579369791403236328496233046421066136200220175787851857409162050489711781820400187282939943446186224328009837323764931814789848119452713007440220765680910376203999203492023906626264491909167985461515778839060397720759279378852241294301017458086862263369284725851403039615558564330385450688652213114813638408384778263790459607186876728509763471271988890680478243230394718650525660978150729861141430305816927924971409161059417185352275887504477592218301158780701975535722241400019548102005661773589781499532325208589753463547007786690406429016763808161740550405117670093673202804549339027992491867306539931640720492238474815280619166900933805732120816350707634351669869625020969023162859350071874190579161241536897514808261904847946571736601005892476655445840838334790544144817684255327207315586349347605137419779525190365032198020108764738368682531025183377533908861426184800374008082238104076468878471647552945326947661700424461063311238021134588694532200116564076327023074292426051582811070387018345324567635625951430032037432740780879056283663406965030844225855967039271869461158513793386475699748568670079823960604393478850861649260304945061743412365828352144806726676841807083754862211408236579802961200027441324438432402331257403545019352428776430880232850855886089962774458164680857875115807014743763867976955049991643998284357290415378143438847303484261903388841494031366139854257635577105335580206622185577060082551288893332226436281984838613239570676191409638533832374343758830859233722284644287996245605476932428998432652677378373173288063210753211238680604674708428051166488709084770291208161104912555598322366244868556651402684641209694982590565519216188104341226838996283071654868525536914850299539675503954938371853405900096187489473992880432496373165753803673586710175783994818471798498246948060532081996066183434012476096639519778021441199752546704080608499344178256285092726523709898651539462193004607364507926212975917698293892367015170992091531567814439791248475706237804600009918293321306880570046591458387208088016887445835557926258465124763087148566313528934166117490617526671492672176128330845273936469244582892571388877839056300482483799839692029222215486145902373478222682521639957440801727144146179559226175083889020074169926238300282286249284182671243405751424188569994272331606998712986882771820617214453142574944015066139463169197629181506579745526236191224848063890033669074365989226349564114665503062965960199720636202603521917776740668777463549375318899587866282125469797102065747232721372918144666659421872003474508942830911535189271114287108376159222380276605327823351661555149369375778466670145717971901227117812780450240026384758788339396817962950690798817121690686929538248529830023476068454114178139110648560236549754227497231007615131870024053910510913817843721791422528587432098524957878034683703337818421444017138688124249984418618129271198533315382567321870421530631197748535214670955334626336610864667332292409879849256691109516143618601548909740241913509623043612196128165950518666022030715613684732364660868905014263913906515063908199378852318365059897299125404479443425166774299659811849233151555272883274028352688442408752811283289980625912673699546247341543333500147231430612750390307397135252069338173843322950701049061867539433130784798015655130384758155685236218010419650255596181934986315913233036096461905990236112681196023441843363334594927631946101716652913823717182394299216272538461776065694542297877071383198817036964588689811863210976900355735884624464835706291453052757101278872027965364479724025405448132748391794128826423835171949197209797145936887537198729130831738033911016128547415377377715951728084111627597186384924222802373441925469991983672192131287035585307966942713416391033882754318613643490100943197409047331014476299861725424423355612237435715825933382804986243892498222780715951762757847109475119033482241412025182688713728193104253478196128440176479531505057110722974314569915223451643121848657575786528197564843508958384722923534559464521215831657751471298708225909292655638836651120681943836904116252668710044560243704200663709001941185557160472044643696932850060046928140507119069261393993902735534545567470314903886022024639948260501762431969305640666366626090207048887438898907498152865444381862917382901051820869936382661868303915273264581286782806601337500096593364625146091723180312930347877421234679118454791311109897794648216922505629399956793483801699157439700537542134485874586856047286751065423341893839099110586465595113646061055156838541217459801807133163612573079611168343863767667307354583494789788316330129240800836356825939157113130978030516441716682518346573675934198084958947940983292500086389778563494693212473426103062713745077286156922596628573857905533240641849018451328284632709269753830867308409142247659474439973348130810986399417379789657010687026734161967196591599588537834822988270125605842365589539690306474965584147981310997157542043256395776070485100881578291408250777738559790129129407309462785944505859412273194812753225152324801503466519048228961406646890305102510916237770448486230229488966711380555607956620732449373374027836767300203011615227008921843515652121379215748206859356920790214502277133099987729459596952817044582181956080965811702798062669891205061560742325686842271306295009864421853470810407128917646906550836129916694778023822502789667843489199409657361704586786242554006942516693979292624714524945408858422726153755260071904336329196375777502176005195800693847635789586878489536872122898557806826518192703632099480155874455575175312736471421295536494084385586615208012115079075068553344489258693283859653013272046970694571546959353658571788894862333292465202735853188533370948455403336565356988172582528918056635488363743793348411845580168331827676834646291995605513470039147876808640322629616641560667508153710646723108461964247537490553744805318226002710216400980584497526023035640038083472053149941172965736785066421400842696497103241919182121213206939769143923368374709228267738708132236680086924703491586840991153098315412063566123187504305467536983230827966457417620806593177265685841681837966106144963432544111706941700222657817358351259821080769101961052229263879745049019254311900620561906577452416191913187533984049343976823310298465893318373015809592522829206820862230332585280119266496314441316442773003237792274712330696417149945532261035475145631290668854345426869788447742981777493710117614651624183616680254815296335308490849943006763654806102940094693750609845588558043970485914449584445079978497045583550685408745163316464118083123079704389849190506587586425810738422420591191941674182490452700288263983057950057341711487031187142834184499153456702915280104485145176055306971441761368582384102787659324662689978418319620312262421177391477208004883578333569204533935953254564897028558589735505751235129536540502842081022785248776603574246366673148680279486052445782673626230852978265057114624846595914210278122788941448163994973881884622768244851622051817076722169863265701654316919742651230041757329904473537672536845792754365412826553581858046840069367718605020070547247548400805530424951854495267247261347318174742180078574693465447136036975884118029408039616746946288540679172138601225419503819704538417268006398820656328792839582708510919958839448297775647152026132871089526163417707151642899487953564854553553148754978134009964854498635824847690590033116961303766127923464323129706628411307427046202032013368350385425360313636763575212604707425311209233402837482949453104727418969287275572027615272268283376741393425652653283068469997597097750005560889932685025049212884068274139881631540456490350775871680074055685724021758685439053228133770707415830756269628316955687424060527726485853050611356384851965918968649596335568216975437621430778665934730450164822432964891270709898076676625671517269062058815549666382573829274182082278960684488222983394816670984039024283514306813767253460126007269262969468672750794346190439996618979611928750519442356402644303271737341591281496056168353988188569484045342311424613559925272330064881627466723523751234311893442118885085079358163848994487544756331689213869675574302737953785262542329024881047181939037220666894702204258836895840939998453560948869946833852579675161882159410981624918741813364726965123980677561947912557957446471427868624053750576104204267149366084980238274680575982591331006919941904651906531171908926077949119217946407355129633864523035673345588033313197080365457184791550432654899559705862888286866606618021882248602144999973122164138170653480175510438406624412822803616648904257377640956326482825258407669045608439490325290526337532316509087681336614242398309530806549661879381949120033919489494065132398816642080088395554942237096734840072642705701165089075196155370186264797456381187856175457113400473810762763014953309735174180655479112660938034311378532532883533352024934365979129341284854970946826329075830193072665337782559314331110963848053940859283988907796210479847919686876539987477095912788727475874439806779824968278272200926449944559380414608770641941810440758269805688038949654616587983904660587645341810289907194293021774519976104495043196841503455514044820928933378657363052830619990077748726922998608279053171691876578860908941817057993404890218441559791092676862796597583952483926734883634745651687016166240642424241228961118010615682342539392180052483454723779219911228595914191877491793823340010078128326506710281781396029120914720100947878752551263372884222353869490067927664511634758101193875319657242121476038284774774571704578610417385747911301908583877890152334343013005282797038580359815182929600305682612091950943737325454171056383887047528950563961029843641360935641632589408137981511693338619797339821670761004607980096016024823096943043806956620123213650140549586250615282588033022908385812478469315720323233601899469437647726721879376826431828382603564520699468630216048874528424363593558622333506235945002890558581611275341783750455936126130852640828051213873177490200249552738734585956405160830583053770732533971552620444705429573538361113677523169972740292941674204423248113875075631319078272188864053374694213842169928862940479635305150560788126366206497231257579019598873041195626227343728900516561111094111745277965482790471250581999077498063821559376885546498822938985408291325129076478386322494781016753491693489288104203015610283386143827378160946341335383578340765314321417150655877547820252454780657301342277470616744241968952613164274104695474621483756288299771804186785084546965619150908695874251184435837306590951460980451247409411373899927822492983367796011015387096129749705566301637307202750734759922943792393824427421186158236161317886392553095117188421298508307238259729144142251579403883011359083331651858234967221259621812507058113759495525022747274674369887131926670769299199084467161228738858457584622726573330753735572823951616964175198675012681745429323738294143824814377139861906716657572945807804820559511881687188075212971832636442155336787751274766940790117057509819575084563565217389544179875074523854455200133572033332379895074393905312918212255259833790909463630202185353848854825062897715616963860712382771725621313460549401770413581731931763370136332252819127547191443450920711848838366818174263342949611870091503049165339464763717766439120798347494627397822171502090670190302469762151278521956142070806461631373236517853976292092025500288962012970141379640038055734949269073535145961208674796547733692958773628635660143767964038430796864138563447801328261284589184898528048048844180821639423974014362903481665458114454366460032490618763039502356402044530748210241366895196644221339200757479128683805175150634662569391937740283512075666260829890491877287833852178522792045771846965855278790447562192663992008409302075673925363735628390829817577902153202106409617373283598494066652141198183810884515459772895164572131897797907491941013148368544639616904607030107596818933741217575988165127000761262789169510406315857637534787420070222051070891257612361658026806815858499852631465878086616800733264676830206391697203064894405628195406190685242003053463156621891327309069687353181641094514288036605995220248248886711554429104721929134248346438705368508648749099178812670565665387191049721820042371492740164460943459845392536706132210616533085662021188968234005752675486101476993688738209584552211571923479686888160853631615862880150395949418529489227074410828207169303387818084936204018255222271010985653444817207470756019245915599431072949578197878590578940052540122867517142511184356437184053563024181225473266093302710397968091064939272722683035410467632591355279683837705019855234621222858410557119921731717969804339317707750755627056047831779844447637560254637033369247114220815519973691371975163241302748712199863404548248524570118553342675264715978310731245663429805221455494156252724028915333354349341217862037007260315279870771872491234494477147909520734761385425485311552773301030342476835865496093722324007154518129732692081058424090557725645803681462234493189708138897143299831347617799679712453782310703739151473878692119187566700319321281896803322696594459286210607438827416919465162267632540665070881071030394178860564893769816734159025925194611823642945652669372203155504700213598846292758012527715422016629954863130324912311029627923723899766416803497141226527931907636326136814145516376656559839788489381733082668779901962886932296597379951931621187215455287394170243669885593888793316744533363119541518404088283815193421234122820030950313341050704760159987985472529190665222479319715440331794836837373220821885773341623856441380700541913530245943913502554531886454796252260251762928374330465102361057583514550739443339610216229675461415781127197001738611494279501411253280621254775810512972088465263158094806633687670147310733540717710876615935856814098212967730759197382973441445256688770855324570888958320993823432102718224114763732791357568615421252849657903335093152776925505845644010552192644505312073756287744998163646332835816140330175813967359427327690448920361880386754955751806890058532927201493923500525845146706982628548257883267398735220457228239290207144822219885587102896991935873074277815159757620764023951243860202032596596250212578349957710085626386118233813318509014686577064010676278617583772772895892746039403930337271873850536912957126715066896688493880885142943609962012966759079225082275313812849851526902931700263136328942095797577959327635531162066753488651317323872438748063513314512644889967589828812925480076425186586490241111127301357197181381602583178506932244007998656635371544088454866393181708395735780799059730839094881804060935959190907473960904410150516321749681412100765719177483767355751000733616922386537429079457803200042337452807566153042929014495780629634138383551783599764708851349004856973697965238695845994595592090709058956891451141412684505462117945026611750166928260250950770778211950432617383223562437601776799362796099368975191394965033358507155418436456852616674243688920371037495328425927131610537834980740739158633817967658425258036737206469351248652238481341663808061505704829059890696451936440018597120425723007316410009916987524260377362177763430621616744884930810929901009517974541564251204822086714586849255132444266777127863728211331536224301091824391243380214046242223349153559516890816288487989988273630445372432174280215755777967021666317047969728172483392841015642274507271779269399929740308072770395013581545142494049026536105825409373114653104943382484379718606937214444600826798002471229489405761853892203425608302697052876621377373594394224114707074072902725461307358541745691419446487624357682397065703184168467540733466346293673983620004041400714054277632480132742202685393698869787607009590048684650626771363070979821006557285101306601010780633743344773073478653881742681230743766066643312775356466578603715192922768440458273283243808212841218776132042460464900801054731426749260826922155637405486241717031027919996942645620955619816454547662045022411449404749349832206807191352767986747813458203859570413466177937228534940031631599544093684089572533438702986717829770373332806801764639502090023941931499115009105276821119510999063166150311585582835582607179410052528583611369961303442790173811787412061288182062023263849861515656451230047792967563618345768105043341769543067538041113928553792529241347339481050532025708728186307291158911335942014761872664291564036371927602306283840650425441742335464549987055318726887926424102147363698625463747159744354943443899730051742525110877357886390946812096673428152585919924857640488055071329814299359911463239919113959926752576359007446572810191805841807342227734721397723218231771716916400108826112549093361186780575722391018186168549108500885272274374212086524852372456248697662245384819298671129452945515497030585919307198497105414181636968976131126744027009648667545934567059936995464500558921628047976365686133316563907395703272034389175415267500915011198856872708848195531676931681272892143031376818016445477367518353497857924276463354162433601125960252109501612264110346083465648235597934274056868849224458745493776752120324703803035491157544831295275891939893680876327685438769557694881422844311998595700727521393176837831770339130423060958999137314684569010422095161967070506420256733873446115655276175992727151877660010238944760539789516945708802728736225121076224091810066700883474737605156285533943565843756271241244457651663064085939507947550920463932245202535463634444791755661725962187199279186575490857852950012840229035061514937310107009446151011613712423761426722541732055959202782129325725947146417224977321316381845326555279604270541871496236585252458648933254145062642337885651464670604298564781968461593663288954299780722542264790400616019751975007460545150060291806638271497016110987951336633771378434416194053121445291855180136575558667615019373029691932076120009255065081583275508499340768797252369987023567931026804136745718956641431852679054717169962990363015545645090044802789055701968328313630718997699153166679208958768572290600915472919636381673596673959975710326015571920237348580521128117458610065152598883843114511894880552129145775699146577530041384717124577965048175856395072895337539755822087777506072339445587895905719156736")
			result := bigint1.Multiply(*bigint1)

			Expect(result.Digits()).To(BeNumerically("==", 39457))
			Expect(result.Modulo(biginteger.OfUint64(1000000)).String()).To(Equal("173696"))
			Expect(result.String()).To(Equal("4014132182036063039166060606038876734377151027041418995582553806466901368729258341160522025816124993789993695822370688304664951507188547123423222707028462645744504495099771236039772731479526109211090004311794404240339124784547422345216014091269896242057728789320308388075464535693734707938810361158822889129810698160881991361391615727898487922111925954357866231992015566112008236187570239498407927683927919775842288663953197267327615522763990546974302225808428095414868354189887477547130070092514353241631780155439824875566969089406890190453301219401944755635782011004657947505574085800780957240830989405542027471697598070708223889360327246648360250590423193485405740024818660966396088809547901243541112991057294564493176030196570968415627054119386294217737236705829632322635830102002662416478438021760275473366229548390019965427781460843016919698486005106981924856484135189126309511355753691653992328843645226535251273257107207655144521985279194687670618137133365061767680692604326802596487142424879424052336554494321865741845750005414127789039482660494669119722147859831209117695710217533300873521325189619906400613838224465802145638763375814734209939839362619843254361854200935168819098750162527330741262315680533040564205957680744473526587521186729552763984974171967043329744088725132917968485124011517373906219196506452837501009780475231936674958303684325425540235711585231786086215390697180427689333406191129678187071298761090238199532583274628668255539551286659571639463666278761674873780891713295358304373630381316642367910442138797034280478845037624720681884127045266971966132114983342843060225291800998287388520216891035412250834037533732042360063710863391716379631799881707736609524346067431213693134871211594265410252790425524534272056805082763168832785380055839924774513993312872314696842207051788925531326197927999177991421171128255733182296142518727575876101593203121573373469932821398066544389206962941141912977252572232867135437096743144409526348896093972589490165188215640760729511639332407683269230949843736635138894075691687022126305813977840524433219874210358515920200231755056725620223215331388220131477362331119420435167846194881185458654851583930079821808678398456988892271572381964480265039471649297678899125791655447634896125360321191630695480343247812206582942215664613839158662648407403549691008421277331933733296802727847751289740183232533826912134138874411587537072148434489199610825499112354149659416715243020346182605006236549464809827596184003972133641052969706618571561687011704990178047065972378213930473698742510372763853852826260631953919612454996794566066156884728715668680906027006860236291939118231676592907769781022362456768586214151241004159221927524018370878848684658164187227120264996078166702143742949273872723291189478841760522549189798371468818876307021543260191247813116898406629847737387213350101952387350770822381283290336805039909429093246363236388892554180007123150803933422140214159097502754441156418937483080165530589247735955574364496286938362823593135322482901755392013176188071632369230705598053524807232330332244850755770773086552102546479856865241592742983108951587550272842151337427500985790733257446974778514717423322164753741064203343873911505959859894843401346764198084574888570021582049897987375412534723203744213413544913921005682952291246029696644695148625577845540165449161494518865963573605166219797587702952678929547091472170854505976772092211830161122361692991516152705668722970188580451007767844433986239967364518420942072166146801482167949547691800161112961158906423261625463823863487993193561114681790843138704920497219883815661976658831769131758793119824969001836745498182398838340648927301879152764226110118671378678796232654817311851742957558547201689875931549991515035340972187105911864648544236836260504847533198372681672129543186083369951735634161845926675644115399774010625408990786348280910523027081645203785639678961312201192456238696520344039542130839856702031559650739069883838674250667916543006817208656138057336626345130349434110433864479062478550035029469824739027938916240527879813120208815523515779675977148726193656755629019016267380677895569115426761646230900288879599082370222089623761745116821667630909113297053145987339042926308730773608693851954656779816437471443388320939306664827788631871296953668897075687483930099227986839760048957379131637875344314107802027446160539266234293408604986954900983474482565227972953603366716227879694033469601687058461710870334469254188547991362387356445617844755230507550536963112039082535792838287188699545369099960024542135080095877346531176357217501893061794845596185695372386978011683132185583747282044356978640614686144305148838937966059778731252315662517317356452764543529368758499128496173025949468596830641128630126982569680227673151994682929383128359400800730647033009742194969397867732106855292305985907659610050250864697979637508638430748544527754238950354389404571972113035427133452477480922100499433629582296479031248198113154187002474014136277395295167051445941127918696437363960957340926932413952086407326427815908876047942365777403903909229226252129092284179163733102953047969448532689603563619318243124090419945712029201139189832010659468563854681499997644274767030324988966550143346180490169487404817781232236796354488293240003979784459018374622874519797289160250924013077191644744667079831907446016014675047387387849879634538532924759700314058964008310063071179192227444805709827095742202129386405849272260737870769456295985734252291272707948263365184386038334687999156723323731035874348731014319014672359686757028278160882218782815645825735630974163177508691319192847566774051552945532503104644040097190905665204611410283248812046204637327671846235824873113373034123535851718551400885999338937594799593962005964839702111477601581045672064416926126176740128698981250198894737842534454809505075861201705184249049902802765451776105523827736706727962349617080352862427145533868638674647621635903522376531004186274084864319795690579816037144710982537578574501401964606734516938171383983632188696122779521156237080602734546665241462878930312520760735314268127853512114206413184372047800360825636519029106214813659202770513059287824118965441521216286434750082511096818238637831486072817169107328121066318430424141551101448255356027051112429885809472561769503290478901969977310426678486943015192324838295069433027716466240050601698187879827464980534758753486354416047446209361934916329537938495521099464116817873898064615746425334307993580485350546795688460002265981773421773567346574241551369109003764817123662223617694812212158169860165209587231204816186164040908062427057433374354333512098587837751590927585021517620279297240368736722183517705003024583189908598536152163526107431442854174391559901463745686622770317539213510723170810141972866483922634430545450911424091655970619224266853425281308205026202294271525889978948848786451889351657070521050683440519622278045532953343858160953573822408291017268710562303350845615817951524349026567175895021778929616302540029650084377784855751579600269157963675920313876164533388292536102158206357715729658280040433132612140272022103947915774047681394936859812993227042855516834600928710152852398138594784429182314097140687415536884945999890119934480542332095414567022677110183124088577034207184035223269746212524708320333384846963664614426420999338980357883890451385386203823384441967107049845111443729321714754258385305719682104264623984894166226966604533114583914334886768314527778332346443260144764225200780823816350140331007279882497628037111262945933483440250886848466844285259801201047044164028789036672797521410884417114481113064890410508926775148007503192306944872367174516769239565732199135780783025193750190215687733235002528390196251747313108762067901840785138565711849300657314346796967787701918147432435368880447718801737606561365196160911276454420767194347998793760385338214424990342842862610858256411298097379924586198393594857153193649535420869601733255913603748735108955767125860412798934582791136590805835437006214619180833467538076124961206622745210375182783108944606061279552864482480673711218660245187076021181567794358666390037640555141608240333854073338043534429832182940110490657817714431406115312869550131360990262744133923687718517425890292711139888287121121169572479514569137506538532660065609376222907442241190395156802180983994401113294687869686532511825311652162444723355095362226595495851115378621729550619060175934795677858252326354523788281947715246424141202934670245380209752183345079219497823320697779797568188823897392439665907151941584911252454491393065641220012500307075134569170546657950628513232706438968999472833971666066853821669755513289360283904818899339170381382830023193685137019576555367344397397500345907203142598676169121614999422779372579464626191089018971078076898761801244103195495210458801583769766231991406931405001468005925383399201077190857244032943731898662908108327191861690909470351138983031879129637914513608632467782348730881613284583111745366299071443307263549292599904843844855174554634025101202206924639936318383190055108648734604528810643754584159097664780111072742553271812943448210019770733791851140909523998843075407964209315192369659421520580000397483091476492606940468048281634901478780014425466903183012195313405040260181937874135100391211424280568522832217343218449566212577839857429517773391325234970123982710656842132962072713670531585628640580476302740569446074057129250670136192085185275185779311649428558491838627836713048047320570250150645544212994809623734780024674994219129671408354527116135589032292260877830511009195805857592537799609204241901732311725470142447799886353484515819626358794478314752255102283719737644618375299683266816023955504186817402439459985266551465059469043798811380521824185863846994738759453069525025684128069242698184909729161080931680706454988296074915640851308149520496483541940056350562485344038487968200442973606174886536585643269441491537565521167279510356916025121335081257328911536240083436352408573331163048584890056692764867498736724967285937871063837364337621755449191928472249494309408397935142908533639574359432700608342105925301578912251377855985573358344672634864198752515613452910223366635904982715715461770104747504089416677887790075318553635109801839887000776413486686719326133277853012682629220681412084427939358117833274566986711663155710308663779187708049545831766986000316809626987410586036153344353518702949856385414013504665443220335274655069539917576038076933068673399231542660016785346281934674190287133321393865039012585301516288471252445626727183613790287375264863440341382950289536738341081070823467072641662638160920689079423415565415544093523000909042459655499558347303116788318768790020786196449106937448960496335000065625994308769918218117532397756112154058514439232714731547625428029301926759401192751887850518150985316736590769370275349296845281095433853324460820026096959856238442741949740512644749005458158427886614712986993755017155932619085888623668510044429325611225280467072107378374969832028261176123924402988656792912890438942424881406263936739195642564927755818346129962830564952862008315127847743704432224501586614191621134121414084766616850004752855737790449381521488307878119389635677152517700443598625113019383248895522656924874676078083365846741835696039642087136059892798010612944067324441233534104611388434095357600417585086903206201518724858670506731668441063355877998945850297008155748109516501339708308398677435237718856664102366195317472530863499414804202191902618201135135419339881515962422512857025124530919065190851292486672103057931987171154027738347902521013817932635496408110387972031749255469418772947038989633740538053684839862500994756773716881213362483372906037927970516028503893249938370998090355966343379594543393857619439969073876416575475031357342100897959643626247141091139113346430912218421332426397617431021590033646367648111866473215333561424765576016929343958665819514782393467237369306387879965393961506126345841089840597707284135690717573248058222433716097790974306041013186967353434100413872393337984581403537366546678381981115735252874653295049715936996306827362887519103246761552934739345207395509749551977003395029962758237343028420333856980432072114541711649988833361686311038725610601607452260800276604537588086629832393662151908732471530331390503598243861105289349754646633675165958622507210044676919618551144191211291948005031985175482748170959930115956149816059343618725964259144545635518370891105083518566650346326164332880670949594501221008865226782171403017037306385834766757756908497060699036637472403967682742063471456284052859300793318493890600834155472228151930479075648450763095792020687833295869010571410169415467895893923766234397059568374301410196033568385224669534494422044604529886447812303243369815606920089024378304637233100419682131167147683138237103836396146863068481497571075313218664281831894773101811973809037521256466043980717523127265465567042185267821091598837152282635562943239683700658828560977961472165705594943816999433425426766733718476472367893845215684320740612582945401570535852174911754718954037420050461252619221549665725031769139664740112317369998809909948202566349256451101570496666251245369829370047942181073261802638686307812145729359707078844999721963284844362680955941477223282704692602722551714620971927977590786466216404705670985017117826604611455959782225613566586016236155091134079620941423490044733734944396550570831455615770556197519856885379072617046894057442615170419480314764677588198245805514991110708278940489929815984675675250952479312601790676807548216870577371326887521518616329453883171165575520660362288398780799764281234277833869649816926108430437360963685319738454724877162529834153355494976546785024962619724067057664631770594248598642761728500729576216980475739090008569984843050074693176669966849426076793075347697354084406617389968364623754360445986017960849084658844161935174736658188484708144038671414964986538071697177283348104836430965020294109632015437996877591862615018792846939194762191455847664016969317280909644729562757288803470163283362836582801704682422551223524506336412190781232068031185091769869744931436653174905514147761661795615257091424626870921426153772807334619515077604027952734976964435745649982854244359572650204443485611522379011929767598037111863487441562070878839417195534319655201269956226091676001307080771515175957150512405356532822536302497166554993412071224143949961699966003179640440528339832860296716347191119277100624886651043402734594260451659728157939427399723866960225757348939016841178417724932792842710235644959285595810814178897596521733762549230792922096541997316437207126341990215143563343225996579081761345505417456543760556902238498414192686706350887441417030898964773106560825204937573907394584813121610631479220095886125576521265559115853565408661656648710665753742826354099309715810718586688386166739695360105153732730810856363115705196930450948164992978356889200164509197423055011582422464983146526410945082677782335972739021354538247235307628778236390018163762505814714880859010783891710645715508502642865818878885105510899076614850325368534907284181813661027044834773040190276150126054961895874190384311374655847147296725596621786811776219607455852197837077254961182259087277750214673688771038775480187275571252382373713598978744791329047432403983932004244483211652502611752940211118201058337234584034864574227810200730590127781709506138200593936779546014630761206017320716612138785694203059538064755958660196119388582888366460140254991004358012838108699804776204688584225575662975980043946364442061016378143365057147725242731478437916914945007628740732858651528582124707039037955659921200984589653967290356671822766500551504147096062559238712097470401466568328488127128944142333691791644053321208072117162615223122987279758153437976713979494564324839309455758207968141493094108989459309112061797912505115997599779734316634642787662861730837554573462339745191748434848718212564814399088890050558563286079628077673083829772294639480998765110246644620219192425600828996562323995526174157897858707146118489491816556359795756236851923837194287030905959361479970840400143320428366606885025067331691738040202055787012703056951782823888375506208988341157408228503942543079543770940443923976231093230318814331402474723753800935580269149469917368514574090096348578363321547574189770258467807477633852131070551000640054998215176863242889754758529656993218524750841275204785416446508661614647568979170309980255454798241426276915609241289998999441374750551762499989889776202922935559056865380485853520295522188744880920554250457015419395009003698380752089875202158738238319050876794599845412793967035367858531763921137171562917168105470402413558860177750713812621593283779321401400233370444782744252825896616155710291863815870331853585385758668507067576817342453134743338486664001732164359176960876199706001844975434858293165037162767244454671714475115808657642963552747840387966052546961337581079959622912645265343467678764190406571477253426299736077244260618237119651858084687036618695794075478789311587417059711343431227303350973799387878820501939565505465074594025436036966165326241011378797146265125327338463869344259066761066708758833017239984197008948737740667243072361128617406867859023220628124060993563955305881007343422807152945373290736658535872277975738221441943388251274447917399816259299528136009365562772905944119535864755915645636481389721332489402998780438636631270258475271207701504449959513826917567118882270665697324654384539229093273515371145657367473351701477176058287971525132123630231302192818809015653271446364720664677553548552591939892622312036148679132676222630440973623016172572796318062245769564016065503400855132972119620506711200765168117939786675101936469765690013861859011813740916497238562803149602498800743366839594464102222803276025264020852530869628324622203235960629552097202049867177844618660210541628779338833715233540816837614040165594121213747804550323106645804853319343703095281130779452288398474053206030303124075718699502769043354788987366303158458622216238559388852586287397244823587628905287478185055937738386553460517964970178405297818574001552698590403422567780250925573684712440051738280186677970204727172629630250807588480950618513205146384314670376008562962889264152845320978121184751157549620353870769614740394880734501759269207166413065178816343866013426191134590442802370423292149028056236608844109615255164543338823924604724010260098700884832237411150539327434122637626938602730735567837232664475826678490431572548909373831937962691028377457710962109223347737888551476979770134673578428196780307228525530719932119728519681020056415065579357762138062341940715505658351032187943326481465363443863803288110372057213933086575266979834114281470948733317432538033803753683322267351759794373856980502790076050489035615460200962483659715876640167205197093480143268848588887523572869448200676376187943746855377974486709054344584706188400881524241541391111150657380625806644492114953522924820276748689110132199293929664946288230813053178306354648926761437335954106994213287933139740635598749442553854302259023579655742203516047488166774689447569867936622820611204174605992950298769916077128982710510288196427378103368565164358036180004462488782696183306161348948232055760805161621292732118579698583726535923141524929496406363413864269242809640344703501492582452690207020817139987977758629694318006602700285855738200063021223919333172965177631394627287558397039465292625282839307496623059783681956962283709629525150084752829826413171823267961454311822679808362195466731957355942359238281197674840013554610317457007073447347555573247289794213606244753195388483334301123060067869123641067367400662684880401594438208468350024707708829070222659764366228295547652120614526659158515804151065386675507811098963031107165735090412267764271530335591618530735740446073375852583244380837916461995678501069928298475649775376625135157948147430845819708807270087123568447382546399604201642056708879576050534869014874330086255052538159839286402227545108717169973795449513568196146492612661017844433986693504400636793705628179607690520396302365359326058011744993452664586783672366992937395096218263188368979803013281224849637296442141604562461162491487916232984232571386163189599947711025232461545410107189286669897464249267488386815630421031187491074169125825170807382257645180821743080058968318717203308067308430306008719215999449549275186652855329514480288672236238137561214814466389304839692226820467175658389499720385721411105875447062795842543723282133190827930422906881951079517628186936042926475582499991844669144045926786553864972953423185548601733151216289576168751452442756075554834451826997459883888370003537652029989534356334645187913576065296792644524200824584132664698775130541572175057149540461321494098316209986410309373087546472581337568958627121819007771825027522508287259305852577876388306401427150201948576417401656477372167355512124351307696967324497195496192739279215969258543281938589966654331010867898501990694317165024492189477378166203114027287886864811921598100732112244472230470554383307370299418471731149022017754359690346403013127778918768231606030509998499642148046196935810238064412531126176156308402893998480786811137872744116447605557524431862992124455086194694346958998213418949062771097472347167350139668646981167262664987259407795730228841737389025744020061321542917633986016330050107026361576525861361343063325677323542710062288653622008100246365528887395634280352106633520268540358179909328076856361340557918726583703265608526215873247374438128953980193021781235856634443870531703734512403790773617204413172216046274108987822929386890747041850246872279320141016139210667033177852792102956417783343165185457063261014219019520235689446075505595529429106951281454091113243566908586495716555268930016123745659467401133529140474108004345232048212474448247768159489601013130975476694099832314918715403248611252016649922870384793692622229702610106986985893119003277409080679262329717165143031612328332615183711602787127836569009303906567838743875732384487702190656722551392684192185152900208659556329503223694369365936502102766611755792835591109286320654357979406231288936371342647544766199694127510725974843267121938848266208727537109341946427283767460259820749824899325101096764583474599804744627964518125872088655466147810233418924212562023675468901778045807940836323107298408868655304836421435151585464449138540150345465287526413345767386761993053273799594296798277380929249849357227606453580704529153148918715666037355647559732620792511746446385929050322349254324570411407740388768669446180113069843465434315221996105127890401599662006197214161055084961947355022344753536314097871510439160892202614670191975302380638766097843290952324948930552090545562940368121762216843274341603413427737554334409079094656810056060243733369584954911218729530084716394133205309124662246096755881559721492929181900212738251392916283098708192082807933814206925026586812173246109020886948080388323334532916853674293858738746575892137769457705381830230677106564567006150503545480458411511998028515408053429026809142780692764408220248254464937173642843127962447327355407015168381094802887654612384524560495973919708044119957466092904213848780764642494233732293224036130949318033995088929763380074431871523703593426236859383367085483569982629686960150819843615701884209216590737774519614344445579276002239274776296217803220678929087832467329479669456393936850538058021066265509788330120339443730684041783493433403368660115853967277552792249908020009855627048227361655544622682459131193333057521100975332158433023062932561369835916252724835530963063956451775313220617962849596175095470955736265891458562873902748194876719191960799278465072154943455540417890854465225477767693389605168449190478542196589501573826166025865606146645171339224009294572597815717484786765737153288478556654851482881403515032894695820372415060115302875192073646119765406141480092567413252348512968515582746901168207643774717589960607598404329437564395139994329635730111110633681700603685437336109634272210103033418054450224938083739879836153523722555040316183597103138650374921578649988715021842250779725214483676183162913274020607114407980836315188919652318618301784597885748980354018002765408849399964518595757380248305671577492554441816003221000192908285620784557292986217457457008937003087353143312855574041832420112317625519541622257673910321225464856548382355266230023024517109008815119619894335099424334371497881198051261886997463187492354570040820824145462052230431910145690364685647900970770134523462958114721686909957771284518586588615757296975240405197323461312166300837418411100951691362769269009598174972554545066600282223191709144658531113305195792079519199933519361197220534355669335573638622367318682415058083685744459089503504693694026138488909769903440659982352197168694607468610546527603693350254013923871222214907366013402318469821065219011342887014544279931938226403296808841138321268966015543190588816680865180553680311946308586320183367554216227804022237489563325384863487502698263215652914982484331951463332308738138198850029274360944596748910049423934828422357287796788094130695149200540222035121940850938877624298622532954796851896147160661437050220616565271922732552946446710677232516460729461776038659142224476257345847794450942373578079859499485126730131326576524516304890314331913177850990987024168111201230369088888814690398172882036590299111655540010027485735699714746027888996894491074636860280327147225823944248442261836656573904886459040779317311388300746153183682936598566496118810851117210155851497355418607451942417218342405303407466882340739425972867428295651119808720504249639132122862955038307890969493994195104113046524043749160794961436697257609684147927251684959704460610863296104831651532986040825318902300447440658625380101877503980395998227932398085723381177857528436566512859064891923925275343188259160781978747236009670847930228634801078413556144904608590374066578382162840734094704556882588766886646505226864738082307119224613333238133939596846371751338552884133868538305671501850677644701334294545522736725909403627900425428575515256280775226097449131464565340866109649126743632594122749398221333024430536383625578216179820036875258724363654014144693138078275083473221516269734620056229410319211192339379373705969871361716416377679106916388237216722287776659174301797860870108669335310887042573584964084549830853806105001556741627463785217549742961939001285500174633919141246947357557108556599045483102378087992187545103227928717679344616781130238768651483461192743956507585364231475371294215456739427454883014263789532061322035968793855933561987325692227761095404352330484826414411503719090531177372003299897016197626979927395457229994791267749526153303628872678038332800066583619933258004913183349079048604629414316357319191327195558231837077806252686775830347316627050885149851971577583062612869863587283067660531257911700069173571826354880947317637083284349872412544075913402708943189703819014885007669840239278005529613445594638832571533109906638102771308609778307918887972144572710048156727332459584973070036821333151613153819961560585678263712112030470268059876228517029425596227162809113710893297470852860095341562560503380819402738344574108370654720306234445795829592390942617361234054693439157976577350694947971977807804891223325574434005726664365864102644111316060149623095131182327136402417631294301589329195653465200427192696211199078526072648627433162854506513539229770773150234458156216226928518387866328495285429966654858464869353026956392941429258494349958271486167595912715257764828267182760461364253896726815111881010426744887647580456490504967345425456474016622973826347763934613338287211436182257186176588003701140172055173997504727844175707668889600207567306051631200903890130181471560130625198949120552559448558746812629141017684709710427004048742389198612176603676356004472896668928059680765558510094787743457924782484751630770082376667308234073330513906096857478178233683870995690359783554071103130624719949126370238910969540603249868402536861477476283095376112069159050062207820661304599450914308949364720434475443492964833264370780845446960946001797745772495904818876793482774452693415669479497667787475534772460122493492137810993808000151111691566207109773752171706281820506088282730034274961444848649370619977405543694383494372828173907950266259467252174774397940298169228566724046348868983056956512001432266855278032438198722850314961928002201732958195918019411712018092436685408928571063807663069906690683161055648082570690897534741031264692365697991501782530491220185233054814248544978607008450497309502054372363671000481526486566864005283493893364400968868076536108775865121526619434205088620757585531968138561890894217464218137576596629621478833379373398964800435850437577640968460377641335406648327717013081773794010794014952875540770031810880177389097917673211290120018726648863589367709501261388823514002523445274937813308340503316275556971834095957070311151685663096650735368840013991029460386047281639940987620333442175681255921413498908635648547978900644835092929548172477565064649311738623373240808627814703749594470924250253523961355950685704430632687972782682608805157573492086682988610293226161060994095786201637476231248277843213382720254902308399876150114954386985250400461709342977255489015322283778501777263653261611258177983369074785018325317279141863782640024236415716586455649381512343173971977193215178101228014653700412479738519336676622766507435407640526480204490442902692196646910526472471764603114393090554086553127756176783454160035662301219800609443889189639503729673575343707967975070419543470415317588945540317651342669175155501366866991136532252978456931698975357335541024815635056389334249470576140186655637589083443486141479466860843188722766352909711195732712796762394563304967709413436958934023108929991080515716005090622028689543836100237553062641924208312610445690723253952243872122829998636223014392769452303784254871986926656774292506382959408998515669855260604591417042824181753974969549598117789414318297442937172941430209155229749282461303079594444929813302739394766566525697098691402533281916843264333759399142217857657800583323560684781493900290697958129276960694733882334672163494056850950141579853354856735690088854960442960294027198017859950777639407901068888088712399238217925269809500311232705165587641320754998776731203952540771780852157239190565359903762414761448208478958513982081094679040994558395347872876020222383814294343590776799037308083155653657107998627712323242404749063269476117901423559246446246049043875239453357276008842608058146113940343526263311302027595586177944864081379933037042171087385719199524237547368354426768791245302511373540084062608180152747935881087224444118033523826554478381262051914523602485027600874489145801713632787713497584751472006327570249613430553152241216076230132010422320276409137336534770453910403780621671522392365990692185624618798127690726849905698065462190562548843867811239017679715179217870401334571969427022770776317456491630491922052986374561038895211523979527427416760633732814942980365064600870622485988343876070384800055371145004676234932204325781711685395163200160695744715969286104365799253691170283772945984214730340355595947699698099014443315017119481178618492015118130740993696491711908414895922891067254054340531857393960808153425278566229517436788370382552922294592734092702611580967299358770050314659048255709559925129488294676708030505260000072877607457539673939927320349188082988967058165417599841852725089899200883920793567706237354669531323071242303361987718500437774770836733074952878390038891081494672745424420160412738654295861782690124679189579057467108678949518611668443667084341533197526296153404267552816832061728386036482614796835743851679123120074530857086709990170706464291074762677725912366546486070021273493408424692371122397873634380087044920399569414776671962284365618374494404783009454907098151420965665235024860054584674458829994719840088578484926907349049795048661188990225467479501168878164556044747702179375760301749116431544796543423102855910576078960509904137472289682236128170228966241929030873619387884286498641050901436813025342891314442008308026895818520425411382325363996463203822969014960406971167617803490743246959796489176709522670239269012944596424741201659344241742588809441910058483052316791067695730906881611834731035613909961818300238959553674977192273324582263936445381352462640344317152685672375636106063538302866212782300594878262565506262507016585942350299700482594486537247614137407731144094911598703420964140653507907973000042395425123178245693161546538944817163474529961240214053746847925700387598211104540059955453332404207376202292184231507885591109535596794571000113817633180093490914619710272097789262841684102208298399555207745134500720393760778388467657647832652533774955213115034905031331170372907717242425891871787083064105096470470807217453134164901845190381303646624497955153487250601970472944597069341107901218626561752688823236509798434123508741666811161583229221522141134524591417586902457552653169145320821793704344419502972752134467994934869096407841923356492882008114296882064839353539886265242831143174995160762690880593232309324669792464735759361400066588228775630758091397420766422347053339669946408397898479748702174241725385759515809588332981191183232863126414025028855857625276026871740845329254485736651573547968942396001683261003382924583419836405578403022925041315498400193585323516044788132412164872883956764107603580761015608253210334039670084993337626961638658264147568821163860874793294101103518067608262144888951715539977598455880517499923088750685836655515716367990472782289753588289643252576984143053751695117510520372288167030240744170016627428431866080813209936498367200961001217572321533582911525242084857121897399133912838543725576002777205065137875354041868752613190374824508337624872736746075527919904054365009325447965772617868469787526900954662228938151701419250336137578079213918409114453283963284518704329695562584886018310263047577013189386033198634495681666587191729216364260315068566525433195460645621272540637136201750601801280350987654877503290379842996692304871979737361237005004450111017209863421434448142080922058731517285624486661835457169958250859601191760247804917587030250960118942516181532668642908379165546572246588236181452870059164059291834098650100893185050300749827811324812752780635973751170231793524648742605639519261065842348084780078197623044666908304947840322375238921647068490029703328925224630287711874385761906275853624862447995733481952579457361161723698658868601102154939831901682740902639859568215016543057699249042977532851613233475279198188446078303777764933793897598006364490095043111364019162853307148803445148714481253986183475235409501451199658719852163307602237212857517328324954494638744581412459025744214264604305408829548141050865942584202478886159597134834055563783789931764838601652428918764415643317948502222059711181053351828207618627913195923795822057302474405998878770616459666538646685572019993889307039142440734270135086457562108234055555913602152888317026235605805447403780314423093360140310738052221104158155449012986496338389532742495169987871687962131277199254025275976117059067166378905169772703115868726807285909921002990680153114176973572430721426455006622038399407192717389485597717430090700963070646135905401723517178560389162691795785833726826368236732660871254754437950905868026860040340401366364436635582586576320987349504904768756226289709224527496531073627879634192454883225759117756134433519663815666388742633604712659979989937746235144098665378461260105924345684318065367788517213904114086891434976034802130574195431083921990953774040907906287164994177909873401433541401406281055944976806366738199196301246628425039477700999584927413239653022006401820331668136835304320862451446420568841630605717804812802517138613388353327527224213450160244942641523828884896780006961744412491117507024662800007956065079359572506235985413451355939632323084876933410583780998867789105653437541591182243656510664393423174423513282895091974044604025713441956635265666081664340113582242832446270925321986436636522995921352610302765934147247681239731855805017623392487563837956337980740647643258859259625230968166051374546932818593962922126306950302112224088870480350891943437215741441855212365932127415525357096827287867943215371873982557655890505256276438799243174077765235070642189254085980721618799647345053650888725401936469806052224188055382756611213257843744940816681770265825784014212608576257356134558715010464453147873634657049530995900860251300368631097592822266898046690705758214212040260587867843995752972455054166387060338521223203534544057498013093445693457199722427763876742236301319402507069545999249384877104695916545027807772467657018619860020828516804627461416534256181881280343208607886177969725316473276995379689412004699194032637410762029497719056835213138650072080295239993255516147171615294023290810597837885201934456664948583259896437704235526839153177878227781877522693174171993336453452050588505768631373387093831751846908443618641574847737753198886730852451782615672785238559503699166184109989475413644759088959291864469293601685700989840873477790483712933516853540988879876691953083772843738010931900826214618175555453671792496843479184959158753655785508410179144331715650847643903821042480226896684782295815749440163147413334944594136460082846748144188409391473422207520617588508610707264101538667653966477921939715922123508830124397049331940593216432845859736931979042092526037821186592843691795635380307576664123718289563347610171048717452247269745322673025738183653056369137030817543511840883567616269178877715702684682190762738439841101367629049613105677104357025677819148511813743087943572307059694143041857016616090920629207137120293080951224799013902087430936964531424771387224605271712394656792951401368133375304662485933127812299850490866890006627921273025453571831182453606726322191198851514821830006179920236866719614945157540193073071059029243928716806066986661745930649013493720457426416759547001036479949983232186950068072545318925600134489641539080502465063196759686205863933348252226827948448786945594506227633824991438061606112010641705478589713365261833710839481822514036991180735281005731150240138695153195639379922197685245849723704580849473986791339573714845783220497423105067071787512857119523806835886839723435656244732180788098126274270975050350319628682405469001773598347856825372536550018928016247538635999112637584873671673810004186533820777439275174703078694582022014840957205857731786850095893923052384394230612850272582984531441037446961202447936669331101648380773106691320617177061861323519654225570263051930591865874337327543075655360160373622977305945734477519210788855269935472318112232779399101592519558252467149084114518621740613293716930129261051303165246811458541017379343302392231344541558403725830432995219203720573437678829112607567232912841276423846446917501077375612727454000827786281412121825259006542529277939897815231636500332847040481095233386397000667544071303013369884726442536693319934157425008544450572343068081204804361117582244301891415246895239142374977327742276519573650394798160420938696928927504422600899900578460921357302715960126670268613664406309255898668806195870578996366536709926807407399180325362544275565838974676261850665812318570934173696"))
		})

		It("Should multiply 4014... (39k digits) ...73696 by itself", func() {
			bigint1, _ := biginteger.Of("4014132182036063039166060606038876734377151027041418995582553806466901368729258341160522025816124993789993695822370688304664951507188547123423222707028462645744504495099771236039772731479526109211090004311794404240339124784547422345216014091269896242057728789320308388075464535693734707938810361158822889129810698160881991361391615727898487922111925954357866231992015566112008236187570239498407927683927919775842288663953197267327615522763990546974302225808428095414868354189887477547130070092514353241631780155439824875566969089406890190453301219401944755635782011004657947505574085800780957240830989405542027471697598070708223889360327246648360250590423193485405740024818660966396088809547901243541112991057294564493176030196570968415627054119386294217737236705829632322635830102002662416478438021760275473366229548390019965427781460843016919698486005106981924856484135189126309511355753691653992328843645226535251273257107207655144521985279194687670618137133365061767680692604326802596487142424879424052336554494321865741845750005414127789039482660494669119722147859831209117695710217533300873521325189619906400613838224465802145638763375814734209939839362619843254361854200935168819098750162527330741262315680533040564205957680744473526587521186729552763984974171967043329744088725132917968485124011517373906219196506452837501009780475231936674958303684325425540235711585231786086215390697180427689333406191129678187071298761090238199532583274628668255539551286659571639463666278761674873780891713295358304373630381316642367910442138797034280478845037624720681884127045266971966132114983342843060225291800998287388520216891035412250834037533732042360063710863391716379631799881707736609524346067431213693134871211594265410252790425524534272056805082763168832785380055839924774513993312872314696842207051788925531326197927999177991421171128255733182296142518727575876101593203121573373469932821398066544389206962941141912977252572232867135437096743144409526348896093972589490165188215640760729511639332407683269230949843736635138894075691687022126305813977840524433219874210358515920200231755056725620223215331388220131477362331119420435167846194881185458654851583930079821808678398456988892271572381964480265039471649297678899125791655447634896125360321191630695480343247812206582942215664613839158662648407403549691008421277331933733296802727847751289740183232533826912134138874411587537072148434489199610825499112354149659416715243020346182605006236549464809827596184003972133641052969706618571561687011704990178047065972378213930473698742510372763853852826260631953919612454996794566066156884728715668680906027006860236291939118231676592907769781022362456768586214151241004159221927524018370878848684658164187227120264996078166702143742949273872723291189478841760522549189798371468818876307021543260191247813116898406629847737387213350101952387350770822381283290336805039909429093246363236388892554180007123150803933422140214159097502754441156418937483080165530589247735955574364496286938362823593135322482901755392013176188071632369230705598053524807232330332244850755770773086552102546479856865241592742983108951587550272842151337427500985790733257446974778514717423322164753741064203343873911505959859894843401346764198084574888570021582049897987375412534723203744213413544913921005682952291246029696644695148625577845540165449161494518865963573605166219797587702952678929547091472170854505976772092211830161122361692991516152705668722970188580451007767844433986239967364518420942072166146801482167949547691800161112961158906423261625463823863487993193561114681790843138704920497219883815661976658831769131758793119824969001836745498182398838340648927301879152764226110118671378678796232654817311851742957558547201689875931549991515035340972187105911864648544236836260504847533198372681672129543186083369951735634161845926675644115399774010625408990786348280910523027081645203785639678961312201192456238696520344039542130839856702031559650739069883838674250667916543006817208656138057336626345130349434110433864479062478550035029469824739027938916240527879813120208815523515779675977148726193656755629019016267380677895569115426761646230900288879599082370222089623761745116821667630909113297053145987339042926308730773608693851954656779816437471443388320939306664827788631871296953668897075687483930099227986839760048957379131637875344314107802027446160539266234293408604986954900983474482565227972953603366716227879694033469601687058461710870334469254188547991362387356445617844755230507550536963112039082535792838287188699545369099960024542135080095877346531176357217501893061794845596185695372386978011683132185583747282044356978640614686144305148838937966059778731252315662517317356452764543529368758499128496173025949468596830641128630126982569680227673151994682929383128359400800730647033009742194969397867732106855292305985907659610050250864697979637508638430748544527754238950354389404571972113035427133452477480922100499433629582296479031248198113154187002474014136277395295167051445941127918696437363960957340926932413952086407326427815908876047942365777403903909229226252129092284179163733102953047969448532689603563619318243124090419945712029201139189832010659468563854681499997644274767030324988966550143346180490169487404817781232236796354488293240003979784459018374622874519797289160250924013077191644744667079831907446016014675047387387849879634538532924759700314058964008310063071179192227444805709827095742202129386405849272260737870769456295985734252291272707948263365184386038334687999156723323731035874348731014319014672359686757028278160882218782815645825735630974163177508691319192847566774051552945532503104644040097190905665204611410283248812046204637327671846235824873113373034123535851718551400885999338937594799593962005964839702111477601581045672064416926126176740128698981250198894737842534454809505075861201705184249049902802765451776105523827736706727962349617080352862427145533868638674647621635903522376531004186274084864319795690579816037144710982537578574501401964606734516938171383983632188696122779521156237080602734546665241462878930312520760735314268127853512114206413184372047800360825636519029106214813659202770513059287824118965441521216286434750082511096818238637831486072817169107328121066318430424141551101448255356027051112429885809472561769503290478901969977310426678486943015192324838295069433027716466240050601698187879827464980534758753486354416047446209361934916329537938495521099464116817873898064615746425334307993580485350546795688460002265981773421773567346574241551369109003764817123662223617694812212158169860165209587231204816186164040908062427057433374354333512098587837751590927585021517620279297240368736722183517705003024583189908598536152163526107431442854174391559901463745686622770317539213510723170810141972866483922634430545450911424091655970619224266853425281308205026202294271525889978948848786451889351657070521050683440519622278045532953343858160953573822408291017268710562303350845615817951524349026567175895021778929616302540029650084377784855751579600269157963675920313876164533388292536102158206357715729658280040433132612140272022103947915774047681394936859812993227042855516834600928710152852398138594784429182314097140687415536884945999890119934480542332095414567022677110183124088577034207184035223269746212524708320333384846963664614426420999338980357883890451385386203823384441967107049845111443729321714754258385305719682104264623984894166226966604533114583914334886768314527778332346443260144764225200780823816350140331007279882497628037111262945933483440250886848466844285259801201047044164028789036672797521410884417114481113064890410508926775148007503192306944872367174516769239565732199135780783025193750190215687733235002528390196251747313108762067901840785138565711849300657314346796967787701918147432435368880447718801737606561365196160911276454420767194347998793760385338214424990342842862610858256411298097379924586198393594857153193649535420869601733255913603748735108955767125860412798934582791136590805835437006214619180833467538076124961206622745210375182783108944606061279552864482480673711218660245187076021181567794358666390037640555141608240333854073338043534429832182940110490657817714431406115312869550131360990262744133923687718517425890292711139888287121121169572479514569137506538532660065609376222907442241190395156802180983994401113294687869686532511825311652162444723355095362226595495851115378621729550619060175934795677858252326354523788281947715246424141202934670245380209752183345079219497823320697779797568188823897392439665907151941584911252454491393065641220012500307075134569170546657950628513232706438968999472833971666066853821669755513289360283904818899339170381382830023193685137019576555367344397397500345907203142598676169121614999422779372579464626191089018971078076898761801244103195495210458801583769766231991406931405001468005925383399201077190857244032943731898662908108327191861690909470351138983031879129637914513608632467782348730881613284583111745366299071443307263549292599904843844855174554634025101202206924639936318383190055108648734604528810643754584159097664780111072742553271812943448210019770733791851140909523998843075407964209315192369659421520580000397483091476492606940468048281634901478780014425466903183012195313405040260181937874135100391211424280568522832217343218449566212577839857429517773391325234970123982710656842132962072713670531585628640580476302740569446074057129250670136192085185275185779311649428558491838627836713048047320570250150645544212994809623734780024674994219129671408354527116135589032292260877830511009195805857592537799609204241901732311725470142447799886353484515819626358794478314752255102283719737644618375299683266816023955504186817402439459985266551465059469043798811380521824185863846994738759453069525025684128069242698184909729161080931680706454988296074915640851308149520496483541940056350562485344038487968200442973606174886536585643269441491537565521167279510356916025121335081257328911536240083436352408573331163048584890056692764867498736724967285937871063837364337621755449191928472249494309408397935142908533639574359432700608342105925301578912251377855985573358344672634864198752515613452910223366635904982715715461770104747504089416677887790075318553635109801839887000776413486686719326133277853012682629220681412084427939358117833274566986711663155710308663779187708049545831766986000316809626987410586036153344353518702949856385414013504665443220335274655069539917576038076933068673399231542660016785346281934674190287133321393865039012585301516288471252445626727183613790287375264863440341382950289536738341081070823467072641662638160920689079423415565415544093523000909042459655499558347303116788318768790020786196449106937448960496335000065625994308769918218117532397756112154058514439232714731547625428029301926759401192751887850518150985316736590769370275349296845281095433853324460820026096959856238442741949740512644749005458158427886614712986993755017155932619085888623668510044429325611225280467072107378374969832028261176123924402988656792912890438942424881406263936739195642564927755818346129962830564952862008315127847743704432224501586614191621134121414084766616850004752855737790449381521488307878119389635677152517700443598625113019383248895522656924874676078083365846741835696039642087136059892798010612944067324441233534104611388434095357600417585086903206201518724858670506731668441063355877998945850297008155748109516501339708308398677435237718856664102366195317472530863499414804202191902618201135135419339881515962422512857025124530919065190851292486672103057931987171154027738347902521013817932635496408110387972031749255469418772947038989633740538053684839862500994756773716881213362483372906037927970516028503893249938370998090355966343379594543393857619439969073876416575475031357342100897959643626247141091139113346430912218421332426397617431021590033646367648111866473215333561424765576016929343958665819514782393467237369306387879965393961506126345841089840597707284135690717573248058222433716097790974306041013186967353434100413872393337984581403537366546678381981115735252874653295049715936996306827362887519103246761552934739345207395509749551977003395029962758237343028420333856980432072114541711649988833361686311038725610601607452260800276604537588086629832393662151908732471530331390503598243861105289349754646633675165958622507210044676919618551144191211291948005031985175482748170959930115956149816059343618725964259144545635518370891105083518566650346326164332880670949594501221008865226782171403017037306385834766757756908497060699036637472403967682742063471456284052859300793318493890600834155472228151930479075648450763095792020687833295869010571410169415467895893923766234397059568374301410196033568385224669534494422044604529886447812303243369815606920089024378304637233100419682131167147683138237103836396146863068481497571075313218664281831894773101811973809037521256466043980717523127265465567042185267821091598837152282635562943239683700658828560977961472165705594943816999433425426766733718476472367893845215684320740612582945401570535852174911754718954037420050461252619221549665725031769139664740112317369998809909948202566349256451101570496666251245369829370047942181073261802638686307812145729359707078844999721963284844362680955941477223282704692602722551714620971927977590786466216404705670985017117826604611455959782225613566586016236155091134079620941423490044733734944396550570831455615770556197519856885379072617046894057442615170419480314764677588198245805514991110708278940489929815984675675250952479312601790676807548216870577371326887521518616329453883171165575520660362288398780799764281234277833869649816926108430437360963685319738454724877162529834153355494976546785024962619724067057664631770594248598642761728500729576216980475739090008569984843050074693176669966849426076793075347697354084406617389968364623754360445986017960849084658844161935174736658188484708144038671414964986538071697177283348104836430965020294109632015437996877591862615018792846939194762191455847664016969317280909644729562757288803470163283362836582801704682422551223524506336412190781232068031185091769869744931436653174905514147761661795615257091424626870921426153772807334619515077604027952734976964435745649982854244359572650204443485611522379011929767598037111863487441562070878839417195534319655201269956226091676001307080771515175957150512405356532822536302497166554993412071224143949961699966003179640440528339832860296716347191119277100624886651043402734594260451659728157939427399723866960225757348939016841178417724932792842710235644959285595810814178897596521733762549230792922096541997316437207126341990215143563343225996579081761345505417456543760556902238498414192686706350887441417030898964773106560825204937573907394584813121610631479220095886125576521265559115853565408661656648710665753742826354099309715810718586688386166739695360105153732730810856363115705196930450948164992978356889200164509197423055011582422464983146526410945082677782335972739021354538247235307628778236390018163762505814714880859010783891710645715508502642865818878885105510899076614850325368534907284181813661027044834773040190276150126054961895874190384311374655847147296725596621786811776219607455852197837077254961182259087277750214673688771038775480187275571252382373713598978744791329047432403983932004244483211652502611752940211118201058337234584034864574227810200730590127781709506138200593936779546014630761206017320716612138785694203059538064755958660196119388582888366460140254991004358012838108699804776204688584225575662975980043946364442061016378143365057147725242731478437916914945007628740732858651528582124707039037955659921200984589653967290356671822766500551504147096062559238712097470401466568328488127128944142333691791644053321208072117162615223122987279758153437976713979494564324839309455758207968141493094108989459309112061797912505115997599779734316634642787662861730837554573462339745191748434848718212564814399088890050558563286079628077673083829772294639480998765110246644620219192425600828996562323995526174157897858707146118489491816556359795756236851923837194287030905959361479970840400143320428366606885025067331691738040202055787012703056951782823888375506208988341157408228503942543079543770940443923976231093230318814331402474723753800935580269149469917368514574090096348578363321547574189770258467807477633852131070551000640054998215176863242889754758529656993218524750841275204785416446508661614647568979170309980255454798241426276915609241289998999441374750551762499989889776202922935559056865380485853520295522188744880920554250457015419395009003698380752089875202158738238319050876794599845412793967035367858531763921137171562917168105470402413558860177750713812621593283779321401400233370444782744252825896616155710291863815870331853585385758668507067576817342453134743338486664001732164359176960876199706001844975434858293165037162767244454671714475115808657642963552747840387966052546961337581079959622912645265343467678764190406571477253426299736077244260618237119651858084687036618695794075478789311587417059711343431227303350973799387878820501939565505465074594025436036966165326241011378797146265125327338463869344259066761066708758833017239984197008948737740667243072361128617406867859023220628124060993563955305881007343422807152945373290736658535872277975738221441943388251274447917399816259299528136009365562772905944119535864755915645636481389721332489402998780438636631270258475271207701504449959513826917567118882270665697324654384539229093273515371145657367473351701477176058287971525132123630231302192818809015653271446364720664677553548552591939892622312036148679132676222630440973623016172572796318062245769564016065503400855132972119620506711200765168117939786675101936469765690013861859011813740916497238562803149602498800743366839594464102222803276025264020852530869628324622203235960629552097202049867177844618660210541628779338833715233540816837614040165594121213747804550323106645804853319343703095281130779452288398474053206030303124075718699502769043354788987366303158458622216238559388852586287397244823587628905287478185055937738386553460517964970178405297818574001552698590403422567780250925573684712440051738280186677970204727172629630250807588480950618513205146384314670376008562962889264152845320978121184751157549620353870769614740394880734501759269207166413065178816343866013426191134590442802370423292149028056236608844109615255164543338823924604724010260098700884832237411150539327434122637626938602730735567837232664475826678490431572548909373831937962691028377457710962109223347737888551476979770134673578428196780307228525530719932119728519681020056415065579357762138062341940715505658351032187943326481465363443863803288110372057213933086575266979834114281470948733317432538033803753683322267351759794373856980502790076050489035615460200962483659715876640167205197093480143268848588887523572869448200676376187943746855377974486709054344584706188400881524241541391111150657380625806644492114953522924820276748689110132199293929664946288230813053178306354648926761437335954106994213287933139740635598749442553854302259023579655742203516047488166774689447569867936622820611204174605992950298769916077128982710510288196427378103368565164358036180004462488782696183306161348948232055760805161621292732118579698583726535923141524929496406363413864269242809640344703501492582452690207020817139987977758629694318006602700285855738200063021223919333172965177631394627287558397039465292625282839307496623059783681956962283709629525150084752829826413171823267961454311822679808362195466731957355942359238281197674840013554610317457007073447347555573247289794213606244753195388483334301123060067869123641067367400662684880401594438208468350024707708829070222659764366228295547652120614526659158515804151065386675507811098963031107165735090412267764271530335591618530735740446073375852583244380837916461995678501069928298475649775376625135157948147430845819708807270087123568447382546399604201642056708879576050534869014874330086255052538159839286402227545108717169973795449513568196146492612661017844433986693504400636793705628179607690520396302365359326058011744993452664586783672366992937395096218263188368979803013281224849637296442141604562461162491487916232984232571386163189599947711025232461545410107189286669897464249267488386815630421031187491074169125825170807382257645180821743080058968318717203308067308430306008719215999449549275186652855329514480288672236238137561214814466389304839692226820467175658389499720385721411105875447062795842543723282133190827930422906881951079517628186936042926475582499991844669144045926786553864972953423185548601733151216289576168751452442756075554834451826997459883888370003537652029989534356334645187913576065296792644524200824584132664698775130541572175057149540461321494098316209986410309373087546472581337568958627121819007771825027522508287259305852577876388306401427150201948576417401656477372167355512124351307696967324497195496192739279215969258543281938589966654331010867898501990694317165024492189477378166203114027287886864811921598100732112244472230470554383307370299418471731149022017754359690346403013127778918768231606030509998499642148046196935810238064412531126176156308402893998480786811137872744116447605557524431862992124455086194694346958998213418949062771097472347167350139668646981167262664987259407795730228841737389025744020061321542917633986016330050107026361576525861361343063325677323542710062288653622008100246365528887395634280352106633520268540358179909328076856361340557918726583703265608526215873247374438128953980193021781235856634443870531703734512403790773617204413172216046274108987822929386890747041850246872279320141016139210667033177852792102956417783343165185457063261014219019520235689446075505595529429106951281454091113243566908586495716555268930016123745659467401133529140474108004345232048212474448247768159489601013130975476694099832314918715403248611252016649922870384793692622229702610106986985893119003277409080679262329717165143031612328332615183711602787127836569009303906567838743875732384487702190656722551392684192185152900208659556329503223694369365936502102766611755792835591109286320654357979406231288936371342647544766199694127510725974843267121938848266208727537109341946427283767460259820749824899325101096764583474599804744627964518125872088655466147810233418924212562023675468901778045807940836323107298408868655304836421435151585464449138540150345465287526413345767386761993053273799594296798277380929249849357227606453580704529153148918715666037355647559732620792511746446385929050322349254324570411407740388768669446180113069843465434315221996105127890401599662006197214161055084961947355022344753536314097871510439160892202614670191975302380638766097843290952324948930552090545562940368121762216843274341603413427737554334409079094656810056060243733369584954911218729530084716394133205309124662246096755881559721492929181900212738251392916283098708192082807933814206925026586812173246109020886948080388323334532916853674293858738746575892137769457705381830230677106564567006150503545480458411511998028515408053429026809142780692764408220248254464937173642843127962447327355407015168381094802887654612384524560495973919708044119957466092904213848780764642494233732293224036130949318033995088929763380074431871523703593426236859383367085483569982629686960150819843615701884209216590737774519614344445579276002239274776296217803220678929087832467329479669456393936850538058021066265509788330120339443730684041783493433403368660115853967277552792249908020009855627048227361655544622682459131193333057521100975332158433023062932561369835916252724835530963063956451775313220617962849596175095470955736265891458562873902748194876719191960799278465072154943455540417890854465225477767693389605168449190478542196589501573826166025865606146645171339224009294572597815717484786765737153288478556654851482881403515032894695820372415060115302875192073646119765406141480092567413252348512968515582746901168207643774717589960607598404329437564395139994329635730111110633681700603685437336109634272210103033418054450224938083739879836153523722555040316183597103138650374921578649988715021842250779725214483676183162913274020607114407980836315188919652318618301784597885748980354018002765408849399964518595757380248305671577492554441816003221000192908285620784557292986217457457008937003087353143312855574041832420112317625519541622257673910321225464856548382355266230023024517109008815119619894335099424334371497881198051261886997463187492354570040820824145462052230431910145690364685647900970770134523462958114721686909957771284518586588615757296975240405197323461312166300837418411100951691362769269009598174972554545066600282223191709144658531113305195792079519199933519361197220534355669335573638622367318682415058083685744459089503504693694026138488909769903440659982352197168694607468610546527603693350254013923871222214907366013402318469821065219011342887014544279931938226403296808841138321268966015543190588816680865180553680311946308586320183367554216227804022237489563325384863487502698263215652914982484331951463332308738138198850029274360944596748910049423934828422357287796788094130695149200540222035121940850938877624298622532954796851896147160661437050220616565271922732552946446710677232516460729461776038659142224476257345847794450942373578079859499485126730131326576524516304890314331913177850990987024168111201230369088888814690398172882036590299111655540010027485735699714746027888996894491074636860280327147225823944248442261836656573904886459040779317311388300746153183682936598566496118810851117210155851497355418607451942417218342405303407466882340739425972867428295651119808720504249639132122862955038307890969493994195104113046524043749160794961436697257609684147927251684959704460610863296104831651532986040825318902300447440658625380101877503980395998227932398085723381177857528436566512859064891923925275343188259160781978747236009670847930228634801078413556144904608590374066578382162840734094704556882588766886646505226864738082307119224613333238133939596846371751338552884133868538305671501850677644701334294545522736725909403627900425428575515256280775226097449131464565340866109649126743632594122749398221333024430536383625578216179820036875258724363654014144693138078275083473221516269734620056229410319211192339379373705969871361716416377679106916388237216722287776659174301797860870108669335310887042573584964084549830853806105001556741627463785217549742961939001285500174633919141246947357557108556599045483102378087992187545103227928717679344616781130238768651483461192743956507585364231475371294215456739427454883014263789532061322035968793855933561987325692227761095404352330484826414411503719090531177372003299897016197626979927395457229994791267749526153303628872678038332800066583619933258004913183349079048604629414316357319191327195558231837077806252686775830347316627050885149851971577583062612869863587283067660531257911700069173571826354880947317637083284349872412544075913402708943189703819014885007669840239278005529613445594638832571533109906638102771308609778307918887972144572710048156727332459584973070036821333151613153819961560585678263712112030470268059876228517029425596227162809113710893297470852860095341562560503380819402738344574108370654720306234445795829592390942617361234054693439157976577350694947971977807804891223325574434005726664365864102644111316060149623095131182327136402417631294301589329195653465200427192696211199078526072648627433162854506513539229770773150234458156216226928518387866328495285429966654858464869353026956392941429258494349958271486167595912715257764828267182760461364253896726815111881010426744887647580456490504967345425456474016622973826347763934613338287211436182257186176588003701140172055173997504727844175707668889600207567306051631200903890130181471560130625198949120552559448558746812629141017684709710427004048742389198612176603676356004472896668928059680765558510094787743457924782484751630770082376667308234073330513906096857478178233683870995690359783554071103130624719949126370238910969540603249868402536861477476283095376112069159050062207820661304599450914308949364720434475443492964833264370780845446960946001797745772495904818876793482774452693415669479497667787475534772460122493492137810993808000151111691566207109773752171706281820506088282730034274961444848649370619977405543694383494372828173907950266259467252174774397940298169228566724046348868983056956512001432266855278032438198722850314961928002201732958195918019411712018092436685408928571063807663069906690683161055648082570690897534741031264692365697991501782530491220185233054814248544978607008450497309502054372363671000481526486566864005283493893364400968868076536108775865121526619434205088620757585531968138561890894217464218137576596629621478833379373398964800435850437577640968460377641335406648327717013081773794010794014952875540770031810880177389097917673211290120018726648863589367709501261388823514002523445274937813308340503316275556971834095957070311151685663096650735368840013991029460386047281639940987620333442175681255921413498908635648547978900644835092929548172477565064649311738623373240808627814703749594470924250253523961355950685704430632687972782682608805157573492086682988610293226161060994095786201637476231248277843213382720254902308399876150114954386985250400461709342977255489015322283778501777263653261611258177983369074785018325317279141863782640024236415716586455649381512343173971977193215178101228014653700412479738519336676622766507435407640526480204490442902692196646910526472471764603114393090554086553127756176783454160035662301219800609443889189639503729673575343707967975070419543470415317588945540317651342669175155501366866991136532252978456931698975357335541024815635056389334249470576140186655637589083443486141479466860843188722766352909711195732712796762394563304967709413436958934023108929991080515716005090622028689543836100237553062641924208312610445690723253952243872122829998636223014392769452303784254871986926656774292506382959408998515669855260604591417042824181753974969549598117789414318297442937172941430209155229749282461303079594444929813302739394766566525697098691402533281916843264333759399142217857657800583323560684781493900290697958129276960694733882334672163494056850950141579853354856735690088854960442960294027198017859950777639407901068888088712399238217925269809500311232705165587641320754998776731203952540771780852157239190565359903762414761448208478958513982081094679040994558395347872876020222383814294343590776799037308083155653657107998627712323242404749063269476117901423559246446246049043875239453357276008842608058146113940343526263311302027595586177944864081379933037042171087385719199524237547368354426768791245302511373540084062608180152747935881087224444118033523826554478381262051914523602485027600874489145801713632787713497584751472006327570249613430553152241216076230132010422320276409137336534770453910403780621671522392365990692185624618798127690726849905698065462190562548843867811239017679715179217870401334571969427022770776317456491630491922052986374561038895211523979527427416760633732814942980365064600870622485988343876070384800055371145004676234932204325781711685395163200160695744715969286104365799253691170283772945984214730340355595947699698099014443315017119481178618492015118130740993696491711908414895922891067254054340531857393960808153425278566229517436788370382552922294592734092702611580967299358770050314659048255709559925129488294676708030505260000072877607457539673939927320349188082988967058165417599841852725089899200883920793567706237354669531323071242303361987718500437774770836733074952878390038891081494672745424420160412738654295861782690124679189579057467108678949518611668443667084341533197526296153404267552816832061728386036482614796835743851679123120074530857086709990170706464291074762677725912366546486070021273493408424692371122397873634380087044920399569414776671962284365618374494404783009454907098151420965665235024860054584674458829994719840088578484926907349049795048661188990225467479501168878164556044747702179375760301749116431544796543423102855910576078960509904137472289682236128170228966241929030873619387884286498641050901436813025342891314442008308026895818520425411382325363996463203822969014960406971167617803490743246959796489176709522670239269012944596424741201659344241742588809441910058483052316791067695730906881611834731035613909961818300238959553674977192273324582263936445381352462640344317152685672375636106063538302866212782300594878262565506262507016585942350299700482594486537247614137407731144094911598703420964140653507907973000042395425123178245693161546538944817163474529961240214053746847925700387598211104540059955453332404207376202292184231507885591109535596794571000113817633180093490914619710272097789262841684102208298399555207745134500720393760778388467657647832652533774955213115034905031331170372907717242425891871787083064105096470470807217453134164901845190381303646624497955153487250601970472944597069341107901218626561752688823236509798434123508741666811161583229221522141134524591417586902457552653169145320821793704344419502972752134467994934869096407841923356492882008114296882064839353539886265242831143174995160762690880593232309324669792464735759361400066588228775630758091397420766422347053339669946408397898479748702174241725385759515809588332981191183232863126414025028855857625276026871740845329254485736651573547968942396001683261003382924583419836405578403022925041315498400193585323516044788132412164872883956764107603580761015608253210334039670084993337626961638658264147568821163860874793294101103518067608262144888951715539977598455880517499923088750685836655515716367990472782289753588289643252576984143053751695117510520372288167030240744170016627428431866080813209936498367200961001217572321533582911525242084857121897399133912838543725576002777205065137875354041868752613190374824508337624872736746075527919904054365009325447965772617868469787526900954662228938151701419250336137578079213918409114453283963284518704329695562584886018310263047577013189386033198634495681666587191729216364260315068566525433195460645621272540637136201750601801280350987654877503290379842996692304871979737361237005004450111017209863421434448142080922058731517285624486661835457169958250859601191760247804917587030250960118942516181532668642908379165546572246588236181452870059164059291834098650100893185050300749827811324812752780635973751170231793524648742605639519261065842348084780078197623044666908304947840322375238921647068490029703328925224630287711874385761906275853624862447995733481952579457361161723698658868601102154939831901682740902639859568215016543057699249042977532851613233475279198188446078303777764933793897598006364490095043111364019162853307148803445148714481253986183475235409501451199658719852163307602237212857517328324954494638744581412459025744214264604305408829548141050865942584202478886159597134834055563783789931764838601652428918764415643317948502222059711181053351828207618627913195923795822057302474405998878770616459666538646685572019993889307039142440734270135086457562108234055555913602152888317026235605805447403780314423093360140310738052221104158155449012986496338389532742495169987871687962131277199254025275976117059067166378905169772703115868726807285909921002990680153114176973572430721426455006622038399407192717389485597717430090700963070646135905401723517178560389162691795785833726826368236732660871254754437950905868026860040340401366364436635582586576320987349504904768756226289709224527496531073627879634192454883225759117756134433519663815666388742633604712659979989937746235144098665378461260105924345684318065367788517213904114086891434976034802130574195431083921990953774040907906287164994177909873401433541401406281055944976806366738199196301246628425039477700999584927413239653022006401820331668136835304320862451446420568841630605717804812802517138613388353327527224213450160244942641523828884896780006961744412491117507024662800007956065079359572506235985413451355939632323084876933410583780998867789105653437541591182243656510664393423174423513282895091974044604025713441956635265666081664340113582242832446270925321986436636522995921352610302765934147247681239731855805017623392487563837956337980740647643258859259625230968166051374546932818593962922126306950302112224088870480350891943437215741441855212365932127415525357096827287867943215371873982557655890505256276438799243174077765235070642189254085980721618799647345053650888725401936469806052224188055382756611213257843744940816681770265825784014212608576257356134558715010464453147873634657049530995900860251300368631097592822266898046690705758214212040260587867843995752972455054166387060338521223203534544057498013093445693457199722427763876742236301319402507069545999249384877104695916545027807772467657018619860020828516804627461416534256181881280343208607886177969725316473276995379689412004699194032637410762029497719056835213138650072080295239993255516147171615294023290810597837885201934456664948583259896437704235526839153177878227781877522693174171993336453452050588505768631373387093831751846908443618641574847737753198886730852451782615672785238559503699166184109989475413644759088959291864469293601685700989840873477790483712933516853540988879876691953083772843738010931900826214618175555453671792496843479184959158753655785508410179144331715650847643903821042480226896684782295815749440163147413334944594136460082846748144188409391473422207520617588508610707264101538667653966477921939715922123508830124397049331940593216432845859736931979042092526037821186592843691795635380307576664123718289563347610171048717452247269745322673025738183653056369137030817543511840883567616269178877715702684682190762738439841101367629049613105677104357025677819148511813743087943572307059694143041857016616090920629207137120293080951224799013902087430936964531424771387224605271712394656792951401368133375304662485933127812299850490866890006627921273025453571831182453606726322191198851514821830006179920236866719614945157540193073071059029243928716806066986661745930649013493720457426416759547001036479949983232186950068072545318925600134489641539080502465063196759686205863933348252226827948448786945594506227633824991438061606112010641705478589713365261833710839481822514036991180735281005731150240138695153195639379922197685245849723704580849473986791339573714845783220497423105067071787512857119523806835886839723435656244732180788098126274270975050350319628682405469001773598347856825372536550018928016247538635999112637584873671673810004186533820777439275174703078694582022014840957205857731786850095893923052384394230612850272582984531441037446961202447936669331101648380773106691320617177061861323519654225570263051930591865874337327543075655360160373622977305945734477519210788855269935472318112232779399101592519558252467149084114518621740613293716930129261051303165246811458541017379343302392231344541558403725830432995219203720573437678829112607567232912841276423846446917501077375612727454000827786281412121825259006542529277939897815231636500332847040481095233386397000667544071303013369884726442536693319934157425008544450572343068081204804361117582244301891415246895239142374977327742276519573650394798160420938696928927504422600899900578460921357302715960126670268613664406309255898668806195870578996366536709926807407399180325362544275565838974676261850665812318570934173696")
			result := bigint1.Multiply(*bigint1)

			Expect(result.Digits()).To(BeNumerically("==", 78914))
			Expect(result.Modulo(biginteger.OfUint64(1000000)).String()).To(Equal("300416"))

			Expect(result.String()).To(Equal("16113257174857604736195721184520050106440238745496695174763712504960718342823532757033064753153854111107217980956972311094371313807551987645270427615679688389330444133499923357612295105472898569728180212454103335089943629049973591907780692062057296552903280552403216694239075203891327355460071271019202307459542499755131303645773654911843755650169520801985361294149607500956837794546801458249510087206015145950849590956971307083110865729968620669481023248557683622154370297881298560466161255588004643805849285632501063457275606556201834157077069866350431709941244149430340376804687891750978965251971591273251528666837509921501820059959278974772955137066499435490409471962069956797070262724882828175396608842375949253862079040514343986517421482079629239951733227576757861140019256666088017481676787816929626050753513514197997219880636635517074110811110090917718919814526080394093014634680319034280918429077805893913799277691975372698510166677624523615113420055448312509654324275129878416504650858425229867602462924196108810701050280869433809450437796353064226137994659354596448860467864613898876433352470791463814942785502343393473062938759017813105282015317357353936794982542364760489912329987944237397830950164269372928014639443070854650822815514356651960897669142478258221975915064919190375531839883496687104268095384564100781761981553851400836112905388609963306865737671824507168252726831850638694385083478775216005992042813943577935491651062098935775915832229588820317867889367035902796321364702417695064373176131935749829552011220191660916305138249322188182375990720523193993581101039054797188073579289719367633920101802417328162711625452959719023888275970325181461602877510137836219269753293290552689428123837471997147144676279554641219338382818600605281421478609887659834378915030776762175009582951300819040556136027426184867045130244663144953753637822529952539573899772230761617078236015566109248692819139423308222674016049972253136330137814775383241846559753843852858175253058610850847829770607304611716124700187691915857017825082140844518926052368333149254496169541987863381443150983371435049841095258420360369083090103295668638334162824653920097285055010378770414275994926509061040584889124837133768897043968307127952138470872615300450165518102683181254005665179713318415328532389713394405555327504617945823008727675119463672887489448804028948671392868411958445081586719122974263349688790928372140143387287713663300809089802349683864622705829708886636287744223635122172790989703032669329884394037058500143828581214950116716860528315076226477760817964144324748729943067269966484491920136287559912723443616218059824705057505221252584287772642059873903818958395542054498311855040504093132827278983759273901952743640374066833443129339180348794901720439121062595011815290536015800793598971588035725982691325656250421658993992750217769570407199156371342499713209335403268331926928188925916061040172477977503083591547036451896298938818110969658258921342134546078979813635214571553241776199842879454073915459154070443752256871790251807688080294183234647980216888629693731884483776952715044743111932510417226055752254983593141687746322886248817949397875616943177089463482273724996741914742471298389193274629178895751123671079093083823859862474218781527382088357620133485554026665741444384924962487491098346254831515003190218011693146521567017873913667600010316023793869684989408495246498043176139308009095514244188928518617472341476450718550585757427314736099010854358912574655409776135568224967205992618968276136929876717108452204609387075353849590890760173812701782860785041170295447608868245841325130970084833035387972625018732615336881865889102346618250302128107123931292918068737118344338313545322644116076475684035459816775808681366204773537518457467006185593938507894394275319133020246665548824289815220177922934028682011073709243695540220528797188216035063505220992550596941134052174060707304493901052826266164853364626816451469256847109850298847543378453288094059009149293012009178558874974148460461498624809990016284948394870890941344591524747761613600260225600426568174570070045347056341629382557166933434677222614214042092138757715890532132676913787639898243620223799361106862973768594284274043455832279049845160465092789220767646714112126942588742429095504819910231371604370996080060100553100946043950173081477095804496368756478770825917893792367184587264849049607501408296584301947812898746674446463661390110084018682359861055784278997703386598118473386824411884425900743923891632847704500432374995763635114514642211793321534245359832625617218907285949846958338171959464305088922422874915346835163452079196022399867597797151875270498642358977767549548630912843719329065442626372490328088854902525445385704617185212944852828890344516939681736013428678388093318717099393029176054922082896708075278499249327762753380790290331022167959046721733707106603767071755800178878035216340716251277802299688464456855269524478174569713858648568287605824625725224181933891729771334259789489866255517134580570061673257153156886259038274884109335731435506427977689573524134432895773467575553414576194129483412582388516519520272077075593453453309712391026933218291003752842366964826392346249973686284107809735832796709401614378658616463280761013417942978636282203203768333403742037291804901258207555712830451794416027875524272256905586115783507382582329090294154456685241393226382863194328140271376402662152619926621396091859827640193154615302600514682313044672676692772785579861546932337862203629144500485144214220920620504964274906403793380400759683806624256781626139736423909814831026013256846856732586450576393224614214184139324906423745800829875492653101863235542609771582492888109814529526633626102763297644073625487204227577620158078985084218756865575287350465291900491632578645150534642450631134519553694744183118306801392108677266438369422557597062553304608113117614574553892362663424648068025355016607192728787088941248719091144590764576296086737213349155511388871075453626679342507908498924421125135480705061389573563521574374588408724970911316664858091288764884218544916059184427168410036945535917998929621927004264578666500199086043729403586345112128718657515248368913071819431167132071087863595825163699832829880486628736325970480429825203771271646270144067855164136771945088863705208212338393890892339089043039622313258084119434970801031873990546488861948620686361636593121938334101829346708396663536095623553387705273038258582009950145521716721385433756599974426047141252426463813477065837402142547417568394251282485080521381636783698643044180238590302051740235382131516867839209321760950551092781040664950530906393659250310390577375443140554770765626200352591328118731507419233418008195259251202450519516278233172857112414456035849540457132279290528779704622758659983804569324757625808281249658694396781554316393777828373578209339939865847464428662506942945448243995933556838514100572270345840987535866122955269932208146609982766788167050411745972787651653016830080882640689084736114935616204292161454447315719447396478960552299510032150563495711998967213924924261204026578607397951827921171653388147547784366789196624634279261586879091209254605209668016874142841942044300834991040592909973350993633342914690979764928561803595111360263398245785164180812243072377098789529799839095692375445307278510309423391121075487096507316989382380761738773733743090629079470189911572246134970794415725705486860803608401970618647845484848405095069817751467957166796088955017941564110523693010872203757767024879518970026039482286777736274679858798696107948697965132605360449712694500763776264736131595747176257264378385555977015978818794653743451939227551940799300001934352848833919345758688837936939611845519565591621589084295505702543557086645415628615548218821988212723318533933604986615890456772824146420122586057681135080279024654828927716024728152813567431485547958876602767207539002416214122639446318806319412056639265073362903063172479681260716477067202620345532212559767329554760893005200949248589965677383129354497347728742647165682572570830498173141872137976455293717860729652965012971044610125868454358096653555176931556520771773929908690059517059679987457436987262273750125465393042661670626373191795950905579659840066947170524728018719322437663864972649660652137470599555666725142812915016465847178614075083593932991376427399472684942137512658720737576268521632975069656124582815035489601079449210318670149965079822501528403351161013574511977014405691468766512419504137272146292370380373892514366935480048430496357163782041581264343926465653205116056894985701820364874697437073617253179973102868355707947323913830472063964568410687731010929002697131294207196369741812326452371052844029770704542504800744176471756458557387995227003004056068123350592349849587206726996688593251553075952578822619000716568816542363735887809870302621493375113294539382060254211753861190579294066643035755863693886173098320164296643861308375363938474140779306576386862513978017021672012349400669793627412856034471490094553955169805529451669174298412950150850254478017871621041042877793040926570905903164946327555564732964165632731458678836786879470211438595161203901036913107683861647374778139311469037814816324300061989127809519904585473356398821301193866485955287121496624107719970626795200120612559172719663877136294462161209594384286352697882825627400553624338731938607743807283836524648520871728361188100583654791521748175350234812324308781244644236743805389248502737707917442558321040993437679132877756762600166714592838947363886750701164331035496359131715165127547014177732469198380246299456330112601395030069175646690896410971416061054996413437172166657367048279714069164791297818538002601798312594280154770925053182262415422985703712465529129799871565112736700493670503669629821964125203833106740860770221699714111237565303800775313708475720318625744021711808350507918559225700066888577325078229777176428642727359192085627266868386014920582364239766577915812515837243944733041783991166353117328810061472981278928493910380360774950642679147041223721830715711114326628526241728691111734569574696245111443633205635880837641007758654042286705647453479999072309460597885203327685021648811551182074145033328233201616359885504451416192056916527319522027737938591828249533719115542830306058565485211866949730442037319065066071188409576398902742104181290340140616670025082738059212514316121681002615690045127548530828647823304591628098751184492725043199694814145677783128539654189327014333783344975807331793174911313476525515914413508275735185947084418163235857253639989714502299800067382389222874573666120004840281905834903812928827507388499434728083133585286344082739738277228630707692330125653826106743658263598367244205547244387614444022112201170132642471027877617528525762191306557861756797917007678629902625642613170745673376693840513847761678572020301677291983585349974686542644198832714093326084038219973754116482309069016131785074659874280289235059026942725935819706701264935694618990246100418847293138592147165981983424047719116492694892235758328431720451252036781309472254377089996160754641336772142283555310193420129993736373981798185394494154194139886942349242884626627695145637944651910943603426272643785626820287153824743828680675054672201523835454936459420842216710840149966230539754111445212749302013658874747849766102631275922326141257253236652323420478625229355132520998054618429633390639603645105517752154225544172732952958578411368480162071005931282394982840379690668301921845417282761681137892533293546818211897208982306403321131455163634527922634152763973788064494818757730826876361111776378220459670837674335147607697923788944553603008514400693141822208370222302934755965208223389626362892852356434100402946973137876120906766964705251209501129944948274772756921857524766203657641947430633360270472652172204782577924887475034606642791115695067774174405268772117626165844023601108210166117197535320930606130805193178048851920401760657496337034998270148512837588042928028622232452754359801549386555345474344625618614094089096644783763800716806373447645421500142419998544123558509893246899097116566708377658077107507327035081904307991991804550468411665752604222917418632477529295669433166643468995557602151108168766247593811449013793031185360045458384923153541010492084351735063472929842302958211409368168642445907157133838642642050106728111930620127281883253759699201702830749729908587790210094767508520773344999474531150558975569804372813443859628142380952778669798006518295889883537870330971806443619576437084653576372046893840431410810467495063074899177549549765478392280872067262246060586960433199487292947701140059142402567666028387027031621158753180603384271166423681775122558262777830937328618753236095429230853812978176943468141241867925928254686687934543986925205764645522504900928622639777377394448744246991363914993688426260612772294511379613492269800764229636361194516354685142214078555837385692055399550422747160388448783164023242944736877692198867883877669349766312031505125342618158587210294291436922785517171028415724984971992733256660037585894691423027160168174776733210546671445905994571491511245722087019385561304668966565541797027372221594683858311680472505528784156180135135295163148128960053379321047428127200942639207331483131676808339277292350979841644863865262431367620729236393766911198365504690219264398643762805037959161005322767126589546436393494528280998155568520940615684411170898745949146339642462665245225875915483337523239846667738061865684234325478942262312888744735900989685455179980222685827919034227143955530808587837496136249094143183636735676511026650730664478561686206976693106435935328909596273665055251535911353011753557396394500483383650726181911969944305122311277206823917016603494044358500694045656404123315985364233665313930406742090854323516907820090276739560961781799613769742324750124257366617554771514704068720774975981354092705291779785573461627141290804137112211162778000875495473271506505948553224936749021250254200745669159231422325722591263257285960217792315983245785388316685918678332276256179389412608582512143643525027920388751702540281806469505098821104548696487939879098120753148180351040504232405186865870370945692747824515295495985348147056139382763254511400707577720875010160843109225374666599234135016877288629586849700550686964793318551903174420820697937866159134119533305676160673753057646067527696151320677761828558616835796317170118374144153627484133391835151755708064987669218505826732979913542865821890443424015012080550160589522217652796485347544336138700304614075911139873257197441021033462976008958444873245411245589106803355019013832898237929930955627502445932328063984227125923131561847625308482419414394443180601555767609711311087559198853606290167019685786823450056251079222395625169006309086118708033677325320203968586919924589922396898377885207461888541962337333694782816053864489692980813353048191466455298284215791571374747092387552370322210166900202993218652112193418570555180862819103676436529727333227276997148004931559155948050806990503660123335204883403362684526022568624337406217240436346330596308266093209060476400900501736461866833734721684113743098974014126545713298784200033212311046363906402437682722468862706147524520180612617517346471922273350966161722761120740476082741542044983423078444274728318013662953684403965367951809405919848230074632428438325015596151549239910119214080283697321324049390268426265222078512272478813187092360177680944815011008988016504613066344857982241662543766424957465151032435033534463426010594770563961281488267268706418995653918080411295807051697482712993395741124043554924346788670478671492757704038405330601520669980771391901776576185882543330315948219174420947487068531614185091707373599505760732217627065777133041579974407907125837267208533840797090136198805538070774030210878094190371249371981055588377353074995281634320002647054578003162225500029811960672186465007672830325748726949784333913664200225697626474535621483841413515083504956184266271390877306176727347330187484973333510980344690047872625386793499544875810821414768769738319399150839803968318470766231632066769327531130762627074644649130204786526005866060615281804824177101685018240184206997772524405532698271623424841611142284230900889020763447071101301954756708647611568216578737176223564135806223063683265381070325052445480230145258070042568105305065663020094945608606900973587424360221749554142756969249273715864071681578590887604607999042808968297707048268999904082324268495094010801466557042728178029022640042656810991230993139513842129055175128969653358917210438390434139581500910536274792698594317132875843870957016735530216907362394319532641556843104482934284778876656959548567764456783001879599562908635415898858311456240476773620585555344590807987152068127668034689956648669994206542664048016307836214901135295834613801153247643658100630582271741201570270808437105436504503782537814629870284465026158841615547423364048845500201014957454037204462955674604050862312104524203011680544045023525994623675686459359580261888128947840824378535198059322177257652535615624320463088199043094202026934409205713886851708201930462163612755000123066997643804232730216335291210029549280678608166384319463515639076481804131675284791777199019909696371867051271795266010914817358426119890461015898635574787165094043006162336604870524471371030174464620625751915194411249905503358242066953762631713700993178551080186525919789366321304271278658121485437460836319374638161490369536587323427553071630452891109544915285201702418157370604574958286027703724696369341853223943354718492097642593325444780743868682970310024765609342494739337628395429116030436220089921399103503716732260334940757373972466026379170046718770502699655978592143388941803817875573206584577577603892391126288150975341062435633328037381001802715313087203002468520135941520006398941857294348070151548835655640238493071426585614661812047850027350324883599555435147330551128542723022115461936064387066014087443730152420872248360256662077740235321589378801234643609372780257755415288985401052406609444416005122225069117331581393952818400980803060769874033536024289125951007264097707218662139987799114107904456207965031325490656774371474437202342939183047411473204076423028352099634365966445782122377780452520863519485660273948679885652203147136004430583530694582657742915805516220169326452255513822070023244436143824921724913303547015418964621126706036448073154025401300997767860171083779303498348861662171771299770713060016422703281719390757161516092714122653517584650372813692153497455758662296062679485058414842327320972418377941910931615029993490538826447984237212288016849023779243483629587494002312740123191551083737294496001489710037223907390326017200678242963390317536036900487922851084841071020445052673521569041849682094721554711885952823695333367457973451374071507965491842933431248116892356100677869900544469506155489044772623906721391011334771041849306813657556859289338043411410296431800234178209798153476408805616712392786919272431114217659988166354945717434473388793031586970236229449561923828149750160029515851499861560916475158670776410533410181557030008096663858289664193808670816065113813423016410165414766617058586516916789912325953004524732091073687411082706909950614219618351130273043299160671063108199868361774557412371658406616928625622725012299903278156664855803025238357295957200884782636823136854622027534983585287738898023386839841138896619944868698001451104719025046111282838758216645148438720171851110428013877174809618445155736491520745729904720523246556082724800657398928746839986438734915688453853039675042133490295145835386108884299321330650391655985424533914299037741683253428192908870925509652667658720126312233896557590053248026040311443618758685047355720712880590961516101607474539010722315420569807051586702456270515362309735232500200398034728604816704978661170296603654583138773523007460054065887516544434413978946601727159875210083810599676843227835932997285740075397993950267721847892905773590352149708586854711084070957541870896790448349675571527983600679021827411806196431508606605969536436343226414585771659800318665646821206734303194351407823299097340039654693057669905441536067495132192561427322039826304425550632355235766422944903543748768406311183323429859160263077281059067063248484091782775322273967164738287644688805953240905112974691226747087985799681406450161694403055589456505489259726361778091269407893861879115080331745633199046493657013524098669031832553860726864006228326588225290480484363676746575187810597009227047178615664399401875183734811954887981036159493353024161758214677966253334568190436415308366399849861034093873192270315844217201697813565209227337091759421552741838380721225606983843897755873943319444831582473784736220178233524448000663248800112782060451282188812640156154534166043185637728275224494656800145269147487405326240002454665000695350583164896744668898174403818052498491766451252091944255872679811481985535039482455050455114737962741723328374895244156421090398805091089776965119798256551260850589125447691956216152918842947655600441563954414760451586493780560358926857593880024662215889771865252165836560273970906476521950202528404602062371283000022934716165041128187914152610524896941707407221753994073228711011145256526889738814957168837416585906704837778090333197266280914917945945418229627397793786806002529254662841997250796048088450047647248782961895413432451577808479412730119918960143752820968994109839574774089984333382990160173005735275712968900714241583007704100702696114918567285391258373735580858083185400275043534067425514617729293065225083065016733022499743958009250880484226772853839405891976634374451593155524122914576700631711226175307934685061830339933940197781852436555195159952457463688337113591009088624615674787480267757603756882108927501376052642298370655244132134911337010635291718402515130463057503796408035756626618758229190034319667523278624054122493005440901330596622753645084644435916157179929071940972394533211845731003027323788864924307030660560273990663454776490353239452862217976344138514017237639777544435530340448465845330889676210760847933105703466837802583640868648626653518478370621469617789239366231059973418609304773802723978620069251384058928856015008671960380623315276332139257723382777916908967716057739042180059105533590805921459852181484932716837281748023189616105120850957219928671171047656045718333753333201081263487594704015652256666517609131276951943326706814998312986442880105780040222898275053679867299529606852364008146563759436214715122703952853062118924288988578673629523855435447039521741601387312606740279431319187009388258503892695013482280709270016346298923463764536962469408008329522396337912597075507853288017978606385728031731376637902345914248172233563526143889766575011501195896803301206934608655199652201051725601854881870016223767962783448614069323186962054616839865634102258722003082123534608119116978858510974144530914880517413209529977234247126418955499335670889797843516811173269701124266478072289144158172170058027436109228769666460024676518776942773638176257307346480441232216734116646862579436484792517463431685313507664324678092937414866101763842977508792047135407893725728770395624491987041705014209767472443715360218851354843938158665704027052484247158563405027573197855266010063522766245080868652696701033448085444727360587333217001015659602666871277068263875317549527758833818799437637194264268394294152028110178420332313740496374360793709716016025741446465222550498799198857458840421759849924646878725715009057083172419529990585268317917087338913277667223972661829336659946045031402992138450343238200083130633535502440052517958462675785379083211078960116239980319816558332556653640586541156610434903731793085155409104037836577371366643295531460320799726425627431413593554865877739847314988733197652277539707821953382657715634982270009019772973735652986835756612579095390433959508208464608403883513643367891308048407279059342412470690352746052804908198666619124905026626319823012548652133983376231411923140170876953043236297201274258025177427713744042675336219594821253250026610516838138758362337965377356843447590506268539901309408265380139979067746112527830734778056287470928537011125974065558816539178486412879081281222825773986147363150406559308622984068941751289876731549276217748046673546519980971212975819169420603754588373729989869776284783413505832742993334238580227141137574232396158336021212115635511232982413636078832197634529841968004383062171378533140721762021871096776116621898371554621322797737446997788605781451604598548724421764773857754599877311934510777240314723679522899957904387294804556219674841814846736380027905208188812013352249461619658053965880002012474740979288303265403133118708869080601517534459014292583491460903732740245714952518259444880880818851318235708285881799008716962687420026764312233140315452307628668248906073259243590974746867509063087656977613308082374548157016477776958582035446252836625359602247626889565913650781861401644522434466173284122311185023659568773339287105121482423021663597503239308989176568573904795888044635584325923678062519144755387462619489775957317974195520913033294091034350975673707620042205668975891921595816359950324644444581990161404780658892021549388832922106671632312197769637117429011132294851387723034376126907003843483646405998759538814052771466109729290657890275958776510633926455391818545131485492568531609949064791153985569842510448419351858051835125793688915203476242990681248673955896407167303955578392568786032616002407877588048711182183976026437504427746637662154033677569758091472328769042694220235681040477953950780207569724373104757230054952903127595433356180167988606726619460206922988502249405894222673342388306613098068804407343475831095303847205026390718476648561502801809297514393759346232143636534735151570427674949999408568792174522063283498623426197687985007338395089431714656156537042324932284618713440762708146569505542089566691646795549955942443965412195366260640765470463871746592338124169819121836392516582760089406364245822692935413677999189114050808508660078607838433788078035136979874346288501165533200267581565753564571576288365988708218183958333920668908740128382368454046726771824777517988623209605047285482962079847969483702776467096565438563647223490440649103288541673534346816535851354884831438539782862218809387337282430845353605022731569094938028229982123328582739462059037197660583120853858916012044881188595082236836147345560860288864977588966362403291015086171869155861949454628475344897593375423399657314856625652822735935026684788257703744431630037173209641928488322494726093996423373830544655277722884659356482525328646222066150881155665422619502958765262888120310375211613738970168308106868070201017352551665735043443610938605426902041410944403396602792591980181988675592160649987159949298138864644373475895736188369785621768061787916321908685230415943318360486991964484571480648913183219982821376086288833654283541529337173401961031487676949314448037642152051634193142900499024736312203614499423450223236752951075279932749410708038651938262017436268605551863691567016009440850649855698665556210575604663969880416797062717451726379280792416012425338297754311371869404126278929569165110116030788352539098738719623400802704990472695130598275763113885343971559709986758620040736055548866474046640816411952238404626496881447203273108623399094616035945119772337024355276436738693456888293165355826434926391271333729504044063445831448669036595540877908661132169676560299095476213215388737285738004972743688161677762349660411694287783856203941127059859312944850963364559361319469658052465582307435296549645813927762573630459036269077365748291824735094157064696066417157357687452756922422734618089537080483812020397880192888112739026738849852858927798810801173023649735126213215656863460596939102663432774567603731861529885289108985811134006723934675188624941786361815183784197684787009338249356171263739486981293216945696329972072143487114568702612488994452421322618785297950072405440191628778621761740455752432820192828723351439266880321392827729219342631011185615261908061235417092587695612425659629846624885146489663681517105394965523331030899408755061844651325180185154415355391058573671769603913366320966890717389936815997897891545530304828295132305363372828969357225783226685887270407465087032049336754441959752812086809279387739494677550499684763743269211201297872487477996121679168362371620047038998901780165524878700627775042977873789692447724149050374155120173035106881707523076203002299094957386603771469279512581695197476298287765043653450666623693268494765381940248336042797849155215927313578595859576229239549320528314495148277476557175508609059643130401797542196284968006361790409340734622800399759879808423268016226579741492649201559392842574755186542512660509177093789828541462005640222676558162948188709711872363603590001323278775655506961432310455985831926926161591654290867873686815250096601492847655802591440963339478017971156779042181997252271499320073210311769002562064035772958277566933192880926534275793171545734587055781576305143499648224402862501463233720739689004249011234741345403248171278102987789292540127007948647485858916625445299030715390189123191261854084073796009130586338070718068926955566033142472702507558624471361973654606927668165914780609749546167941341127080027658038886484545337200051540889065702764152362831234672751403716193538457512697070537301094445531979465701812190845825260146008237049458489019298079361626821557609625971976076456048763664229379402490111384453698080760559969998025364317619120903789556968006695211979237801443277816161870440177924236444506800394646699400032765399602751263018097599539443635394805185560013784503078004923858771074419280471801441164107307680405305043815633003437772322147886808773380069528050388798331285988651871763171927239902716799240581972257606807021547480133987817869861555401564376448432643374426830266532162759178167772704523004064442183592473355597913704317674684669166811344099154121249191783821932534110460569622673646939198182584460312380365855875249025167261890418507214673859199578857601210879250090503373475767074759127881029789555723990463670460830874180567393561079385355022940392600093120117575308160619399203025833248535585072637266792372564314595865996867915218927399721260533577078598445480682642448232529442658479791611110749667840570619840361603542460525128148359953045007952160663982580928747569674210151407749707912000147369434231594392099137345584025158026682891889386790762156538563945965678224934282497873798232603310545345600568899154417204081512224601146974945363474969502147623744166763353723437192642367340754629517465007382776482630057299130645835145359744168679110926189332917764869997549950702200119851569275361670303004870709801455058646752380730319824501961709057858208776676075920343284725825975459865094592442261759956067305329136799401375202438005556384222770257781221892006356331034697264382723071442572859702771747947123897262563463400118642092977260226800080125548801659233697128380536274880650286568935848152158363273468492345583610895254102221223913585813398097546588251514748212799285558170492401930024376999106167834448471893919415314638072855407892802697378859727076035214218314222359868836009010336867574666811528698196353683505618114897782586077112995682865073790515778942146602255142038486158192145656558116548768738658828177036634985629485867962418160411632255098526160400151956994694380086620365155803503021311520766402394013192722110718344588748306445977841813737792590050821609664191914661850253143648599819021733915985456138919433579597142748915425651772060468106657964238131856660986578439668355822813994345080227453658109104125581716218623345696804515629408477829791677717698610922738414386444311384119473908956508912064971476574526231628094998962690103884201649816495206161712055278940565704537264352800318345290041266171239768586071858227336937478586709470235168044422445603881353365753039629114917667176403032147108196885302800431529355968903007916746777321515025539695011225409506764671762835029837263893300157157333645553305238043634294485062503736735065995610398324058260791586605631299384582255720698183803251487120172461157348913330780378013252197377741653188380389943126879431066478528423159421665623028119477979077229955950946492586001459608516384505872266225612436289353841525845750263957105653137368265215097262988678027008801196270932267601761411006782162880982394894925055901427908573509194919541969395746737919766729089131639758092739919121314700980071802325359366163102200000272514413074525256877463800707830899507132296735232932426888755249857369258117795452438558511747719280093755921463732453461927845542823841885349194464214733027560215202389902651636739108367277613979791230639298590177071959745687251655270442323883810421329898174739926647462003082203957601420159212066277823404662773294144114657310089154215044884990087869185092142223850327774215635437418691830396166032977803192329225925876736304144029379203141939085202045519345591710729062885681507750119337444695916763040559613354016125836390129353287168715823255241535831215474783634061626022023950862411896848961833151593511444042184604354526917242811184934449761229533699096084448562333152495711884885973147029650996476623670827829319820803551559006280068813592993493680035437107702693544881508819730138894359666058987389493762511987292635875832498507768517138267468038487907169274233670952303100107204263410157800238207487350024298008748223304519969287151952624844621951684553357324979317860416175535818087686697316633750960237572138162387545687218691270287003842213349228813338862217345678985889870305172505932337802840561169601331878149715341622945195654387920807172042900775317677101919056106716783909574362145577191839023308952350574681913934728311927661917987713960687358680291484972281651198207659633987400754262830682668954435511455121727388019711727694048164247234122871545337674599205008048387227962747919786800719369538436285587662801217999760196555619272914490573910772761039585805866051478943136795024861067124316660052669155679117322987411984718926517816922176359811026741217430701022209405616337887755952644580226801972844111843543548504677056097373822234217214898859048537463160664405503973633797288384959241004580321766996681294299624662294270204987383789659086701847429832026767230048135590659175705348784757911930240558200229618495501402029339575903572758237507202546673704063207584833784803077201463704617955744596747132281163642735668198503639763459724772119553483895980562914816439356362959386228509691565297857787448008561687070331844707862506761344548290920600524522638471871683535647841286422235515692813343971890620905505261707674457790210127681215380234402191204155554614811085116922545837037385095033017992884253838015465339253268829976338935409543985718333270000888178749984982852672190680599651708960943790367467213760104019262197448592503641718895932487669135080288241957741881903553072178800320613863517571360520316224539256722127794531019951763881657930024865079712675770701383112095001918557385638774629771408431603264806680575619001853546569525571062230140159499745782694677141570187045275785077103770692064202599949527258525665594371179435504521935002034448674443948126749421960720434647116500204871990907321600238879901308357451626639454117133900049254046349060603252843130957471207171273557005104647246287206799335738425325309939970069290245444896845140504830337548997991664915468815751732733748587803164696447112476233845245764114983422483991682215726945729703862241915339957217115563530958820401034016326116971972696048808581541743861456127624589906997346760366094766902049973258807483070722571539040503275191424974522995868709944629298052514294061380477692552536825950444353753794388979003107516149046021533307351097727704482286501035988770951186121919366081418706982127097453682312272202393109238731395649062890670144847849051067861435801958410934699694534737200770467503963981568061271915024712172778894311060354811492339863524911611769346482128769059095661096339704546032450198336896156516684286178780923092400190904725229219634773503751436642973703772973706126053236404713797375017807223992548659549990832510530875292173097152666731536721347138096605575560163195020749322780936505574390949344506393666725066324268047726260993530332565380345689210667976918203284605219524076560268261667922609341588312348753324469620221291840493754464776609537905498124052931749528154865019973713935195108424214354476330948530385185889359477821191767117079624146577044446108741805443251245276292519698982880387726753657514408300532589916250583644213917483095627434929285443706158652953956963045084564742371463975517463788798389022272820031669775093771039248006980147125961374266265567029191922463125548586946146737596380532231994671632839957813159006266654106092851204656700406327435535605099144657132816916783241857417125459142053754027702519159966500878221656543810127474758157519897845269164632764259748411881256469322268852791554839998920121745805775436188136868162488357228897962196967373903876946010153890851106123060357168276533332029296583153179511687892012639566044546886135084021723033912387343029586707865032061634126963381060209001391135573100082066232122305559019854262000346057530020907053205493125314224720681493499427484565839252677401469979240258999379222716286115993598958371981150461365827236255467330356035412918452471562119957989135636514041163383733274411056403058814996219632263685944320693897230455825840502435826665055057813711287208051599256674561876638082178596585913164749487905320168768921599080819779482459876913816737071117186645768598250315766366313046075989693932206832946859104425826492403668099968523570152747234459353533063967646254944748730445159563303723988867041749370168230285150413841047058917161531356446751027454169090985707464252249267138605130250475544647871602306024951603897855380278332579583991912595437523742581460415363924240006244532936838163852927106140905860773078462428380028991951957312411082228584131993038676994733707765690760543212988792899274548857395062190813058631178647348515268393935704298068002101226931093316388729908843100173280001842519815301249107407296506147874152254816123053250872967416805629953194331710904277351350572440832584586652655380493202734183799553576602382985038408856215516583361614787488061574445701500788376868905522451835596730161970580312314510825885184528859100983160758389139916428035681723568665819046588232253953648310183926389241321427539603337522236206982481837231825079551720155142012096372390378529335435963262717094079500192749505742564665704113757438481659987288903197462085131530022276972382561010640062500458681995623937998538150642745981592453854426762250370557819327627103126605944111198642296935623120442045332172214028282416626649684486123713903692057605856740527804342653586677114253931425814505438794883846728994584201263458149032148101752533446610832563503011066461300221332769162685304465234634647730470145171653447215376691388791986616628429166572346461781720635028044672440513337428949074052204756193915709747578709056368614790044629013529911646796485010051955746524070186997212587359085975232793463666254211333132962554088429930686869913510621313247707948361880059027765582418911247402157580466303678515179938255395112865718569477868794159566712289089590672040385513549930663362283899132302666335083039761567001085147884115449449209592081776993967186413721178719604679671410896390827143014340534781822537840557366423118836589695339011956069599114399248671706946173215901291402082217249058481393149180850001695762849737640405191453021641178581735518463134755263680874060194929247342765852007433618568515422077535971058655431912685128718726266708031977556238893266046168015350273559635428925393880123486455724762099615363582185886504358563754054523384162297213966550453957415900964069337712235484320463105899303396219238002385327913444747436175312019938235691920235872405835652612989186969627327847392552710049457799923484043606176477217039145216149691534241624074377916740088041038994971078287468213179031338919123512131288069387945995937838409862529045742402399666772625910250152039684812577241582493982224237359506638352401364246479057710331976659975343134761046220525981107343060432974228647211279574038798747011559085863842870891936004946918991292389567573310455822006799748528405843610248651522870427262501707886244325625347561638201284400671795683538825198931606761426398928665926447401643342672122800721360996836198003169008636690279309065756658668702169070679822708915041233949124477147372040953403088475891237919916256197587641307507519468506688159237263839382134782588304846527775724433043653745856715252996284057386273568276147178672162628443434032383838463301373402816434806954359179155589932623794654326396074257539914789017247021466351958629003708577499551491804349629756524143190860737377620885715933762899814906407514835619445559994569893971511359133574530824855869724760307136573379017511362058720612056022000204507230052779044361624973260071504261016233944884290857486724921551658178815230490384824663837943991239139467820687154685578467368204010988778022841121526354941839247012819773217058573582801506358636123338629207375718768278273806045050869811005234481252729295649725618306963935065993235305388582389503576079089983112733132250430793475051157115017292521436689993220520881824783656938696858341723743726298085726930703812688620125494883073722654502945991862766651917880005315659670590695523676858235597726166307190029856953078356933148516302271041969125576456863658908248173612720894653472074192580775470888030917372422015442846829110905440386233178080143443799897537625085221053713334560984417745731930368579661785693531892708546181657001311921450313341529444392017221528860696592260632590979809353309870928080069505302055858695262666487217096137100671493181834671589177420642531234204009714174640931039129666175308492307316176557830592797155518896517590303408929673067747011026047132234679280646478272739419924963732822676891776959915071138967851412528239808392043203471374185013703935850775256796384432568429188702758435447130979360882290342298772254793928322894132620793740540716387131173261486578188219772109763445626361160878747405249873114743769988904687074092210449066703180940848156401703453401177717698754388382129341660449617934634706985007016661558218532120190291565284039520140134288963718230454196274306111970524102654580004605822299369603851957077921165417342759372699686089205888774732349894550064556463362462708487795741043210675330136469301509820571208337619369824048205747532487237780868387424120351377131636999407847242918098463324069882501562498942688380764685616480179312098222767578553495612397464313114063810378117541216125723617568929947972789595933724643270879842337581270272752670073059837241773030410006243989075683417654542746873324192820227793790761157483854410602778612779681404247266755174274322298105020796113643818299547265770406137004374721379198835263610195944145332172436339775352612462944623853498661607475464672546856711436807319427498270678217278152218888931455155071175858332307476379415733912122317739369813403211487906788463125371346644870222866646959907031291609372320354086226775529076780450448636303033616185733252638683393587229904005780732936675333085045501806679961195860589975745560507473279119629726886708609381304965276123381054238794311547182607760149278982390110600429988492039125015312885980683066195748361653945966532445778806358730871322629656024533199628313157197997684466319316741457914276897520486199495520643579106235798378921202558746850497277996964535514446034982179043095136685599360539282267308534336275351981368311591965318223587400140845829901671240954413780319904565107240143555388408463012801588172340481985673722204371357650242870082794659739029388792082300950238519098495212570155165239867470189909113718691860524432896548998567458928897978651533251515494444517953593328065524548994385447678923927981612032495805525688065556768984011986579777787790608084116787055780449812367582973543799505983903053748660555862987577131665651847817885530701363255343301543692552704443669681659181290553039394780984796279307805159405367781254724983406751038719066211290735600915924838623666444636322042978055573634645494869670237985186313637573106176323913315980693374486115328697547259818297806639329207176267148703238588559569020488483984299015651083527147655051202466430549105913942767726655893699658476689326167694177152379571544854749101806304838169179259217995627774518113404449739183495647951266088160481203738399072340689673293049414378265842639924094677844441781225875793849215103708975724813971846876139483276675698026104149766179881572200688013081611801005802615053565900295602079773973835927994398119059827714517092576239396499509684627211508850404592585188479606106926222459195988607863986961716180606270666451652695611765053076881418583881455989013879645746623504924159162797704180300210563496189442733911744352586551504321938717866407028700139242721425982555186456738375236034781488220687561840275378623738710354954386665899726768309859811525416344180702049233144402840906617089529411297064585361796989231049123424938755060999696975985515675159659442503694194995313165308680670671756951234475454305168080834793697693165913070280292639105387318288458773681380414830521693384475423754495681675815567882481573279416814119975855238958365392003350329301815289493890919197172008527278505423685129904600178036867805721983843555419142622527276668198739522612450812734021416572988453946625487323750465043879546439177821936468673216119047807552683923769153961379672946231041335441140122315712448121032399001776913511221502376643137059555341282512255302681198336570184545438872323751675754921805078240635475012203767172496882439513247663393302157214694933519253196609901946662608272734539353968936933403095091666554841385624555493139959971008819316677638341715690163673712887764385879589970763603888629080832475766570712580388760158626624629943646550637655873994929925598244797845350169483791545764893476328605458001136172290880756681220393832413575909211337386686201229433494377372788420275297127749406431284970696469792021977337523333910948062407426686089848058540756241818481445833597455397689893272633737406773745406874142823143889312372436458174651101060677198604108441896561831802986252408193273976638669465820339101701834118326400342126946135595515304331408850973160648940035008520892707195372421048071601099686357581429165048843101295516400641071583101195383882368011539853983271072569092643793195898607670965074448649049809417814833453955992940589274789533516277011076026869648457362628426908071588247985244174098166079431306230237319104107967951168331148755053804992116646451258520724232950778814535136683353824135745629635709138068042782295318122908402230705723741104710240705764284103215525758155483648029307444251387257134918029005615728026491268821017061113605728422275100623513213651906864073171676426371797431835954524605033748812516462701209172438179072704974436051361274225275771843152970224204368156078174165217807890394084748581074126131458212041248197901726023316960675820129904924604174656258641892793341842164735295566263211600699139050920771261383577588371428183925479822529058641168397382878035777351780312786885109003850787092176060285728742125432057845479606515156478051468822339852949643479075172310691040442714656743885471462142906480109869031192981576679460246310010262844036359285011053241199043322191375573548431730159745710350007197796142292461940070594571719001342886963741977837222196262822034738632330875177006446720331774902363858828945780953082049933863666511841644637044761181992365822428015637869083076032752089973271552979051503889095282126255830559697654446479394031629428853066241642864923069160882460324080053333947331757702929758522550903972474250614222477250899565911633438596807227436276339415629242482224259116569799104952521788175417134561624334063808818121525965737030491233785379464393888073679759122737961485257588784173286173918564366871472489728740023585659410401906339257680928624791796154338186433077083998927732403809685098979421149743341729040424731373155543095155200320388025028799011424045409138110801204440488872634516790077721029045133005455228806149855545926974180666285457130942348234641047934555776255887476957596659739722577806061665557190618998414132432757629301353096312291901273849473601482736467433729076571828399544830855900020618890291035531415572864457217393949463575378833427928191451381915929313364261624561088778574315325893328013393873078190223500665932516952844790798798337342286027458799883316157699080328558984652665576200490426301489389218335244987791204822933478173122724209341253991256396963779858207140454510660344153357892978749280164703166556384662602432800529586980439364314635086509366433853708930174778376447277507019818055358381315649867000546194532872414481939283939249054610259828178407079245832965980654691215679060083537644284916906588225060460963882395734393445636931980620165776119896672702429199134095467710833346229257728414618789143249730428897659504659820752847129566264529340495163716895226412439620396528751208379991661100308482909182985765025961528322833435564675863461434064006684909246335229220960197875877874423955697653658399014908158937028592249659954801700680842755228763922295182056772139856170403867160525732641990061303872081444488219241642733028634431932983768697488869399186635045507690311019378655215811533898108592299256817078329087005754540628623211232967412082795161552655487294257045053333061658565513482181149260470505229098943746448039951330219832820562582463536879928316931764839902628224781712845835969574363167375822407505864712364765012381459441527452029284145439120953961966995815879070750470233418257819930075695729403737255085204370561633009520765003956471588990477697233404253968330375412316261355256388688508082385745702537953273972070992506492848067330938989777725700336125455648413737474113668554865215940413528299577877594496689645536796172088692291279321372748873992702608127899853613313744480536713351645769645782122867277275707606833388412158554643584839639802565448635157061193774320330623934694370651616571452821618979375312730027270169830087117300380253803030121179339868905820689163514874230738838368075928354932849791071617862435874355742044866133742066738849435084355870277201755256078587765142099802075171396691186947224192079303883072365675633608506955414119653252950247821092055375314998137550519402422824928530914428093127117434601100896551782834085183567184056608856653480036730331104108629148761203384157351118411884824672192324107409050741539067370489973781743971063462643272771423021916057444072860452230903226294507879571089206387994256277101287217542823105014674564478191955544906423908893786526179878211556869306952561835029159838751738257123286394406704961437113279430641425630846052656503663568097366023482951861554068247827732559076582510507525847674665157596618526904002136910217579918917412256752451318312865159825507795331618730186001913662449974706936782766867399386487250034968267508838834323332380135087274522546889066185564975041734300365710764898220357163786019535837065855545125855017565690368746360664322016300631633872096285982124781977352059430451298178395495033310367016467980399528487197561187553600134242284143260159196557915131249188031598899481003401173655417171919206401903089542055595458199998167088187015488672167814959478294092832219884009609970033036343826692187135874051350558645501237426905964935739816508975227067283665577786650599594809964255658914186575776546417921261162545755966118499647282255805583571242438809854106436303762718615752019932549938835696826979817405088913561820067935145884375644396449226017661537801907417787716198921303681125894838521772515280260545587183003371332172031170600933341931325036892965155715945368272306039575574806505383201597336979861741584170535810466483748281048146772256322197406786135067643083214085478527905190742923453657293910422574159203434906454024891486405846728008618353898532827854737923026834459590176406523718385292854620550042543214693903526759748610081888534750715875679186164538434249508573145655512510602485781475867493584233680010980235324836672293520404985767104434131555909673691517549050313642070343908189090315612732585531181203470305009349390139393009312219358454673613323194470727377759823288381123950794834746746043152317250336056706312742820820634468280060720281444516327470308879951019058139242016896208495086163014988404619508578165637878281124639674780701578382741748333708964852273057803546312696387179297539395121991705667270417778056641651453616307012263291872304741307568605029098982856153889675401983328771319233748155488286316922102725747438537629551818099691391011862020093895210621544385228758706430896653140994936840335722068041307431587487455085885402461485007490759332125198120917088420414537684651581674949990832898728479935891377150776237137363393324601846582103801360655141911838900080724355966145873798929995646480867720102302929238023432265692692853800123119974104144562710920394321092715790831406768981112046919872994769419196743302418012060041022011526893927135872115614121987942060957108035430514039432169870414616572278651200605508847797748266350316280753975859462836109484196571976705645233068930619673984710686766399476588322651456958437191231563977435005899944324784168232310085753627659399288976264602480278855100016563938262563093670209708870098992890880822826303997756891187710934846739139708115379862353624288516989135601807613090874746974590672982135595152231531970543119909560083513912593656162058316521204648564520655593189716303735703443200324235560444794650396127204469964473065277635609551214888581223371736468029900323319361737380276033468985441202016098313762160946795076717717270172154618525190472795494229768001215786352684203236947372634478520586256998538066557156295754902262453588422016681016717554273959039761154934719259993311228604701786878817089405764894863249266428717149516476870075642487139623684600026344928488834569624953333671764015237481457969689971377229263874173541943556305595691941285384560982193674533615785553246628839325558001303001491509639758488305997111312198874791517212753991675078126102643458732872281475050424097334121930619737336789244282706402553026965061696823682583781641182736932429894663642900951361522634840736868754499281592500985367619698717469986296421225873123812301820316260337666292066497454535672185202629111300657868714232657921005464227218907362430474333792111525680909192224848751832863527243773491657078698757415619947648648791120515336984287493542502507467005610056763345418873745056497134890593280716569599184837594143208349459849360164805471841825440069805160666128236184843346969616224057273858390493337478116277168719579147981992850628366644801986410861643434793677484955497595618721454578915421183924106352000729258716304884474559953016788380585724333174865413456523963062687648290829557625447107931205441007065161322649056954507260948368478021792176306947397696662838145746914543454226219060286748561747965829249138884180100276280223156492882109177682407695372732329227452896573880556486830115072116949319490848415221331875988446386792223063287381885972345262799144026136688925717402045634643250982446852857787436400545348244774911893886171233649618037026932050286238291275212856885311282174490349803429898294517018466485892462561675375116352393900661499656289114371351462511241033469500857936799052360120232603318206789515597059112180307981276336255152800578419420505023079562586782268881742333563045412060700832435947787989589391079481587860562225849119280981107449436575581289460933760656707830922234801700365580712903964174389259979601697300038803776761347420971071412748976033826228688291071511290662101053204755261550516014033078054233734516131820450814435776752207954056596734378102382229725968767154537716444829064284235607865722259863864667989108312657323285810826133095367062818163030909814477867629836366804533512743838323258815952024151360611941263307450865911118553089635627175391745578720327388432456045415369175663291270851159654484849427885440662264998194038224711728311857007875930616604133588545231582391413261855886702563234642028729263296121121371556400184468648739630688217048702702725035881210884707483350975711465291697789027754342813806178725411408209645656061074178193510125415669733599029861825147625024626216138444358024177573141411161237818512529773222458008149856643930047612506758184170799774286956144169701235110566519233239716661311656151979636917474650199085543616379445936022309655167129873162697272142645131690338056439553884002466664875159571684303970220897596970815854036128775915903029641616221998325492540509056745253447497600240278599555236896638567702391153962376003679433632810423815606159058314759125595491041060404028009808962128863541928463293565915515550551887725254245893728273350867674307978288682053929131005953850884939288686905639656877050966934058594660795661246239609904689041430593334754721093749932185605583040422325139217538691185817250096543961037406899189196227562620148932853791925004500620854724027211220199908537045200538742836375701813089094802961091983029812393320875430744122061545312609265134697700941573776331344891193321348950726440228287253573796814714779427836565831663174826541336928701214140890915980409442241104197221806055702644113161010651048759833681946450285309719104721318855793266532472666671407645561958305766268427654941497838929905342043805957200656371695428098570069697714052648769105190671345947421563945617269224313172098483417367034037400542137425113283213202927890240800561130690070953278122904257346792393153351509310396852653674130181898961356707115110016372496303780310431540749975647572499583641375405384204153280672336589108804130821936756781341162212527760993913759946655433701423763799729717640688010103070450663903928968620852566797573497869648955629610925741314583244065954329981454661315841030329704037275872117082475843710394930126401216589590965467716227687525338121456604437471430988168627070436145502176645900532889183588880389086540704990517977404458803771968729631200528217786192113042906508309242203371069809966017178444269969716595395525425665848266293728391087583780859197554642032803692285974585543512663189271912067599995237330987334394628298377974777637074090509646193842897862362558822554105527610042881603866155580487980296468770088475285519884429350363631016391818482128700239917162325859905579523952121341021944994261352892240075162125280344304677084334650827924456434476636270642782169402335101985900340090373111961518392931684583570374059329704266428698605928137811185246277729925840251232747465416555420271059935184831102029774902908376309298375478470359112871466624208817930505281570494878559005035337750796131230657437387451465749419696616950203267250652525427502542637529317953042297115901233387322186747724449797804941763420498951173050579319780005464456019197008823037516660855765474777416100384260084713400025264166885583522161305294521162436317092637653457157437899950542631841908924776774656313548671767813274993117130364036112657400602724800981820459747261071941735483212603300722681619657284390355082878751899646765980354492042576407126571749250734150410247951920844723441591918207757981454828634588928662341639550360885848172008129012511302221351967722005710096804348473649589491950285626115731784846984049670310033117301848296319001943688166429719778331248881956475405698982163140285468448385122368356146233081694189681974607899181707332386490872693189094130145376614984202101207132461763214879256198729082113204734748362070967115381141814810715936656274732734968909213489752922286435336952351331687836392049318581080512171001941938085891743650669946745589289247782646817353014639390233246816767906750030700491356939729162407964812781179686737653872543570624659217109495787503545561765225888940756663158778646673839695349360953079891078665422250612371708708346285949571058394939852504213261701786075688102546218677159506278220710783899361491046177855469532991321006444283722305109971683052816028276090277453774572901281155219047977222352276991261152592525358115766871661990017791723870991620653139684866601057285507908407916561637935541336397370638168801838642280342517145368090270489701568782628447928982737559177904622744070282700322403217795661375243725630077067088312589444827635538618233290677935535750204922294466080191714927611642215348797490863948029682694730974647560389624342086411779884113146891964197909607996217084580329164687206186909029040110242890791159037952126315408912436479193661096015660237028635354224073925291383541330264142224467051089702421614983534432284268430306537296478319016469719193230687433305189109831273699575247590174727002404774665201413989271587253097211945015214136188550906747047045412618574459257376528867417433621946826019584001979169697294218026055574150997707459349227247008945623680996565545633999162323093917091981987337683926144449849309045689590384324317220655411419536067249787242316192743375646924489742685762654217540058023830922386802680449656452642643872737602489370508189022948692584363089967460446640006459663654283135634676973116276772653314189030520223157860751329763846379492363522435030376718837369109029513908493350487587917689607387804140874729990121666131730030092499419378571637143546095282365653098334199939083036904274690868642103868054165370614089325709199752466330865469993424328586775066153331460045555670628028156233446499344568677894244871054504622990880631905005679457547221648616817529907698258708124302601384869544516117374272755349155744463528505655019542421626961082887911952200827378136232918887211315352949296471318265488544021128764874637211378287645940604804997429385989328433983685269229621969883124789854610329870758923312504834614335219873081959066447168996503080589924718185736110636234708241040819259088949623593847206972894770356049744713263076076330389958983979886952137890020042037702645485260900662916027025361405845005910914870924682692260606773857138563047526946853183984433336342612810332003086131061587495297407500615029631283154533986826944627054524962023286096674239903750000760809248360953948247849013915273476890263515165677014897209319925759407886027412138666283309417342322911107973767578302144712593184863422480168530874862071684132713653868794435985015770060381831619074539185516290948899881381228578071853143070199528782323307621630460490380389354473765731988838953866724430538978250145515707137332666681125618396138859826790736549588126672493637365889025233307769219516592910980421557127639460829831913981839223302528331589893381411658553459391995527155652964528531027985033261070315282448538786861170984462075153308744160364364504143737251393398766452582379931160572069489038524992226029161214658409043547319499057496500110052504629390308665742173031893776010913637183431744104516042030421372408276499566060173353390456771925321495755247566573432418276013509170208490113230664562064089311249241176011208548858491923081528706608747050954359878631963867215938697019384196165672084371036585003853092265031087883899180229660827574456163517339676955625427000711619049223530311471528735720866441160486179727893908280144248345029280715007084334222312368056324026655415210116764532233772123746509573174694019418149015784440646903372454361716752853725991946028442951429027515305276524811697200073271405180159590906147412757022436359972447207889805909378531532705928842273911356566316200020886396526389060546675632811590590062202108845691975137732159650996024872105714325556701110261376264385285417858630300576770195825529144548853133381225186882539671753259689222989744528484544961060802253597372543602562858586139980205788382841571269789045631628256360814400632709593024378244350239751677874945301452021511684129245541074394346348017005819240045736794186406548653655047181878506210122249330373800411731580853978471330343870098050120718685556697317497354173913492205919902680830336326578533930986809230611735629868115947124064840614968352530515707988179105899165390727317884568390430701811115378595491225088114016381911983332310866436135047595308339901186634846596445509551905500594084503751152330734882711329611027275075656276650760147694074621456564641422040086733473610662910322377072991516046725643356671635434331980462151069736008589278571617727853479197797400851521814321014926049517697732087391523710071471539496532007321165321589137767026492118999105306307943367521653780396849078886363297737049287238800484995485369594332809267353448584747236690391721505746740984960453131493661361394010867924701049617040414772139254305096400885900325389867419275142655435005620029110571275080714144603435745568915626657190889642503096718823173773094816381573065455195415373830301497591399674528266720122140434826232718098683093541856597260649678795606874103072640291195044499765230581103357644980449460291038980252460307581149649376181346808736353135196618591434390882577741795884381659253572042051544687597168434786697393766539538537897425613981915416306385636657832811390820684513935230706389052243540666781949466494992208967452831339939636750912155634977622219647569275899391848869642742227370492280823645219973095563806341124477462624696702043104218958374942896009057395329186232009294554932120827012925525455437178200567874188561537098316103756403711314034464960093144705149930818753520620009463665797454176108391119870880771817816015071656385828911170107294638270451561603099681454972697075550288702114060243770814120365091445587355451898097681482300757747530778251897538230332556393922884938791570078773342173490178176118442447932443396278405281812764261885975154951079407032270818160391867919914514847930646406873857622459495973444074545074269402353922209089430317559014914174638145529991354159143765333754730233295179144089072382273905532348285075744348805188246497034789308796800937509967765542315905717585103796163184741412969340167416386957941114906255744015407072022864935728021461522052197109450341526309110417367662164133233379085236981614341374173889787429659973524300058370840327894471536330686870427626459586863348664811081566750745206970105669433074414666954021943745476139834700812393339681832873003536377295763707491625734812015286541137709718223574978107723141179286450807506522534098138048224081788092506127034467389807626466667469852655963712914474420929368841882009698623368017953445378549898817525848377220916668608899777258318020172530749861752204954875951105577500684971923209453366806419176631849578284614685673476128492742246344271434050268582330191509654291382301721518288594078166396303084722703293588578953459240898565319590435980726358347600342707123949654321040905733680820781015405425586227807091434080585327226040650324495511444407830841384586362137389188960784051621622538934137113396142579265144103902712108436238478162293249045939990611112723269291729977267475563922912210428604685314529896270098142151937576374137328147441401997970407522657417953944807364407594748604361321816894544090670079397376452382015603493816411619551420709366192228166812636507448466413296171938314285167945679747018177651398997480279847925039452289986033324396180741414564511351132876890278231851416617736889254930838155526374808370519320345870883131035587267829531760633533607427218976199621217384636959422816814729067547090527066091899063862792590250206610260742633409490649915752462159968964966946148771085587309647759583237054596701645413216396228583450966876408886586847976157182940396398356701755419467102149489603442245634573187243103869596931032115692033648372808373067799139385134115794618014794273342630300208487732167359060812729100574941554496718903135470408577904067397042005717535132541601235253289338978693042209047558505184048560273247200524525954380031766331260309210402442687358012003133331853067900340383351423237251108708828162351168720425898042070969185531111489927544981331461169951497456970114302935647016958157291153633601743599092529916252403699817346200659172702187378926252719301651336078019315115262695198003652938335207026026689986757124610126889186820156032549632120747054550198638011172545303402734648795284859783781497086450648917445069382779217452114120351149468818484973222425284518800115631444330260317564302841456939579934696340193375839423700409519947843050301152360629240004925667482439236160250470068226518725483768957945368621731325585404460423493992318465624869023685656001313936223534136780229671521863580408795519193743254050912140028593860213909641113947222477339527708342881823288940747372852107768414817028163056407132402640266090100496723934042312697828590429982051621485570922009109160940167121417335633249544546322894871002985662285806688190565423095002700790964145028601466870080616876220555669090527079783919450450330196955611564105364215162266456481080502445772631536648728780176538669411556846494142230202517088143339935465120166250495962069159007171387424030932595873075561873173885287689942011768378729646802181738316530418947812140936824695116226720386362661226031266981574594024376558439451900514088068233599598965741366293433766491033656162595466719317002902029318310866589481126381075360194743943459295890045648274300065742202095593744585849076538944499539089991045189463493757393672125199202123320155411938334173265634292837709780221453416313763996090676668656838482617860646133418142957581230100700872765767836596056279562222274686161048565210300907380802636790493471401976098091869538384171829811705296944904430676536725764908630904856273520519580738288510223043271189389024622612287335320784689283280440364895730984377452703387756321256028470514422959568599413862302269961172242095875563792955294086001932385166000220309707396477793092643205454263709192441135846095304377136624947649547056497693864818713202462728976032034452318283189100527668039113709858863396673872212454034967113493705197754117034635062379286674882217821483991129840397654110807020893008798925745164283670110996897711375021620138102891175734772561163182381763738526902773160578788954366980519037451852624320453573197868603529588311291154820669460521954911363838414859460819927644195729312064563594826954631558402419984962664690166607844694643951225244904128288427405283958140676584665974420758480358764097101094638311071033976411421032633377978464970552276876149304390224397712908718115962498337854763537337041424595338110745716125546190401425692333739033185912083060909825683462853511256012049060828772354093845797080667047381791957224487466679569426640595729278383071827604273920781417585884965781195747464591953109557172562679554189084451808807742596248945450159190841861396597947174442815820241187174595619776799322568661322148817172915221568731129105138099106661326066381638029078570993677641994753262597497503311693339173636625549784681787125106095848766983079303882089591139755108505853330378843099699739483749208340555860472392887564146780941182320541145468277398237458270951283529333475884757417966043443045104601775213420363268870063648065519256018776753196156743625385229372170066533622184454019905913928435079196463126999208047061857482161640345432525389224293314933957917802545453386017545917695788233677443116800617650956227142084505740146615555708869516114764358621928247205531632875021020125103764651138873407018768679311090560897923593534670358096536084463418834586945384933874887796802930698681958364321425822286621743609458585785968656112354995850470917703695886826990972196013065838494388768518963154570550984676499254799793460316407337591997304570623096449781567592983051599191713720769495528457915926642280595080117954445815527377382221714966716669165488735799843644682184142359338802550913277911664535102354510286723822211250445464990169739506906115915356803532834048881501000744566430043546070993994692943013773187386132932104390711557542054064548604514284607351919016381843444416643313069397068919866491354773222345204925237038901436463023602315802803610314178811597428124914919119736479608883679515507723766201474754871260233775443055787173514411284548418981913420388863580479580241109378905187438628652098131624091316432267383078491590319805106727379401350407412803047142090169063512613920806511403849867745335846521901247512891844469861283787651413486868610995162457214052249970297637742408112548598068903711496107889918242325131837064751665085339319931917901488924292503273830763353534339856142200294887742549795752938994601219858670779258091875228170394132261912300860984341806456729122753771635520976181473694842081650957692854943781215297552623457352800576323192604333140870706269833809505947490735748040860353955999327118958469044522217534687536860079167780651827798049896709184672018358908777045182582194405133233174527002719574295221715218104447000355466264145715593294545416779132191945573906369484244081204288776313162428114773790361703495345143854243967624640487564155096086322751487294738034934866266255021011496237480246668746608926710661556131581296097672298361222015661386393432232552236224105055992897538809271471500070555433579491379187067399633702431869060169378317937961262640774376916737190458013613821418942186135410522363852200290152251963967320521909527032810323612105910807677385639306954696617228512904079232186396076941266611766571596351325789663164813429715040394618564607845020528366779753679634975395040730411888884673298388430449225041091945857122218880400366793244036739630143863200506054793716068417460254102118175221818854383360129416175486349570191497744873810496346861293270644646178688762192053873004379568597300011175331334275011644418381191256399490700770030535038867964031815667328160462185257096688618082756120515614171161499280696389203259875578493870985118628107700681615210145880020814443512284339917427442169965731890676356775874864753885803213593435183275879295866811801595440322201708211861669752345481569176631699642762911059568718139425419562702900169612407242892094746225981552738553763924442457346875914432194556421698266643837404766067270958122447322726983656811158919101817483532877661943041601312315917829041511177135532563394843365281005322481636627461794934231416637800205740048277012866543702522701804274555874968316255594014675566424490346799655856833220095158058554832992111244378861402855077010588833985337871332581054407784945839206561474056977774485775694714143476360636390623852050916197660653346387715602946763117077239494031809154234314105763721869408797529528503124768626487882804949886619527575243080750332878080388687705534958569905311850508499615280208584271790435172392921225264680429245300833521394139272095160035170431496844830578005242146936737400802624851975938567386299202833586392710714197789704323667477185180888822688354102163598285412644144218252378436717173820965231211332752419552141604556731263315414146069758127701751130758657557682142079615263636581650430343313639827743435026071643280244602491537692982178288286941339726651286586899310064828400125362684772585569700689552239188272151623961083592971360468839153493958836686129200556069280103448711541576229087641339940378738264531928355180393746593185129446182746464124194982240920651774664456797241789690474557550098983595945624994050329237271719615624944968045664222469681940699887464326589035274940713414736552410289088334372960661824427263878155000891868069057790038413259531856854975647508602930224963343655199131794024644527636976805944159218541329495966384566258420223938828450819530573415182798906126856851999916205669706719682516516720397674922642501493226682184182670123167826898821311457214579392855073202129825356706318552711924320079613759974614577981502088477155333689696992142088946890545474155279376740953084675314856836871716356093652532572661581792310488714086403778451479727173987065132092679054532210415557362702796089495731629862071579532687380181749306395480631029941936341067656361547034371537659534706071826293844728612343576578613571782511280751326424492442096852157100146299694731189163404163904959911721784238979159191979929678409809710223091429420804517671127182309976094771022763155349812580570538031428231310626819216341111395713911085937633440811719611236527194011296921369121780925248764968579432632208035576609039336824794816805403298939259168606072524363971252867471817430171458829296330880687039296490939494275020002810992391083334873994685669635698301816697210321119314781904791803484900655521684897880987727148336976005671716586701860139573074356390153249917230498427764829449639863866037364941250672701747782342847511157647873888049073543206262742809426947999614383741822894901650787155662884935742945713661071391798298798227592838081178839540452934807378001800047769349050445721775765774647426217417646144344425672141752606281987480282829367080778770221419176008101496858674352656993462914400134654815349284538953305408100802161100434227470355317124401685295927923163954094789535592105449486064687251677372480266065861113658918755210048450207344916339043799535770722187626979166146756682066045250623334940941167515156504752147658954801555141771222184602726077359791427567956634212593223580972437891201565811000822797129865808865275324554879231368131275584185609420152054332273588208485600178142355325718273687206284306667710783381779755922926468821894900037314911123316641837772024296573187375408481165850664330220289433841835966445965942162213596642691960607357255067445445501033908809215511666112694388740887871017173632824703768275986719074077976356516470233502200559195290952210853183267643716527511627104922315050867652572043812616250182059581702891037212814920102239333611227357651186276379150673920680082014814259315417463213534787905670646512698208669821624707421720138997946148398548725110228081355882372762921119023616947735756059368894437424445248004153469058785265735620772349231137532854406367470905183065659680825784644875933271868518921119275689676301589640270906655865675362884522280052739870680194871827368595567412473579929749996735937934241057634316159800859435873214689908355167611770051320718150876673150469295532928666224223546231342868131373734926466118683030256466387722238035263580189766795317849800586034531873077479117047785935093238621142764664288120472045981262250762754551400119546043691248558995220054311954776982626670013756583807212030750931494460469813228407223235875683132121612573482375907407842286306608806954725886973013042452380797730220521165782096488927915309029566497211174750117424926892016747201896372809215145133448018760388598890742338506490021468573272636889957172086409151885265365626972205604952212149840532483781169450314075232499386029233943989419187432561794218826496877109273254597963236258281854816886684389335757483177626694737079518943920781885348018560630442020759995248521709107828661970735906712465820255475170515495501433083773513336093850749065680831534336354389742124869398405602308913985033056854898644583312535332524326499137387203650107299993809500432559730753862605349934298300416"))
		})

		It("340282366920938463463374607431768211455 mul 340282366920938463463374607431768211456", func() {
			bigint1, _ := biginteger.Of("340282366920938463463374607431768211455")
			bigint2, _ := biginteger.Of("340282366920938463463374607431768211456")
			result := bigint1.Multiply(*bigint2)

			Expect(result.String()).To(
				Equal("115792089237316195423570985008687907852929702298719625575994209400481361428480"))
		})

		It("[1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1] * ([0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1] - 1)", func() {
			a := biginteger.OfUint64Array([]uint64{
				1, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
				1,
			})

			b := biginteger.OfUint64Array([]uint64{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
				1,
			}).Subtract(biginteger.One())

			result := a.Multiply(b)

			expected := biginteger.OfUint64Array([]uint64{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
				1,
			}).Subtract(biginteger.One())

			Expect(result.IsEqualTo(expected)).To(BeTrue())
		})

		It("[1, 0, 0, 0, 0, 0, 0, 0, 1] * ([0, 0, 0, 0, 0, 0, 0, 0, 1] - 1)", func() {
			a := biginteger.OfUint64Array([]uint64{
				1, 0, 0, 0, 0, 0, 0, 0,
				1,
			})

			b := biginteger.OfUint64Array([]uint64{
				0, 0, 0, 0, 0, 0, 0, 0,
				1,
			}).Subtract(biginteger.One())

			result := a.Multiply(b)

			expected := biginteger.OfUint64Array([]uint64{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
				1,
			}).Subtract(biginteger.One())

			Expect(result.IsEqualTo(expected)).To(BeTrue())
		})

		It("[1, 0, 0, 0, 1] * ([0, 0, 0, 0, 1] - 1)", func() {
			a := biginteger.OfUint64Array([]uint64{
				1, 0, 0, 0, 1,
			})

			b := biginteger.OfUint64Array([]uint64{
				0, 0, 0, 0, 1,
			}).Subtract(biginteger.One())

			result := a.Multiply(b)

			expected := biginteger.OfUint64Array([]uint64{
				0, 0, 0, 0, 0, 0, 0, 0,
				1,
			}).Subtract(biginteger.One())

			Expect(result.IsEqualTo(expected)).To(BeTrue())
		})

		It("[1, 0, 1] * ([0, 0, 1] - 1)", func() {
			a := biginteger.OfUint64Array([]uint64{
				1, 0, 1,
			})

			b := biginteger.OfUint64Array([]uint64{
				0, 0, 1,
			}).Subtract(biginteger.One())

			result := a.Multiply(b)

			expected := biginteger.OfUint64Array([]uint64{
				0, 0, 0, 0,
				1,
			}).Subtract(biginteger.One())

			Expect(result.IsEqualTo(expected)).To(BeTrue())
		})

		It("[18446744073709551608, 3] * [18446744073709551593, 5]", func() {
			a := biginteger.OfUint64Array([]uint64{
				18446744073709551608, 3,
			})

			b := biginteger.OfUint64Array([]uint64{
				18446744073709551593, 5,
			})

			result := a.Multiply(b)

			Expect(result.String()).To(Equal("8166776806102523120538446408043099848888"))

			expected := biginteger.OfUint64Array([]uint64{
				184, 18446744073709551476, 23,
			})

			Expect(result.IsEqualTo(expected)).To(BeTrue())
		})

		It("[18446744073709551608, 3] * [18446744073709551593, 5]", func() {
			a := biginteger.OfUint64Array([]uint64{
				18446744073709551608, 3,
			})

			b := biginteger.OfUint64Array([]uint64{
				18446744073709551593, 5,
			})

			result := a.Multiply(b)

			Expect(result.String()).To(Equal("8166776806102523120538446408043099848888"))

			expected := biginteger.OfUint64Array([]uint64{
				184, 18446744073709551476, 23,
			})

			Expect(result.IsEqualTo(expected)).To(BeTrue())
		})

		It("[3689348814741910309 1] * 18446744073709551593", func() {
			a := biginteger.OfUint64Array([]uint64{3689348814741910309, 1})
			b := biginteger.OfUint64Array([]uint64{18446744073709551593})

			Expect(a.String()).To(Equal("22136092888451461925"))

			result := a.Multiply(b)

			Expect(result.String()).To(Equal("408338840305126155384975626637062596525"))

			expectedResult := biginteger.OfUint64Array([]uint64{
				7378697629483820973, 3689348814741910281, 1,
			})

			Expect(result.IsEqualTo(expectedResult)).To(BeTrue())
		})

		It("[1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1] * ([0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1] - 1)", func() {
			a := biginteger.OfUint64Array([]uint64{
				1, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
				1,
			})

			b := biginteger.OfUint64Array([]uint64{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
				1,
			}).Subtract(biginteger.One())

			result := a.Multiply(b)

			expected := biginteger.OfUint64Array([]uint64{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
				1,
			}).Subtract(biginteger.One())

			Expect(result.IsEqualTo(expected)).To(BeTrue())
		})
	})

	Context("Divide", func() {
		It("Should divide 2 by 2", func() {
			bigint1, _ := biginteger.Of("2")
			bigint2, _ := biginteger.Of("2")

			Expect(bigint1.Divide(*bigint2).String()).To(Equal("1"))
		})

		It("Should divide 1 by 2", func() {
			bigint1, _ := biginteger.Of("1")
			bigint2, _ := biginteger.Of("2")

			Expect(bigint1.Divide(*bigint2).String()).To(Equal("0"))
		})

		It("Should divide 4 by 2", func() {
			bigint1, _ := biginteger.Of("4")
			bigint2, _ := biginteger.Of("2")
			divide := bigint1.Divide(*bigint2)

			Expect(divide.String()).To(Equal("2"))
		})

		It("Should divide -4 by 2", func() {
			bigint1, _ := biginteger.Of("-4")
			bigint2, _ := biginteger.Of("2")

			Expect(bigint1.Divide(*bigint2).String()).To(Equal("-2"))
		})

		It("Should divide -4 by -2", func() {
			bigint1, _ := biginteger.Of("-4")
			bigint2, _ := biginteger.Of("-2")

			Expect(bigint1.Divide(*bigint2).String()).To(Equal("2"))
		})

		It("Should divide 4 by -2", func() {
			bigint1, _ := biginteger.Of("4")
			bigint2, _ := biginteger.Of("-2")

			Expect(bigint1.Divide(*bigint2).String()).To(Equal("-2"))
		})

		It("Should divide 8 by 4", func() {
			bigint1, _ := biginteger.Of("8")
			bigint2, _ := biginteger.Of("4")

			Expect(bigint1.Divide(*bigint2).String()).To(Equal("2"))
		})

		It("Should divide 80 by 10", func() {
			bigint1, _ := biginteger.Of("80")
			bigint2, _ := biginteger.Of("10")

			Expect(bigint1.Divide(*bigint2).String()).To(Equal("8"))
		})

		It("Should divide 4294967296 by 10", func() {
			bigint1, _ := biginteger.Of("4294967296")
			bigint2, _ := biginteger.Of("10")
			result := bigint1.Divide(*bigint2)

			Expect(result.String()).To(Equal("429496729"))
		})

		It("Should divide 4294967296 by 4294967296", func() {
			bigint1, _ := biginteger.Of("4294967296")
			bigint2, _ := biginteger.Of("4294967296")

			Expect(bigint1.Divide(*bigint2).String()).To(Equal("1"))
		})

		It("Should divide 18446744073709551615 by 1", func() {
			bigint1, _ := biginteger.Of("18446744073709551615")
			bigint2, _ := biginteger.Of("1")
			result := bigint1.Divide(*bigint2)

			Expect(result.String()).To(Equal("18446744073709551615"))
		})

		It("Should divide 18446744073709551615 by 2", func() {
			bigint1, _ := biginteger.Of("18446744073709551615")
			bigint2, _ := biginteger.Of("2")
			divide := bigint1.Divide(*bigint2)

			Expect(divide.String()).To(Equal("9223372036854775807"))
		})

		It("Should divide 18446744073709551615 by 10", func() {
			bigint1, _ := biginteger.Of("18446744073709551615")
			bigint2, _ := biginteger.Of("10")
			divide := bigint1.Divide(*bigint2)

			Expect(divide.String()).To(Equal("1844674407370955161"))
		})

		It("Should divide 18446744073709551615 by 18446744073709551616", func() {
			bigint1, _ := biginteger.Of("18446744073709551615")
			bigint2, _ := biginteger.Of("18446744073709551616")

			Expect(bigint1.Divide(*bigint2).String()).To(Equal("0"))
		})

		It("Should divide 18446744073709551616 by 18446744073709551615", func() {
			bigint1, _ := biginteger.Of("18446744073709551616")
			bigint2, _ := biginteger.Of("18446744073709551615")

			divide := bigint1.Divide(*bigint2)

			Expect(divide.String()).To(Equal("1"))
		})

		It("Should divide 18446744073709551616 by 2", func() {
			bigint1, _ := biginteger.Of("18446744073709551616")
			bigint2, _ := biginteger.Of("2")
			divide := bigint1.Divide(*bigint2)

			Expect(divide.String()).To(Equal("9223372036854775808"))
		})

		It("Should divide 18446744073709551616 by 18446744073709551616", func() {
			bigint1, _ := biginteger.Of("18446744073709551616")
			bigint2, _ := biginteger.Of("18446744073709551616")

			Expect(bigint1.Divide(*bigint2).String()).To(Equal("1"))
		})

		It("Should divide 18446744073709551616 by 10", func() {
			bigint1, _ := biginteger.Of("18446744073709551616")
			bigint2, _ := biginteger.Of("10")
			result := bigint1.Divide(*bigint2)

			Expect(result.String()).To(Equal("1844674407370955161"))
		})

		It("Should divide 36893488147419103230 by 10", func() {
			bigint1, _ := biginteger.Of("36893488147419103230")
			bigint2, _ := biginteger.Of("10")
			result := bigint1.Divide(*bigint2)

			Expect(result.String()).To(Equal("3689348814741910323"))
		})

		It("Should divide 73786976294838206464 by 10", func() {
			bigint1, _ := biginteger.Of("73786976294838206464")
			bigint2, _ := biginteger.Of("10")
			result := bigint1.Divide(*bigint2)

			Expect(result.String()).To(Equal("7378697629483820646"))
		})

		It("Should divide 340282366920938463444927863358058659840 by 10", func() {
			bigint1, _ := biginteger.Of("340282366920938463444927863358058659840")
			bigint2, _ := biginteger.Of("10")
			result := bigint1.Divide(*bigint2)

			Expect(result.String()).To(Equal("34028236692093846344492786335805865984"))
		})

		It("Should divide 115792089237316195423570985008687907853269984665640564039457584007913129639936 by 10", func() {
			bigint1, _ := biginteger.Of("115792089237316195423570985008687907853269984665640564039457584007913129639936")
			bigint2, _ := biginteger.Of("10")
			result := bigint1.Divide(*bigint2)

			Expect(result.String()).To(Equal("11579208923731619542357098500868790785326998466564056403945758400791312963993"))
		})

		It("Should divide 10 by 115792089237316195423570985008687907853269984665640564039457584007913129639936", func() {
			bigint1, _ := biginteger.Of("10")
			bigint2, _ := biginteger.Of("115792089237316195423570985008687907853269984665640564039457584007913129639936")
			result := bigint1.Divide(*bigint2)

			Expect(result.String()).To(Equal("0"))
		})

		It("Should divide 7234724982749223422342342348234982734729384727349827498233423423474924 by 23049828420348234290384203840283402840923842320234242342094823", func() {
			bigint1, _ := biginteger.Of("7234724982749223422342342348234982734729384727349827498233423423474924")
			bigint2, _ := biginteger.Of("23049828420348234290384203840283402840923842320234242342094823")
			result := bigint1.Divide(*bigint2)

			Expect(result.String()).To(Equal("313873268"))
		})

		It("Should divide by BZ algorithm 1", func() {
			a := biginteger.Two().
				Power(biginteger.OfUint64(64 * 64)).
				Subtract(biginteger.One())
			b := biginteger.Two().Power(biginteger.OfUint64(64 * 32)).
				Subtract(biginteger.One())

			result := a.Divide(b)

			expected := biginteger.OfUint64Array([]uint64{
				1,
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 1,
			})

			Expect(result.IsEqualTo(expected)).To(BeTrue())

			Expect(result.Digits()).To(Equal(uint64(617)))
			Expect(result.String()).To(Equal("32317006071311007300714876688669951960444102669715484032130345427524655138867890893197201411522913463688717960921898019494119559150490921095088152386448283120630877367300996091750197750389652106796057638384067568276792218642619756161838094338476170470581645852036305042887575891541065808607552399123930385521914333389668342420684974786564569494856176035326322058077805659331026192708460314150258592864177116725943603718461857357598351152301645904403697613233287231227125684710820209725157101726931323469678542580656697935045997268352998638215525166389437335543602135433229604645318478604952148193555853611059596230657"))

			modulo1Mio := result.Modulo(biginteger.OfUint64(1_000_000))
			Expect(modulo1Mio.String()).To(Equal("230657"))
		})

		It("[18446744073709551615,18446744073709551615] / [18446744073709551615]", func() {
			a := biginteger.OfUint64Array([]uint64{18446744073709551615, 18446744073709551615})
			b := biginteger.OfUint64(18446744073709551615)

			result := a.Divide(b)

			Expect(result.Digits()).To(Equal(uint64(20)))
			Expect(result.String()).To(Equal("18446744073709551617"))

			modulo1Mio := result.Modulo(biginteger.OfUint64(1_000_000))
			Expect(modulo1Mio.String()).To(Equal("551617"))
		})

		It("Should divide by BZ algorithm 1a", func() {
			a := biginteger.Two().Power(biginteger.OfUint64(64).Multiply(biginteger.OfUint64(32))).Subtract(biginteger.One())
			b := biginteger.Two().Power(biginteger.OfUint64(64).Multiply(biginteger.OfUint64(16))).Subtract(biginteger.One())

			result := a.Divide(b)

			Expect(result.Digits()).To(Equal(uint64(309)))
			Expect(result.String()).To(Equal("179769313486231590772930519078902473361797697894230657273430081157732675805500963132708477322407536021120113879871393357658789768814416622492847430639474124377767893424865485276302219601246094119453082952085005768838150682342462881473913110540827237163350510684586298239947245938479716304835356329624224137217"))

			modulo1Mio := result.Modulo(biginteger.OfUint64(1_000_000))
			Expect(modulo1Mio.String()).To(Equal("137217"))
		})

		It("Should divide by BZ algorithm 1b", func() {
			a := biginteger.Two().Power(biginteger.OfUint64(64).Multiply(biginteger.OfUint64(16))).Subtract(biginteger.One())
			b := biginteger.Two().Power(biginteger.OfUint64(64).Multiply(biginteger.OfUint64(8))).Subtract(biginteger.One())

			result := a.Divide(b)

			Expect(result.Digits()).To(Equal(uint64(155)))
			Expect(result.String()).To(Equal("13407807929942597099574024998205846127479365820592393377723561443721764030073546976801874298166903427690031858186486050853753882811946569946433649006084097"))

			modulo1Mio := result.Modulo(biginteger.OfUint64(1_000_000))
			Expect(modulo1Mio.String()).To(Equal("84097"))
		})

		It("Should divide by BZ algorithm 1c", func() {
			a := biginteger.Two().Power(biginteger.OfUint64(64).Multiply(biginteger.OfUint64(8))).Subtract(biginteger.One())
			b := biginteger.Two().Power(biginteger.OfUint64(64).Multiply(biginteger.OfUint64(4))).Subtract(biginteger.One())

			result := a.Divide(b)

			Expect(result.Digits()).To(Equal(uint64(78)))
			Expect(result.String()).To(Equal("115792089237316195423570985008687907853269984665640564039457584007913129639937"))

			modulo1Mio := result.Modulo(biginteger.OfUint64(1_000_000))
			Expect(modulo1Mio.String()).To(Equal("639937"))
		})

		It("Should divide by BZ algorithm 1d", func() {
			a := biginteger.Two().Power(biginteger.OfUint64(64).Multiply(biginteger.OfUint64(4))).Subtract(biginteger.One())
			b := biginteger.Two().Power(biginteger.OfUint64(64).Multiply(biginteger.OfUint64(2))).Subtract(biginteger.One())

			result := a.Divide(b)

			Expect(result.Digits()).To(Equal(uint64(39)))
			Expect(result.String()).To(Equal("340282366920938463463374607431768211457"))

			modulo1Mio := result.Modulo(biginteger.OfUint64(1_000_000))
			Expect(modulo1Mio.String()).To(Equal("211457"))
		})

		It("Should divide by BZ algorithm 1e", func() {
			a := biginteger.Two().Power(biginteger.OfUint64(64).Multiply(biginteger.OfUint64(2))).Subtract(biginteger.One())
			b := biginteger.Two().Power(biginteger.OfUint64(64).Multiply(biginteger.OfUint64(1))).Subtract(biginteger.One())

			result := a.Divide(b)

			Expect(result.Digits()).To(Equal(uint64(20)))
			Expect(result.String()).To(Equal("18446744073709551617"))

			modulo1Mio := result.Modulo(biginteger.OfUint64(1_000_000))
			Expect(modulo1Mio.String()).To(Equal("551617"))
		})

		It("Should divide by BZ algorithm 2", func() {
			a := biginteger.Two().Power(biginteger.OfUint64(64).Multiply(biginteger.OfUint64(64)))
			b := biginteger.Two().Power(biginteger.OfUint64(64).Multiply(biginteger.OfUint64(32)))

			result := a.Divide(b)

			Expect(result.IsEqualTo(b)).To(BeTrue())
		})

		It("Should divide by BZ algorithm 3", func() {
			a := biginteger.Two().Power(biginteger.OfUint64(64).Multiply(biginteger.OfUint64(63)))
			b := biginteger.Two().Power(biginteger.OfUint64(64).Multiply(biginteger.OfUint64(31)))

			result := a.Divide(b)

			expected := biginteger.Two().Power(biginteger.OfUint64(64).Multiply(biginteger.OfUint64(32)))

			Expect(result.IsEqualTo(expected)).To(BeTrue())
		})

		It("Should divide by (not really) BZ algorithm 4", func() {
			a := biginteger.Two().Power(biginteger.OfUint64(64).Multiply(biginteger.OfUint64(3)))
			b := biginteger.Two().Power(biginteger.OfUint64(64).Multiply(biginteger.OfUint64(1)))

			result := a.Divide(b)

			expected := biginteger.Two().Power(biginteger.OfUint64(64).Multiply(biginteger.OfUint64(2)))

			Expect(result.IsEqualTo(expected)).To(BeTrue())
		})

		It("Should divide by (not really) BZ algorithm 5", func() {
			a := biginteger.Two().Power(biginteger.OfUint64(64)).Subtract(biginteger.One())
			b := biginteger.Two().Power(biginteger.OfUint64(32)).Subtract(biginteger.One())

			result := a.Divide(b)

			expected := biginteger.OfUint64(4294967297)

			Expect(result.IsEqualTo(expected)).To(BeTrue())
		})

		It("Should divide by (not really) BZ algorithm 6", func() {
			a := biginteger.Two().Power(biginteger.OfUint64(64 * 2)).Subtract(biginteger.One())
			b := biginteger.Two().Power(biginteger.OfUint64(64)).Subtract(biginteger.One())

			result := a.Divide(b)

			expected, _ := biginteger.Of("18446744073709551617")

			Expect(result.IsEqualTo(*expected)).To(BeTrue())
		})

		It("Should divide by (not really) BZ algorithm 7", func() {
			a := biginteger.Two().Power(biginteger.OfUint64(64 * 3)).Subtract(biginteger.One())
			b := biginteger.Two().Power(biginteger.OfUint64(64)).Subtract(biginteger.One())

			result := a.Divide(b)

			expected, _ := biginteger.Of("340282366920938463481821351505477763073")

			Expect(result.IsEqualTo(*expected)).To(BeTrue())
		})

		It("Should divide by (not really) BZ algorithm 8", func() {
			a := biginteger.Two().Power(biginteger.OfUint64(64 * 3)).Subtract(biginteger.One())
			b := biginteger.Two().Power(biginteger.OfUint64(64 * 2)).Subtract(biginteger.One())

			result := a.Divide(b)

			expected, _ := biginteger.Of("18446744073709551616")

			Expect(result.IsEqualTo(*expected)).To(BeTrue())
		})

		It("Should divide by (not really) BZ algorithm 9", func() {
			a := biginteger.Two().Power(biginteger.OfUint64(64 * 4)).Subtract(biginteger.One())
			b := biginteger.Two().Power(biginteger.OfUint64(64 * 1)).Subtract(biginteger.One())

			result := a.Divide(b)

			expected, _ := biginteger.Of("6277101735386680764176071790128604879584176795969512275969")

			Expect(result.IsEqualTo(*expected)).To(BeTrue())
		})

		It("Should divide by (not really) BZ algorithm 10", func() {
			a := biginteger.Two().Power(biginteger.OfUint64(64 * 4)).Subtract(biginteger.One())
			b := biginteger.Two().Power(biginteger.OfUint64(64 * 2)).Subtract(biginteger.One())

			result := a.Divide(b)

			expected, _ := biginteger.Of("340282366920938463463374607431768211457")

			Expect(result.IsEqualTo(*expected)).To(BeTrue())
		})

		It("[2249300428400506475 327839218333391898 12018134535777490305] / [11595963786453332908 3312835375285772574]", func() {
			a := biginteger.OfUint64Array([]uint64{
				2249300428400506475, 327839218333391898, 12018134535777490305,
			})

			Expect(a.String()).To(Equal("4089559265808638403608548000220092746857740622404617647723"))

			b := biginteger.OfUint64Array([]uint64{
				3312835375285772574, 11595963786453332908,
			})

			Expect(b.String()).To(Equal("213907776256708591358813130932743151902"))

			quotient, remainder := a.DivMod(b)

			expectedQuotient := biginteger.OfUint64Array([]uint64{
				671585036966626273, 1,
			})

			expectedRemainder := biginteger.OfUint64Array([]uint64{
				14127552491733795085, 3052621451255280110,
			})

			Expect(quotient.String()).To(Equal("19118329110676177889"))
			Expect(remainder.String()).To(Equal("56310926665221989276999289752316952845"))

			Expect(quotient.IsEqualTo(expectedQuotient)).To(BeTrue())
			Expect(remainder.IsEqualTo(expectedRemainder)).To(BeTrue())
		})

		It("[9350441601238114644 12018134535777490305 327839218333391898 2249300428400506475] / [11595963786453332908 3312835375285772574]", func() {
			a := biginteger.OfUint64Array([]uint64{
				9350441601238114644, 12018134535777490305, 327839218333391898, 2249300428400506475,
			})

			b := biginteger.OfUint64Array([]uint64{
				11595963786453332908, 3312835375285772574,
			})

			quotient, remainder := a.DivMod(b)

			Expect(a.String()).To(Equal("14119087622518823676962868251197053621424793707571764183901859319376275382612"))
			Expect(b.String()).To(Equal("61111026326228183616332404080143512492"))
			expectedQuotient := biginteger.OfUint64Array([]uint64{
				6742305324661190591, 12524700037052152845,
			})

			expectedRemainder := biginteger.OfUint64Array([]uint64{
				0,
			})

			Expect(quotient.IsEqualTo(expectedQuotient)).To(BeTrue())
			Expect(remainder.IsEqualTo(expectedRemainder)).To(BeTrue())
			Expect(quotient.String()).To(Equal("231039936183481602043341757937109938111"))
			Expect(remainder.String()).To(Equal("0"))
		})

		It("[184 18446744073709551476 23] / [5 18446744073709551593]", func() {
			a := biginteger.OfUint64Array([]uint64{184, 18446744073709551476, 23})
			b := biginteger.OfUint64Array([]uint64{5, 18446744073709551593})

			Expect(a.String()).To(Equal("8166776806102523120538446408043099848888"))
			Expect(b.String()).To(Equal("340282366920938463039099493736448524293"))

			quotient, remainder := a.DivMod(b)

			Expect(quotient.String()).To(Equal("24"))
			Expect(remainder.String()).To(Equal("7600058558368335265856"))

			expectedQuotient := biginteger.OfUint64Array([]uint64{
				24,
			})

			expectedRemainder := biginteger.OfUint64Array([]uint64{
				64, 412,
			})

			Expect(quotient.IsEqualTo(expectedQuotient)).To(BeTrue())
			Expect(remainder.IsEqualTo(expectedRemainder)).To(BeTrue())
		})
	})

	Context("Modulo", func() {
		It("Should return 0 for 1 mod 1", func() {
			bigint1, _ := biginteger.Of("1")
			bigint2, _ := biginteger.Of("1")

			Expect(bigint1.Modulo(*bigint2).String()).To(Equal("0"))
		})

		It("Should return 0 for 2 mod 1", func() {
			bigint1, _ := biginteger.Of("2")
			bigint2, _ := biginteger.Of("1")

			Expect(bigint1.Modulo(*bigint2).String()).To(Equal("0"))
		})

		It("Should return 1 for 1 mod 2", func() {
			bigint1, _ := biginteger.Of("1")
			bigint2, _ := biginteger.Of("2")

			Expect(bigint1.Modulo(*bigint2).String()).To(Equal("1"))
		})

		It("Should return 0 for 0 mod 2", func() {
			bigint1, _ := biginteger.Of("0")
			bigint2, _ := biginteger.Of("2")

			Expect(bigint1.Modulo(*bigint2).String()).To(Equal("0"))
		})

		It("Should return 0 for 2 mod 2", func() {
			bigint1, _ := biginteger.Of("2")
			bigint2, _ := biginteger.Of("2")

			Expect(bigint1.Modulo(*bigint2).String()).To(Equal("0"))
		})

		It("Should return 0 for 4 mod 2", func() {
			bigint1, _ := biginteger.Of("4")
			bigint2, _ := biginteger.Of("2")
			modulo := bigint1.Modulo(*bigint2)

			Expect(modulo.String()).To(Equal("0"))
		})

		It("Should return 0 for 8 mod 4", func() {
			bigint1, _ := biginteger.Of("8")
			bigint2, _ := biginteger.Of("4")

			Expect(bigint1.Modulo(*bigint2).String()).To(Equal("0"))
		})

		It("Should return -15 for -15 mod -16", func() {
			bigint1, _ := biginteger.Of("-15")
			bigint2, _ := biginteger.Of("-16")

			Expect(bigint1.Modulo(*bigint2).String()).To(Equal("-15"))
		})

		It("Should return -15 for -15 mod 16", func() {
			bigint1, _ := biginteger.Of("-15")
			bigint2, _ := biginteger.Of("16")

			Expect(bigint1.Modulo(*bigint2).String()).To(Equal("-15"))
		})

		It("Should return 15 for 15 mod -16", func() {
			bigint1, _ := biginteger.Of("15")
			bigint2, _ := biginteger.Of("-16")

			Expect(bigint1.Modulo(*bigint2).String()).To(Equal("15"))
		})

		It("Should return 1 for 17 mod 16", func() {
			bigint1, _ := biginteger.Of("17")
			bigint2, _ := biginteger.Of("16")

			Expect(bigint1.Modulo(*bigint2).String()).To(Equal("1"))
		})

		It("Should return -1 for -17 mod -16", func() {
			bigint1, _ := biginteger.Of("-17")
			bigint2, _ := biginteger.Of("-16")

			Expect(bigint1.Modulo(*bigint2).String()).To(Equal("-1"))
		})

		It("Should return -1 for -17 mod 16", func() {
			bigint1, _ := biginteger.Of("-17")
			bigint2, _ := biginteger.Of("16")

			Expect(bigint1.Modulo(*bigint2).String()).To(Equal("-1"))
		})

		It("Should return 1 for 17 mod -16", func() {
			bigint1, _ := biginteger.Of("17")
			bigint2, _ := biginteger.Of("-16")

			Expect(bigint1.Modulo(*bigint2).String()).To(Equal("1"))
		})

		It("Should return 0 for 4294967296 mod 4294967296", func() {
			bigint1, _ := biginteger.Of("4294967296")
			bigint2, _ := biginteger.Of("4294967296")

			Expect(bigint1.Modulo(*bigint2).String()).To(Equal("0"))
		})

		It("Should return 0 for 18446744073709551615 mod 1", func() {
			bigint1, _ := biginteger.Of("18446744073709551615")
			bigint2, _ := biginteger.Of("1")

			Expect(bigint1.Modulo(*bigint2).String()).To(Equal("0"))
		})

		It("Should return 1 for 18446744073709551615 mod 2", func() {
			bigint1, _ := biginteger.Of("18446744073709551615")
			bigint2, _ := biginteger.Of("2")

			Expect(bigint1.Modulo(*bigint2).String()).To(Equal("1"))
		})

		It("Should return 18446744073709551615 for 18446744073709551615 mod 18446744073709551616", func() {
			bigint1, _ := biginteger.Of("18446744073709551615")
			bigint2, _ := biginteger.Of("18446744073709551616")

			Expect(bigint1.Modulo(*bigint2).String()).To(Equal("18446744073709551615"))
		})

		It("Should return 1 for 18446744073709551616 mod 18446744073709551615", func() {
			bigint1, _ := biginteger.Of("18446744073709551616")
			bigint2, _ := biginteger.Of("18446744073709551615")

			Expect(bigint1.Modulo(*bigint2).String()).To(Equal("1"))
		})

		It("Should return 1 for 18446744073709551616 mod 10", func() {
			bigint1, _ := biginteger.Of("18446744073709551616")
			bigint2, _ := biginteger.Of("10")
			result := bigint1.Modulo(*bigint2)

			Expect(result.String()).To(Equal("6"))
		})

		It("Should return 1 for 36893488147419103230 mod 10", func() {
			bigint1, _ := biginteger.Of("36893488147419103230")
			bigint2, _ := biginteger.Of("10")
			result := bigint1.Modulo(*bigint2)

			Expect(result.String()).To(Equal("0"))
		})

		It("Should return 1 for 36893488147419103232 mod 10", func() {
			bigint1, _ := biginteger.Of("36893488147419103232")
			bigint2, _ := biginteger.Of("10")
			result := bigint1.Modulo(*bigint2)

			Expect(result.String()).To(Equal("2"))
		})

		It("Should return 0 for 340282366920938463444927863358058659840 mod 10", func() {
			bigint1, _ := biginteger.Of("340282366920938463444927863358058659840")
			bigint2, _ := biginteger.Of("10")
			result := bigint1.Modulo(*bigint2)

			Expect(result.String()).To(Equal("0"))
		})

		It("Should return 40 for 340282366920938463444927863358058659840 mod 100", func() {
			bigint1, _ := biginteger.Of("340282366920938463444927863358058659840")
			bigint2, _ := biginteger.Of("100")
			result := bigint1.Modulo(*bigint2)

			Expect(result.String()).To(Equal("40"))
		})

		It("Should mod 7234724982749223422342342348234982734729384727349827498233423423474924 by 23049828420348234290384203840283402840923842320234242342094823", func() {
			bigint1, _ := biginteger.Of("7234724982749223422342342348234982734729384727349827498233423423474924")
			bigint2, _ := biginteger.Of("23049828420348234290384203840283402840923842320234242342094823")
			result := bigint1.Modulo(*bigint2)

			Expect(result.String()).To(Equal("9615245427589791313307081038888134199181203328816147362583360"))
		})
	})

	Context("DivMod", func() {
		It("Should divide 2 by 2", func() {
			bigint1, _ := biginteger.Of("2")
			bigint2, _ := biginteger.Of("2")

			quotient, remainder := bigint1.DivMod(*bigint2)

			Expect(quotient.String()).To(Equal("1"))
			Expect(remainder.String()).To(Equal("0"))
		})

		It("Should divide 4 by 2", func() {
			bigint1, _ := biginteger.Of("4")
			bigint2, _ := biginteger.Of("2")

			quotient, remainder := bigint1.DivMod(*bigint2)

			Expect(quotient.String()).To(Equal("2"))
			Expect(remainder.String()).To(Equal("0"))
		})

		It("Should divide 8 by 4", func() {
			bigint1, _ := biginteger.Of("8")
			bigint2, _ := biginteger.Of("4")

			quotient, remainder := bigint1.DivMod(*bigint2)

			Expect(quotient.String()).To(Equal("2"))
			Expect(remainder.String()).To(Equal("0"))
		})

		It("Should divide 80 by 10", func() {
			bigint1, _ := biginteger.Of("80")
			bigint2, _ := biginteger.Of("10")

			quotient, remainder := bigint1.DivMod(*bigint2)

			Expect(quotient.String()).To(Equal("8"))
			Expect(remainder.String()).To(Equal("0"))
		})

		It("Should divide 4294967296 by 10", func() {
			bigint1, _ := biginteger.Of("4294967296")
			bigint2, _ := biginteger.Of("10")

			quotient, remainder := bigint1.DivMod(*bigint2)

			Expect(quotient.String()).To(Equal("429496729"))
			Expect(remainder.String()).To(Equal("6"))
		})

		It("Should divide 4294967296 by 4294967296", func() {
			bigint1, _ := biginteger.Of("4294967296")
			bigint2, _ := biginteger.Of("4294967296")

			quotient, remainder := bigint1.DivMod(*bigint2)

			Expect(quotient.String()).To(Equal("1"))
			Expect(remainder.String()).To(Equal("0"))
		})

		It("Should divide 18446744073709551615 by 1", func() {
			bigint1, _ := biginteger.Of("18446744073709551615")
			bigint2, _ := biginteger.Of("1")

			quotient, remainder := bigint1.DivMod(*bigint2)

			Expect(quotient.String()).To(Equal("18446744073709551615"))
			Expect(remainder.String()).To(Equal("0"))
		})

		It("Should divide 18446744073709551615 by 2", func() {
			bigint1, _ := biginteger.Of("18446744073709551615")
			bigint2, _ := biginteger.Of("2")

			quotient, remainder := bigint1.DivMod(*bigint2)

			Expect(quotient.String()).To(Equal("9223372036854775807"))
			Expect(remainder.String()).To(Equal("1"))
		})

		It("Should divide 18446744073709551615 by 18446744073709551616", func() {
			bigint1, _ := biginteger.Of("18446744073709551615")
			bigint2, _ := biginteger.Of("18446744073709551616")

			quotient, remainder := bigint1.DivMod(*bigint2)

			Expect(quotient.String()).To(Equal("0"))
			Expect(remainder.String()).To(Equal("18446744073709551615"))
		})

		It("Should divide 18446744073709551616 by 18446744073709551615", func() {
			bigint1, _ := biginteger.Of("18446744073709551616")
			bigint2, _ := biginteger.Of("18446744073709551615")

			quotient, remainder := bigint1.DivMod(*bigint2)

			Expect(quotient.String()).To(Equal("1"))
			Expect(remainder.String()).To(Equal("1"))
		})

		It("Should divide 18446744073709551616 by 2", func() {
			bigint1, _ := biginteger.Of("18446744073709551616")
			bigint2, _ := biginteger.Of("2")

			quotient, remainder := bigint1.DivMod(*bigint2)

			Expect(quotient.String()).To(Equal("9223372036854775808"))
			Expect(remainder.String()).To(Equal("0"))
		})

		It("Should divide 18446744073709551616 by 18446744073709551616", func() {
			bigint1, _ := biginteger.Of("18446744073709551616")
			bigint2, _ := biginteger.Of("18446744073709551616")

			quotient, remainder := bigint1.DivMod(*bigint2)

			Expect(quotient.String()).To(Equal("1"))
			Expect(remainder.String()).To(Equal("0"))
		})

		It("Should divide 18446744073709551616 by 10", func() {
			bigint1, _ := biginteger.Of("18446744073709551616")
			bigint2, _ := biginteger.Of("10")

			quotient, remainder := bigint1.DivMod(*bigint2)

			Expect(quotient.String()).To(Equal("1844674407370955161"))
			Expect(remainder.String()).To(Equal("6"))
		})

		It("Should divide 36893488147419103230 by 10", func() {
			bigint1, _ := biginteger.Of("36893488147419103230")
			bigint2, _ := biginteger.Of("10")

			quotient, remainder := bigint1.DivMod(*bigint2)

			Expect(quotient.String()).To(Equal("3689348814741910323"))
			Expect(remainder.String()).To(Equal("0"))
		})

		It("Should divide 73786976294838206464 by 10", func() {
			bigint1, _ := biginteger.Of("73786976294838206464")
			bigint2, _ := biginteger.Of("10")

			quotient, remainder := bigint1.DivMod(*bigint2)

			Expect(quotient.String()).To(Equal("7378697629483820646"))
			Expect(remainder.String()).To(Equal("4"))
		})

		It("Should divide 340282366920938463444927863358058659840 by 10", func() {
			bigint1, _ := biginteger.Of("340282366920938463444927863358058659840")
			bigint2, _ := biginteger.Of("10")

			quotient, remainder := bigint1.DivMod(*bigint2)

			Expect(quotient.String()).To(Equal("34028236692093846344492786335805865984"))
			Expect(remainder.String()).To(Equal("0"))
		})

		It("Should divide 115792089237316195423570985008687907853269984665640564039457584007913129639936 by 10", func() {
			bigint1, _ := biginteger.Of("115792089237316195423570985008687907853269984665640564039457584007913129639936")
			bigint2, _ := biginteger.Of("10")

			quotient, remainder := bigint1.DivMod(*bigint2)

			Expect(quotient.String()).To(Equal("11579208923731619542357098500868790785326998466564056403945758400791312963993"))
			Expect(remainder.String()).To(Equal("6"))
		})
	})

	Context("Negate", func() {
		It("Should negate 0", func() {
			bigint, _ := biginteger.Of("0")

			Expect(bigint.Negate().String()).To(Equal("0"))
		})

		It("Should negate 1", func() {
			bigint, _ := biginteger.Of("1")

			Expect(bigint.Negate().String()).To(Equal("-1"))
		})

		It("Should negate -1", func() {
			bigint, _ := biginteger.Of("-1")

			Expect(bigint.Negate().String()).To(Equal("1"))
		})

		It("Should negate 10", func() {
			bigint, _ := biginteger.Of("10")

			Expect(bigint.Negate().String()).To(Equal("-10"))
		})

		It("Should negate -10", func() {
			bigint, _ := biginteger.Of("-10")

			Expect(bigint.Negate().String()).To(Equal("10"))
		})

		It("Should negate 10000000000", func() {
			bigint, _ := biginteger.Of("10000000000")

			Expect(bigint.Negate().String()).To(Equal("-10000000000"))
		})

		It("Should negate -10000000000", func() {
			bigint, _ := biginteger.Of("-10000000000")

			Expect(bigint.Negate().String()).To(Equal("10000000000"))
		})

		It("Should negate 18446744073709551616", func() {
			bigint, _ := biginteger.Of("18446744073709551616")

			Expect(bigint.Negate().String()).To(Equal("-18446744073709551616"))
		})

		It("Should negate -18446744073709551616", func() {
			bigint, _ := biginteger.Of("-18446744073709551616")

			Expect(bigint.Negate().String()).To(Equal("18446744073709551616"))
		})
	})

	Context("Power", func() {
		It("Should return 1 for 1 pow 1", func() {
			bigint1, _ := biginteger.Of("1")
			bigint2, _ := biginteger.Of("1")

			Expect(bigint1.Power(*bigint2).String()).To(Equal("1"))
		})

		It("Should return 1 for 2 pow 0", func() {
			bigint1, _ := biginteger.Of("2")
			bigint2, _ := biginteger.Of("0")

			Expect(bigint1.Power(*bigint2).String()).To(Equal("1"))
		})

		It("Should return 2 for 2 pow 1", func() {
			bigint1, _ := biginteger.Of("2")
			bigint2, _ := biginteger.Of("1")

			Expect(bigint1.Power(*bigint2).String()).To(Equal("2"))
		})

		It("Should return 4 for 2 pow 2", func() {
			bigint1, _ := biginteger.Of("2")
			bigint2, _ := biginteger.Of("2")

			Expect(bigint1.Power(*bigint2).String()).To(Equal("4"))
		})

		It("Should return 8 for 2 pow 3", func() {
			bigint1, _ := biginteger.Of("2")
			bigint2, _ := biginteger.Of("3")

			Expect(bigint1.Power(*bigint2).String()).
				To(Equal("8"))
		})

		It("Should return 16 for 2 pow 4", func() {
			bigint1, _ := biginteger.Of("2")
			bigint2, _ := biginteger.Of("4")

			Expect(bigint1.Power(*bigint2).String()).To(Equal("16"))
		})

		It("Should return 32 for 2 pow 5", func() {
			bigint1, _ := biginteger.Of("2")
			bigint2, _ := biginteger.Of("5")

			Expect(bigint1.Power(*bigint2).String()).To(Equal("32"))
		})

		It("Should return 64 for 2 pow 6", func() {
			bigint1, _ := biginteger.Of("2")
			bigint2, _ := biginteger.Of("6")

			Expect(bigint1.Power(*bigint2).String()).To(Equal("64"))
		})

		It("Should return 10000000000 for 10 pow 10", func() {
			bigint1, _ := biginteger.Of("10")
			bigint2, _ := biginteger.Of("10")

			Expect(bigint1.Power(*bigint2).String()).To(Equal("10000000000"))
		})

		It("Should return 100000000000000000000 for 10 pow 20", func() {
			bigint1, _ := biginteger.Of("10")
			bigint2, _ := biginteger.Of("20")
			result := bigint1.Power(*bigint2)

			Expect(result.String()).To(Equal("100000000000000000000"))
		})

		It("Should return 4294967296 for 2 pow 32", func() {
			bigint1, _ := biginteger.Of("2")
			bigint2, _ := biginteger.Of("32")

			Expect(bigint1.Power(*bigint2).String()).To(Equal("4294967296"))
		})

		It("Should return 18446744073709551616 for 2 pow 64", func() {
			bigint1, _ := biginteger.Of("2")
			bigint2, _ := biginteger.Of("64")

			Expect(bigint1.Power(*bigint2).String()).To(Equal("18446744073709551616"))
		})

		It("Should return (huge result) for 16258 pow 16258", func() {
			bigint1, _ := biginteger.Of("16258")
			bigint2, _ := biginteger.Of("16258")
			result := bigint1.Power(*bigint2)

			Expect(result.Digits()).To(BeNumerically("==", 68464))

			modulo := result.Modulo(biginteger.OfUint64(1_000_000))
			Expect(modulo.String()).To(Equal("662784"))
		})

		It("Should return (huge result) for (2 pow 14) pow (2 pow 14)", func() {
			two := biginteger.Two()
			fourteen := biginteger.OfUint64(14)

			bigint1 := two.Power(fourteen)
			result := bigint1.Power(bigint1)

			Expect(result.Digits()).To(BeNumerically("==", 69050))

			modulo := result.Modulo(biginteger.OfUint64(1_000_000))
			Expect(modulo.String()).To(Equal("827136"))
		})

		It("123 pow 12345", func() {
			bigint1, _ := biginteger.Of("123")
			bigint2, _ := biginteger.Of("12345")
			result := bigint1.Power(*bigint2)

			Expect(result.String()).To(Equal("756137399546316460705162409626542886222066220152418320848613295496552630922260433617779137550814204108921516070029961165178507081553817710836188698253584651316669837036834833252496653933690421089849557803264443462704938872773167647892456537965911434710565471031231875715831869056051328923375088502201750555868975402087039063687550268372655656918889419492136333736351829045696809517932178713759973802905135142564957314354572551132056108167353134230815792644051764346023660582608167324644164610819471476569696381419487538525877692987536061192218986335000303767183495228880222291166067108749835149655423492288197712194897339730070373350569489963389338839480076391350573539405336164485992215660768005758183352365720371371225021037497897473582848795909103878616963940057517825497879962459590471839075679388267523315875321697117575411671368027486838866990044816560935603845135615915034198986546588006175197859614038735285007703248207587909769981698743980214589915688447023194897697360458380568181597563807522841543034135496541442657246172694815178485356768997631875156311069480789289402409542527036358878511113891052600700454431746781485269715186980445066791124217132198419993712164168228259983597516993192212319028971985430138569527577719862661739450698177749917560615314239725823820873515989759674106409862888869933262066556629952493076667050974825582817082453003164299337176971921179196670622921109740892002751442405534923702043168084069998587129994562780007815756785850242558095739067872134844449773038638270026945731926820894468092616294164890018960172855877640840351546284018469942672692227611871422447425200831826988606806013017825552508816937756526798472606627618842967320671344118207417664870601378429038024136134311979026029770277507837566326761662156926612157393162440804257417274488704297593634930391340867046177823198726785956536958196243202370970862793569371735463576946093945539530736362094699543216016726202731689296887503570792898016764733399775980141937775621390829885113003071540126358245066660933411576410001620855356379231894680679149161429399656246902446926675967817499565138186497191940585672358930074285017909905347228658856491909714234911222424607399612196432093562108948549968787413986971624439668386831819385761860721288639914858133936158164888933119286316163990003841947117469507307336086977467012781033903143612646137928107264683992559029856718307096893729253652108499183762954677551848009266499030913572236013166758120273412404655422406608310284551351442383346651778185537631193518462923109035278374294356237823853016753486482723833152604493544977815391630412216977811986955225877904494641651655111317764071019117781165352566390464725614166612089628990581273519969794670116897599233367778000838998824854056185110354628622212002307240703957543529363678785471529974823877997025170594999456284636611299197922215808380239624314211930715801351803007123329222540301452210343289501392732347062772600590671711159119159073086190882014781429195699295975778011928906389230843239420858120365250088359692541927840205057806383911420587409838382682007225254756279127881637672372728115227870439174902393654062150715243133960566000764358088827139328861599896283293637759767664074216328030583480577478779703694202929800360882253421607183735676091366926399384407179048684158341417857933841229327765688486125668977082660136231668164226463828748464219992824922551727056499611267498699868156769325346395650511057193873964830079152391745158332188561803046640568066491301631123282537496809760581070367158484895159779217184774812929802240043178618748205722141200951091463866478297333659177502532654395724826717798107969752932571649806740904928948333584717610424802988425568959541958567908664411023910241339925778980437357940294555994492784917406606557896476701324290975307219631940568720391195383645233048330994562267735755243453213090036245854962724177096183625746633873441516102640618500678637031129439456694130226683034603801911732247661902442280435395882408391654306377808372714719645960888915151919846188134203849619056284992976672927620844371149848455222692468480915217007404366894135302399504779291471547229309400156475476151961236274411448587753105898433305032032473570110765181701724751380613678026099532408297538039786872204736461524601650353795140497565832464149977377116409964138265403302487926595626365176255689232142964303875627870128082248370287964952088372674349708997719108334250343467509500246740690789147865308855923393024283161183656345668890911344969161561824642096135344099593185172223206762707564769001156036503194326141467692521840125434185524135588044366612575830809180313704485900700045146640315148073716039245289673153919781983398203117895688219284463604449388910906574058957602603731385871172222368990274959121940324402429837281442802277111679029585216901739597986782531763256284858099995185876455356371872617748647158155026428618863514818233515957538449346904698775128365362584498203234411766879794212771298110519777799488692220796444171859351539547880458470331419312598441057700673875284975009259706370528319634956668797147042564994375336988685870224452962851103592761109528493861758944034057153264941807788730985377210120951332456721960545245619271977783868659413436730002010016692914684745810419707070523112847990121036625603401385346219670844303568760144488891400026541711640169119373199464258066597382268261157815180226794250432239218577953754577822633733365031749911586185540602543227693139945087304204378798918558247048547775407459462521011111145047757366839511324626869046692694184283750576425814952459076891111025233204341768903841341517048607168806321567677204821988383861987729017575554324212575026843105061655335977804389272966614486368107641498013976245694016082986420594501759357362265034029971354039895300471567808457978690256164763458828506456627909929275917135664707854181683323866246653217449328486702557428696720968396306384512202391419753835321371800420081971385430452393662297434015851479364370592335340820257036756130582001760604898287792462216102732257016712079783905517628870297386910910202915020308497952167211556837610078506980432770850161297459136963608918618718237769309873153124104857900845483750109857262664163822219410010509000061976596167545455165429224971392212970795321700217772034695918052298130666359995247646912876648503079731195400530847230710631350424827929743006105885882628130212140657128620914151239140434544693015982415363960265518564231049158928808319008306991634081643357242872022412850789788829838038557661058249690125172448887649298165179907811453136928918680585485877370291961507277299538046233045214222781265859432069685155362359698038857275020216015419146438506032629245052580462950536703556968966476807117293709736329198948038160011605346529257739622356309605600839385931214981577689447649554101112327080469915305165690463695892529397060327658136426135901791592873141935923186283346605105577083371728750765955329592151227318337981237263360657624542108668741553135433121206423844869616850173159568040155060347588738446039793467461541100535608333907486918676942734296624174364414758078647783005120465227543700970454041641084645237747876841743139778527481477014081364041299290542817049921326468152244599682420195945140330878438954859006960987704832981075926581127482396533646953022233083133750645742276385068170498554373300880019620118153073646341472696987565753195273146815166603513219192210391644302765561517282012949028286448282211568193513868099792442947973326157237148443501782083263342592503495204823143714317334893495556834864277323334139740777794652082048986654549406522472489760348415154246608201718376005988049992185531194550326156007076677243023113562797692487770168943907170271394808452513250705075353267028609023983349364215676815686664961146552749691469626949564598948581943117754052844206743255202296092110090304006116588560513294210624865742730453481147151786259223212628215718126754413392854833208844492400396627400070778420601446540629592153072527365730868506243700991741541594067416163110214611565019630827928217647162344932494667086996254442270997306113728079687255599037619301421327915492949628010953501439402625414889776582988752854872095182663752775978693689687308351379290144724264698880580989762604095312816079481274649460106918782876861024379655149540582453843052014079843012878641668382553581190109700208423796842290515426116930671106702210430367506924574612464230836445092994707304907645595652811652670877236072706931949587058517359745666139400226499218958998427852704301439902574587145157923225812268770873821144379306522787385285748302041890008449359134841135356504337965572075530010688611212298259569530919011177408682359532304699728095250203552658008023631226386373546967339471597035862042187265162111900420550332936019546535610288862554955105587109924780054379089103288125279837930195226179697961214763359077061566635823741693840193904765957237340644981321176167102808491628168817827089391080191921202180435665079168711371805208624215487620986627883994048781951368128795240276657806663761016986205268700302729125030394146892329294620090245781631997712065375017238377662668715813614987007278276785126225010376251671143993062195890378507466805013567734109943697589159016193815827820957107574917873021084220204165463779430894446403701744536766875335131737933865577673681730087815783736633416721922653904451444638711996439888408020240152932198360931672785179445654258153264809344282944673138662146069212976679445516182679784805700787773316702777109375750428564693492720045980192554776647333263617846238338480224752606238951154365356240465671008821620275243243572468300475862444558337873990703087658338448480461317543390040975627365946197457149651780002294618478502918644747609692525842202555562406166299726497382838560405702701949734367687583503965501846299408044616549748687241094272988801847966441505126820853434552706740488956748984107022252625574838124304764399323057830913476448117923946814436073983392392249050018790243135215471020206333667803869998530720673949459646805368114593578882575544222311505140433516570610834423945035978301261474638524633880525508460775846949394554019902839397010437256417846826929066588175663217138453564962550310664523723160083833748595053087618322722423821709596131874542825741481411377013619074313009367911815801182312247928230161114488074654340463582677104941248460869824914927591168643583235966511386643265855778017108869200174503080591361723925520428706006038161176527978813909579942361846386591413973041014419061946979517693957756897402824078108006190280215306821757165533593464396764490821329953899959188193451332858154559147107201513950073134851254400621251592996396115947297515814716625090311214618700275605511262799531626203852264304030436531956475648583698356020560348215286096300237349195902985765878658958434945830605432858284422407356058919932634599133317055389182883774573733877496549685781380956309524874716257839776784756872427940351901911320853999488954627980034543426440686912849219760032016714422767368252056017121819355887354491274871880149714078119216066716514460869483092882874458738222040901978970698798015305134882034788566728499695354322924468183974931572026924209894727316565627380034411782409345331201723727846787145510613132075240183913183114432945251304810859908523499936015266753121743726476745669901302654647416555252295363618582734484823928260238595582484212188309776706234685105723499914410889635382526732175789635664851028988155525819449183523209911889243280205592716939269794144257477742910474873290739660018178677955312506486141417119689006693574500292486699188719173570855659291625865363013950788028324338935929910131482575012386461763224857776486358691784703383581239221827300782936221006175526607497001185755386692401827982591149302619929944747399536296205021186502345736147495321534995966876896719458168267286307836084466048198711215969754975284836677696793844231454332063598646296672784989517351268087131145727841717331999198129736212510282120996024918640991061645730761971948871183759348919390234222275115930656121478305775342569425256268896373922699652069183710099179078008858656992706529682280487661295861526119302256967664432695072122608372649356379017738449000070189263860218949717630550568321143951821991456687358424458391182929861814783496369547954023469513943327426520837442908712657964056243164942284861273805076655340619950783318530295775257530788025473539168950360639013254212513420504786503604526478526712317833383752375773972977265422193578106072127156566844137216925715019796578862038465382679386232477898035447553358505255229734700795803008458392092913645457472686639406709101778455323593116763722943816681203853563512676352022089949719139062205320311821193797935993286668132176163346167477256149047421864303668549510681975750430973164488613504875660583542386714871788598565141060081401341147265011387271430859045651897660388612685477666447729622623334294039540385622212807780866186809828181356625841724650448083523491697680990197605572255947868765551549569226805813233847862704384527915827029720441436304284902379099595887374246667853560273409773362087215477591434160980573026705172429337402799987445668118804271024294853872511041608253107653725566856945773240222711662789059406441439354725272910813984251766397160093933146195673147965455816721243297657223320061336522029838094427751224578739467818762734306356986950267962473375074314782076895133911434055121886960594458723637017191197875970436966564644933706219284176595013913887322486311339277054618498246051387691604991181048488135742143455577908275597033780154017813905406844760603266663649643511124947936500274370044888385855449071891997303492344589880113565985711814070723065698185645560239729825634861512998174222670304435240541738635532710165928839157110610135122821665618950548914238003473603450453638070327413485958548140843827266751056488904902386786190681483318296869786449808110731244212227576154026196819771895742281191987898446800772162374102950715937994266811544387341115485821831198600891252772827616331193560899537076094200791217605760140582641303682386046390646605959941189806945258638550983076509980913828411902216046194214190636882549044985323657906532996827992470155436149011794046153695734988344763968022089759997497617914748027000860904661943690260842470953673366768396686057871381443352494795465931868714644871476833762219294055941134364692429895266112517183609015911868409654817067249190516525321753061884428855196105383169504802989225172644773563096908366503683949180751578981050351809115048106993598057553794930702274151953927008792456713999775213971676336434450502683993992948510056900290748943632752348506148269134138421710552321957505440459423577635656022576761249613098965396441889576807385059635543942555908794325850774304967687213111658075620984290210998550676645440983195609551723967777153079136545981687065842317068130204688408211188128713998509275178640708597333681139578353156948160605040527016275515439114623904678324733850616356690255595010215913039042036772233024123525347570721670704998751217147778916162877797910486988045143013205475183700880931131188325810690242670256551277291928619177408057771177934198047648166118054853537862387905422697201921265434139723424979035949013629366066434470118766743569157484274187419245809194285745080810299653675858038259015112554489599465758123551278228774899380117493555091267716091909903569464013421802525412139877918828813496565040196304634690054348209162992223500150984371382447411239255952109001225297059790926128471262622207190820726690967767942527235674403790795072477044647886402924737945142032959256952592084899652634150180563839167493638209203042202390254237312286676311802374244826090191111915084381381511269228584121649649265491217741242792606593662993998853530216572231487456228432296401022202423963216738280044190551950152806294143984605585029226602410794979945385207354335909452654380852017607831980975164318555520729239304013039287285318299588429271095897216010340761043440067869066991418240697798484855029116187238975782941362670255241381549439591036579074521092650650233490017851100272446193048722391927963332379366036568736711784902353646591128459387832667332809083007996884430646202865419353969876061285610224468005129843330438804748193514334277134408438813181257612421498518642800689745117879569246501375523930252330652077963136935102015233533658301104410807416419278796020368430474527657706471830249571049179509540119216769167351946200907963981648354189046208188511682045483016949830227151282765620325989844658364396272781407158613841800853966981041400266455215671745232548386272703928303351453391598199674644616603517709983003138772423190037719880749877759145056250779313856146533130548131395164103427879526860931089519953042046767609878806870795720381919915424340314387912828809598764438462842228930697505788032606877982787575510572427006680801675991558916993497451190846132840853815461104036251916434378455160157861482346311165328245766947893255424960465488101492603773514360814326186512539868228324176970210492840382883939871005663066203352227491797258178430992475374390673710292167884406021190556355763082209518662373725770911891374336199248295880593856312872001152972735638774268959112960958928070522133831467191228708692008951554627391942267277503788471340259963478964815729100803088521662211152349355053422178267654836360594150118419114801919382170546692474477718095599194530694710538477281956194173256893223222684212908148761579345379972195927594258050012162929545715615556731329996722415123959599642563843622909675404033857519001528456530723703935800694608889907069966913960081442212884167995183824585723041683367626090903949615821224250650953236044991386052504490853314131480754366057698462699452615339578736395104818083940108714478420116597389357724398519455272001437114525798008591605821136241895264650359060112223297538199425160639643449761182490105683213315467472004899722549498960125024836896271880904855266729362651778241951043966264102905875701766424820594679991896268903453180876398792605325817965165752483024692482854451320680840939445469883297725450029489869361326395295273946658066954525399799967571969483623604260605313043200868911734300818290045070420942630076647309165176313456140900715409811508617703348226695964985690608985189038662779658721815551493342142043399215474395817919223079692789075300433869022143193395486372456698676461495147269425992258200115125752075370260937578323568379209133883187012496218936073941940402856928826465317929898989582968168941010884196345657917394687720987216695857647349335629098847935910924043077005560931184020448715576050257148826431212011167107706843102140507432509991421736211469246243990257736620104202001814181214201549908818077839700085103394579395132923011158058717774666581749839585172058770845203060820780501870182182217057292144598607844381218007219950611907555081002729363820398663162560701883849989171634199121092032809705768984859150188220816898302787496393272973577341742740682588532259601162405682587024845992393235210639766679825676007485476976689603696614723736993771624848029464923040131265738028664282185431370116896129985424406653235358860068870038509091296427794592398867288071622210890783926282868092411297382306821590177831624304076778041274679091792436440823188944379770184840079861453874862464588385148791544736603294733087206617661927533696115686574508110705091485814958179890282615920060274728496405808665766621636028247048410570975311336084207490306010916351568631895884631863399731275026286850423698228484759965067954797801875988449126675273976059720686824296412318416594077730907987171574921886346564069608392982060156884342529095296507114785318693454884013163065722466054005517973115771161501343869519403716611199511050983302361690972898006207372880685539045638386322990238748657236405745796244579141012346600063309670891478472806614420926617731421254314713601874925336554593470178129460202022074359722107904819998419289903849787190316372526851217601102337369251156670110948295954563552744459345842248293093977592863884985210116830469192601817282915516600583117776569097654960700318176058847555341086489750582162619342787757338514285631708913457927889911374819810132417144469529855998376569692813977230643846327736653884893976900227870130814645769410557209549089513350749793369782766077427429368301982227762596948469905047431306548416823599068812226766124063242975464230612405519708380927252733793426537357845899088216021909260413162484855005290450909720628557655466451537047052449248395120243865013934059232722051437497329817219990680198416292055172413330798201894481587111632381140597788912180750704467430380022390782403813673034458283355185925692954212670612299954427707882842446888173899240022510286859734869143999728533095850002284622154773429257413089518994784532135729447209467507172488888927612987310626192589938067498529284186728563844847580145642157652058332678086906186717519120470842847822838030676604399687202944678570836493997318038617901128859143036358678490315934893194540122595713423191532086188234840725304223813731277072728388044137418188134499644957119512668985991135410855294739032269490995669752768363821544970387338705247473795949376712484006589044425581197554581283728118877496855444895614948501335787018648802373028933519241883674167578522688778855705465719458601018125142176199937611196597719101545183778211635075567043840305561255624454022151429319334692660044035384022104893118285684663200962424422697132666196924082674142828761807735637013535000410414769229261880810520409691803942106095220198845715407390151694253849185699578534180416264720636264906761734608122411587646544498352471933193402060916564192770710771969470556273139562265189870432487198796423056261185527550217136663284047421230634809921062918987207296693081721085262695792997371573328524580798527850769220916783723728129935549813986123944351621268998339770565743457891052828504455576529124614548001904588370021753990311082353795880211624812045501833874381425130113700624277911449881922181581561599440171994649037911784415656569279134521146201355619912187958449906856669760539678072443690246511236597083334488166803776107717110055789056397392845025731431387521229913607269797027362033000990814214183666768625496573723557813073825689507503646210602148393795721487845527737763108276001029760153121342446383282954361743726619307380202281624210137831139760292150812755962799715322096268705859478572433452018232245559963823379419015997978020583112090923239961692389879386527419124353560745111457723320337420151957194709360159685936253931881039997984883920529459320731978700466162720139892000904323814728440909966764175950118721784610023043148168363222743675366145953704500058571875296249104311634906136979680056841952114001959504747849956566426696302962482581949267928592123779234007639043514990586684444643126250079754477241298991594860422290447898399342366366744188282632068249167326221844787377656020195008915427382554138971524017296086516297448134657227651393491166339777657181669969583480527387842837356241528132941549760033213480243114692516170729740882895255937700431982758030368876857880947478369836773103336542099529290385658108542229139115840986949203516797054168439329195550704789348116479352577372400006663864848186854236843493397790311411577940917886538744530089416729045520383519626682090403974679092436128456062410844204609071482838468806564132401746220949668762356414990353617825019401047867059168123058431385936313572033351696975843160050960650911597413677291982446673487709462466564974576332195609857535703504200657011610205458692111704384762644574777563644223276171439432910831266455747878532922211739849940413196353113775465530893990446909035759267803074365098568635953547309753676791200292385155093782405368868952510995892832813462038211477349785596983725764346245722856674801010057443524592763730596654831706621839437508763446801052950312256481764476095492812490068441147301280702309629330970820148576131894770323530363605094633562484642772795932314924435133592896246515048124119965713645495372989217805762438432699907044768930718301200318058618256521807479272269365538597925481700534684446420615764593150212595410135174761057559915371488877817748309536326136393367458751756984164146160979874450503602125106973897107325964030086442967071224706922472932746709861391909610533839060343979520740151895541775209878412260913469237261708543045046727414396234527792256065738464231617635874122327001334582567634390949780708549048758054792334563017996883355922582434584674284799088668531111962109502909525488399923616442330625009008206814815516013256109242214233675259617817987537356086137276647775529602246381075240692060868651006021017363505259365494283966831015560927831317839447993467486275434451697801179934747193488960954613778430157676569574162937827430498130962487915335643296911827866519001702637776819020834867826097851637286160801236927256147356459825601148509148590017444846665233064035233479832688265080796319427924231494646706048405707671873304072346141918083084445627143385445648273244832559761945658747982801580479729180403310069863326176000204066826348448400965935640034579692960543685787671065864792487944689852786811104200606961631484130550712127819206213473773003553022258104661089166866059560482278528231604865872615778396351677613694317356406884557846353648995186505111733569443330895372418460061435627578422437224548292069425434211165157854231595998947445758990503171063301556091142127678133634006777427131422196745899620078020389523521152817608270743852852146940073875998180962945059132507157145811425883740692323559322970041430015200614713894821917924666076562622511200098308714840950300173181105278447571434588559964605230038518508905059527747199630289189893946808067711912365123563106998533882798331149612010669723157220086951496354354488347759301824906260776909957924717217020972638444637452330549846581731012207539373685706240142043"))
		})
	})

	Context("ShiftRight", func() {
		It("Should return 0 >> 0", func() {
			bigint, _ := biginteger.Of("0")

			Expect(bigint.ShiftRight(0).String()).
				To(Equal("0"))
		})

		It("Should return 0 >> 1", func() {
			bigint, _ := biginteger.Of("0")

			Expect(bigint.ShiftRight(1).String()).
				To(Equal("0"))
		})

		It("Should return 1 >> 0", func() {
			bigint, _ := biginteger.Of("1")

			Expect(bigint.ShiftRight(0).String()).To(Equal("1"))
		})

		It("Should return 1 >> 1", func() {
			bigint, _ := biginteger.Of("1")

			Expect(bigint.ShiftRight(1).String()).To(Equal("0"))
		})

		It("Should return 2 >> 0", func() {
			bigint, _ := biginteger.Of("2")

			Expect(bigint.ShiftRight(0).String()).To(Equal("2"))
		})

		It("Should return 2 >> 1", func() {
			bigint, _ := biginteger.Of("2")

			Expect(bigint.ShiftRight(1).String()).To(Equal("1"))
		})

		It("Should return 2 >> 2", func() {
			bigint, _ := biginteger.Of("2")

			Expect(bigint.ShiftRight(2).String()).To(Equal("0"))
		})

		It("Should return 18446744073709551615 >> 0", func() {
			bigint, _ := biginteger.Of("18446744073709551615")

			Expect(bigint.ShiftRight(0).String()).
				To(Equal("18446744073709551615"))
		})

		It("Should return 18446744073709551615 >> 1", func() {
			bigint, _ := biginteger.Of("18446744073709551615")

			Expect(bigint.ShiftRight(1).String()).
				To(Equal("9223372036854775807"))
		})

		It("Should return 18446744073709551615 >> 2", func() {
			bigint, _ := biginteger.Of("18446744073709551615")

			Expect(bigint.ShiftRight(2).String()).
				To(Equal("4611686018427387903"))
		})

		It("Should return 18446744073709551615 >> 16", func() {
			bigint1, _ := biginteger.Of("18446744073709551615")

			Expect(bigint1.ShiftRight(16).String()).
				To(Equal("281474976710655"))
		})

		It("Should return 18446744073709551615 >> 32", func() {
			bigint1, _ := biginteger.Of("18446744073709551615")

			Expect(bigint1.ShiftRight(32).String()).
				To(Equal("4294967295"))
		})

		It("Should return 18446744073709551615 >> 63", func() {
			bigint1, _ := biginteger.Of("18446744073709551615")

			Expect(bigint1.ShiftRight(63).String()).
				To(Equal("1"))
		})

		It("Should return 18446744073709551615 >> 64", func() {
			bigint1, _ := biginteger.Of("18446744073709551615")

			Expect(bigint1.ShiftRight(64).String()).
				To(Equal("0"))
		})
	})

	Context("ShiftLeft", func() {
		It("Should return 0 shiftLeft 0", func() {
			bigint, _ := biginteger.Of("0")

			Expect(bigint.ShiftLeft(0).String()).
				To(Equal("0"))
		})

		It("Should return 0 shiftLeft 1000000", func() {
			bigint, _ := biginteger.Of("0")

			Expect(bigint.ShiftLeft(1000000).String()).
				To(Equal("0"))
		})

		It("Should return 1 shiftLeft 0", func() {
			bigint, _ := biginteger.Of("1")

			Expect(bigint.ShiftLeft(0).String()).
				To(Equal("1"))
		})

		It("Should return 2 shiftLeft 0", func() {
			bigint, _ := biginteger.Of("2")

			Expect(bigint.ShiftLeft(0).String()).
				To(Equal("2"))
		})

		It("Should return 3 shiftLeft 0", func() {
			bigint, _ := biginteger.Of("3")

			Expect(bigint.ShiftLeft(0).String()).
				To(Equal("3"))
		})

		It("Should return 2 shiftLeft 1", func() {
			bigint, _ := biginteger.Of("2")

			Expect(bigint.ShiftLeft(1).String()).
				To(Equal("4"))
		})

		It("Should return 2 shiftLeft 2", func() {
			bigint, _ := biginteger.Of("2")

			Expect(bigint.ShiftLeft(2).String()).
				To(Equal("8"))
		})

		It("Should return 2 shiftLeft 3", func() {
			bigint, _ := biginteger.Of("2")

			Expect(bigint.ShiftLeft(3).String()).
				To(Equal("16"))
		})

		It("Should return 1 shiftLeft 32", func() {
			bigint, _ := biginteger.Of("1")
			result := bigint.ShiftLeft(32)

			Expect(result.String()).
				To(Equal("4294967296"))
		})

		It("Should return 1 shiftLeft 62", func() {
			bigint, _ := biginteger.Of("1")
			result := bigint.ShiftLeft(62)

			Expect(result.String()).
				To(Equal("4611686018427387904"))
		})

		It("Should return 1 <  < 63", func() {
			bigint, _ := biginteger.Of("1")
			result := bigint.ShiftLeft(63)

			Expect(result.String()).
				To(Equal("9223372036854775808"))
		})

		It("Should return 1 shiftLeft 64", func() {
			bigint, _ := biginteger.Of("1")

			result := bigint.ShiftLeft(64)

			Expect(result.String()).
				To(Equal("18446744073709551616"))
		})

		It("Should return 1 shiftLeft 65", func() {
			bigint, _ := biginteger.Of("1")

			Expect(bigint.ShiftLeft(65).String()).
				To(Equal("36893488147419103232"))
		})

		It("Should return 9223372036854775808 shiftLeft 1", func() {
			bigint, _ := biginteger.Of("9223372036854775808")
			result := bigint.ShiftLeft(1)

			Expect(result.String()).
				To(Equal("18446744073709551616"))
		})
	})

	Context("BigLength", func() {
		It("Should return len 0", func() {
			bigint, _ := biginteger.Of("0")

			Expect(bigint.BitLength()).
				To(BeNumerically("==", 1))
		})

		It("Should return len 1", func() {
			bigint, _ := biginteger.Of("1")

			Expect(bigint.BitLength()).
				To(BeNumerically("==", 1))
		})

		It("Should return len 2", func() {
			bigint, _ := biginteger.Of("2")

			Expect(bigint.BitLength()).
				To(BeNumerically("==", 2))
		})

		It("Should return len 3", func() {
			bigint, _ := biginteger.Of("3")

			Expect(bigint.BitLength()).
				To(BeNumerically("==", 2))
		})

		It("Should return len 4", func() {
			bigint, _ := biginteger.Of("4")

			Expect(bigint.BitLength()).
				To(BeNumerically("==", 3))
		})

		It("Should return len 7", func() {
			bigint, _ := biginteger.Of("7")

			Expect(bigint.BitLength()).
				To(BeNumerically("==", 3))
		})

		It("Should return len 8", func() {
			bigint, _ := biginteger.Of("8")

			Expect(bigint.BitLength()).
				To(BeNumerically("==", 4))
		})

		It("Should return len 15", func() {
			bigint, _ := biginteger.Of("15")

			Expect(bigint.BitLength()).
				To(BeNumerically("==", 4))
		})

		It("Should return len 16", func() {
			bigint, _ := biginteger.Of("16")

			Expect(bigint.BitLength()).
				To(BeNumerically("==", 5))
		})

		It("Should return len 31", func() {
			bigint, _ := biginteger.Of("31")

			Expect(bigint.BitLength()).
				To(BeNumerically("==", 5))
		})

		It("Should return len 4294967295", func() {
			bigint, _ := biginteger.Of("4294967295")

			Expect(bigint.BitLength()).
				To(BeNumerically("==", 32))
		})

		It("Should return len 4294967296", func() {
			bigint, _ := biginteger.Of("4294967296")

			Expect(bigint.BitLength()).
				To(BeNumerically("==", 33))
		})

		It("Should return len 18446744073709551615", func() {
			bigint, _ := biginteger.Of("18446744073709551615")

			Expect(bigint.BitLength()).
				To(BeNumerically("==", 64))
		})

		It("Should return len 18446744073709551616", func() {
			bigint, _ := biginteger.Of("18446744073709551616")

			Expect(bigint.BitLength()).
				To(BeNumerically("==", 65))
		})

		It("Should return len 340282366920938463463374607431768211455", func() {
			bigint, _ := biginteger.Of("340282366920938463463374607431768211455")

			Expect(bigint.BitLength()).
				To(BeNumerically("==", 128))
		})

		It("Should return len 340282366920938463463374607431768211456", func() {
			bigint, _ := biginteger.Of("340282366920938463463374607431768211456")

			Expect(bigint.BitLength()).
				To(BeNumerically("==", 129))
		})
	})

	Context("Digits", func() {
		It("Should return 0 for 0", func() {
			bigint, _ := biginteger.Of("0")

			Expect(bigint.Digits()).To(BeNumerically("==", 1))
		})

		It("Should return digits of 1", func() {
			bigint, _ := biginteger.Of("1")

			Expect(bigint.Digits()).To(BeNumerically("==", 1))
		})

		It("Should return digits of -1", func() {
			bigint, _ := biginteger.Of("1")

			Expect(bigint.Digits()).To(BeNumerically("==", 1))
		})

		It("Should return digits of 2", func() {
			bigint, _ := biginteger.Of("2")

			Expect(bigint.Digits()).To(BeNumerically("==", 1))
		})

		It("Should digits of 3", func() {
			bigint, _ := biginteger.Of("3")

			Expect(bigint.Digits()).To(BeNumerically("==", 1))
		})

		It("Should digits of 4", func() {
			bigint, _ := biginteger.Of("4")

			Expect(bigint.Digits()).To(BeNumerically("==", 1))
		})

		It("Should digits of 7", func() {
			bigint, _ := biginteger.Of("7")

			Expect(bigint.Digits()).To(BeNumerically("==", 1))
		})

		It("Should digits of 8", func() {
			bigint, _ := biginteger.Of("8")

			Expect(bigint.Digits()).To(BeNumerically("==", 1))
		})

		It("Should digits of 15", func() {
			bigint, _ := biginteger.Of("15")

			Expect(bigint.Digits()).To(BeNumerically("==", 2))
		})

		It("Should digits of 16", func() {
			bigint, _ := biginteger.Of("16")

			Expect(bigint.Digits()).To(BeNumerically("==", 2))
		})

		It("Should digits of 31", func() {
			bigint, _ := biginteger.Of("31")

			Expect(bigint.Digits()).To(BeNumerically("==", 2))
		})

		It("Should digits of 4294967295", func() {
			bigint, _ := biginteger.Of("4294967295")

			Expect(bigint.Digits()).To(BeNumerically("==", 10))
		})

		It("Should digits of 4294967296", func() {
			bigint, _ := biginteger.Of("4294967296")

			Expect(bigint.Digits()).To(BeNumerically("==", 10))
		})

		It("Should digits of 18446744073709551615", func() {
			bigint, _ := biginteger.Of("18446744073709551615")

			Expect(bigint.Digits()).To(BeNumerically("==", 20))
		})

		It("Should digits of 18446744073709551616", func() {
			bigint, _ := biginteger.Of("18446744073709551616")

			Expect(bigint.Digits()).To(BeNumerically("==", 20))
		})

		It("Should digits of 10000000000", func() {
			bigint, _ := biginteger.Of("10000000000")

			Expect(bigint.Digits()).To(BeNumerically("==", 11))
		})

		It("Should digits of 100000000000000000000", func() {
			bigint, _ := biginteger.Of("100000000000000000000")

			Expect(bigint.Digits()).To(BeNumerically("==", 21))
		})

		It("Should digits of 340282366920938463463374607431768211456", func() {
			bigint, _ := biginteger.Of("340282366920938463463374607431768211456")

			Expect(bigint.Digits()).To(BeNumerically("==", 39))
		})

		It("Should digits of 115792089237316195423570985008687907853269984665640564039457584007913129639936", func() {
			bigint, _ := biginteger.Of("115792089237316195423570985008687907853269984665640564039457584007913129639936")

			Expect(bigint.Digits()).To(BeNumerically("==", 78))
		})
	})

	Context("CompareTo", func() {
		It("Should return a negative value for 1 < 2", func() {
			bigint1, _ := biginteger.Of("1")
			bigint2, _ := biginteger.Of("2")

			Expect(bigint1.CompareTo(*bigint2)).To(BeNumerically("<", 0))
		})

		It("Should return a positive value for 2 < 1", func() {
			bigint1, _ := biginteger.Of("2")
			bigint2, _ := biginteger.Of("1")

			Expect(bigint1.CompareTo(*bigint2)).To(BeNumerically(">", 0))
		})

		It("Should return zero for 2 < 2", func() {
			bigint1, _ := biginteger.Of("2")
			bigint2, _ := biginteger.Of("2")

			Expect(bigint1.CompareTo(*bigint2)).To(BeZero())
		})

		It("Should return a negative value for -1 < 1", func() {
			bigint1, _ := biginteger.Of("-1")
			bigint2, _ := biginteger.Of("1")

			Expect(bigint1.CompareTo(*bigint2)).To(BeNumerically("<", 0))
		})

		It("Should return a positive value for 1 < -1", func() {
			bigint1, _ := biginteger.Of("1")
			bigint2, _ := biginteger.Of("-1")

			Expect(bigint1.CompareTo(*bigint2)).To(BeNumerically(">", 0))
		})

		It("Should return zero for -1 < -1", func() {
			bigint1, _ := biginteger.Of("-1")
			bigint2, _ := biginteger.Of("-1")

			Expect(bigint1.CompareTo(*bigint2)).To(BeZero())
		})

		It("Should return a negative value for -2 < -1", func() {
			bigint1, _ := biginteger.Of("-2")
			bigint2, _ := biginteger.Of("-1")

			Expect(bigint1.CompareTo(*bigint2)).To(BeNumerically("<", 0))
		})

		It("Should return a positive value for -1 < -2", func() {
			bigint1, _ := biginteger.Of("-1")
			bigint2, _ := biginteger.Of("-2")

			Expect(bigint1.CompareTo(*bigint2)).To(BeNumerically(">", 0))
		})

		It("Should return zero for -2 < -2", func() {
			bigint1, _ := biginteger.Of("-2")
			bigint2, _ := biginteger.Of("-2")

			Expect(bigint1.CompareTo(*bigint2)).To(BeZero())
		})

		It("Should return a negative value for -2 < 1", func() {
			bigint1, _ := biginteger.Of("-2")
			bigint2, _ := biginteger.Of("1")

			Expect(bigint1.CompareTo(*bigint2)).To(BeNumerically("<", 0))
		})

		It("Should return a positive value for 1 < -2", func() {
			bigint1, _ := biginteger.Of("1")
			bigint2, _ := biginteger.Of("-2")

			Expect(bigint1.CompareTo(*bigint2)).To(BeNumerically(">", 0))
		})

		It("Should return a negative value for -2 < 2", func() {
			bigint1, _ := biginteger.Of("-2")
			bigint2, _ := biginteger.Of("2")

			Expect(bigint1.CompareTo(*bigint2)).To(BeNumerically("<", 0))
		})

		It("Should return a positive value for 2 < -2", func() {
			bigint1, _ := biginteger.Of("2")
			bigint2, _ := biginteger.Of("-2")

			Expect(bigint1.CompareTo(*bigint2)).To(BeNumerically(">", 0))
		})

		It("Should return zero for 0 < 0", func() {
			bigint1, _ := biginteger.Of("0")
			bigint2, _ := biginteger.Of("0")

			Expect(bigint1.CompareTo(*bigint2)).To(BeZero())
		})

		It("Should return a negative value for 0 < 1", func() {
			bigint1, _ := biginteger.Of("0")
			bigint2, _ := biginteger.Of("1")

			Expect(bigint1.CompareTo(*bigint2)).To(BeNumerically("<", 0))
		})

		It("Should return a positive value for 1 < 0", func() {
			bigint1, _ := biginteger.Of("1")
			bigint2, _ := biginteger.Of("0")

			Expect(bigint1.CompareTo(*bigint2)).To(BeNumerically(">", 0))
		})

		It("Should return a positive value for 0 < -1", func() {
			bigint1, _ := biginteger.Of("0")
			bigint2, _ := biginteger.Of("-1")

			Expect(bigint1.CompareTo(*bigint2)).To(BeNumerically(">", 0))
		})

		It("Should return a negative value for -1 < 0", func() {
			bigint1, _ := biginteger.Of("-1")
			bigint2, _ := biginteger.Of("0")

			Expect(bigint1.CompareTo(*bigint2)).To(BeNumerically("<", 0))
		})

		It("Should return a negative value for 18446744073709551616 < 18446744073709551615", func() {
			bigint1, _ := biginteger.Of("18446744073709551616")
			bigint2, _ := biginteger.Of("18446744073709551615")

			Expect(bigint1.CompareTo(*bigint2)).To(BeNumerically(">", 0))
		})

		It("Should return a positive value for 18446744073709551615 < 18446744073709551616", func() {
			bigint1, _ := biginteger.Of("18446744073709551615")
			bigint2, _ := biginteger.Of("18446744073709551616")

			Expect(bigint1.CompareTo(*bigint2)).To(BeNumerically("<", 0))
		})

		It("Should return zero for 18446744073709551615 < 18446744073709551615", func() {
			bigint1, _ := biginteger.Of("18446744073709551615")
			bigint2, _ := biginteger.Of("18446744073709551615")

			Expect(bigint1.CompareTo(*bigint2)).To(BeZero())
		})

		It("Should return zero for 18446744073709551616 < 18446744073709551616", func() {
			bigint1, _ := biginteger.Of("18446744073709551616")
			bigint2, _ := biginteger.Of("18446744073709551616")

			Expect(bigint1.CompareTo(*bigint2)).To(BeZero())
		})

		It("Should return a positive value for 36893488147419103232 < 18446744073709551615", func() {
			bigint1, _ := biginteger.Of("36893488147419103232")
			bigint2, _ := biginteger.Of("18446744073709551615")

			Expect(bigint1.CompareTo(*bigint2)).To(BeNumerically(">", 0))
		})

		It("Should return a negative value for 18446744073709551615 < 36893488147419103232", func() {
			bigint1, _ := biginteger.Of("18446744073709551615")
			bigint2, _ := biginteger.Of("36893488147419103232")

			Expect(bigint1.CompareTo(*bigint2)).To(BeNumerically("<", 0))
		})

		It("Should return a positive value for 36893488147419103232 < 18446744073709551616", func() {
			bigint1, _ := biginteger.Of("36893488147419103232")
			bigint2, _ := biginteger.Of("18446744073709551616")

			Expect(bigint1.CompareTo(*bigint2)).To(BeNumerically(">", 0))
		})

		It("Should return a negative value for 18446744073709551616 < 36893488147419103232", func() {
			bigint1, _ := biginteger.Of("18446744073709551616")
			bigint2, _ := biginteger.Of("36893488147419103232")

			Expect(bigint1.CompareTo(*bigint2)).To(BeNumerically("<", 0))
		})

		It("Should return a positive value for 73786976294838206464 < 18446744073709551615", func() {
			bigint1, _ := biginteger.Of("73786976294838206464")
			bigint2, _ := biginteger.Of("18446744073709551615")

			Expect(bigint1.CompareTo(*bigint2)).To(BeNumerically(">", 0))
		})

		It("Should return a negative value for 18446744073709551615 < 73786976294838206464", func() {
			bigint1, _ := biginteger.Of("18446744073709551615")
			bigint2, _ := biginteger.Of("73786976294838206464")

			Expect(bigint1.CompareTo(*bigint2)).To(BeNumerically("<", 0))
		})

		It("Should return positive value for 73786976294838206464 < 36893488147419103231", func() {
			bigint1, _ := biginteger.Of("73786976294838206464")
			bigint2, _ := biginteger.Of("36893488147419103231")

			Expect(bigint1.CompareTo(*bigint2)).To(BeNumerically(">", 0))
		})

		It("Should return negative value for 36893488147419103231 < 73786976294838206464", func() {
			bigint1, _ := biginteger.Of("36893488147419103231")
			bigint2, _ := biginteger.Of("73786976294838206464")

			Expect(bigint1.CompareTo(*bigint2)).To(BeNumerically("<", 0))
		})

		It("Should return positive value for 73786976294838206464 < 36893488147419103232", func() {
			bigint1, _ := biginteger.Of("73786976294838206464")
			bigint2, _ := biginteger.Of("36893488147419103232")

			Expect(bigint1.CompareTo(*bigint2)).To(BeNumerically(">", 0))
		})

		It("Should return negative value for 36893488147419103232 < 73786976294838206464", func() {
			bigint1, _ := biginteger.Of("36893488147419103232")
			bigint2, _ := biginteger.Of("73786976294838206464")

			Expect(bigint1.CompareTo(*bigint2)).To(BeNumerically("<", 0))
		})
	})

	Context("IsLessThan", func() {
		It("Should return true for 1 < 2", func() {
			bigint1, _ := biginteger.Of("1")
			bigint2, _ := biginteger.Of("2")

			Expect(bigint1.IsLessThan(*bigint2)).To(BeTrue())
		})

		It("Should return false for 2 < 1", func() {
			bigint1, _ := biginteger.Of("2")
			bigint2, _ := biginteger.Of("1")

			Expect(bigint1.IsLessThan(*bigint2)).To(BeFalse())
		})

		It("Should return false for 2 < 2", func() {
			bigint1, _ := biginteger.Of("2")
			bigint2, _ := biginteger.Of("2")

			Expect(bigint1.IsLessThan(*bigint2)).To(BeFalse())
		})

		It("Should return true for -1 < 1", func() {
			bigint1, _ := biginteger.Of("-1")
			bigint2, _ := biginteger.Of("1")

			Expect(bigint1.IsLessThan(*bigint2)).To(BeTrue())
		})

		It("Should return false for 1 < -1", func() {
			bigint1, _ := biginteger.Of("1")
			bigint2, _ := biginteger.Of("-1")

			Expect(bigint1.IsLessThan(*bigint2)).To(BeFalse())
		})

		It("Should return false for -1 < -1", func() {
			bigint1, _ := biginteger.Of("-1")
			bigint2, _ := biginteger.Of("-1")

			Expect(bigint1.IsLessThan(*bigint2)).To(BeFalse())
		})

		It("Should return true for -2 < -1", func() {
			bigint1, _ := biginteger.Of("-2")
			bigint2, _ := biginteger.Of("-1")

			Expect(bigint1.IsLessThan(*bigint2)).To(BeTrue())
		})

		It("Should return false for -1 < -2", func() {
			bigint1, _ := biginteger.Of("-1")
			bigint2, _ := biginteger.Of("-2")

			Expect(bigint1.IsLessThan(*bigint2)).To(BeFalse())
		})

		It("Should return false for -2 < -2", func() {
			bigint1, _ := biginteger.Of("-2")
			bigint2, _ := biginteger.Of("-2")

			Expect(bigint1.IsLessThan(*bigint2)).To(BeFalse())
		})

		It("Should return true for -2 < 1", func() {
			bigint1, _ := biginteger.Of("-2")
			bigint2, _ := biginteger.Of("1")

			Expect(bigint1.IsLessThan(*bigint2)).To(BeTrue())
		})

		It("Should return false for 1 < -2", func() {
			bigint1, _ := biginteger.Of("1")
			bigint2, _ := biginteger.Of("-2")

			Expect(bigint1.IsLessThan(*bigint2)).To(BeFalse())
		})

		It("Should return false for -2 < 2", func() {
			bigint1, _ := biginteger.Of("-2")
			bigint2, _ := biginteger.Of("2")

			Expect(bigint1.IsLessThan(*bigint2)).To(BeTrue())
		})

		It("Should return false for 2 < -2", func() {
			bigint1, _ := biginteger.Of("2")
			bigint2, _ := biginteger.Of("-2")

			Expect(bigint1.IsLessThan(*bigint2)).To(BeFalse())
		})

		It("Should return false for 0 < 0", func() {
			bigint1, _ := biginteger.Of("0")
			bigint2, _ := biginteger.Of("0")

			Expect(bigint1.IsLessThan(*bigint2)).To(BeFalse())
		})

		It("Should return true for 0 < 1", func() {
			bigint1, _ := biginteger.Of("0")
			bigint2, _ := biginteger.Of("1")

			Expect(bigint1.IsLessThan(*bigint2)).To(BeTrue())
		})

		It("Should return false for 1 < 0", func() {
			bigint1, _ := biginteger.Of("1")
			bigint2, _ := biginteger.Of("0")

			Expect(bigint1.IsLessThan(*bigint2)).To(BeFalse())
		})

		It("Should return false for 0 < -1", func() {
			bigint1, _ := biginteger.Of("0")
			bigint2, _ := biginteger.Of("-1")

			Expect(bigint1.IsLessThan(*bigint2)).To(BeFalse())
		})

		It("Should return true for -1 < 0", func() {
			bigint1, _ := biginteger.Of("-1")
			bigint2, _ := biginteger.Of("0")

			Expect(bigint1.IsLessThan(*bigint2)).To(BeTrue())
		})

		It("Should return false for 18446744073709551616 < 18446744073709551615", func() {
			bigint1, _ := biginteger.Of("18446744073709551616")
			bigint2, _ := biginteger.Of("18446744073709551615")

			Expect(bigint1.IsLessThan(*bigint2)).To(BeFalse())
		})

		It("Should return true for 18446744073709551615 < 18446744073709551616", func() {
			bigint1, _ := biginteger.Of("18446744073709551615")
			bigint2, _ := biginteger.Of("18446744073709551616")

			Expect(bigint1.IsLessThan(*bigint2)).To(BeTrue())
		})

		It("Should return false for 18446744073709551615 < 18446744073709551615", func() {
			bigint1, _ := biginteger.Of("18446744073709551615")
			bigint2, _ := biginteger.Of("18446744073709551615")

			Expect(bigint1.IsLessThan(*bigint2)).To(BeFalse())
		})

		It("Should return false for 18446744073709551616 < 18446744073709551616", func() {
			bigint1, _ := biginteger.Of("18446744073709551616")
			bigint2, _ := biginteger.Of("18446744073709551616")

			Expect(bigint1.IsGreaterThan(*bigint2)).To(BeFalse())
		})

		It("Should return false for 36893488147419103232 < 18446744073709551615", func() {
			bigint1, _ := biginteger.Of("36893488147419103232")
			bigint2, _ := biginteger.Of("18446744073709551615")

			Expect(bigint1.IsLessThan(*bigint2)).To(BeFalse())
		})

		It("Should return true for 18446744073709551615 < 36893488147419103232", func() {
			bigint1, _ := biginteger.Of("18446744073709551615")
			bigint2, _ := biginteger.Of("36893488147419103232")

			Expect(bigint1.IsLessThan(*bigint2)).To(BeTrue())
		})

		It("Should return false for 36893488147419103232 < 18446744073709551616", func() {
			bigint1, _ := biginteger.Of("36893488147419103232")
			bigint2, _ := biginteger.Of("18446744073709551616")

			Expect(bigint1.IsLessThan(*bigint2)).To(BeFalse())
		})

		It("Should return true for 18446744073709551616 < 36893488147419103232", func() {
			bigint1, _ := biginteger.Of("18446744073709551616")
			bigint2, _ := biginteger.Of("36893488147419103232")

			Expect(bigint1.IsLessThan(*bigint2)).To(BeTrue())
		})

		It("Should return false for 73786976294838206464 < 18446744073709551615", func() {
			bigint1, _ := biginteger.Of("73786976294838206464")
			bigint2, _ := biginteger.Of("18446744073709551615")

			Expect(bigint1.IsLessThan(*bigint2)).To(BeFalse())
		})

		It("Should return true for 18446744073709551615 < 73786976294838206464", func() {
			bigint1, _ := biginteger.Of("18446744073709551615")
			bigint2, _ := biginteger.Of("73786976294838206464")

			Expect(bigint1.IsLessThan(*bigint2)).To(BeTrue())
		})

		It("Should return false for 73786976294838206464 < 36893488147419103231", func() {
			bigint1, _ := biginteger.Of("73786976294838206464")
			bigint2, _ := biginteger.Of("36893488147419103231")

			Expect(bigint1.IsLessThan(*bigint2)).To(BeFalse())
		})

		It("Should return true for 36893488147419103231 < 73786976294838206464", func() {
			bigint1, _ := biginteger.Of("36893488147419103231")
			bigint2, _ := biginteger.Of("73786976294838206464")

			Expect(bigint1.IsLessThan(*bigint2)).To(BeTrue())
		})

		It("Should return false for 73786976294838206464 < 36893488147419103232", func() {
			bigint1, _ := biginteger.Of("73786976294838206464")
			bigint2, _ := biginteger.Of("36893488147419103232")

			Expect(bigint1.IsLessThan(*bigint2)).To(BeFalse())
		})

		It("Should return true for 36893488147419103232 < 73786976294838206464", func() {
			bigint1, _ := biginteger.Of("36893488147419103232")
			bigint2, _ := biginteger.Of("73786976294838206464")

			Expect(bigint1.IsLessThan(*bigint2)).To(BeTrue())
		})
	})

	Context("IsGreaterThan", func() {
		It("Should return false for 1 > 2", func() {
			bigint1, _ := biginteger.Of("1")
			bigint2, _ := biginteger.Of("2")

			Expect(bigint1.IsGreaterThan(*bigint2)).To(BeFalse())
		})

		It("Should return true for 2 > 1", func() {
			bigint1, _ := biginteger.Of("2")
			bigint2, _ := biginteger.Of("1")

			Expect(bigint1.IsGreaterThan(*bigint2)).To(BeTrue())
		})

		It("Should return false for 2 > 2", func() {
			bigint1, _ := biginteger.Of("2")
			bigint2, _ := biginteger.Of("2")

			Expect(bigint1.IsGreaterThan(*bigint2)).To(BeFalse())
		})

		It("Should return false for -1 > 1", func() {
			bigint1, _ := biginteger.Of("-1")
			bigint2, _ := biginteger.Of("1")

			Expect(bigint1.IsGreaterThan(*bigint2)).To(BeFalse())
		})

		It("Should return true for 1 > -1", func() {
			bigint1, _ := biginteger.Of("1")
			bigint2, _ := biginteger.Of("-1")

			Expect(bigint1.IsGreaterThan(*bigint2)).To(BeTrue())
		})

		It("Should return false for -1 > -1", func() {
			bigint1, _ := biginteger.Of("-1")
			bigint2, _ := biginteger.Of("-1")

			Expect(bigint1.IsGreaterThan(*bigint2)).To(BeFalse())
		})

		It("Should return false for -2 > -1", func() {
			bigint1, _ := biginteger.Of("-2")
			bigint2, _ := biginteger.Of("-1")

			Expect(bigint1.IsGreaterThan(*bigint2)).To(BeFalse())
		})

		It("Should return true for -1 > -2", func() {
			bigint1, _ := biginteger.Of("-1")
			bigint2, _ := biginteger.Of("-2")

			Expect(bigint1.IsGreaterThan(*bigint2)).To(BeTrue())
		})

		It("Should return true for 18446744073709551616 > 18446744073709551615", func() {
			bigint1, _ := biginteger.Of("18446744073709551616")
			bigint2, _ := biginteger.Of("18446744073709551615")

			Expect(bigint1.IsGreaterThan(*bigint2)).To(BeTrue())
		})

		It("Should return false for 18446744073709551615 > 18446744073709551616", func() {
			bigint1, _ := biginteger.Of("18446744073709551615")
			bigint2, _ := biginteger.Of("18446744073709551616")

			Expect(bigint1.IsGreaterThan(*bigint2)).To(BeFalse())
		})

		It("Should return false for 18446744073709551615 > 18446744073709551615", func() {
			bigint1, _ := biginteger.Of("18446744073709551615")
			bigint2, _ := biginteger.Of("18446744073709551615")

			Expect(bigint1.IsGreaterThan(*bigint2)).To(BeFalse())
		})

		It("Should return false for 18446744073709551616 > 18446744073709551616", func() {
			bigint1, _ := biginteger.Of("18446744073709551616")
			bigint2, _ := biginteger.Of("18446744073709551616")

			Expect(bigint1.IsGreaterThan(*bigint2)).To(BeFalse())
		})

		It("Should return true for 36893488147419103232 > 18446744073709551615", func() {
			bigint1, _ := biginteger.Of("36893488147419103232")
			bigint2, _ := biginteger.Of("18446744073709551615")

			Expect(bigint1.IsGreaterThan(*bigint2)).To(BeTrue())
		})

		It("Should return false for 18446744073709551615 > 36893488147419103232", func() {
			bigint1, _ := biginteger.Of("18446744073709551615")
			bigint2, _ := biginteger.Of("36893488147419103232")

			Expect(bigint1.IsGreaterThan(*bigint2)).To(BeFalse())
		})

		It("Should return true for 36893488147419103232 > 18446744073709551616", func() {
			bigint1, _ := biginteger.Of("36893488147419103232")
			bigint2, _ := biginteger.Of("18446744073709551616")

			Expect(bigint1.IsGreaterThan(*bigint2)).To(BeTrue())
		})

		It("Should return false for 18446744073709551616 > 36893488147419103232", func() {
			bigint1, _ := biginteger.Of("18446744073709551616")
			bigint2, _ := biginteger.Of("36893488147419103232")

			Expect(bigint1.IsGreaterThan(*bigint2)).To(BeFalse())
		})

		It("Should return true for 73786976294838206464 > 18446744073709551615", func() {
			bigint1, _ := biginteger.Of("73786976294838206464")
			bigint2, _ := biginteger.Of("18446744073709551615")

			Expect(bigint1.IsGreaterThan(*bigint2)).To(BeTrue())
		})

		It("Should return false for 18446744073709551615 > 73786976294838206464", func() {
			bigint1, _ := biginteger.Of("18446744073709551615")
			bigint2, _ := biginteger.Of("73786976294838206464")

			Expect(bigint1.IsGreaterThan(*bigint2)).To(BeFalse())
		})

		It("Should return true for 73786976294838206464 > 36893488147419103231", func() {
			bigint1, _ := biginteger.Of("73786976294838206464")
			bigint2, _ := biginteger.Of("36893488147419103231")

			Expect(bigint1.IsGreaterThan(*bigint2)).To(BeTrue())
		})

		It("Should return false for 36893488147419103231 > 73786976294838206464", func() {
			bigint1, _ := biginteger.Of("36893488147419103231")
			bigint2, _ := biginteger.Of("73786976294838206464")

			Expect(bigint1.IsGreaterThan(*bigint2)).To(BeFalse())
		})

		It("Should return true for 73786976294838206464 > 36893488147419103232", func() {
			bigint1, _ := biginteger.Of("73786976294838206464")
			bigint2, _ := biginteger.Of("36893488147419103232")

			Expect(bigint1.IsGreaterThan(*bigint2)).To(BeTrue())
		})

		It("Should return false for 36893488147419103232 > 73786976294838206464", func() {
			bigint1, _ := biginteger.Of("36893488147419103232")
			bigint2, _ := biginteger.Of("73786976294838206464")

			Expect(bigint1.IsGreaterThan(*bigint2)).To(BeFalse())
		})
	})

	Context("IsEqualTo", func() {
		It("Should return true for 1 == 1", func() {
			bigint1, _ := biginteger.Of("1")
			bigint2, _ := biginteger.Of("1")

			Expect(bigint1.IsEqualTo(*bigint2)).To(BeTrue())
		})

		It("Should return false for 1 == 2", func() {
			bigint1, _ := biginteger.Of("1")
			bigint2, _ := biginteger.Of("2")

			Expect(bigint1.IsEqualTo(*bigint2)).To(BeFalse())
		})

		It("Should return false for 1 == -1", func() {
			bigint1, _ := biginteger.Of("1")
			bigint2, _ := biginteger.Of("-1")

			Expect(bigint1.IsEqualTo(*bigint2)).To(BeFalse())
		})

		It("Should return true for -1 == -1", func() {
			bigint1, _ := biginteger.Of("-1")
			bigint2, _ := biginteger.Of("-1")

			Expect(bigint1.IsEqualTo(*bigint2)).To(BeTrue())
		})
	})

	Context("Log2", func() {
		It("Should return 0 for 1", func() {
			bigint, _ := biginteger.Of("1")

			Expect(bigint.Log2()).To(BeNumerically("==", 0.0))
		})

		It("Should return 1 for 2", func() {
			bigint, _ := biginteger.Of("2")

			Expect(bigint.Log2()).To(BeNumerically("==", 1.0))
		})

		It("Should return 1 for 3", func() {
			bigint, _ := biginteger.Of("3")

			Expect(bigint.Log2()).To(BeNumerically("==", 1.584962500721156))
		})

		It("Should return 2 for 4", func() {
			bigint, _ := biginteger.Of("4")

			Expect(bigint.Log2()).To(BeNumerically("==", 2.0))
		})

		It("Should return 2 for 5", func() {
			bigint, _ := biginteger.Of("5")

			Expect(bigint.Log2()).To(BeNumerically("==", 2.321928094887362))
		})

		It("Should return 2 for 6", func() {
			bigint, _ := biginteger.Of("6")

			Expect(bigint.Log2()).To(BeNumerically("==", 2.584962500721156))
		})

		It("Should return 2 for 7", func() {
			bigint, _ := biginteger.Of("7")

			Expect(bigint.Log2()).To(BeNumerically("==", 2.807354922057604))
		})

		It("Should return 3 for 8", func() {
			bigint, _ := biginteger.Of("8")

			Expect(bigint.Log2()).To(BeNumerically("==", 3.0))
		})

		It("Should return 3 for 9", func() {
			bigint, _ := biginteger.Of("9")

			Expect(bigint.Log2()).To(BeNumerically("==", 3.1699250014423126))
		})

		It("Should return 3 for 10", func() {
			bigint, _ := biginteger.Of("10")

			Expect(bigint.Log2()).To(BeNumerically("==", 3.321928094887362))
		})

		It("Should return 3 for 11", func() {
			bigint, _ := biginteger.Of("11")

			Expect(bigint.Log2()).To(BeNumerically("==", 3.4594316186372973))
		})

		It("Should return 3 for 12", func() {
			bigint, _ := biginteger.Of("12")

			Expect(bigint.Log2()).To(BeNumerically("==", 3.584962500721156))
		})

		It("Should return 5 for 32", func() {
			bigint, _ := biginteger.Of("32")

			Expect(bigint.Log2()).To(BeNumerically("==", 5.0))
		})

		It("Should return 5 for 33", func() {
			bigint, _ := biginteger.Of("33")

			Expect(bigint.Log2()).To(BeNumerically("==", 5.044394119358453))
		})

		It("Should return 6 for 64", func() {
			bigint, _ := biginteger.Of("64")

			Expect(bigint.Log2()).To(BeNumerically("==", 6.0))
		})

		It("Should return 6 for 65", func() {
			bigint, _ := biginteger.Of("65")

			Expect(bigint.Log2()).To(BeNumerically("==", 6.022367813028454))
		})

		It("Should return 9 for 1023", func() {
			bigint, _ := biginteger.Of("1023")

			Expect(bigint.Log2()).To(BeNumerically("==", 9.99859042974533))
		})

		It("Should return 10 for 1024", func() {
			bigint, _ := biginteger.Of("1024")

			Expect(bigint.Log2()).To(BeNumerically("==", 10.0))
		})

		It("Should return 13 for 12345", func() {
			bigint, _ := biginteger.Of("12345")

			Expect(bigint.Log2()).To(BeNumerically("==", 13.591639216030144))
		})

		It("Should return 16 for 123456", func() {
			bigint, _ := biginteger.Of("123456")

			Expect(bigint.Log2()).To(BeNumerically("==", 16.913637428049103))
		})

		It("Should return 20 for 1234567", func() {
			bigint, _ := biginteger.Of("1234567")

			Expect(bigint.Log2()).To(BeNumerically("==", 20.23557370304651))
		})

		It("Should return 23 for 12345678", func() {
			bigint, _ := biginteger.Of("12345678")

			Expect(bigint.Log2()).To(BeNumerically("==", 23.55750273280064))
		})

		It("Should return 26 for 123456789", func() {
			bigint, _ := biginteger.Of("123456789")

			Expect(bigint.Log2()).To(BeNumerically("==", 26.879430932860473))
		})

		It("Should return 30 for 1234567890", func() {
			bigint, _ := biginteger.Of("1234567890")

			Expect(bigint.Log2()).To(BeNumerically("==", 30.201359027747838))
		})

		It("Should return 33 for 12345678901", func() {
			bigint, _ := biginteger.Of("12345678901")

			Expect(bigint.Log2()).To(BeNumerically("==", 33.52328712275206))
		})

		It("Should return 36 for 123456789012", func() {
			bigint, _ := biginteger.Of("123456789012")

			Expect(bigint.Log2()).To(BeNumerically("==", 36.84521521766279))
		})

		It("Should return 40 for 1234567890123", func() {
			bigint, _ := biginteger.Of("1234567890123")

			Expect(bigint.Log2()).To(BeNumerically("==", 40.16714331255366))
		})

		It("Should return 43 for 12345678901234", func() {
			bigint, _ := biginteger.Of("12345678901234")

			Expect(bigint.Log2()).To(BeNumerically("==", 43.489071407441486))
		})

		It("log2(100000000000000000000)", func() {
			bigint, _ := biginteger.Of("100000000000000000000")

			Expect(bigint.Log2()).To(BeNumerically("==", 66.438561897747))
		})

		It("log2([0, 1])", func() {
			bigint := biginteger.OfUint64Array([]uint64{0, 1})

			Expect(bigint.Log2()).To(BeNumerically("==", 64.0))
		})
		
		It("log2([0, 2])", func() {
			bigint := biginteger.OfUint64Array([]uint64{0, 2})

			Expect(bigint.Log2()).To(BeNumerically("==", 65.0))
		})
	})

	Context("Log10", func() {
		It("Should return 0 for 1", func() {
			bigint, _ := biginteger.Of("1")

			Expect(bigint.Log10()).To(BeNumerically("==", 0))
		})

		It("Should return 0 for 9", func() {
			bigint, _ := biginteger.Of("9")

			Expect(bigint.Log10()).To(BeNumerically("==", 0.954242509439325))
		})

		It("Should return 1 for 10", func() {
			bigint, _ := biginteger.Of("10")

			Expect(bigint.Log10()).To(BeNumerically("==", 1.0))
		})

		It("Should return 1 for 99", func() {
			bigint, _ := biginteger.Of("99")

			Expect(bigint.Log10()).To(BeNumerically("==", 1.99563519459755))
		})

		It("Should return 2 for 100", func() {
			bigint, _ := biginteger.Of("100")

			Expect(bigint.Log10()).To(BeNumerically("==", 2.0))
		})

		It("Should return 2 for 999", func() {
			bigint, _ := biginteger.Of("999")

			Expect(bigint.Log10()).To(BeNumerically("==", 2.999565488225983))
		})

		It("Should return 3 for 1000", func() {
			bigint, _ := biginteger.Of("1000")

			Expect(bigint.Log10()).To(BeNumerically("==", 3.0))
		})

		It("Should return 3 for 9999", func() {
			bigint, _ := biginteger.Of("9999")

			Expect(bigint.Log10()).To(BeNumerically("==", 3.9999565683801928))
		})

		It("Should return 4 for 10000", func() {
			bigint, _ := biginteger.Of("10000")

			Expect(bigint.Log10()).To(BeNumerically("==", 4.0))
		})

		It("Should return 4 for 99999", func() {
			bigint, _ := biginteger.Of("99999")

			Expect(bigint.Log10()).To(BeNumerically("==", 4.999995657033466))
		})

		It("Should return 5 for 100000", func() {
			bigint, _ := biginteger.Of("100000")

			Expect(bigint.Log10()).To(BeNumerically("==", 5))
		})

		It("Should return 5 for 999999", func() {
			bigint, _ := biginteger.Of("999999")

			Expect(bigint.Log10()).To(BeNumerically("==", 5.999999565705301))
		})

		It("Should return 6 for 1000000", func() {
			bigint, _ := biginteger.Of("1000000")

			Expect(bigint.Log10()).To(BeNumerically("==", 6))
		})

		It("Should return 20 for 100000000000000000000", func() {
			bigint, _ := biginteger.Of("100000000000000000000")

			Expect(bigint.Log10()).To(BeNumerically("==", 20.0))
		})
	})

	Context("Log", func() {
		It("753457234543.log(2)", func() {
			bigint, _ := biginteger.Of("753457234543")
			base, _ := biginteger.Of("2")

			Expect(bigint.Log(*base)).To(BeNumerically("==", 39.45473467202251))
		})

		It("753457234543.log(3)", func() {
			bigint, _ := biginteger.Of("753457234543")
			base, _ := biginteger.Of("3")

			Expect(bigint.Log(*base)).To(BeNumerically("==", 24.8931660238464))
		})

		It("753457234543.log(4)", func() {
			bigint, _ := biginteger.Of("753457234543")
			base, _ := biginteger.Of("4")

			Expect(bigint.Log(*base)).To(BeNumerically("==", 19.727367336011255))
		})

		It("753457234543.log(5)", func() {
			bigint, _ := biginteger.Of("753457234543")
			base, _ := biginteger.Of("5")

			Expect(bigint.Log(*base)).To(BeNumerically("==", 16.992229328245617))
		})

		It("753457234543.log(6)", func() {
			bigint, _ := biginteger.Of("753457234543")
			base, _ := biginteger.Of("6")

			Expect(bigint.Log(*base)).To(BeNumerically("==", 15.26317486656591))
		})

		It("753457234543.log(7)", func() {
			bigint, _ := biginteger.Of("753457234543")
			base, _ := biginteger.Of("7")

			Expect(bigint.Log(*base)).To(BeNumerically("==", 14.054060055614492))
		})

		It("753457234543.log(8)", func() {
			bigint, _ := biginteger.Of("753457234543")
			base, _ := biginteger.Of("8")

			Expect(bigint.Log(*base)).To(BeNumerically("==", 13.151578224007503))
		})
	})
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "biginteger Test Suite")
}
