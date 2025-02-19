package digits_test

import (
	"github.com/borisskert/go-biginteger"
	. "github.com/borisskert/go-biginteger/digits"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Digits", func() {
	Context("Add", func() {
		It("0 + 0", func() {
			a := Zero().AsDigits()
			b := Zero().AsDigits()

			result := a.Add(b)

			Expect(result.Length()).To(Equal(uint(1)))
			Expect(result.DigitAt(0)).To(Equal(Digit(0)))
			Expect(result.IsNegative()).To(BeFalse())
			Expect(result.IsZero()).To(BeTrue())
		})

		It("1 + -1", func() {
			a := Digit(1).AsDigits()
			b := Digit(1).AsDigits().Negative()

			result := a.Add(b)

			Expect(result.Length()).To(Equal(uint(1)))
			Expect(result.DigitAt(0)).To(Equal(Digit(0)))
			Expect(result.IsNegative()).To(BeFalse())
			Expect(result.IsZero()).To(BeTrue())
		})

		It("-1 + 1", func() {
			a := Digit(1).AsDigits().Negative()
			b := Digit(1).AsDigits()

			result := a.Add(b)

			Expect(result.Length()).To(Equal(uint(1)))
			Expect(result.DigitAt(0)).To(Equal(Digit(0)))
			Expect(result.IsNegative()).To(BeFalse())
			Expect(result.IsZero()).To(BeTrue())
		})

		It("-1 + -1", func() {
			a := Digit(1).AsDigits().Negative()
			b := Digit(1).AsDigits().Negative()

			result := a.Add(b)

			Expect(result.Length()).To(Equal(uint(1)))
			Expect(result.DigitAt(0)).To(Equal(Digit(2)))
			Expect(result.IsNegative()).To(BeTrue())
			Expect(result.IsZero()).To(BeFalse())
		})

		It("[0x1] + [0x1] ==> [0x2]", func() {
			a := Digit(1).AsDigits()
			b := Digit(1).AsDigits()

			result := a.Add(b)

			Expect(result.Length()).To(Equal(uint(1)))
			Expect(result.DigitAt(0)).To(Equal(Digit(2)))
		})

		It("[0x1] + [0x2] ==> [0x3]", func() {
			a := Digit(1).AsDigits()
			b := Digit(2).AsDigits()

			result := a.Add(b)

			Expect(result.Length()).To(Equal(uint(1)))
			Expect(result.DigitAt(0)).To(Equal(Digit(3)))
		})

		It("[0x1] + -[0x2] ==> -[0x1]", func() {
			a := Digit(1).AsDigits()
			b := Digit(2).AsDigits().Negate()

			result := a.Add(b)

			Expect(result.Length()).To(Equal(uint(1)))
			Expect(result.DigitAt(0)).To(Equal(Digit(1)))
			Expect(result.IsNegative()).To(BeTrue())
		})

		It("-[0x1] + -[0x2] ==> -[0x3]", func() {
			a := Digit(1).AsDigits().Negate()
			b := Digit(2).AsDigits().Negate()

			result := a.Add(b)

			Expect(result.Length()).To(Equal(uint(1)))
			Expect(result.DigitAt(0)).To(Equal(Digit(3)))
			Expect(result.IsNegative()).To(BeTrue())
		})

		It("[0x1] + -[0x2, 0x3] ==> -[0x1, 0x3]", func() {
			a := Digit(1).AsDigits()
			b := Digit(2).AsDigits().
				Add(Digit(3).AsDigits().LeftShiftBits(64)).
				Negate()

			result := a.Add(b)

			Expect(result.Length()).To(Equal(uint(2)))
			Expect(result.DigitAt(0)).To(Equal(Digit(1)))
			Expect(result.DigitAt(1)).To(Equal(Digit(3)))
			Expect(result.IsNegative()).To(BeTrue())
		})

		It("[0x1] + -[0x0, 0x3] ==> -[0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF, 0x2]", func() {
			a := Digit(1).AsDigits()
			b := Digit(0).AsDigits().
				Add(Digit(3).AsDigits().LeftShiftBits(64)).
				Negate()

			result := a.Add(b)

			Expect(result.Length()).To(Equal(uint(2)))
			Expect(result.DigitAt(0)).To(Equal(Digit(18446744073709551615)))
			Expect(result.DigitAt(1)).To(Equal(Digit(2)))
			Expect(result.IsNegative()).To(BeTrue())
		})

		It("-[0x1, 0x2] + [0x0, 0x3] ==> [0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF]", func() {
			a := Digit(1).AsDigits().
				Add(Digit(2).AsDigits().LeftShiftBits(64)).Negate()
			b := Digit(0).AsDigits().
				Add(Digit(3).AsDigits().LeftShiftBits(64))

			result := a.Add(b)

			Expect(result.Length()).To(Equal(uint(1)))
			Expect(result.DigitAt(0)).To(Equal(Digit(18446744073709551615)))
			Expect(result.IsNegative()).To(BeFalse())
		})

		It("-[0x1, 0x2] + [0x1, 0x2]", func() {
			a := OfUint64Array([]uint64{1, 2}).Negate()
			b := OfUint64Array([]uint64{1, 2})

			result := a.Add(b)

			Expect(result.Length()).To(Equal(uint(1)))
			Expect(result.DigitAt(0)).To(Equal(Digit(0)))
			Expect(result.IsNegative()).To(BeFalse())
			Expect(result.IsZero()).To(BeTrue())
		})

		It("[0x1, 0x2] + -[0x1, 0x2]", func() {
			a := OfUint64Array([]uint64{1, 2})
			b := OfUint64Array([]uint64{1, 2}).Negate()

			result := a.Add(b)

			Expect(result.Length()).To(Equal(uint(1)))
			Expect(result.DigitAt(0)).To(Equal(Digit(0)))
			Expect(result.IsNegative()).To(BeFalse())
			Expect(result.IsZero()).To(BeTrue())
		})

		It("Should add 2 digit numbers (Example a, p1)", func() {
			t0 := biginteger.OfUint64Array([]uint64{
				6742305324661190591, 12524700037052152845,
			})
			x1p := biginteger.OfUint64Array([]uint64{
				8597934400591931070, 13063502066075112160,
			})

			p1 := biginteger.OfUint64Array([]uint64{
				15340239725253121661,
				7141458029417713389,
				1,
			})

			result := t0.Add(x1p)

			Expect(result.IsEqualTo(p1)).To(BeTrue())
		})

		It("Should add 2 digit numbers (Example a, q1)", func() {
			t0q := OfUint64Array([]uint64{
				11595963786453332908, 3312835375285772574,
			})
			x1q := OfUint64Array([]uint64{
				18281896215638159356, 12893435606261567004,
			})

			q1 := OfUint64Array([]uint64{
				11431115928381940648, 16206270981547339579,
			})

			result := t0q.Add(x1q)

			Expect(result.IsEqualTo(q1)).To(BeTrue())
		})
	})

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

		It("[0x5, 0x0, 0x1] - [0x3, 0x1, 0x2]", func() {
			a := Digit(5).AsDigits().
				Add(Digit(0).AsDigits().LeftShiftBits(64)).
				Add(Digit(1).AsDigits().LeftShiftBits(64 * 2))

			b := Digit(3).AsDigits().
				Add(Digit(1).AsDigits().LeftShiftBits(64)).
				Add(Digit(2).AsDigits().LeftShiftBits(64 * 2))

			result, borrow := a.SubtractUnderflow(b)

			Expect(borrow).To(BeTrue())
			Expect(result.Length()).To(Equal(uint(3)))
			Expect(result.DigitAt(0)).To(Equal(Digit(2)))
			Expect(result.DigitAt(1)).To(Equal(Digit(18446744073709551615)))
			Expect(result.DigitAt(2)).To(Equal(Digit(18446744073709551614)))
		})
	})

	Context("SubtractAndBorrow", func() {
		It("1 - 2", func() {
			a := Digit(1).AsDigits()
			b := Digit(2).AsDigits()

			result, borrow := a.SubtractAndBorrow(b)

			Expect(result.Length()).To(Equal(uint(1)))
			Expect(result.DigitAt(0)).To(Equal(Digit(1)))
			Expect(result.IsNegative()).To(BeTrue())
			Expect(borrow).To(BeTrue())
		})

		It("2 - 1", func() {
			a := Digit(2).AsDigits()
			b := Digit(1).AsDigits()

			result, borrow := a.SubtractAndBorrow(b)

			Expect(result.Length()).To(Equal(uint(1)))
			Expect(result.DigitAt(0)).To(Equal(Digit(1)))
			Expect(result.IsNegative()).To(BeFalse())
			Expect(borrow).To(BeFalse())
		})

		It("1 - (2 pow 64 - 1)", func() {
			a := Digit(1).AsDigits()
			b, _ := Digit(1).AsDigits().LeftShiftBits(64).Decrement()

			result, borrow := a.SubtractAndBorrow(b.Trim())

			Expect(result.IsNegative()).To(BeTrue())
			Expect(result.Length()).To(Equal(uint(1)))
			Expect(result.DigitAt(0)).To(Equal(Digit(18446744073709551614)))
			Expect(borrow).To(BeTrue())
		})

		It("1 - 2 pow 64", func() {
			a := Digit(1).AsDigits()
			b := Digit(1).AsDigits().LeftShiftBits(64)

			result, borrow := a.SubtractAndBorrow(b)

			Expect(result.IsNegative()).To(BeTrue())
			Expect(result.Length()).To(Equal(uint(1)))
			Expect(result.DigitAt(0)).To(Equal(Digit(18446744073709551615)))
			Expect(borrow).To(BeTrue())
		})

		It("2 - 2 pow 64", func() {
			a := Digit(2).AsDigits()
			b := Digit(1).AsDigits().LeftShiftBits(64)

			result, borrow := a.SubtractAndBorrow(b)

			Expect(result.IsNegative()).To(BeTrue())
			Expect(result.Length()).To(Equal(uint(1)))
			Expect(result.DigitAt(0)).To(Equal(Digit(18446744073709551614)))
			Expect(borrow).To(BeTrue())
		})

		It("2 pow 128 - 1", func() {
			a := Digit(1).AsDigits().LeftShiftBits(128)
			b := Digit(1).AsDigits()

			result, borrow := a.SubtractAndBorrow(b)

			Expect(result.IsNegative()).To(BeFalse())
			Expect(result.Length()).To(Equal(uint(2)))
			Expect(result.DigitAt(0)).To(Equal(Digit(18446744073709551615)))
			Expect(result.DigitAt(1)).To(Equal(Digit(18446744073709551615)))
			Expect(borrow).To(BeFalse())
		})
	})

	Context("SubtractInPlace", func() {
		It("1 - 2", func() {
			a := Digit(1).AsDigits()
			b := Digit(2).AsDigits()

			borrow := a.SubtractInPlace(b)

			Expect(a.Length()).To(Equal(uint(1)))
			Expect(a.DigitAt(0)).To(Equal(Digit(1)))
			Expect(a.IsNegative()).To(BeTrue())
			Expect(borrow).To(BeTrue())
		})

		It("2 - 1", func() {
			a := Digit(2).AsDigits()
			b := Digit(1).AsDigits()

			borrow := a.SubtractInPlace(b)

			Expect(a.Length()).To(Equal(uint(1)))
			Expect(a.DigitAt(0)).To(Equal(Digit(1)))
			Expect(a.IsNegative()).To(BeFalse())
			Expect(borrow).To(BeFalse())
		})

		It("1 - (2 pow 64 - 1)", func() {
			a := Digit(1).AsDigits()
			b, _ := Digit(1).AsDigits().LeftShiftBits(64).Decrement()

			borrow := a.SubtractInPlace(b.Trim())

			Expect(a.IsNegative()).To(BeTrue())
			Expect(a.Length()).To(Equal(uint(1)))
			Expect(a.DigitAt(0)).To(Equal(Digit(18446744073709551614)))
			Expect(borrow).To(BeTrue())
		})

		It("1 - 2 pow 64", func() {
			a := Digit(1).AsDigits()
			b := Digit(1).AsDigits().LeftShiftBits(64)

			borrow := a.SubtractInPlace(b)

			Expect(a.IsNegative()).To(BeTrue())
			Expect(a.Length()).To(Equal(uint(1)))
			Expect(a.DigitAt(0)).To(Equal(Digit(18446744073709551615)))
			Expect(borrow).To(BeTrue())
		})
	})

	Context("Subtract", func() {
		It("1 - 2", func() {
			a := Digit(1).AsDigits()
			b := Digit(2).AsDigits()

			result := a.Subtract(b)

			Expect(result.Length()).To(Equal(uint(1)))
			Expect(result.DigitAt(0)).To(Equal(Digit(1)))
			Expect(result.IsNegative()).To(BeTrue())
		})

		It("1 - 1", func() {
			a := Digit(1).AsDigits()
			b := Digit(1).AsDigits()

			result := a.Subtract(b)

			Expect(result.Length()).To(Equal(uint(1)))
			Expect(result.DigitAt(0)).To(Equal(Digit(0)))
			Expect(result.IsNegative()).To(BeFalse())
			Expect(result.IsZero()).To(BeTrue())
		})

		It("-1 - -1", func() {
			a := Digit(1).AsDigits().Negate()
			b := Digit(1).AsDigits().Negate()

			result := a.Subtract(b)

			Expect(result.Length()).To(Equal(uint(1)))
			Expect(result.DigitAt(0)).To(Equal(Digit(0)))
			Expect(result.IsNegative()).To(BeFalse())
			Expect(result.IsZero()).To(BeTrue())
		})

		It("1 - -2", func() {
			a := Digit(1).AsDigits()
			b := Digit(2).AsDigits().Negate()

			result := a.Subtract(b)

			Expect(result.Length()).To(Equal(uint(1)))
			Expect(result.DigitAt(0)).To(Equal(Digit(3)))
			Expect(result.IsNegative()).To(BeFalse())
		})

		It("-1 - -2", func() {
			a := Digit(1).AsDigits().Negate()
			b := Digit(2).AsDigits().Negate()

			result := a.Subtract(b)

			Expect(result.Length()).To(Equal(uint(1)))
			Expect(result.DigitAt(0)).To(Equal(Digit(1)))
			Expect(result.IsNegative()).To(BeFalse())
		})

		It("-24 - 3", func() {
			a := Digit(24).AsDigits().Negate()
			b := Digit(3).AsDigits()

			result := a.Subtract(b)

			Expect(result.Length()).To(Equal(uint(1)))
			Expect(result.DigitAt(0)).To(Equal(Digit(27)))
			Expect(result.IsNegative()).To(BeTrue())
		})

		It("2 - 1", func() {
			a := Digit(2).AsDigits()
			b := Digit(1).AsDigits()

			result := a.Subtract(b)

			Expect(result.Length()).To(Equal(uint(1)))
			Expect(result.DigitAt(0)).To(Equal(Digit(1)))
			Expect(result.IsNegative()).To(BeFalse())
		})

		It("1 - (2 pow 64 - 1)", func() {
			a := Digit(1).AsDigits()
			b, _ := Digit(1).AsDigits().LeftShiftBits(64).Decrement()

			result := a.Subtract(b.Trim())

			Expect(result.Length()).To(Equal(uint(1)))
			Expect(result.DigitAt(0)).To(Equal(Digit(18446744073709551614)))
			Expect(result.IsNegative()).To(BeTrue())
		})

		It("1 - 2 pow 64", func() {
			a := Digit(1).AsDigits()
			b := Digit(1).AsDigits().LeftShiftBits(64)

			result := a.Subtract(b)

			Expect(result.Length()).To(Equal(uint(1)))
			Expect(result.DigitAt(0)).To(Equal(Digit(18446744073709551615)))
			Expect(result.IsNegative()).To(BeTrue())
		})

		It("2 pow 64 - 1", func() {
			a := Digit(1).AsDigits().LeftShiftBits(64)
			b := Digit(1).AsDigits()

			result := a.Subtract(b)

			Expect(result.Length()).To(Equal(uint(1)))
			Expect(result.DigitAt(0)).To(Equal(Digit(18446744073709551615)))
			Expect(result.IsNegative()).To(BeFalse())
		})

		It("310 - 36", func() {
			a := Digit(310).AsDigits()
			b := Digit(36).AsDigits()

			result := a.Subtract(b)

			Expect(result.Length()).To(Equal(uint(1)))
			Expect(result.DigitAt(0)).To(Equal(Digit(274)))
			Expect(result.IsNegative()).To(BeFalse())
		})

		It("-[0x1, 0x2] - -[0x1, 0x2]", func() {
			a := OfUint64Array([]uint64{1, 2}).Negate()
			b := OfUint64Array([]uint64{1, 2}).Negate()

			result := a.Subtract(b)

			Expect(result.Length()).To(Equal(uint(1)))
			Expect(result.DigitAt(0)).To(Equal(Digit(0)))
			Expect(result.IsNegative()).To(BeFalse())
			Expect(result.IsZero()).To(BeTrue())
		})

		It("[0x1, 0x2] - [0x1, 0x2]", func() {
			a := OfUint64Array([]uint64{1, 2})
			b := OfUint64Array([]uint64{1, 2})

			result := a.Subtract(b)

			Expect(result.Length()).To(Equal(uint(1)))
			Expect(result.DigitAt(0)).To(Equal(Digit(0)))
			Expect(result.IsNegative()).To(BeFalse())
			Expect(result.IsZero()).To(BeTrue())
		})

		It("-[0x1, 0x2] - [0x1, 0x2]", func() {
			a := OfUint64Array([]uint64{1, 2}).Negate()
			b := OfUint64Array([]uint64{1, 2})

			result := a.Subtract(b)

			Expect(result.Length()).To(Equal(uint(2)))
			Expect(result.DigitAt(0)).To(Equal(Digit(2)))
			Expect(result.DigitAt(1)).To(Equal(Digit(4)))
			Expect(result.IsNegative()).To(BeTrue())
			Expect(result.IsZero()).To(BeFalse())
		})

		It("[0x1, 0x2] - -[0x1, 0x2]", func() {
			a := OfUint64Array([]uint64{1, 2})
			b := OfUint64Array([]uint64{1, 2}).Negate()

			result := a.Subtract(b)

			Expect(result.Length()).To(Equal(uint(2)))
			Expect(result.DigitAt(0)).To(Equal(Digit(2)))
			Expect(result.DigitAt(1)).To(Equal(Digit(4)))
			Expect(result.IsNegative()).To(BeFalse())
			Expect(result.IsZero()).To(BeFalse())
		})

		It("Should subtract 2 digit numbers (Example a, qm1)", func() {
			t0q := OfUint64Array([]uint64{
				11595963786453332908, 3312835375285772574,
			})
			x1q := OfUint64Array([]uint64{
				18281896215638159356, 12893435606261567004,
			})

			qm1 := OfUint64Array([]uint64{
				6685932429184826448, 9580600230975794430,
			}).Negative()

			result := t0q.Subtract(x1q)

			Expect(result.IsEqualTo(qm1)).To(BeTrue())
		})

		It("Should subtract 2 digit numbers (Example a, qm1mx1q)", func() {
			qm1 := OfUint64Array([]uint64{
				6685932429184826448, 9580600230975794430,
			}).Negative()
			x1q := OfUint64Array([]uint64{
				18281896215638159356, 12893435606261567004,
			})

			qm1mx1q := OfUint64Array([]uint64{
				6521084571113434188, 4027291763527809819, 1,
			}).Negative()

			result := qm1.Subtract(x1q)

			Expect(result.IsEqualTo(qm1mx1q)).To(BeTrue())
		})

		It("Should subtract 2 digit numbers (Example a, pm1)", func() {
			t0 := OfUint64Array([]uint64{
				6742305324661190591, 12524700037052152845,
			})
			x1 := OfUint64Array([]uint64{
				8597934400591931070, 13063502066075112160,
			})

			pm1 := OfUint64Array([]uint64{
				1855629075930740479, 538802029022959315,
			}).Negate()

			result := t0.Subtract(x1)

			Expect(result.IsEqualTo(pm1)).To(BeTrue())
		})

		It("Should subtract 2 digit numbers (Example a, pm1mx1)", func() {
			pm1 := OfUint64Array([]uint64{
				1855629075930740479, 538802029022959315,
			}).Negate()
			x1 := OfUint64Array([]uint64{
				8597934400591931070, 13063502066075112160,
			})

			pm1mx1 := OfUint64Array([]uint64{
				10453563476522671549, 13602304095098071475,
			}).Negate()

			result := pm1.Subtract(x1)

			Expect(result.IsEqualTo(pm1mx1)).To(BeTrue())
		})

		// [327839218333391898, 12018134535777490305] - [17963987839374373744, 12018134535777490304]
		It("[327839218333391898, 12018134535777490305] - [17963987839374373744, 12018134535777490304]", func() {
			a := OfUint64Array([]uint64{327839218333391898, 12018134535777490305})
			b := OfUint64Array([]uint64{17963987839374373744, 12018134535777490304})

			result := a.Subtract(b)

			Expect(result.Length()).To(Equal(uint(1)))
			Expect(result.DigitAt(0)).To(Equal(Digit(810595452668569770)))
			Expect(result.IsNegative()).To(BeFalse())
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

	Context("MultiplyByDigit", func() {
		It("1 * 1", func() {
			a := Digit(1).AsDigits()
			b := Digit(1)

			result := a.MultiplyByDigit(b)

			Expect(result.Length()).To(Equal(uint(1)))
			Expect(result.DigitAt(0)).To(Equal(Digit(1)))
		})

		It("2 * 2", func() {
			a := Digit(2).AsDigits()
			b := Digit(2)

			result := a.MultiplyByDigit(b)

			Expect(result.Length()).To(Equal(uint(1)))
			Expect(result.DigitAt(0)).To(Equal(Digit(4)))
		})

		It("2 pow 64  *  2", func() {
			a := Digit(1).AsDigits().LeftShiftBits(64)
			b := Digit(2)

			result := a.MultiplyByDigit(b)

			Expect(result.Length()).To(Equal(uint(2)))
			Expect(result.DigitAt(0)).To(Equal(Digit(0)))
			Expect(result.DigitAt(1)).To(Equal(Digit(2)))
		})

		It("2 pow (64 * 2)  *  2", func() {
			a := Digit(1).AsDigits().LeftShiftBits(64 * 2)
			b := Digit(2)

			result := a.MultiplyByDigit(b)

			Expect(result.Length()).To(Equal(uint(3)))
			Expect(result.DigitAt(0)).To(Equal(Digit(0)))
			Expect(result.DigitAt(1)).To(Equal(Digit(0)))
			Expect(result.DigitAt(2)).To(Equal(Digit(2)))
		})

		It("2 pow (64 * 2)  *  18446744073709551615", func() {
			a := Digit(1).AsDigits().LeftShiftBits(64 * 2)
			b := Digit(18446744073709551615)

			result := a.MultiplyByDigit(b)

			Expect(result.Length()).To(Equal(uint(3)))
			Expect(result.DigitAt(0)).To(Equal(Digit(0)))
			Expect(result.DigitAt(1)).To(Equal(Digit(0)))
			Expect(result.DigitAt(2)).To(Equal(Digit(18446744073709551615)))
		})

		It("18446744073709551615 * 2 pow (64 * 2)  *  2", func() {
			a := Digit(18446744073709551615).AsDigits().LeftShiftBits(64 * 2)
			b := Digit(2)

			result := a.MultiplyByDigit(b)

			Expect(result.Length()).To(Equal(uint(4)))
			Expect(result.DigitAt(0)).To(Equal(Digit(0)))
			Expect(result.DigitAt(1)).To(Equal(Digit(0)))
			Expect(result.DigitAt(2)).To(Equal(Digit(18446744073709551614)))
			Expect(result.DigitAt(3)).To(Equal(Digit(1)))
		})

		It("[18446744073709551615, 18446744073709551615] * 18446744073709551615", func() {
			a := Digit(18446744073709551615).AsDigits().
				Add(Digit(18446744073709551615).AsDigits().LeftShiftBits(64))
			b := Digit(18446744073709551615)

			result := a.MultiplyByDigit(b)

			Expect(result.Length()).To(Equal(uint(3)))
			Expect(result.DigitAt(0)).To(Equal(Digit(1)))
			Expect(result.DigitAt(1)).To(Equal(Digit(18446744073709551615)))
			Expect(result.DigitAt(2)).To(Equal(Digit(18446744073709551614)))
		})

		It("[11579921336388089544, 3] * 3312835375285772574", func() {
			a := OfUint64Array([]uint64{
				11579921336388089544, 3,
			})
			b := Digit(3312835375285772574)

			result := a.MultiplyByDigit(b)

			Expect(result.Length()).To(Equal(uint(2)))
			Expect(result.DigitAt(0)).To(Equal(Digit(17963987839374373744)))
			Expect(result.DigitAt(1)).To(Equal(Digit(12018134535777490304)))
		})
	})

	Context("Split2", func() {
		It("of [0x1] =(1)=> [[0x1], []]", func() {
			a := Digit(1).AsDigits()

			a1, a0 := a.Split2(1)

			Expect(a0.Length()).To(Equal(uint(1)))
			Expect(a0.DigitAt(0)).To(Equal(Digit(1)))

			Expect(a1.Length()).To(Equal(uint(1)))
			Expect(a1.DigitAt(0)).To(Equal(Digit(0)))
		})

		It("of -[0x1] =(1)=> [[0x1], []]", func() {
			a := Digit(1).AsDigits().Negative()

			a1, a0 := a.Split2(1)

			Expect(a0.Length()).To(Equal(uint(1)))
			Expect(a0.DigitAt(0)).To(Equal(Digit(1)))
			Expect(a0.IsNegative()).To(BeFalse())

			Expect(a1.Length()).To(Equal(uint(1)))
			Expect(a1.DigitAt(0)).To(Equal(Digit(0)))
			Expect(a1.IsNegative()).To(BeFalse())
		})

		It("of [0x1] =(2)=> [[0x1], []]", func() {
			a := Digit(1).AsDigits()

			a1, a0 := a.Split2(2)

			Expect(a0.Length()).To(Equal(uint(1)))
			Expect(a0.DigitAt(0)).To(Equal(Digit(1)))

			Expect(a1.Length()).To(Equal(uint(1)))
			Expect(a1.DigitAt(0)).To(Equal(Digit(0)))
		})

		It("of [0x1] =(3)=> [[0x1], []]", func() {
			a := Digit(1).AsDigits()

			a1, a0 := a.Split2(3)

			Expect(a0.Length()).To(Equal(uint(1)))
			Expect(a0.DigitAt(0)).To(Equal(Digit(1)))

			Expect(a1.Length()).To(Equal(uint(1)))
			Expect(a1.DigitAt(0)).To(Equal(Digit(0)))
		})

		It("of [0x1, 0x2] =(1)=> [[0x1], [0x2]]", func() {
			a := Digit(1).AsDigits().
				Add(Digit(2).AsDigits().LeftShiftBits(64))

			a1, a0 := a.Split2(1)

			Expect(a0.Length()).To(Equal(uint(1)))
			Expect(a0.DigitAt(0)).To(Equal(Digit(1)))

			Expect(a1.Length()).To(Equal(uint(1)))
			Expect(a1.DigitAt(0)).To(Equal(Digit(2)))
		})

		It("of [0x1, 0x2] =(2)=> [[0x1, 0x2], []]", func() {
			a := Digit(1).AsDigits().
				Add(Digit(2).AsDigits().LeftShiftBits(64))

			a1, a0 := a.Split2(2)

			Expect(a0.Length()).To(Equal(uint(2)))
			Expect(a0.DigitAt(0)).To(Equal(Digit(1)))
			Expect(a0.DigitAt(1)).To(Equal(Digit(2)))

			Expect(a1.Length()).To(Equal(uint(1)))
			Expect(a1.DigitAt(0)).To(Equal(Digit(0)))
		})

		It("of [0x1, 0x2] =(3)=> [[0x1, 0x2], []]", func() {
			a := Digit(1).AsDigits().
				Add(Digit(2).AsDigits().LeftShiftBits(64))

			a1, a0 := a.Split2(3)

			Expect(a0.Length()).To(Equal(uint(2)))
			Expect(a0.DigitAt(0)).To(Equal(Digit(1)))
			Expect(a0.DigitAt(1)).To(Equal(Digit(2)))

			Expect(a1.Length()).To(Equal(uint(1)))
			Expect(a1.DigitAt(0)).To(Equal(Digit(0)))
		})

		It("of [0x1, 0x2, 0x3] =(1)=> [0x1, [0x2, 0x3]]", func() {
			a := Digit(1).AsDigits().
				Add(Digit(2).AsDigits().LeftShiftBits(64)).
				Add(Digit(3).AsDigits().LeftShiftBits(64 * 2))

			a1, a0 := a.Split2(1)

			Expect(a0.Length()).To(Equal(uint(1)))
			Expect(a0.DigitAt(0)).To(Equal(Digit(1)))

			Expect(a1.Length()).To(Equal(uint(2)))
			Expect(a1.DigitAt(0)).To(Equal(Digit(2)))
			Expect(a1.DigitAt(1)).To(Equal(Digit(3)))
		})

		It("of [0x1, 0x2, 0x3] =(2)=> [[0x1, 0x2], [0x3]]", func() {
			a := Digit(1).AsDigits().
				Add(Digit(2).AsDigits().LeftShiftBits(64)).
				Add(Digit(3).AsDigits().LeftShiftBits(64 * 2))

			a1, a0 := a.Split2(2)

			Expect(a0.Length()).To(Equal(uint(2)))
			Expect(a0.DigitAt(0)).To(Equal(Digit(1)))
			Expect(a0.DigitAt(1)).To(Equal(Digit(2)))

			Expect(a1.Length()).To(Equal(uint(1)))
			Expect(a1.DigitAt(0)).To(Equal(Digit(3)))
		})

		It("of [0x1, 0x2, 0x3] =(3)=> [[0x1, 0x2, 0x3], []]", func() {
			a := Digit(1).AsDigits().
				Add(Digit(2).AsDigits().LeftShiftBits(64)).
				Add(Digit(3).AsDigits().LeftShiftBits(64 * 2))

			a1, a0 := a.Split2(3)

			Expect(a0.Length()).To(Equal(uint(3)))
			Expect(a0.DigitAt(0)).To(Equal(Digit(1)))
			Expect(a0.DigitAt(1)).To(Equal(Digit(2)))
			Expect(a0.DigitAt(2)).To(Equal(Digit(3)))

			Expect(a1.Length()).To(Equal(uint(1)))
			Expect(a1.DigitAt(0)).To(Equal(Digit(0)))
		})
	})

	Context("Split3", func() {
		It("of [0x1] =(1)=> [[0x1], [], []]", func() {
			a := Digit(1).AsDigits()

			a2, a1, a0 := a.Split3(1)

			Expect(a0.Length()).To(Equal(uint(1)))
			Expect(a0.DigitAt(0)).To(Equal(Digit(1)))

			Expect(a1.Length()).To(Equal(uint(1)))
			Expect(a1.DigitAt(0)).To(Equal(Digit(0)))

			Expect(a2.Length()).To(Equal(uint(1)))
			Expect(a2.DigitAt(0)).To(Equal(Digit(0)))
		})

		It("of -[0x1] =(1)=> [[0x1], [], []]", func() {
			a := Digit(1).AsDigits().Negative()

			a2, a1, a0 := a.Split3(1)

			Expect(a0.Length()).To(Equal(uint(1)))
			Expect(a0.DigitAt(0)).To(Equal(Digit(1)))
			Expect(a0.IsNegative()).To(BeFalse())

			Expect(a1.Length()).To(Equal(uint(1)))
			Expect(a1.DigitAt(0)).To(Equal(Digit(0)))
			Expect(a1.IsNegative()).To(BeFalse())

			Expect(a2.Length()).To(Equal(uint(1)))
			Expect(a2.DigitAt(0)).To(Equal(Digit(0)))
			Expect(a2.IsNegative()).To(BeFalse())
		})

		It("of [0x1] =(2)=> [[0x1], [], []]", func() {
			a := Digit(1).AsDigits()

			a2, a1, a0 := a.Split3(2)

			Expect(a0.Length()).To(Equal(uint(1)))
			Expect(a0.DigitAt(0)).To(Equal(Digit(1)))

			Expect(a1.Length()).To(Equal(uint(1)))
			Expect(a1.DigitAt(0)).To(Equal(Digit(0)))

			Expect(a2.Length()).To(Equal(uint(1)))
			Expect(a2.DigitAt(0)).To(Equal(Digit(0)))
		})

		It("of [0x1] =(3)=> [[0x1], [], []]", func() {
			a := Digit(1).AsDigits()

			a2, a1, a0 := a.Split3(3)

			Expect(a0.Length()).To(Equal(uint(1)))
			Expect(a0.DigitAt(0)).To(Equal(Digit(1)))

			Expect(a1.Length()).To(Equal(uint(1)))
			Expect(a1.DigitAt(0)).To(Equal(Digit(0)))

			Expect(a2.Length()).To(Equal(uint(1)))
			Expect(a2.DigitAt(0)).To(Equal(Digit(0)))
		})

		It("of [0x1, 0x2] =(1)=> [[0x1], [0x2], []]", func() {
			a := Digit(1).AsDigits().
				Add(Digit(2).AsDigits().LeftShiftBits(64))

			a2, a1, a0 := a.Split3(1)

			Expect(a0.Length()).To(Equal(uint(1)))
			Expect(a0.DigitAt(0)).To(Equal(Digit(1)))

			Expect(a1.Length()).To(Equal(uint(1)))
			Expect(a1.DigitAt(0)).To(Equal(Digit(2)))

			Expect(a2.Length()).To(Equal(uint(1)))
			Expect(a2.DigitAt(0)).To(Equal(Digit(0)))
		})

		It("of [0x1, 0x2] =(2)=> [[0x1, 0x2], [], []]", func() {
			a := Digit(1).AsDigits().
				Add(Digit(2).AsDigits().LeftShiftBits(64))

			a2, a1, a0 := a.Split3(2)

			Expect(a0.Length()).To(Equal(uint(2)))
			Expect(a0.DigitAt(0)).To(Equal(Digit(1)))
			Expect(a0.DigitAt(1)).To(Equal(Digit(2)))

			Expect(a1.Length()).To(Equal(uint(1)))
			Expect(a1.DigitAt(0)).To(Equal(Digit(0)))

			Expect(a2.Length()).To(Equal(uint(1)))
			Expect(a2.DigitAt(0)).To(Equal(Digit(0)))
		})

		It("of [0x1, 0x2] =(3)=> [[0x1, 0x2], [], []]", func() {
			a := Digit(1).AsDigits().
				Add(Digit(2).AsDigits().LeftShiftBits(64))

			a2, a1, a0 := a.Split3(3)

			Expect(a0.Length()).To(Equal(uint(2)))
			Expect(a0.DigitAt(0)).To(Equal(Digit(1)))
			Expect(a0.DigitAt(1)).To(Equal(Digit(2)))

			Expect(a1.Length()).To(Equal(uint(1)))
			Expect(a1.DigitAt(0)).To(Equal(Digit(0)))

			Expect(a2.Length()).To(Equal(uint(1)))
			Expect(a2.DigitAt(0)).To(Equal(Digit(0)))
		})

		It("of [0x1, 0x2, 0x3] =(1)> [[0x1], [0x2], [0x3]]", func() {
			a := Digit(1).AsDigits().
				Add(Digit(2).AsDigits().LeftShiftBits(64)).
				Add(Digit(3).AsDigits().LeftShiftBits(64 * 2))

			a2, a1, a0 := a.Split3(1)

			Expect(a0.Length()).To(Equal(uint(1)))
			Expect(a0.DigitAt(0)).To(Equal(Digit(1)))

			Expect(a1.Length()).To(Equal(uint(1)))
			Expect(a1.DigitAt(0)).To(Equal(Digit(2)))

			Expect(a2.Length()).To(Equal(uint(1)))
			Expect(a2.DigitAt(0)).To(Equal(Digit(3)))
		})

		It("of [0x1, 0x2, 0x3] =(2)> [[0x1, 0x2], [0x3], []]", func() {
			a := Digit(1).AsDigits().
				Add(Digit(2).AsDigits().LeftShiftBits(64)).
				Add(Digit(3).AsDigits().LeftShiftBits(64 * 2))

			a2, a1, a0 := a.Split3(2)

			Expect(a0.Length()).To(Equal(uint(2)))
			Expect(a0.DigitAt(0)).To(Equal(Digit(1)))
			Expect(a0.DigitAt(1)).To(Equal(Digit(2)))

			Expect(a1.Length()).To(Equal(uint(1)))
			Expect(a1.DigitAt(0)).To(Equal(Digit(3)))

			Expect(a2.Length()).To(Equal(uint(1)))
			Expect(a2.DigitAt(0)).To(Equal(Digit(0)))
		})

		It("of [0x1, 0x2, 0x3] =(3)> [[0x1, 0x2, 0x3], [], []]", func() {
			a := Digit(1).AsDigits().
				Add(Digit(2).AsDigits().LeftShiftBits(64)).
				Add(Digit(3).AsDigits().LeftShiftBits(64 * 2))

			a2, a1, a0 := a.Split3(3)

			Expect(a0.Length()).To(Equal(uint(3)))
			Expect(a0.DigitAt(0)).To(Equal(Digit(1)))
			Expect(a0.DigitAt(1)).To(Equal(Digit(2)))
			Expect(a0.DigitAt(2)).To(Equal(Digit(3)))

			Expect(a1.Length()).To(Equal(uint(1)))
			Expect(a1.DigitAt(0)).To(Equal(Digit(0)))

			Expect(a2.Length()).To(Equal(uint(1)))
			Expect(a2.DigitAt(0)).To(Equal(Digit(0)))
		})

		It("of [0x1, 0x2, 0x3] =(4)> [[0x1, 0x2, 0x3], [], []]", func() {
			a := Digit(1).AsDigits().
				Add(Digit(2).AsDigits().LeftShiftBits(64)).
				Add(Digit(3).AsDigits().LeftShiftBits(64 * 2))

			a2, a1, a0 := a.Split3(4)

			Expect(a0.Length()).To(Equal(uint(3)))
			Expect(a0.DigitAt(0)).To(Equal(Digit(1)))
			Expect(a0.DigitAt(1)).To(Equal(Digit(2)))
			Expect(a0.DigitAt(2)).To(Equal(Digit(3)))

			Expect(a1.Length()).To(Equal(uint(1)))
			Expect(a1.DigitAt(0)).To(Equal(Digit(0)))

			Expect(a2.Length()).To(Equal(uint(1)))
			Expect(a2.DigitAt(0)).To(Equal(Digit(0)))
		})

		It("of [0x1, 0x2, 0x3, 0x4] =(1)=> [[0x1], [0x2], [0x3, 0x4]]", func() {
			a := Digit(1).AsDigits().
				Add(Digit(2).AsDigits().LeftShiftBits(64)).
				Add(Digit(3).AsDigits().LeftShiftBits(64 * 2)).
				Add(Digit(4).AsDigits().LeftShiftBits(64 * 3))

			a2, a1, a0 := a.Split3(1)

			Expect(a0.Length()).To(Equal(uint(1)))
			Expect(a0.DigitAt(0)).To(Equal(Digit(1)))

			Expect(a1.Length()).To(Equal(uint(1)))
			Expect(a1.DigitAt(0)).To(Equal(Digit(2)))

			Expect(a2.Length()).To(Equal(uint(2)))
			Expect(a2.DigitAt(0)).To(Equal(Digit(3)))
			Expect(a2.DigitAt(1)).To(Equal(Digit(4)))
		})

		It("of [0x1, 0x2, 0x3, 0x4, 0x5] =(1)=> [[0x1], [0x2], [0x3, 0x4, 0x5]]", func() {
			a := Digit(1).AsDigits().
				Add(Digit(2).AsDigits().LeftShiftBits(64)).
				Add(Digit(3).AsDigits().LeftShiftBits(64 * 2)).
				Add(Digit(4).AsDigits().LeftShiftBits(64 * 3)).
				Add(Digit(5).AsDigits().LeftShiftBits(64 * 4))

			a2, a1, a0 := a.Split3(1)

			Expect(a0.Length()).To(Equal(uint(1)))
			Expect(a0.DigitAt(0)).To(Equal(Digit(1)))

			Expect(a1.Length()).To(Equal(uint(1)))
			Expect(a1.DigitAt(0)).To(Equal(Digit(2)))

			Expect(a2.Length()).To(Equal(uint(3)))
			Expect(a2.DigitAt(0)).To(Equal(Digit(3)))
			Expect(a2.DigitAt(1)).To(Equal(Digit(4)))
			Expect(a2.DigitAt(2)).To(Equal(Digit(5)))
		})

		It("of [0x1, 0x2, 0x3, 0x4, 0x3] =(2)=> [[0x1, 0x2], [0x3, 0x4], [0x3]]", func() {
			a := Digit(1).AsDigits().
				Add(Digit(2).AsDigits().LeftShiftBits(64)).
				Add(Digit(3).AsDigits().LeftShiftBits(64 * 2)).
				Add(Digit(4).AsDigits().LeftShiftBits(64 * 3))

			a2, a1, a0 := a.Split3(2)

			Expect(a0.Length()).To(Equal(uint(2)))
			Expect(a0.DigitAt(0)).To(Equal(Digit(1)))
			Expect(a0.DigitAt(1)).To(Equal(Digit(2)))

			Expect(a1.Length()).To(Equal(uint(2)))
			Expect(a1.DigitAt(0)).To(Equal(Digit(3)))
			Expect(a1.DigitAt(1)).To(Equal(Digit(4)))

			Expect(a2.Length()).To(Equal(uint(1)))
			Expect(a2.DigitAt(0)).To(Equal(Digit(0)))
		})

		It("of [0x1, 0x2, 0x3, 0x4, 0x5] =(3)=> [[0x1, 0x2, 0x3], [0x4, 0x5], []]", func() {
			a := Digit(1).AsDigits().
				Add(Digit(2).AsDigits().LeftShiftBits(64)).
				Add(Digit(3).AsDigits().LeftShiftBits(64 * 2)).
				Add(Digit(4).AsDigits().LeftShiftBits(64 * 3)).
				Add(Digit(5).AsDigits().LeftShiftBits(64 * 4))

			a2, a1, a0 := a.Split3(3)

			Expect(a0.Length()).To(Equal(uint(3)))
			Expect(a0.DigitAt(0)).To(Equal(Digit(1)))
			Expect(a0.DigitAt(1)).To(Equal(Digit(2)))
			Expect(a0.DigitAt(2)).To(Equal(Digit(3)))

			Expect(a1.Length()).To(Equal(uint(2)))
			Expect(a1.DigitAt(0)).To(Equal(Digit(4)))
			Expect(a1.DigitAt(1)).To(Equal(Digit(5)))

			Expect(a2.Length()).To(Equal(uint(1)))
			Expect(a2.DigitAt(0)).To(Equal(Digit(0)))
		})

		It("of [0x1, 0x2, 0x3, 0x4, 0x5] =(4)=> [[0x1, 0x2, 0x3, 0x4], [0x5], []]", func() {
			a := Digit(1).AsDigits().
				Add(Digit(2).AsDigits().LeftShiftBits(64)).
				Add(Digit(3).AsDigits().LeftShiftBits(64 * 2)).
				Add(Digit(4).AsDigits().LeftShiftBits(64 * 3)).
				Add(Digit(5).AsDigits().LeftShiftBits(64 * 4))

			a2, a1, a0 := a.Split3(4)

			Expect(a0.Length()).To(Equal(uint(4)))
			Expect(a0.DigitAt(0)).To(Equal(Digit(1)))
			Expect(a0.DigitAt(1)).To(Equal(Digit(2)))
			Expect(a0.DigitAt(2)).To(Equal(Digit(3)))
			Expect(a0.DigitAt(3)).To(Equal(Digit(4)))

			Expect(a1.Length()).To(Equal(uint(1)))
			Expect(a1.DigitAt(0)).To(Equal(Digit(5)))

			Expect(a2.Length()).To(Equal(uint(1)))
			Expect(a2.DigitAt(0)).To(Equal(Digit(0)))
		})

		It("of [0x1, 0x2, 0x3, 0x4, 0x5] =(5)> [[0x1, 0x2, 0x3, 0x4, 0x5], [], []]", func() {
			a := Digit(1).AsDigits().
				Add(Digit(2).AsDigits().LeftShiftBits(64)).
				Add(Digit(3).AsDigits().LeftShiftBits(64 * 2)).
				Add(Digit(4).AsDigits().LeftShiftBits(64 * 3)).
				Add(Digit(5).AsDigits().LeftShiftBits(64 * 4))

			a2, a1, a0 := a.Split3(5)

			Expect(a0.Length()).To(Equal(uint(5)))
			Expect(a0.DigitAt(0)).To(Equal(Digit(1)))
			Expect(a0.DigitAt(1)).To(Equal(Digit(2)))
			Expect(a0.DigitAt(2)).To(Equal(Digit(3)))
			Expect(a0.DigitAt(3)).To(Equal(Digit(4)))
			Expect(a0.DigitAt(4)).To(Equal(Digit(5)))

			Expect(a1.Length()).To(Equal(uint(1)))
			Expect(a1.DigitAt(0)).To(Equal(Digit(0)))

			Expect(a2.Length()).To(Equal(uint(1)))
			Expect(a2.DigitAt(0)).To(Equal(Digit(0)))
		})

		It("of [0x1, 0x2, 0x3, 0x4, 0x5, 0x6] =(3)=> [[0x1, 0x2, 0x3], [0x4, 0x5, 0x6], []]", func() {
			a := Digit(1).AsDigits().
				Add(Digit(2).AsDigits().LeftShiftBits(64)).
				Add(Digit(3).AsDigits().LeftShiftBits(64 * 2)).
				Add(Digit(4).AsDigits().LeftShiftBits(64 * 3)).
				Add(Digit(5).AsDigits().LeftShiftBits(64 * 4)).
				Add(Digit(6).AsDigits().LeftShiftBits(64 * 5))

			a2, a1, a0 := a.Split3(3)

			Expect(a0.Length()).To(Equal(uint(3)))
			Expect(a0.DigitAt(0)).To(Equal(Digit(1)))
			Expect(a0.DigitAt(1)).To(Equal(Digit(2)))
			Expect(a0.DigitAt(2)).To(Equal(Digit(3)))

			Expect(a1.Length()).To(Equal(uint(3)))
			Expect(a1.DigitAt(0)).To(Equal(Digit(4)))
			Expect(a1.DigitAt(1)).To(Equal(Digit(5)))
			Expect(a1.DigitAt(2)).To(Equal(Digit(6)))

			Expect(a2.Length()).To(Equal(uint(1)))
			Expect(a2.DigitAt(0)).To(Equal(Digit(0)))
		})

		It("of [0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7] =(3)=> [[0x1, 0x2, 0x3], [0x4, 0x5 0x6], [0x7]]", func() {
			a := Digit(1).AsDigits().
				Add(Digit(2).AsDigits().LeftShiftBits(64)).
				Add(Digit(3).AsDigits().LeftShiftBits(64 * 2)).
				Add(Digit(4).AsDigits().LeftShiftBits(64 * 3)).
				Add(Digit(5).AsDigits().LeftShiftBits(64 * 4)).
				Add(Digit(6).AsDigits().LeftShiftBits(64 * 5)).
				Add(Digit(7).AsDigits().LeftShiftBits(64 * 6))

			a2, a1, a0 := a.Split3(3)

			Expect(a0.Length()).To(Equal(uint(3)))
			Expect(a0.DigitAt(0)).To(Equal(Digit(1)))
			Expect(a0.DigitAt(1)).To(Equal(Digit(2)))
			Expect(a0.DigitAt(2)).To(Equal(Digit(3)))

			Expect(a1.Length()).To(Equal(uint(3)))
			Expect(a1.DigitAt(0)).To(Equal(Digit(4)))
			Expect(a1.DigitAt(1)).To(Equal(Digit(5)))
			Expect(a1.DigitAt(2)).To(Equal(Digit(6)))

			Expect(a2.Length()).To(Equal(uint(1)))
			Expect(a2.DigitAt(0)).To(Equal(Digit(7)))
		})
	})

	Context("SplitEvenOdd", func() {
		It("[1] -> [[1], []]", func() {
			a := Digit(1).AsDigits()

			even, odd := a.SplitEvenOdd()

			Expect(even.Length()).To(Equal(uint(1)))
			Expect(even.DigitAt(0)).To(Equal(Digit(1)))
			Expect(even.IsNegative()).To(BeFalse())

			Expect(odd.Length()).To(Equal(uint(1)))
			Expect(odd.DigitAt(0)).To(Equal(Digit(0)))
			Expect(odd.IsNegative()).To(BeFalse())
		})

		It("-[1] -> [[1], []]", func() {
			a := Digit(1).AsDigits().Negative()

			even, odd := a.SplitEvenOdd()

			Expect(even.Length()).To(Equal(uint(1)))
			Expect(even.DigitAt(0)).To(Equal(Digit(1)))
			Expect(even.IsNegative()).To(BeFalse())

			Expect(odd.Length()).To(Equal(uint(1)))
			Expect(odd.DigitAt(0)).To(Equal(Digit(0)))
			Expect(odd.IsNegative()).To(BeFalse())
		})

		It("[1, 2] -> [[1], [2]]", func() {
			a := OfUint64Array([]uint64{1, 2})

			even, odd := a.SplitEvenOdd()

			Expect(even.Length()).To(Equal(uint(1)))
			Expect(even.DigitAt(0)).To(Equal(Digit(1)))
			Expect(even.IsNegative()).To(BeFalse())

			Expect(odd.Length()).To(Equal(uint(1)))
			Expect(odd.DigitAt(0)).To(Equal(Digit(2)))
			Expect(odd.IsNegative()).To(BeFalse())
		})

		It("-[1, 2] -> [[1], [2]]", func() {
			a := OfUint64Array([]uint64{1, 2}).Negative()

			even, odd := a.SplitEvenOdd()

			Expect(even.Length()).To(Equal(uint(1)))
			Expect(even.DigitAt(0)).To(Equal(Digit(1)))
			Expect(even.IsNegative()).To(BeFalse())

			Expect(odd.Length()).To(Equal(uint(1)))
			Expect(odd.DigitAt(0)).To(Equal(Digit(2)))
			Expect(odd.IsNegative()).To(BeFalse())
		})

		It("[1, 2, 3] -> [[1, 3], [2]]", func() {
			a := OfUint64Array([]uint64{1, 2, 3})

			even, odd := a.SplitEvenOdd()

			Expect(even.Length()).To(Equal(uint(2)))
			Expect(even.DigitAt(0)).To(Equal(Digit(1)))
			Expect(even.DigitAt(1)).To(Equal(Digit(3)))
			Expect(even.IsNegative()).To(BeFalse())

			Expect(odd.Length()).To(Equal(uint(1)))
			Expect(odd.DigitAt(0)).To(Equal(Digit(2)))
			Expect(odd.IsNegative()).To(BeFalse())
		})

		It("-[1, 2, 3] -> [[1, 3], [2]]", func() {
			a := OfUint64Array([]uint64{1, 2, 3}).Negative()

			even, odd := a.SplitEvenOdd()

			Expect(even.Length()).To(Equal(uint(2)))
			Expect(even.DigitAt(0)).To(Equal(Digit(1)))
			Expect(even.DigitAt(1)).To(Equal(Digit(3)))
			Expect(even.IsNegative()).To(BeFalse())

			Expect(odd.Length()).To(Equal(uint(1)))
			Expect(odd.DigitAt(0)).To(Equal(Digit(2)))
			Expect(odd.IsNegative()).To(BeFalse())
		})

		It("[1, 2, 3, 4] -> [[1, 3], [2, 4]]", func() {
			a := OfUint64Array([]uint64{1, 2, 3, 4})

			even, odd := a.SplitEvenOdd()

			Expect(even.Length()).To(Equal(uint(2)))
			Expect(even.DigitAt(0)).To(Equal(Digit(1)))
			Expect(even.DigitAt(1)).To(Equal(Digit(3)))
			Expect(even.IsNegative()).To(BeFalse())

			Expect(odd.Length()).To(Equal(uint(2)))
			Expect(odd.DigitAt(0)).To(Equal(Digit(2)))
			Expect(odd.DigitAt(1)).To(Equal(Digit(4)))
			Expect(odd.IsNegative()).To(BeFalse())
		})

		It("-[1, 2, 3, 4] -> [[1, 3], [2, 4]]", func() {
			a := OfUint64Array([]uint64{1, 2, 3, 4}).Negative()

			even, odd := a.SplitEvenOdd()

			Expect(even.Length()).To(Equal(uint(2)))
			Expect(even.DigitAt(0)).To(Equal(Digit(1)))
			Expect(even.DigitAt(1)).To(Equal(Digit(3)))
			Expect(even.IsNegative()).To(BeFalse())

			Expect(odd.Length()).To(Equal(uint(2)))
			Expect(odd.DigitAt(0)).To(Equal(Digit(2)))
			Expect(odd.DigitAt(1)).To(Equal(Digit(4)))
			Expect(odd.IsNegative()).To(BeFalse())
		})

		It("[1, 2, 3, 4, 5] -> [[1, 3, 5], [2, 4]]", func() {
			a := OfUint64Array([]uint64{1, 2, 3, 4, 5})

			even, odd := a.SplitEvenOdd()

			Expect(even.Length()).To(Equal(uint(3)))
			Expect(even.DigitAt(0)).To(Equal(Digit(1)))
			Expect(even.DigitAt(1)).To(Equal(Digit(3)))
			Expect(even.DigitAt(2)).To(Equal(Digit(5)))
			Expect(even.IsNegative()).To(BeFalse())

			Expect(odd.Length()).To(Equal(uint(2)))
			Expect(odd.DigitAt(0)).To(Equal(Digit(2)))
			Expect(odd.DigitAt(1)).To(Equal(Digit(4)))
			Expect(odd.IsNegative()).To(BeFalse())
		})

		It("-[1, 2, 3, 4, 5] -> [[1, 3, 5], [2, 4]]", func() {
			a := OfUint64Array([]uint64{1, 2, 3, 4, 5}).Negative()

			even, odd := a.SplitEvenOdd()

			Expect(even.Length()).To(Equal(uint(3)))
			Expect(even.DigitAt(0)).To(Equal(Digit(1)))
			Expect(even.DigitAt(1)).To(Equal(Digit(3)))
			Expect(even.DigitAt(2)).To(Equal(Digit(5)))
			Expect(even.IsNegative()).To(BeFalse())

			Expect(odd.Length()).To(Equal(uint(2)))
			Expect(odd.DigitAt(0)).To(Equal(Digit(2)))
			Expect(odd.DigitAt(1)).To(Equal(Digit(4)))
			Expect(odd.IsNegative()).To(BeFalse())
		})
	})

	Context("DivideByDigit", func() {
		It("1 / 1", func() {
			a := Digit(1).AsDigits()
			b := Digit(1)

			result, remainder := a.DivideByDigit(b)

			Expect(result.Length()).To(Equal(uint(1)))
			Expect(result.DigitAt(0)).To(Equal(Digit(1)))
			Expect(remainder).To(Equal(Digit(0)))
		})

		It("2 / 1", func() {
			a := Digit(2).AsDigits()
			b := Digit(1)

			result, remainder := a.DivideByDigit(b)

			Expect(result.Length()).To(Equal(uint(1)))
			Expect(result.DigitAt(0)).To(Equal(Digit(2)))
			Expect(remainder).To(Equal(Digit(0)))
		})

		It("2 / 2", func() {
			a := Digit(2).AsDigits()
			b := Digit(2)

			result, remainder := a.DivideByDigit(b)

			Expect(result.Length()).To(Equal(uint(1)))
			Expect(result.DigitAt(0)).To(Equal(Digit(1)))
			Expect(remainder).To(Equal(Digit(0)))
		})

		It("2 / 3", func() {
			a := Digit(2).AsDigits()
			b := Digit(3)

			result, remainder := a.DivideByDigit(b)

			Expect(result.Length()).To(Equal(uint(1)))
			Expect(result.DigitAt(0)).To(Equal(Digit(0)))
			Expect(remainder).To(Equal(Digit(2)))
		})

		It("2 / 4", func() {
			a := Digit(2).AsDigits()
			b := Digit(4)

			result, remainder := a.DivideByDigit(b)

			Expect(result.Length()).To(Equal(uint(1)))
			Expect(result.DigitAt(0)).To(Equal(Digit(0)))
			Expect(remainder).To(Equal(Digit(2)))
		})

		It("2 / 5", func() {
			a := Digit(2).AsDigits()
			b := Digit(5)

			result, remainder := a.DivideByDigit(b)

			Expect(result.Length()).To(Equal(uint(1)))
			Expect(result.DigitAt(0)).To(Equal(Digit(0)))
			Expect(remainder).To(Equal(Digit(2)))
		})

		It("12 / 6", func() {
			a := Digit(12).AsDigits()
			b := Digit(6)

			result, remainder := a.DivideByDigit(b)

			Expect(result.Length()).To(Equal(uint(1)))
			Expect(result.DigitAt(0)).To(Equal(Digit(2)))
			Expect(result.IsNegative()).To(BeFalse())
			Expect(remainder).To(Equal(Digit(0)))
		})

		It("-12 / 6", func() {
			a := Digit(12).AsDigits().Negate()
			b := Digit(6)

			result, remainder := a.DivideByDigit(b)

			Expect(result.Length()).To(Equal(uint(1)))
			Expect(result.DigitAt(0)).To(Equal(Digit(2)))
			Expect(result.IsNegative()).To(BeTrue())
			Expect(remainder).To(Equal(Digit(0)))
		})

		It("[2, 1] / 6", func() {
			a := Digit(2).AsDigits().
				Add(Digit(1).AsDigits().LeftShiftBits(64))
			b := Digit(6)

			result, remainder := a.DivideByDigit(b)

			Expect(result.Length()).To(Equal(uint(1)))
			Expect(remainder).To(Equal(Digit(0)))
			Expect(result.String()).To(Equal("3074457345618258603"))
		})

		It("-[2, 1] / 6", func() {
			a := OfUint64Array([]uint64{2, 1}).Negate()
			b := Digit(6)

			result, remainder := a.DivideByDigit(b)

			Expect(result.Length()).To(Equal(uint(1)))
			Expect(remainder).To(Equal(Digit(0)))
			Expect(result.String()).To(Equal("-3074457345618258603"))
		})

		It("[0x4, 0x2, 0x3] / 6", func() {
			a := Digit(4).AsDigits().
				Add(Digit(2).AsDigits().LeftShiftBits(64)).
				Add(Digit(3).AsDigits().LeftShiftBits(64 * 2))
			b := Digit(6)

			result, remainder := a.DivideByDigit(b)

			Expect(result.Length()).To(Equal(uint(2)))
			Expect(result.DigitAt(0)).To(Equal(Digit(6148914691236517206)))
			Expect(result.DigitAt(1)).To(Equal(Digit(9223372036854775808)))
			Expect(remainder).To(Equal(Digit(0)))
			Expect(result.String()).To(Equal("170141183460469231737836218407120622934"))
		})

		It("[2, 2] / 2", func() {
			a := OfUint64Array([]uint64{2, 2}).Negate()
			b := Digit(2)

			result, remainder := a.DivideByDigit(b)

			Expect(result.Length()).To(Equal(uint(2)))
			Expect(result.DigitAt(0)).To(Equal(Digit(1)))
			Expect(result.DigitAt(1)).To(Equal(Digit(1)))
			Expect(remainder).To(Equal(Digit(0)))
		})
	})

	Context("DivideByDoubleDigit", func() {
		It("1 / 1", func() {
			a := Digit(1).AsDigits()
			b := Digit(1).AsDoubleDigit()

			result, remainder := a.DivideByDoubleDigit(b)

			Expect(result.Length()).To(Equal(uint(1)))
			Expect(result.DigitAt(0)).To(Equal(Digit(1)))
			Expect(remainder.Low()).To(Equal(Digit(0)))
			Expect(remainder.High()).To(Equal(Digit(0)))
		})

		It("2 / 1", func() {
			a := Digit(2).AsDigits()
			b := Digit(1).AsDoubleDigit()

			result, remainder := a.DivideByDoubleDigit(b)

			Expect(result.Length()).To(Equal(uint(1)))
			Expect(result.DigitAt(0)).To(Equal(Digit(2)))
			Expect(remainder.Low()).To(Equal(Digit(0)))
			Expect(remainder.High()).To(Equal(Digit(0)))
		})

		It("2 / 2", func() {
			a := Digit(2).AsDigits()
			b := Digit(2).AsDoubleDigit()

			result, remainder := a.DivideByDoubleDigit(b)

			Expect(result.Length()).To(Equal(uint(1)))
			Expect(result.DigitAt(0)).To(Equal(Digit(1)))
			Expect(remainder.Low()).To(Equal(Digit(0)))
			Expect(remainder.High()).To(Equal(Digit(0)))
		})

		It("[0, 1] / 2", func() {
			a := OfUint64Array([]uint64{0, 1})
			b := Digit(2).AsDoubleDigit()

			result, remainder := a.DivideByDoubleDigit(b)

			Expect(result.Length()).To(Equal(uint(1)))
			Expect(result.DigitAt(0)).To(Equal(Digit(9223372036854775808)))
			Expect(remainder.Low()).To(Equal(Digit(0)))
			Expect(remainder.High()).To(Equal(Digit(0)))
		})

		It("[0, 2] / 2", func() {
			a := OfUint64Array([]uint64{0, 2})
			b := Digit(2).AsDoubleDigit()

			result, remainder := a.DivideByDoubleDigit(b)

			Expect(result.Length()).To(Equal(uint(2)))
			Expect(result.DigitAt(0)).To(Equal(Digit(0)))
			Expect(result.DigitAt(1)).To(Equal(Digit(1)))
			Expect(remainder.Low()).To(Equal(Digit(0)))
			Expect(remainder.High()).To(Equal(Digit(0)))
		})

		It("[2, 2] / 2", func() {
			a := OfUint64Array([]uint64{2, 2})
			b := Digit(2).AsDoubleDigit()

			result, remainder := a.DivideByDoubleDigit(b)

			Expect(result.Length()).To(Equal(uint(2)))
			Expect(result.DigitAt(0)).To(Equal(Digit(1)))
			Expect(result.DigitAt(1)).To(Equal(Digit(1)))
			Expect(remainder.Low()).To(Equal(Digit(0)))
			Expect(remainder.High()).To(Equal(Digit(0)))
		})

		It("[2, 2] / [1, 1]", func() {
			a := OfUint64Array([]uint64{2, 2})
			b := OfUint64Array([]uint64{1, 1}).AsDoubleDigit()

			result, remainder := a.DivideByDoubleDigit(b)

			Expect(result.Length()).To(Equal(uint(1)))
			Expect(result.DigitAt(0)).To(Equal(Digit(2)))
			Expect(remainder.Low()).To(Equal(Digit(0)))
			Expect(remainder.High()).To(Equal(Digit(0)))
		})

		It("[2, 2] / [2, 2]", func() {
			a := OfUint64Array([]uint64{2, 2})
			b := OfUint64Array([]uint64{2, 2}).AsDoubleDigit()

			result, remainder := a.DivideByDoubleDigit(b)

			Expect(result.Length()).To(Equal(uint(1)))
			Expect(result.DigitAt(0)).To(Equal(Digit(1)))
			Expect(remainder.Low()).To(Equal(Digit(0)))
			Expect(remainder.High()).To(Equal(Digit(0)))
		})

		It("[0, 2, 4] / [2, 2]", func() {
			a := OfUint64Array([]uint64{0, 2, 2})
			b := OfUint64Array([]uint64{2, 2}).AsDoubleDigit()

			result, remainder := a.DivideByDoubleDigit(b)

			Expect(result.Length()).To(Equal(uint(2)))
			Expect(result.DigitAt(0)).To(Equal(Digit(0)))
			Expect(result.DigitAt(1)).To(Equal(Digit(1)))
			Expect(remainder.Low()).To(Equal(Digit(0)))
			Expect(remainder.High()).To(Equal(Digit(0)))
		})

		It("[9350441601238114644 12018134535777490305 327839218333391898 2249300428400506475] / [11595963786453332908 3312835375285772574]", func() {
			a := OfUint64Array([]uint64{9350441601238114644, 12018134535777490305, 327839218333391898, 2249300428400506475})
			b := OfUint64Array([]uint64{11595963786453332908, 3312835375285772574}).AsDoubleDigit()

			result, remainder := a.DivideByDoubleDigit(b)

			Expect(result.Length()).To(Equal(uint(2)))
			Expect(result.DigitAt(0)).To(Equal(Digit(6742305324661190591)))
			Expect(result.DigitAt(1)).To(Equal(Digit(12524700037052152845)))
			Expect(remainder.Low()).To(Equal(Digit(0)))
			Expect(remainder.High()).To(Equal(Digit(0)))
		})

		It("[184 18446744073709551476 23] / [5 18446744073709551593]", func() {
			a := OfUint64Array([]uint64{184, 18446744073709551476, 23})
			b := OfUint64Array([]uint64{5, 18446744073709551593}).AsDoubleDigit()

			result, remainder := a.DivideByDoubleDigit(b)

			Expect(result.Length()).To(Equal(uint(1)))
			Expect(result.DigitAt(0)).To(Equal(Digit(24)))
			Expect(remainder.Low()).To(Equal(Digit(64)))
			Expect(remainder.High()).To(Equal(Digit(412)))
		})
	})

	Context("DivThreeByTwo", func() {
		XIt("[2249300428400506475 327839218333391898 12018134535777490305] / [11595963786453332908 3312835375285772574]", func() {
			// TODO this test is not valid
			aHi := Digit(12018134535777490305)
			aMid := Digit(327839218333391898)
			aLo := Digit(2249300428400506475)

			bHi := Digit(3312835375285772574)
			bLo := Digit(11595963786453332908)

			result, remainder := DivThreeByTwo(aHi, aMid, aLo, bHi, bLo)

			Expect(result.Low()).To(Equal(Digit(11579921336388089531)))
			Expect(result.High()).To(Equal(Digit(3)))
			Expect(remainder.Low()).To(Equal(Digit(11185678565676756935)))
			Expect(remainder.High()).To(Equal(Digit(1810211695536404770)))
		})

		It("[0 2 2] / [2 2]", func() {
			aLo := Digit(0)
			aMid := Digit(2)
			aHi := Digit(2)

			bLo := Digit(2)
			bHi := Digit(2)

			result, remainder := DivThreeByTwo(aHi, aMid, aLo, bHi, bLo)

			Expect(result.Low()).To(Equal(Digit(0)))
			Expect(result.High()).To(Equal(Digit(1)))
			Expect(remainder.Low()).To(Equal(Digit(0)))
			Expect(remainder.High()).To(Equal(Digit(0)))
		})

		It("[0 1 2] / [2 2]", func() {
			aLo := Digit(0)
			aMid := Digit(1)
			aHi := Digit(2)

			bLo := Digit(2)
			bHi := Digit(2)

			result, remainder := DivThreeByTwo(aHi, aMid, aLo, bHi, bLo)

			Expect(result.Low()).To(Equal(Digit(18446744073709551615)))
			Expect(result.High()).To(Equal(Digit(0)))
			Expect(remainder.Low()).To(Equal(Digit(2)))
			Expect(remainder.High()).To(Equal(Digit(1)))
		})

		It("[2 2 1] / [2 2]", func() {
			aLo := Digit(2)
			aMid := Digit(2)
			aHi := Digit(1)

			bLo := Digit(2)
			bHi := Digit(2)

			result, remainder := DivThreeByTwo(aHi, aMid, aLo, bHi, bLo)

			Expect(result.Low()).To(Equal(Digit(9223372036854775808)))
			Expect(result.High()).To(Equal(Digit(0)))
			Expect(remainder.Low()).To(Equal(Digit(2)))
			Expect(remainder.High()).To(Equal(Digit(1)))
		})

		It("[0 4 4] / [2 2]", func() {
			aLo := Digit(0)
			aMid := Digit(4)
			aHi := Digit(4)

			bLo := Digit(2)
			bHi := Digit(2)

			result, remainder := DivThreeByTwo(aHi, aMid, aLo, bHi, bLo)

			Expect(result.Low()).To(Equal(Digit(0)))
			Expect(result.High()).To(Equal(Digit(2)))
			Expect(remainder.Low()).To(Equal(Digit(0)))
			Expect(remainder.High()).To(Equal(Digit(0)))
		})

		It("[1 4 4] / [2 2]", func() {
			aLo := Digit(1)
			aMid := Digit(4)
			aHi := Digit(4)

			bLo := Digit(2)
			bHi := Digit(2)

			result, remainder := DivThreeByTwo(aHi, aMid, aLo, bHi, bLo)

			Expect(result.Low()).To(Equal(Digit(0)))
			Expect(result.High()).To(Equal(Digit(2)))
			Expect(remainder.Low()).To(Equal(Digit(1)))
			Expect(remainder.High()).To(Equal(Digit(0)))
		})

		It("[2 4 4] / [2 2]", func() {
			aLo := Digit(2)
			aMid := Digit(4)
			aHi := Digit(4)

			bLo := Digit(2)
			bHi := Digit(2)

			result, remainder := DivThreeByTwo(aHi, aMid, aLo, bHi, bLo)

			Expect(result.Low()).To(Equal(Digit(0)))
			Expect(result.High()).To(Equal(Digit(2)))
			Expect(remainder.Low()).To(Equal(Digit(2)))
			Expect(remainder.High()).To(Equal(Digit(0)))
		})

		It("[12018134535777490305 327839218333391898 2249300428400506475] / [3312835375285772574 11595963786453332908]", func() {
			aHi := Digit(2249300428400506475)
			aMid := Digit(327839218333391898)
			aLo := Digit(12018134535777490305)
			bHi := Digit(11595963786453332908)
			bLo := Digit(3312835375285772574)

			result, remainder := DivThreeByTwo(aHi, aMid, aLo, bHi, bLo)

			Expect(result.Low()).To(Equal(Digit(3578164791792606809)))
			Expect(result.High()).To(Equal(Digit(0)))
			Expect(remainder.Low()).To(Equal(Digit(10451349232378154515)))
			Expect(remainder.High()).To(Equal(Digit(11363108077110180365)))
		})
	})

	Context("Difference", func() {
		It("1 - 1", func() {
			a := Digit(1).AsDigits()
			b := Digit(1).AsDigits()

			result := a.Difference(b)

			Expect(result.Length()).To(Equal(uint(1)))
			Expect(result.DigitAt(0)).To(Equal(Digit(0)))
			Expect(result.IsNegative()).To(BeFalse())
		})

		It("1 - 2", func() {
			a := Digit(1).AsDigits()
			b := Digit(2).AsDigits()

			result := a.Difference(b)

			Expect(result.Length()).To(Equal(uint(1)))
			Expect(result.DigitAt(0)).To(Equal(Digit(1)))
			Expect(result.IsNegative()).To(BeFalse())
		})

		It("|2 - 1|", func() {
			a := Digit(2).AsDigits()
			b := Digit(1).AsDigits()

			result := a.Difference(b)

			Expect(result.Length()).To(Equal(uint(1)))
			Expect(result.DigitAt(0)).To(Equal(Digit(1)))
			Expect(result.IsNegative()).To(BeFalse())
		})

		It("|(-2) - 1|", func() {
			a := Digit(2).AsDigits().Negate()
			b := Digit(1).AsDigits()

			result := a.Difference(b)

			Expect(result.Length()).To(Equal(uint(1)))
			Expect(result.DigitAt(0)).To(Equal(Digit(3)))
			Expect(result.IsNegative()).To(BeFalse())
		})

		It("|(-2) - (-1)|", func() {
			a := Digit(2).AsDigits().Negate()
			b := Digit(1).AsDigits().Negate()

			result := a.Difference(b)

			Expect(result.Length()).To(Equal(uint(1)))
			Expect(result.DigitAt(0)).To(Equal(Digit(1)))
			Expect(result.IsNegative()).To(BeFalse())
		})

		It("|2 - (-1)|", func() {
			a := Digit(2).AsDigits()
			b := Digit(1).AsDigits().Negate()

			result := a.Difference(b)

			Expect(result.Length()).To(Equal(uint(1)))
			Expect(result.DigitAt(0)).To(Equal(Digit(3)))
			Expect(result.IsNegative()).To(BeFalse())
		})

		It("[0x1] - [0x0, 0x1]", func() {
			a := Digit(1).AsDigits()
			b := Digit(1).AsDigits().
				LeftShiftBits(64)

			result := a.Difference(b)

			Expect(result.Length()).To(Equal(uint(1)))
			Expect(result.DigitAt(0)).To(Equal(Digit(18446744073709551615)))
			Expect(result.IsNegative()).To(BeFalse())
		})

		It("-[0x1] - [0x0, 0x1]", func() {
			a := Digit(1).AsDigits().Negate()
			b := Digit(1).AsDigits().
				LeftShiftBits(64)

			result := a.Difference(b)

			Expect(result.Length()).To(Equal(uint(2)))
			Expect(result.DigitAt(0)).To(Equal(Digit(1)))
			Expect(result.DigitAt(1)).To(Equal(Digit(1)))
			Expect(result.IsNegative()).To(BeFalse())
		})

		It("[0x1] - -[0x0, 0x1]", func() {
			a := Digit(1).AsDigits()
			b := Digit(1).AsDigits().
				LeftShiftBits(64).
				Negate()

			result := a.Difference(b)

			Expect(result.Length()).To(Equal(uint(2)))
			Expect(result.DigitAt(0)).To(Equal(Digit(1)))
			Expect(result.DigitAt(1)).To(Equal(Digit(1)))
			Expect(result.IsNegative()).To(BeFalse())
		})

		It("-[0x1] - -[0x0, 0x1]", func() {
			a := Digit(1).AsDigits().Negate()
			b := Digit(1).AsDigits().
				LeftShiftBits(64).
				Negate()

			result := a.Difference(b)

			Expect(result.Length()).To(Equal(uint(1)))
			Expect(result.DigitAt(0)).To(Equal(Digit(18446744073709551615)))
			Expect(result.IsNegative()).To(BeFalse())
		})
	})

	Context("RightShiftBits", func() {
		It("[2, 2] by one", func() {
			a := OfUint64Array([]uint64{2, 2})

			result := a.RightShiftBits(1)

			Expect(result.Length()).To(Equal(uint(2)))
			Expect(result.DigitAt(0)).To(Equal(Digit(1)))
			Expect(result.DigitAt(1)).To(Equal(Digit(1)))
		})
	})
})
