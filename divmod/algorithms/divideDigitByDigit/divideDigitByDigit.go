package divideDigitByDigit

import (
	"github.com/borisskert/go-biginteger/digits"
)

type DivideDigitByDigit struct {
}

func (d DivideDigitByDigit) DivMod(numerator digits.Digits, denominator digits.Digits) (digits.Digits, digits.Digits) {
	quotient, remainder := numerator.AsDigit().Divide(denominator.AsDigit())
	return quotient.AsDigits(), remainder.AsDigits()
}
