package biginteger_test

import (
	"github.com/borisskert/go-biginteger"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("my test", func() {

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
