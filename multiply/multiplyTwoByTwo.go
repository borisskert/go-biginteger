package multiply

import "github.com/borisskert/go-biginteger/digits"

func MultiplyTwoByTwo(a, b digits.Digits) digits.Digits {
	return multiplyTwoByTwoAbs(
		a.AsDoubleDigit(),
		b.AsDoubleDigit(),
	).
		Sign(a.IsNegative() != b.IsNegative()).
		Trim()
}

func multiplyTwoByTwoAbs(a, b digits.DoubleDigit) digits.Digits {
	hi, lo := a.Multiply(b)

	result := digits.OfUint64Array([]uint64{
		uint64(lo.Low()),
		uint64(lo.High()),
		uint64(hi.Low()),
		uint64(hi.High()),
	})

	return result
}
