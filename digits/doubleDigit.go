package digits

import (
	"math/bits"
)

type DoubleDigit struct {
	hi Digit
	lo Digit
}

func DoubleDigitOf(hi, lo Digit) DoubleDigit {
	return DoubleDigit{
		hi: hi,
		lo: lo,
	}
}

func (d DoubleDigit) High() Digit {
	return d.hi
}

func (d DoubleDigit) Low() Digit {
	return d.lo
}

func (d DoubleDigit) Divide64(b Digit) (Digit, Digit) {
	q, r := bits.Div64(uint64(d.hi), uint64(d.lo), uint64(b))
	return Digit(q), Digit(r)
}

func (d DoubleDigit) DivideByDigit(b Digit) (DoubleDigit, Digit) {
	if b.High() == 0 {
		quotient, r := d.DivideByHalfDigit(b.Low())
		return quotient, r.AsDigit()
	}

	if uint64(b) == 0 {
		panic("division by zero")
	}

	// Step 1: Divide the high 64-bit part
	qHi, r := bits.Div64(0, uint64(d.High()), uint64(b))

	// Step 2: Divide the lower 64-bit part using the remainder from step 1
	qLo, r := bits.Div64(r, uint64(d.Low()), uint64(b))

	return DoubleDigitOf(Digit(qHi), Digit(qLo)), Digit(r)
}

func (d DoubleDigit) DivideByHalfDigit(b HalfDigit) (DoubleDigit, HalfDigit) {
	if b == 0 {
		panic("Division by zero")
	}

	v := b.AsDigit()

	qHi, r := bits.Div64(0, uint64(d.High()), uint64(v))
	qLo, r := bits.Div64(r, uint64(d.Low()), uint64(v))

	return DoubleDigitOf(Digit(qHi), Digit(qLo)), HalfDigit(r)
}

func (a DoubleDigit) Divide(b DoubleDigit) (DoubleDigit, DoubleDigit) {
	// Handle division by zero
	if b.hi == 0 && b.lo == 0 {
		panic("division by zero")
	}

	// If the divisor is larger than the dividend, the quotient is 0 and the remainder is the dividend.
	if a.IsLessThan(b) {
		return DoubleDigit{hi: 0, lo: 0}, a
	}

	// If the divisor is a single uint64 (Hi == 0), use a simpler algorithm.
	if b.hi == 0 {
		q, r := a.DivideByDigit(b.Low())
		return q, r.AsDoubleDigit()
	}

	// Perform long division for the general case.
	var quotient DoubleDigit
	remainder := a

	for !remainder.IsLessThan(b) {
		// Calculate how many times we can subtract the divisor from the remainder.
		shift := max(int(remainder.LeadingZeros())-int(b.LeadingZeros()), 0)
		shiftedDivisor := b.LeftShift(uint(shift))

		if shiftedDivisor.IsLessThan(remainder) || shiftedDivisor.IsEqual(remainder) {
			quotient, _ = quotient.Add(DoubleDigit{lo: 1}.LeftShift(uint(shift)))
			remainder, _ = remainder.Subtract(shiftedDivisor)
		} else {
			shift--
			shiftedDivisor = b.LeftShift(uint(shift))
			quotient, _ = quotient.Add(DoubleDigit{lo: 1}.LeftShift(uint(shift)))
			remainder, _ = remainder.Subtract(shiftedDivisor)
		}
	}

	return quotient, remainder
}

func (d DoubleDigit) IsGreaterThanOrEqual(other DoubleDigit) bool {
	if d.hi > other.hi {
		return true
	}

	if d.hi == other.hi {
		return d.lo >= other.lo
	}

	return false
}

func (d DoubleDigit) Subtract(b DoubleDigit) (DoubleDigit, Digit) {
	lo, carry := bits.Sub64(uint64(d.lo), uint64(b.lo), 0)
	hi, carry := bits.Sub64(uint64(d.hi), uint64(b.hi), carry)

	return DoubleDigit{Digit(hi), Digit(lo)}, Digit(carry)
}

func (d DoubleDigit) AddDigitIgnoreOverflow(b Digit) DoubleDigit {
	sum, _ := d.AddDigit(b)
	return sum
}

func (d DoubleDigit) AddDigit(b Digit) (DoubleDigit, Digit) {
	lo, carry := d.lo.Add(b)
	hi, carry := d.hi.Add(carry)

	return DoubleDigit{hi, lo}, carry
}

func (d DoubleDigit) Add(b DoubleDigit) (DoubleDigit, Digit) {
	lo, carry := bits.Add64(uint64(d.lo), uint64(b.lo), 0)
	hi, carry := bits.Add64(uint64(d.hi), uint64(b.hi), carry)

	return DoubleDigit{Digit(hi), Digit(lo)}, Digit(carry)
}

func (d DoubleDigit) IsLessThan(b DoubleDigit) bool {
	if d.hi < b.hi {
		return true
	}

	if d.hi == b.hi {
		return d.lo < b.lo
	}

	return false
}

func (d DoubleDigit) IsZero() bool {
	return d.hi == 0 && d.lo == 0
}

func (d DoubleDigit) LeadingZeros() uint {
	if d.hi == 0 {
		return uint(bits.LeadingZeros64(uint64(d.lo)))
	}

	return uint(bits.LeadingZeros64(uint64(d.hi)))
}

func (d DoubleDigit) LeftShift(shift uint) DoubleDigit {
	if shift == 0 {
		return d
	}

	if shift >= 128 {
		return DoubleDigitOf(0, 0)
	}

	mod := shift % 64

	if shift >= 64 {
		return DoubleDigitOf(d.lo<<mod, 0)
	}

	return DoubleDigitOf((d.hi<<mod)|(d.lo>>(64-mod)), d.lo<<mod)
}

func (d DoubleDigit) RightShift(shift uint) DoubleDigit {
	if shift == 0 {
		return d
	}

	if shift >= 128 {
		return DoubleDigitOf(0, 0)
	}

	if shift >= 64 {
		return DoubleDigitOf(0, d.hi>>(shift-64))
	}

	mod := shift % 64

	return DoubleDigitOf((d.hi>>mod)|(d.lo<<(64-mod)), d.lo>>mod)
}

func (d DoubleDigit) AsDigits() Digits {
	if d.hi == 0 {
		return Digits{false, []uint64{uint64(d.lo)}}
	}

	return Digits{false, []uint64{uint64(d.lo), uint64(d.hi)}}
}

func (d DoubleDigit) IsEqual(v DoubleDigit) bool {
	return d.hi == v.hi && d.lo == v.lo
}

func (d DoubleDigit) IsOne() bool {
	return d.hi == 0 && d.lo == 1
}

func (d DoubleDigit) SubtractDigit(b Digit) (DoubleDigit, Digit) {
	lo, borrow := d.lo.Subtract(b)
	hi, borrow := d.hi.Subtract(borrow)

	return DoubleDigit{hi, lo}, borrow
}

func (d DoubleDigit) MultiplyIgnoreOverflow(b DoubleDigit) DoubleDigit {
	rHi, rLo := bits.Mul64(uint64(d.lo), uint64(b.lo))

	_, lo := bits.Mul64(uint64(d.lo), uint64(b.hi))
	rHi, _ = bits.Add64(rHi, lo, 0)

	_, lo = bits.Mul64(uint64(d.hi), uint64(b.lo))
	rHi, _ = bits.Add64(rHi, lo, 0)

	return DoubleDigitOf(
		Digit(rHi),
		Digit(rLo),
	)
}

func (d DoubleDigit) Multiply(b DoubleDigit) (DoubleDigit, DoubleDigit) {
	result := [4]uint64{}
	carry := uint64(0)

	hi, lo := bits.Mul64(uint64(d.lo), uint64(b.lo))

	result[0], carry = bits.Add64(result[0], lo, 0)
	result[1], carry = bits.Add64(result[1], hi, carry)
	result[2], carry = bits.Add64(result[2], carry, 0)

	hi, lo = bits.Mul64(uint64(d.lo), uint64(b.hi))

	result[1], carry = bits.Add64(result[1], lo, 0)
	result[2], carry = bits.Add64(result[2], hi, carry)
	result[3], carry = bits.Add64(result[3], carry, 0)

	hi, lo = bits.Mul64(uint64(d.hi), uint64(b.lo))

	result[1], carry = bits.Add64(result[1], lo, 0)
	result[2], carry = bits.Add64(result[2], hi, carry)
	result[3], carry = bits.Add64(result[3], carry, 0)

	hi, lo = bits.Mul64(uint64(d.hi), uint64(b.hi))

	result[2], carry = bits.Add64(result[2], lo, 0)
	result[3], _ = bits.Add64(result[3], hi, carry)

	return DoubleDigitOf(
			Digit(result[3]),
			Digit(result[2]),
		), DoubleDigitOf(
			Digit(result[1]),
			Digit(result[0]),
		)
}

func (d DoubleDigit) MultiplyDigit(b Digit) (DoubleDigit, Digit) { // TODO return hi,lo in this order
	mLo := d.lo.Multiply(b)
	mHi := d.hi.Multiply(b)

	u, v := mHi.High(), mHi.Low()
	w, x := mLo.High(), mLo.Low()

	//   u, v
	// +    w, x
	// -----------
	//   z, y, x

	y, carryMid := v.Add(w)
	z, carryHi := u.Add(carryMid)

	if carryHi != 0 {
		panic("unexpected carry")
	}

	return DoubleDigit{y, x}, z
}

func (d DoubleDigit) IsGreaterThan(doubleDigit DoubleDigit) bool {
	if d.hi > doubleDigit.hi {
		return true
	}

	if d.hi == doubleDigit.hi {
		return d.lo > doubleDigit.lo
	}

	return false
}

func (d DoubleDigit) Decrement() (DoubleDigit, bool) {
	if d.lo == 0 && d.hi == 0 {
		return DoubleDigit{^Digit(0), ^Digit(0)}, true
	}

	if d.lo == 0 {
		return DoubleDigit{d.hi - 1, ^Digit(0)}, false
	}

	return DoubleDigit{d.hi, d.lo - 1}, false
}

func (d DoubleDigit) IsNonZero() bool {
	return d.hi != 0 || d.lo != 0
}
