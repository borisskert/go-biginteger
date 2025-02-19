package multiply

import "github.com/borisskert/go-biginteger/digits"

func Multiply(a, b []uint64) []uint64 {
	product := MultiplySwitch2(
		digits.Wrap(a),
		digits.Wrap(b),
		Multiply2,
	)
	return product.Trim().AsArray()
}

func MultiplySwitch(a, b digits.Digits) digits.Digits {
	m := a.Length()
	n := b.Length()

	if min(m, n) < 2 {
		return SchoolbookMultiply(a, b)
	}

	if min(m, n) < 3 {
		return KaratsubaMultiply(a, b)
	}

	if min(m, n) < 150 {
		return ToomCook3Multiply(a, b)
	}

	return schoenhageStrassenMultiply(a, b)
}

func Multiply2(a, b digits.Digits) digits.Digits {
	fn := func(a, b digits.Digits) digits.Digits {
		return MultiplySwitch2(a, b, Multiply2)
	}

	return earlyExitMultiply(a, b, fn)
}

func MultiplySwitch2(a, b digits.Digits, fn func(digits.Digits, digits.Digits) digits.Digits) digits.Digits {
	m := a.Length()
	n := b.Length()

	if m == 1 && n == 1 {
		return a.AsDigit().
			Multiply(b.AsDigit()).
			AsDigits().
			Sign(a.IsNegative() != b.IsNegative())
	}

	if m == 2 && n == 1 {
		lo, hi := a.AsDoubleDigit().MultiplyDigit(b.AsDigit())

		result := digits.OfUint64Array([]uint64{
			uint64(lo.Low()),
			uint64(lo.High()),
			uint64(hi),
		})

		return result.
			Sign(a.IsNegative() != b.IsNegative()).
			Trim()
	}

	if m == 1 && n == 2 {
		lo, hi := b.AsDoubleDigit().MultiplyDigit(a.AsDigit())

		result := digits.OfUint64Array([]uint64{
			uint64(lo.Low()),
			uint64(lo.High()),
			uint64(hi),
		})

		return result.
			Sign(a.IsNegative() != b.IsNegative()).
			Trim()
	}

	if m == 2 && n == 2 {
		return MultiplyTwoByTwo(a, b)
	}

	if min(m, n) < 2 {
		return SchoolbookMultiply(a, b)
	}

	if min(m, n) < 3 {
		return KaratsubaMultiply2(a, b, fn)
	}

	if n > m {
		a, b = b, a
	}

	if min(m, n) < 150 {
		return ToomCook3Multiply2(a, b, fn)
	}

	return schoenhageStrassenMultiply(a, b)
}
