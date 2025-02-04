package multiply

import "github.com/borisskert/go-biginteger/digits"

func Multiply(a, b []uint64) []uint64 {
	product := multiplySwitch(digits.Wrap(a), digits.Wrap(b))
	return product.Trim().AsArray()
}

func multiplySwitch(a, b digits.Digits) digits.Digits {
	if a.Length() < 2 || b.Length() < 2 {
		return schoolbookMultiply(a, b)
	}

	return karatsubaMultiply(a, b)
}
