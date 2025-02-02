package divideDoubleDigitByDigit

import (
	"github.com/borisskert/go-biginteger/digits"
)

type DivideDoubleDigitByDigit struct {
}

func (d DivideDoubleDigitByDigit) DivMod(numerator digits.Digits, denominator digits.Digits) (digits.Digits, digits.Digits) {
	quotient, remainder := numerator.AsDoubleDigit().DivideByDigit(denominator.AsDigit())
	return quotient.AsDigits(), remainder.AsDigits()
}
