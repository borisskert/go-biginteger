package multiply

import "github.com/borisskert/go-biginteger/digits"

func earlyExitMultiply(a, b digits.Digits, fn func(digits.Digits, digits.Digits) digits.Digits) digits.Digits {
	if a.IsZero() || b.IsZero() {
		return digits.Zero().AsDigits()
	}

	if a.IsOne() {
		return b
	}

	if b.IsOne() {
		return a
	}

	if b.Abs().IsOne() {
		return a.Negate()
	}

	if a.Abs().IsOne() {
		return b.Negate()
	}

	return fn(a, b)
}
