package divideDigitsByDoubleDigit

import (
	"github.com/borisskert/go-biginteger/digits"
)

type DivideDigitsByDoubleDigit struct {
}

func (d DivideDigitsByDoubleDigit) DivMod(numerator digits.Digits, denominator digits.Digits) (digits.Digits, digits.Digits) {
	if numerator.IsNegative() {
		panic("numerator must be positive")
	}

	quotient, remainder := numerator.DivideByDoubleDigit(denominator.AsDoubleDigit())
	return quotient, remainder.AsDigits()
}
