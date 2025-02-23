package digits_test

import (
	. "github.com/borisskert/go-biginteger/digits"
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

	Context("MultiplyDigit", func() {
		It("[1,0] * 1 == [1,0]", func() {
			a := DoubleDigitOf(1, 0)
			b := One()

			hi, lo := a.MultiplyDigit(b)

			Expect(lo.Low()).To(Equal(Digit(0)))
			Expect(lo.High()).To(Equal(Digit(1)))
			Expect(hi).To(Equal(Digit(0)))
		})

		It("[1,0] * 2 == [2,0]", func() {
			a := DoubleDigitOf(1, 0)
			b := Digit(2)

			hi, lo := a.MultiplyDigit(b)

			Expect(lo.Low()).To(Equal(Digit(0)))
			Expect(lo.High()).To(Equal(Digit(2)))
			Expect(hi).To(Equal(Digit(0)))
		})

		It("[3,0] * 2 == [6,0]", func() {
			a := DoubleDigitOf(3, 0)
			b := Digit(2)

			hi, lo := a.MultiplyDigit(b)

			Expect(lo.Low()).To(Equal(Digit(0)))
			Expect(lo.High()).To(Equal(Digit(6)))
			Expect(hi).To(Equal(Digit(0)))
		})

		It("[3,1] * 2 == [6,2]", func() {
			a := DoubleDigitOf(3, 1)
			b := Digit(2)

			hi, lo := a.MultiplyDigit(b)

			Expect(lo.Low()).To(Equal(Digit(2)))
			Expect(lo.High()).To(Equal(Digit(6)))
			Expect(hi).To(Equal(Digit(0)))
		})

		It("[2, 0] * 18446744073709551615 == [0, 18446744073709551614, 1]", func() {
			a := DoubleDigitOf(2, 0)
			b := Digit(18446744073709551615)

			hi, lo := a.MultiplyDigit(b)

			Expect(lo.Low()).To(Equal(Digit(0)))
			Expect(lo.High()).To(Equal(Digit(18446744073709551614)))
			Expect(hi).To(Equal(Digit(1)))
		})

		It("[18446744073709551615, 18446744073709551615] * 18446744073709551615 == [1, 18446744073709551615, 18446744073709551614]", func() {
			a := DoubleDigitOf(18446744073709551615, 18446744073709551615)
			b := Digit(18446744073709551615)

			hi, lo := a.MultiplyDigit(b)

			Expect(lo.Low()).To(Equal(Digit(1)))
			Expect(lo.High()).To(Equal(Digit(18446744073709551615)))
			Expect(hi).To(Equal(Digit(18446744073709551614)))
		})

		It("[12952425778477584591, 1] * 14863705750885344022", func() {
			a := DoubleDigitOf(1, 12952425778477584591)
			b := Digit(14863705750885344022)

			// 14863705750885344022, 0
			// 10436586790720179144, 15771953601206468298
			// ----------------
			// 1, 6853548467895971550, 15771953601206468298

			hi, lo := a.MultiplyDigit(b)

			Expect(lo.Low()).To(Equal(Digit(15771953601206468298)))
			Expect(lo.High()).To(Equal(Digit(6853548467895971550)))
			Expect(hi).To(Equal(Digit(1)))
		})

		It("[12952425778477584591, 2] * 14863705750885344022", func() {
			a := DoubleDigitOf(2, 12952425778477584591)
			b := Digit(14863705750885344022)

			// 29727411501770688044, 0                 -> 1, 11280667428061136428, 0
			// 192521045531472853064731924028641165002 -> 0, 10436586790720179144, 15771953601206468298
			// ----------------------------------------------------------------------------------------
			//											  1, 21717254218781315572, 15771953601206468298
			//											  2,  3270510145071763956, 15771953601206468298

			hi, lo := a.MultiplyDigit(b)

			Expect(lo.Low()).To(Equal(Digit(15771953601206468298)))
			Expect(lo.High()).To(Equal(Digit(3270510145071763956)))
			Expect(hi).To(Equal(Digit(2)))
		})
	})

	Context("Multiply", func() {
		It("[1,0] * [1,0] == [0,0],[1,0]", func() {
			a := DoubleDigitOf(1, 0)
			b := DoubleDigitOf(1, 0)

			hi, lo := a.Multiply(b)

			Expect(lo.Low()).To(Equal(Digit(0)))
			Expect(lo.High()).To(Equal(Digit(0)))
			Expect(hi.Low()).To(Equal(Digit(1)))
			Expect(hi.High()).To(Equal(Digit(0)))
		})

		It("[1,0] * [2,0] == [0,0],[2,0]", func() {
			a := DoubleDigitOf(1, 0)
			b := DoubleDigitOf(2, 0)

			hi, lo := a.Multiply(b)

			Expect(lo.Low()).To(Equal(Digit(0)))
			Expect(lo.High()).To(Equal(Digit(0)))
			Expect(hi.Low()).To(Equal(Digit(2)))
			Expect(hi.High()).To(Equal(Digit(0)))
		})

		It("[1,0] * [3,0] == [0,0],[3,0]", func() {
			a := DoubleDigitOf(1, 0)
			b := DoubleDigitOf(3, 0)

			hi, lo := a.Multiply(b)

			Expect(lo.Low()).To(Equal(Digit(0)))
			Expect(lo.High()).To(Equal(Digit(0)))
			Expect(hi.Low()).To(Equal(Digit(3)))
			Expect(hi.High()).To(Equal(Digit(0)))
		})

		It("[1,0] * [3,1] == [0,3],[1,0]", func() {
			a := DoubleDigitOf(1, 0)
			b := DoubleDigitOf(3, 1)

			hi, lo := a.Multiply(b)

			Expect(lo.Low()).To(Equal(Digit(0)))
			Expect(lo.High()).To(Equal(Digit(1)))
			Expect(hi.Low()).To(Equal(Digit(3)))
			Expect(hi.High()).To(Equal(Digit(0)))
		})

		It("[1,0] * [2^64-1,0] == [2^64-1, 0],[0,0]", func() {
			a := DoubleDigitOf(1, 0)
			b := DoubleDigitOf(18446744073709551615, 0)

			hi, lo := a.Multiply(b)

			Expect(lo.Low()).To(Equal(Digit(0)))
			Expect(lo.High()).To(Equal(Digit(0)))
			Expect(hi.Low()).To(Equal(Digit(18446744073709551615)))
			Expect(hi.High()).To(Equal(Digit(0)))
		})

		It("[1,0] * [2^64-1,1] == [2^64-1, 1],[0,0]", func() {
			a := DoubleDigitOf(1, 0)
			b := DoubleDigitOf(18446744073709551615, 1)

			hi, lo := a.Multiply(b)

			Expect(lo.Low()).To(Equal(Digit(0)))
			Expect(lo.High()).To(Equal(Digit(1)))
			Expect(hi.Low()).To(Equal(Digit(18446744073709551615)))
			Expect(hi.High()).To(Equal(Digit(0)))
		})

		It("[1,0] * [18446744073709551615, 18446744073709551615]", func() {
			a := DoubleDigitOf(1, 0)
			b := DoubleDigitOf(18446744073709551615, 18446744073709551615)

			hi, lo := a.Multiply(b)

			Expect(lo.Low()).To(Equal(Digit(0)))
			Expect(lo.High()).To(Equal(Digit(18446744073709551615)))
			Expect(hi.Low()).To(Equal(Digit(18446744073709551615)))
			Expect(hi.High()).To(Equal(Digit(0)))
		})

		It("[2,0] * [18446744073709551615, 18446744073709551615]", func() {
			a := DoubleDigitOf(2, 0)
			b := DoubleDigitOf(18446744073709551615, 18446744073709551615)

			hi, lo := a.Multiply(b)

			Expect(lo.Low()).To(Equal(Digit(0)))
			Expect(lo.High()).To(Equal(Digit(18446744073709551614)))
			Expect(hi.Low()).To(Equal(Digit(18446744073709551615)))
			Expect(hi.High()).To(Equal(Digit(1)))
		})

		It("[18446744073709551615, 18446744073709551615] * [18446744073709551615, 18446744073709551615]", func() {
			a := DoubleDigitOf(18446744073709551615, 18446744073709551615)
			b := DoubleDigitOf(18446744073709551615, 18446744073709551615)

			hi, lo := a.Multiply(b)

			Expect(lo.Low()).To(Equal(Digit(1)))
			Expect(lo.High()).To(Equal(Digit(0)))
			Expect(hi.Low()).To(Equal(Digit(18446744073709551614)))
			Expect(hi.High()).To(Equal(Digit(18446744073709551615)))
		})

		It("[6742305324661190591, 12524700037052152845] * [11595963786453332908, 3312835375285772574]", func() {
			a := DoubleDigitOf(12524700037052152845, 6742305324661190591)
			b := DoubleDigitOf(3312835375285772574, 11595963786453332908)

			hi, lo := a.Multiply(b)

			Expect(lo.Low()).To(Equal(Digit(9350441601238114644)))
			Expect(lo.High()).To(Equal(Digit(12018134535777490305)))
			Expect(hi.Low()).To(Equal(Digit(327839218333391898)))
			Expect(hi.High()).To(Equal(Digit(2249300428400506475)))
		})

		It("[18446744073709551608, 3] * [18446744073709551593, 5]", func() {
			a := DoubleDigitOf(18446744073709551608, 3)
			b := DoubleDigitOf(18446744073709551593, 5)

			hi, lo := a.Multiply(b)

			Expect(lo.Low()).To(Equal(Digit(15)))
			Expect(lo.High()).To(Equal(Digit(18446744073709551507)))
			Expect(hi.Low()).To(Equal(Digit(191)))
			Expect(hi.High()).To(Equal(Digit(18446744073709551585)))
		})
	})

	Context("DivideByDigit", func() {
		It("[327839218333391898, 12018134535777490305] / 3312835375285772574", func() {
			a := DoubleDigitOf(12018134535777490305, 327839218333391898)
			b := Digit(3312835375285772574)

			q, r := a.DivideByDigit(b)

			Expect(q.Low()).To(Equal(Digit(11579921336388089544)))
			Expect(q.High()).To(Equal(Digit(3)))
			Expect(r).To(Equal(Digit(810595452668569770)))
		})

		It("[18446744073709551545, 5] / 5", func() {
			a := DoubleDigitOf(5, 18446744073709551545)
			b := Digit(5)

			q, r := a.DivideByDigit(b)

			Expect(q.Low()).To(Equal(Digit(3689348814741910309)))
			Expect(q.High()).To(Equal(Digit(1)))
			Expect(r).To(Equal(Digit(0)))
		})
	})

	Context("Divide", func() {
		It("[327839218333391898 12018134535777490305] / [11595963786453332908 3312835375285772574]", func() {
			a := DoubleDigitOf(12018134535777490305, 327839218333391898)
			b := DoubleDigitOf(3312835375285772574, 11595963786453332908)

			q, r := a.Divide(b)

			Expect(q.Low()).To(Equal(Digit(3)))
			Expect(q.High()).To(Equal(Digit(0)))
			Expect(r.Low()).To(Equal(Digit(2433436006392496406)))
			Expect(r.High()).To(Equal(Digit(2079628409920172581)))
		})
	})
})
