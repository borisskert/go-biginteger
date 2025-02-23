package api

import "github.com/borisskert/go-biginteger/digits"

type MultiplyAlgorithm interface {
	Multiply(multiplicand digits.Digits, multiplier digits.Digits) (product digits.Digits)
}
