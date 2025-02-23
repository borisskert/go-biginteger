package multiply_test

import (
	"github.com/borisskert/go-biginteger"
	"github.com/borisskert/go-biginteger/digits"
	"github.com/borisskert/go-biginteger/multiply"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"testing"
)

var _ = Describe("Multiply", func() {
	It("Should multiply 2 digit numbers (Example a, v0)", func() {
		p0 := biginteger.OfUint64Array([]uint64{
			6742305324661190591, 12524700037052152845,
		})
		q0 := biginteger.OfUint64Array([]uint64{
			11595963786453332908, 3312835375285772574,
		})

		expectedResult := biginteger.OfUint64Array([]uint64{
			9350441601238114644,
			12018134535777490305,
			327839218333391898,
			2249300428400506475,
		})

		result := p0.Multiply(q0)

		Expect(result.IsEqualTo(expectedResult)).To(BeTrue())
	})

	It("Should multiply 2 digit numbers (Example a, v1)", func() {
		p1 := digits.OfUint64Array([]uint64{
			15340239725253121661,
			7141458029417713389,
			1,
		})
		q1 := digits.OfUint64Array([]uint64{
			11431115928381940648, 16206270981547339579,
		})

		expectedResult := digits.OfUint64Array([]uint64{
			14427007699956832520,
			6648931555219967304,
			14987701824921030509,
			4033609941983580208,
			1,
		})

		result := multiply.MultiplySwitch(p1, q1)

		Expect(result.IsEqualTo(expectedResult)).To(BeTrue())
	})

	It("Should multiply 2 digit numbers (Example a, vm1)", func() {
		pm1 := digits.OfUint64Array([]uint64{
			1855629075930740479, 538802029022959315,
		}).Negate()
		qm1 := digits.OfUint64Array([]uint64{
			6685932429184826448, 9580600230975794430,
		}).Negative()

		expectedResult := digits.OfUint64Array([]uint64{
			1965498881762243504,
			74293621570320226,
			9445435595678237052,
			279835120120985550,
		})

		result := multiply.MultiplySwitch(pm1, qm1)

		Expect(result.IsEqualTo(expectedResult)).To(BeTrue())
	})

	It("Should multiply 2 digit numbers (Example a, vm2)", func() {
		pm2 := digits.OfUint64Array([]uint64{
			10453563476522671549, 13602304095098071475,
		}).Negative()
		qm2 := digits.OfUint64Array([]uint64{
			6521084571113434188, 4027291763527809819, 1,
		}).Negative()

		expectedResult := digits.OfUint64Array([]uint64{
			10718923615238770716, 3695425665771625431,
		}).Add(digits.OfUint64Array([]uint64{
			0, 14328280640247263727, 2282221183331507844,
		})).Add(digits.OfUint64Array([]uint64{
			0, 8133934727708222756, 4808532877764353134,
		})).Add(digits.OfUint64Array([]uint64{
			0, 0, 6349429345627481825, 2969653995756497575,
		})).Add(digits.OfUint64Array([]uint64{
			0, 0, 10453563476522671549,
		})).Add(digits.OfUint64Array([]uint64{
			0, 0, 0, 13602304095098071475,
		}))

		result := multiply.MultiplySwitch(pm2, qm2)

		Expect(expectedResult.IsEqualTo(result)).To(BeTrue())
	})

	It("Should multiply 4 digit numbers", func() {
		a := biginteger.OfUint64Array([]uint64{
			6742305324661190591, 12524700037052152845,
			8597934400591931070, 13063502066075112160,
		})

		b := biginteger.OfUint64Array([]uint64{
			11595963786453332908, 3312835375285772574,
			18281896215638159356, 12893435606261567004,
		})

		expected := biginteger.OfUint64Array([]uint64{
			9350441601238114644, 12018134535777490305,
			6558593627430686406, 14759991432080105822,
			1616944804242820096, 11667109537258502405,
			11888729491966241882, 9130794139506552212,
		})

		result := a.Multiply(b)

		Expect(result.IsEqualTo(expected)).To(BeTrue())
	})

	It("Should multiply 2 digits by 3 digits", func() {
		a := biginteger.OfUint64Array([]uint64{
			12268717471831606891,
			16727042310808773037,
		})

		b := biginteger.OfUint64Array([]uint64{
			16954562203620167011,
			5262421968781838169,
			1,
		})

		expected := biginteger.OfUint64Array([]uint64{
			14717840216547519073,
			13421059992701675616,
			7163263090314600523,
			3052129763433291490,
			1,
		})

		result := a.Multiply(b)

		Expect(result.IsEqualTo(expected)).To(BeTrue())
	})
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "multiply Test Suite")
}
