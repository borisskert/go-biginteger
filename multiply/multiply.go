package multiply

import "github.com/borisskert/go-biginteger/digits"

func Multiply(a, b []uint64) []uint64 {
	product := MultiplySwitch(digits.Wrap(a), digits.Wrap(b))
	return product.Trim().AsArray()
}

func MultiplySwitch(a, b digits.Digits) digits.Digits {
	if a.Length() < 2 || b.Length() < 2 {
		return SchoolbookMultiply(a, b)
	}

	if a.Length() < 3 || b.Length() < 3 {
		return KaratsubaMultiply(a, b)
	}

	return ToomCook3Multiply(a, b)
}
