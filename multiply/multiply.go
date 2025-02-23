package multiply

import (
	"github.com/borisskert/go-biginteger/digits"
	"github.com/borisskert/go-biginteger/multiply/recursive"
)

var recursiveMultiply = recursive.NewRecursiveMultiplyAlgorithm()

func Multiply(a, b []uint64) []uint64 {
	product := recursiveMultiply.Multiply(
		digits.Wrap(a),
		digits.Wrap(b),
	)
	return product.Trim().AsArray()
}

func MultiplySwitch(a, b digits.Digits) digits.Digits {
	return recursiveMultiply.Multiply(a, b)
}
