package divideDigitsByDigit

import (
	"github.com/borisskert/go-biginteger/digits"
)

type DivideDigitsByDigit struct {
}

func (d DivideDigitsByDigit) DivMod(numerator digits.Digits, denominator digits.Digits) (digits.Digits, digits.Digits) {
	if numerator.IsNegative() {
		panic("numerator must be positive")
	}

	quotient, remainder := numerator.DivideByDigit(denominator.DigitAt(0))
	return quotient, remainder.AsDigits()
}
