package multiply

import "github.com/borisskert/go-biginteger/digits"

func Multiply(a, b []uint64) []uint64 {
	product := schoolbookMultiply(digits.Wrap(a), digits.Wrap(b))
	return product.Trim().AsArray()
}
