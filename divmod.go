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

func divModUint64Arrays(a []uint64, b []uint64) ([]uint64, []uint64) {
	if compareUint64Arrays(b, []uint64{0}) == 0 {
		panic("Division by zero")
	}

	if compareUint64Arrays(a, []uint64{0}) == 0 {
		return []uint64{0}, []uint64{0}
	}

	if compareUint64Arrays(b, []uint64{1}) == 0 {
		return a, []uint64{0}
	}

	if compareUint64Arrays(a, b) < 0 {
		return []uint64{0}, a
	}

	result := []uint64{0}
	remaining := a

	divisor := b
	quotient := []uint64{1}

	for compareUint64Arrays(divisor, remaining) <= 0 {
		divisor = shiftLeftUint64Array(divisor, 1)
		quotient = shiftLeftUint64Array(quotient, 1)
	}

	for compareUint64Arrays(divisor, b) >= 0 {
		if compareUint64Arrays(remaining, divisor) >= 0 {
			remaining = subtractUint64Arrays(remaining, divisor)
			result = addUint64Arrays(result, quotient)
		}

		divisor = shiftRightUint64Array(divisor, 1)
		quotient = shiftRightUint64Array(quotient, 1)
	}

	return normalize(result), normalize(remaining)
}

func normalize(value []uint64) []uint64 {
	if len(value) == 0 {
		return []uint64{0}
	}

	return trimLeadingZeros(value)
}
