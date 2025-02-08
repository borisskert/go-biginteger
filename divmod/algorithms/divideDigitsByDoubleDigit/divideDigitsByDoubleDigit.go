package divideDigitsByDoubleDigit

import (
	"github.com/borisskert/go-biginteger/digits"
	"github.com/borisskert/go-biginteger/divmod/functions"
)

type DivideDigitsByDoubleDigit struct {
}

func (d DivideDigitsByDoubleDigit) DivMod(numerator digits.Digits, denominator digits.Digits) (digits.Digits, digits.Digits) {
	if numerator.IsNegative() {
		panic("numerator must be positive")
	}

	quotient, remainder := functions.DivByDoubleDigit(numerator, denominator.AsDoubleDigit())
	return quotient, remainder.AsDigits()
}
