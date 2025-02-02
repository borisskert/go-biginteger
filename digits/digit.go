package digits

import (
	"fmt"
	"math/bits"
)

type Digit uint64

func (a Digit) High() HalfDigit {
	return HalfDigit(a >> 32)
}

func (a Digit) Low() HalfDigit {
	return HalfDigit(a)
}

func (a Digit) Split() (HalfDigit, HalfDigit) {
	return HalfDigit(a >> 32), HalfDigit(a)
}

func (a Digit) DivideByHalfDigit(b1 HalfDigit) (HalfDigit, HalfDigit) {
	q, r := bits.Div64(0, uint64(a), uint64(b1))
	return HalfDigit(q), HalfDigit(r)
}

func (a Digit) Subtract(d Digit) (Digit, Digit) {
	lo, borrow := bits.Sub64(uint64(a), uint64(d), 0)
	return Digit(lo), Digit(borrow)
}

func (a Digit) Multiply(b Digit) DoubleDigit {
	hi, lo := bits.Mul64(uint64(a), uint64(b))
	return DoubleDigitOf(Digit(hi), Digit(lo))
}

func (a Digit) Add(b Digit) (Digit, Digit) {
	sum, carry := bits.Add64(uint64(a), uint64(b), 0)
	return Digit(sum), Digit(carry)
}

func (a Digit) IsLessThanDoubleDigit(other DoubleDigit) bool {
	if other.hi > 0 {
		return true
	}

	return a < other.lo
}

func (a Digit) DivideToDigits(b Digit) (Digits, Digits) {
	q, r := bits.Div64(0, uint64(a), uint64(b))
	return Digits{false, []uint64{q}}, Digits{false, []uint64{r}}
}

func (a Digit) Divide(b Digit) (Digit, Digit) {
	q, r := bits.Div64(0, uint64(a), uint64(b))
	return Digit(q), Digit(r)
}

func (a Digit) AsDoubleDigit() DoubleDigit {
	return DoubleDigitOf(0, a)
}

func (a Digit) ShiftLeftToDoubleDigit(shift int) DoubleDigit {
	if shift == 0 {
		return a.AsDoubleDigit()
	}

	if shift >= 128 {
		return DoubleDigitOf(0, 0)
	}

	mod := shift % 64

	if shift >= 64 {
		return DoubleDigitOf(a<<mod, 0)
	}

	return DoubleDigitOf(0|(a>>(64-mod)), a<<mod)
}

func (a Digit) AsDigits() Digits {
	return Digits{false, []uint64{uint64(a)}}
}

func (a Digit) IsEqualDoubleDigit(other DoubleDigit) bool {
	return a == other.lo && other.hi == 0
}

func Zero() Digit {
	return Digit(0)
}

func One() Digit {
	return Digit(1)
}

func OneAsDigit() Digit { // TODO later: One() returns Digit
	return 1
}

func (a Digit) Hexadecimal() string {
	return fmt.Sprintf("0x%X", uint64(a))
}
