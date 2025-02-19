package api

import "github.com/borisskert/go-biginteger/digits"

type DivisionAlgorithm interface {
	DivMod(numerator digits.Digits, denominator digits.Digits) (quotient digits.Digits, remainder digits.Digits)
}
