package biginteger_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go-biginteger/biginteger"
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
	})

	Context("Add", func() {
		It("Should add 1 and 2", func() {
			bigint1, _ := biginteger.Of("1")
			bigint2, _ := biginteger.Of("2")

			Expect(bigint1.Add(*bigint2).String()).To(Equal("3"))
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

			Expect(bigint1.Add(*bigint2).String()).To(Equal("20000000000000000000"))
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
	})

	Context("Multiply", func() {
		It("Should multiply 2 by 1", func() {
			bigint1, _ := biginteger.Of("2")
			bigint2, _ := biginteger.Of("1")

			Expect(bigint1.Multiply(*bigint2).String()).To(Equal("2"))
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
	})

	Context("Divide", func() {
		It("Should divide 2 by 2", func() {
			bigint1, _ := biginteger.Of("2")
			bigint2, _ := biginteger.Of("2")

			Expect(bigint1.Divide(*bigint2).String()).To(Equal("1"))
		})

		It("Should divide 4 by 2", func() {
			bigint1, _ := biginteger.Of("4")
			bigint2, _ := biginteger.Of("2")

			Expect(bigint1.Divide(*bigint2).String()).To(Equal("2"))
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

		It("Should divide 18446744073709551615 by 18446744073709551616", func() {
			bigint1, _ := biginteger.Of("18446744073709551615")
			bigint2, _ := biginteger.Of("18446744073709551616")

			Expect(bigint1.Divide(*bigint2).String()).To(Equal("0"))
		})

		It("Should divide 18446744073709551616 by 18446744073709551615", func() {
			bigint1, _ := biginteger.Of("18446744073709551616")
			bigint2, _ := biginteger.Of("18446744073709551615")

			Expect(bigint1.Divide(*bigint2).String()).To(Equal("1"))
		})

		It("Should divide 18446744073709551616 by 2", func() {
			bigint1, _ := biginteger.Of("18446744073709551616")
			bigint2, _ := biginteger.Of("2")

			Expect(bigint1.Divide(*bigint2).String()).To(Equal("9223372036854775808"))
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
	})

	Context("Modulo", func() {
		It("Should return 0 for 1 % 1", func() {
			bigint1, _ := biginteger.Of("1")
			bigint2, _ := biginteger.Of("1")

			Expect(bigint1.Modulo(*bigint2).String()).To(Equal("0"))
		})

		It("Should return 0 for 2 % 1", func() {
			bigint1, _ := biginteger.Of("2")
			bigint2, _ := biginteger.Of("1")

			Expect(bigint1.Modulo(*bigint2).String()).To(Equal("0"))
		})

		It("Should return 1 for 1 % 2", func() {
			bigint1, _ := biginteger.Of("1")
			bigint2, _ := biginteger.Of("2")

			Expect(bigint1.Modulo(*bigint2).String()).To(Equal("1"))
		})

		It("Should return 0 for 0 % 2", func() {
			bigint1, _ := biginteger.Of("0")
			bigint2, _ := biginteger.Of("2")

			Expect(bigint1.Modulo(*bigint2).String()).To(Equal("0"))
		})

		It("Should return 0 for 2 % 2", func() {
			bigint1, _ := biginteger.Of("2")
			bigint2, _ := biginteger.Of("2")

			Expect(bigint1.Modulo(*bigint2).String()).To(Equal("0"))
		})

		It("Should return 0 for 4 % 2", func() {
			bigint1, _ := biginteger.Of("4")
			bigint2, _ := biginteger.Of("2")

			Expect(bigint1.Modulo(*bigint2).String()).To(Equal("0"))
		})

		It("Should return 0 for 8 % 4", func() {
			bigint1, _ := biginteger.Of("8")
			bigint2, _ := biginteger.Of("4")

			Expect(bigint1.Modulo(*bigint2).String()).To(Equal("0"))
		})

		It("Should return 0 for 4294967296 % 4294967296", func() {
			bigint1, _ := biginteger.Of("4294967296")
			bigint2, _ := biginteger.Of("4294967296")

			Expect(bigint1.Modulo(*bigint2).String()).To(Equal("0"))
		})

		It("Should return 0 for 18446744073709551615 % 1", func() {
			bigint1, _ := biginteger.Of("18446744073709551615")
			bigint2, _ := biginteger.Of("1")

			Expect(bigint1.Modulo(*bigint2).String()).To(Equal("0"))
		})

		It("Should return 1 for 18446744073709551615 % 2", func() {
			bigint1, _ := biginteger.Of("18446744073709551615")
			bigint2, _ := biginteger.Of("2")

			Expect(bigint1.Modulo(*bigint2).String()).To(Equal("1"))
		})

		It("Should return 18446744073709551615 for 18446744073709551615 % 18446744073709551616", func() {
			bigint1, _ := biginteger.Of("18446744073709551615")
			bigint2, _ := biginteger.Of("18446744073709551616")

			Expect(bigint1.Modulo(*bigint2).String()).To(Equal("18446744073709551615"))
		})

		It("Should return 1 for 18446744073709551616 % 18446744073709551615", func() {
			bigint1, _ := biginteger.Of("18446744073709551616")
			bigint2, _ := biginteger.Of("18446744073709551615")

			Expect(bigint1.Modulo(*bigint2).String()).To(Equal("1"))
		})

		It("Should return 1 for 18446744073709551616 % 10", func() {
			bigint1, _ := biginteger.Of("18446744073709551616")
			bigint2, _ := biginteger.Of("10")
			result := bigint1.Modulo(*bigint2)

			Expect(result.String()).To(Equal("6"))
		})

		It("Should return 1 for 36893488147419103230 % 10", func() {
			bigint1, _ := biginteger.Of("36893488147419103230")
			bigint2, _ := biginteger.Of("10")
			result := bigint1.Modulo(*bigint2)

			Expect(result.String()).To(Equal("0"))
		})

		It("Should return 1 for 36893488147419103232 % 10", func() {
			bigint1, _ := biginteger.Of("36893488147419103232")
			bigint2, _ := biginteger.Of("10")
			result := bigint1.Modulo(*bigint2)

			Expect(result.String()).To(Equal("2"))
		})

		It("Should return 0 for 340282366920938463444927863358058659840 % 10", func() {
			bigint1, _ := biginteger.Of("340282366920938463444927863358058659840")
			bigint2, _ := biginteger.Of("10")
			result := bigint1.Modulo(*bigint2)

			Expect(result.String()).To(Equal("0"))
		})

		It("Should return 40 for 340282366920938463444927863358058659840 % 100", func() {
			bigint1, _ := biginteger.Of("340282366920938463444927863358058659840")
			bigint2, _ := biginteger.Of("100")
			result := bigint1.Modulo(*bigint2)

			Expect(result.String()).To(Equal("40"))
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

		It("Should return 4294 967296 for 2 pow 32", func() {
			bigint1, _ := biginteger.Of("2")
			bigint2, _ := biginteger.Of("32")

			Expect(bigint1.Power(*bigint2).String()).To(Equal("4294967296"))
		})

		It("Should return 18446744073709551616 for 2 pow 64", func() {
			bigint1, _ := biginteger.Of("2")
			bigint2, _ := biginteger.Of("64")

			Expect(bigint1.Power(*bigint2).String()).To(Equal("18446744073709551616"))
		})
	})

	Context("ShiftRight", func() {
		It("Should return 0 >> 0", func() {
			bigint, _ := biginteger.Of("0")

			Expect(bigint.ShiftRight(biginteger.Zero).String()).
				To(Equal("0"))
		})

		It("Should return 0 >> 1", func() {
			bigint, _ := biginteger.Of("0")

			Expect(bigint.ShiftRight(biginteger.One).String()).
				To(Equal("0"))
		})

		It("Should return 1 >> 0", func() {
			bigint, _ := biginteger.Of("1")

			Expect(bigint.ShiftRight(biginteger.Zero).String()).To(Equal("1"))
		})

		It("Should return 1 >> 1", func() {
			bigint, _ := biginteger.Of("1")

			Expect(bigint.ShiftRight(biginteger.One).String()).To(Equal("0"))
		})

		It("Should return 2 >> 0", func() {
			bigint, _ := biginteger.Of("2")

			Expect(bigint.ShiftRight(biginteger.Zero).String()).To(Equal("2"))
		})

		It("Should return 2 >> 1", func() {
			bigint, _ := biginteger.Of("2")

			Expect(bigint.ShiftRight(biginteger.One).String()).To(Equal("1"))
		})

		It("Should return 2 >> 2", func() {
			bigint, _ := biginteger.Of("2")

			Expect(bigint.ShiftRight(biginteger.Two).String()).To(Equal("0"))
		})

		It("Should return 18446744073709551615 >> 0", func() {
			bigint, _ := biginteger.Of("18446744073709551615")

			Expect(bigint.ShiftRight(biginteger.Zero).String()).
				To(Equal("18446744073709551615"))
		})

		It("Should return 18446744073709551615 >> 1", func() {
			bigint, _ := biginteger.Of("18446744073709551615")

			Expect(bigint.ShiftRight(biginteger.One).String()).
				To(Equal("9223372036854775807"))
		})

		It("Should return 18446744073709551615 >> 2", func() {
			bigint, _ := biginteger.Of("18446744073709551615")

			Expect(bigint.ShiftRight(biginteger.Two).String()).
				To(Equal("4611686018427387903"))
		})

		It("Should return 18446744073709551615 >> 16", func() {
			bigint1, _ := biginteger.Of("18446744073709551615")
			bigint2, _ := biginteger.Of("16")

			Expect(bigint1.ShiftRight(*bigint2).String()).
				To(Equal("281474976710655"))
		})

		It("Should return 18446744073709551615 >> 32", func() {
			bigint1, _ := biginteger.Of("18446744073709551615")
			bigint2, _ := biginteger.Of("32")

			Expect(bigint1.ShiftRight(*bigint2).String()).
				To(Equal("4294967295"))
		})

		It("Should return 18446744073709551615 >> 63", func() {
			bigint1, _ := biginteger.Of("18446744073709551615")
			bigint2, _ := biginteger.Of("63")

			Expect(bigint1.ShiftRight(*bigint2).String()).
				To(Equal("1"))
		})

		It("Should return 18446744073709551615 >> 64", func() {
			bigint1, _ := biginteger.Of("18446744073709551615")
			bigint2, _ := biginteger.Of("64")

			Expect(bigint1.ShiftRight(*bigint2).String()).
				To(Equal("0"))
		})
	})

	Context("ShiftLeft", func() {
		It("Should return 0 << 0", func() {
			bigint, _ := biginteger.Of("0")

			Expect(bigint.ShiftLeft(biginteger.Zero).String()).
				To(Equal("0"))
		})

		It("Should return 0 << 1000000", func() {
			bigint, _ := biginteger.Of("0")
			bigint2, _ := biginteger.Of("1000000")

			Expect(bigint.ShiftLeft(*bigint2).String()).
				To(Equal("0"))
		})

		It("Should return 1 << 0", func() {
			bigint, _ := biginteger.Of("1")

			Expect(bigint.ShiftLeft(biginteger.Zero).String()).
				To(Equal("1"))
		})

		It("Should return 2 << 0", func() {
			bigint, _ := biginteger.Of("2")

			Expect(bigint.ShiftLeft(biginteger.Zero).String()).
				To(Equal("2"))
		})

		It("Should return 3 << 0", func() {
			bigint, _ := biginteger.Of("3")

			Expect(bigint.ShiftLeft(biginteger.Zero).String()).
				To(Equal("3"))
		})

		It("Should return 2 << 1", func() {
			bigint, _ := biginteger.Of("2")
			bigint2, _ := biginteger.Of("1")

			Expect(bigint.ShiftLeft(*bigint2).String()).
				To(Equal("4"))
		})

		It("Should return 2 << 2", func() {
			bigint, _ := biginteger.Of("2")
			bigint2, _ := biginteger.Of("2")

			Expect(bigint.ShiftLeft(*bigint2).String()).
				To(Equal("8"))
		})

		It("Should return 2 << 3", func() {
			bigint, _ := biginteger.Of("2")
			bigint2, _ := biginteger.Of("3")

			Expect(bigint.ShiftLeft(*bigint2).String()).
				To(Equal("16"))
		})

		It("Should return 1 << 32", func() {
			bigint, _ := biginteger.Of("1")
			bigint2, _ := biginteger.Of("32")
			result := bigint.ShiftLeft(*bigint2)

			Expect(result.String()).
				To(Equal("4294967296"))
		})

		It("Should return 1 << 62", func() {
			bigint, _ := biginteger.Of("1")
			bigint2, _ := biginteger.Of("62")
			result := bigint.ShiftLeft(*bigint2)

			Expect(result.String()).
				To(Equal("4611686018427387904"))
		})

		It("Should return 1 << 63", func() {
			bigint, _ := biginteger.Of("1")
			bigint2, _ := biginteger.Of("63")
			result := bigint.ShiftLeft(*bigint2)

			Expect(result.String()).
				To(Equal("9223372036854775808"))
		})

		It("Should return 1 << 64", func() {
			bigint, _ := biginteger.Of("1")
			bigint2, _ := biginteger.Of("64")

			Expect(bigint.ShiftLeft(*bigint2).String()).
				To(Equal("18446744073709551616"))
		})

		It("Should return 1 << 65", func() {
			bigint, _ := biginteger.Of("1")
			bigint2, _ := biginteger.Of("65")

			Expect(bigint.ShiftLeft(*bigint2).String()).
				To(Equal("36893488147419103232"))
		})

		It("Should return 9223372036854775808 << 1", func() {
			bigint, _ := biginteger.Of("9223372036854775808")
			bigint2, _ := biginteger.Of("1")
			result := bigint.ShiftLeft(*bigint2)

			Expect(result.String()).
				To(Equal("18446744073709551616"))
		})
	})

	Context("BigLength", func() {
		It("Should return len 0", func() {
			bigint, _ := biginteger.Of("0")

			Expect(bigint.BitLength().String()).
				To(Equal("1"))
		})

		It("Should return len 1", func() {
			bigint, _ := biginteger.Of("1")

			Expect(bigint.BitLength().String()).
				To(Equal("1"))
		})

		It("Should return len 2", func() {
			bigint, _ := biginteger.Of("2")

			Expect(bigint.BitLength().String()).
				To(Equal("2"))
		})

		It("Should return len 3", func() {
			bigint, _ := biginteger.Of("3")

			Expect(bigint.BitLength().String()).
				To(Equal("2"))
		})

		It("Should return len 4", func() {
			bigint, _ := biginteger.Of("4")

			Expect(bigint.BitLength().String()).
				To(Equal("3"))
		})

		It("Should return len 7", func() {
			bigint, _ := biginteger.Of("7")

			Expect(bigint.BitLength().String()).
				To(Equal("3"))
		})

		It("Should return len 8", func() {
			bigint, _ := biginteger.Of("8")

			Expect(bigint.BitLength().String()).
				To(Equal("4"))
		})

		It("Should return len 15", func() {
			bigint, _ := biginteger.Of("15")

			Expect(bigint.BitLength().String()).
				To(Equal("4"))
		})

		It("Should return len 16", func() {
			bigint, _ := biginteger.Of("16")

			Expect(bigint.BitLength().String()).
				To(Equal("5"))
		})

		It("Should return len 31", func() {
			bigint, _ := biginteger.Of("31")

			Expect(bigint.BitLength().String()).
				To(Equal("5"))
		})

		It("Should return len 4294967295", func() {
			bigint, _ := biginteger.Of("4294967295")

			Expect(bigint.BitLength().String()).
				To(Equal("32"))
		})

		It("Should return len 4294967296", func() {
			bigint, _ := biginteger.Of("4294967296")

			Expect(bigint.BitLength().String()).
				To(Equal("33"))
		})

		It("Should return len 18446744073709551615", func() {
			bigint, _ := biginteger.Of("18446744073709551615")

			Expect(bigint.BitLength().String()).
				To(Equal("64"))
		})

		It("Should return len 18446744073709551616", func() {
			bigint, _ := biginteger.Of("18446744073709551616")

			Expect(bigint.BitLength().String()).
				To(Equal("65"))
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
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Numbers Test Suite")
}
