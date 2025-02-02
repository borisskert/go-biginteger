package divideDigitsByDigit

import (
	"github.com/borisskert/go-biginteger/digits"
	"github.com/borisskert/go-biginteger/divmod/functions"
)

type DivideDigitsByDigit struct {
}

func (d DivideDigitsByDigit) DivMod(numerator digits.Digits, denominator digits.Digits) (digits.Digits, digits.Digits) {
	quotient, remainder := functions.DivByDigit(numerator, denominator.DigitAt(0))
	return quotient, remainder.AsDigits()
}
