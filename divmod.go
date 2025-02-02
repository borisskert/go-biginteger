package biginteger

import (
	"github.com/borisskert/go-biginteger/divmod"
)

func divMod(a BigInteger, b BigInteger) (BigInteger, BigInteger) {
	if b.IsEqualTo(zero) {
		panic("Division by zero")
	}

	if a.IsEqualTo(zero) {
		return zero, zero
	}

	if b.IsEqualTo(one) {
		return a, zero
	}

	if a.IsEqualTo(b) {
		return one, zero
	}

	quotient, remainder := divmod.DivMod(a.value, b.value)

	quotientSign := false
	if a.sign != b.sign {
		quotientSign = true
	}

	remainderSign := false
	if a.sign {
		remainderSign = true
	}

	return BigInteger{
			sign:  quotientSign,
			value: quotient},
		BigInteger{
			sign:  remainderSign,
			value: remainder}
}
