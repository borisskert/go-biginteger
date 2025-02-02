package digits

import "math/bits"

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
		//quotient, r := DivTwoDigitsByHalf(d, b.Low())
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

func (d DoubleDigit) Divide(b DoubleDigit) (DoubleDigit, DoubleDigit) {
	if b.IsZero() {
		panic("Division by zero")
	}

	if d.IsLessThan(b) {
		return Digit(0).AsDoubleDigit(), d
	}

	if d.IsEqual(b) {
		return Digit(1).AsDoubleDigit(), Digit(0).AsDoubleDigit()
	}

	if b.High() == 0 {
		quotient, remainder := d.DivideByDigit(b.Low())
		return quotient, remainder.AsDoubleDigit()
	}

	shift := b.LeadingZeros()
	b = b.LeftShift(shift)
	d = d.LeftShift(shift)

	quot, _ := d.High().Divide(b.High())

	prod := quot.Multiply(b.Low())
	prod, _ = prod.Add(quot.Multiply(b.High()).Low().ShiftLeftToDoubleDigit(64))

	if prod.IsGreaterThanOrEqual(d) {
		quot = quot - 1
		prod = quot.Multiply(b.Low())
		prod, _ = prod.Add(quot.Multiply(b.High()).Low().ShiftLeftToDoubleDigit(64))
	}

	rem, _ := d.Subtract(prod)
	rem = rem.RightShift(shift)

	return quot.AsDoubleDigit(), rem
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

func (d DoubleDigit) AddDigit(b Digit) (DoubleDigit, Digit) {
	lo, carry := d.lo.Add(b)
	hi, carry := d.hi.Add(carry)

	return DoubleDigit{hi, lo}, carry
}

func (d DoubleDigit) Add(b DoubleDigit) (DoubleDigit, Digit) {
	lo, carry := d.lo.Add(b.lo)
	hi, carry := d.hi.Add(carry)
	hi, carry = hi.Add(b.hi)

	return DoubleDigit{hi, lo}, carry
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

func (d DoubleDigit) Multiply(b DoubleDigit) (DoubleDigit, DoubleDigit) {
	// Extract high and low parts
	aLo, aHi := d.lo, d.hi
	bLo, bHi := b.lo, b.hi

	// Perform 64-bit multiplications
	lowLow := aLo.Multiply(bLo) // 64-bit x 64-bit = 128-bit result
	lowHigh := aLo.Multiply(bHi)
	highLow := aHi.Multiply(bLo)
	highHigh := aHi.Multiply(bHi)

	// Summing the middle terms
	mid1, carry1 := lowHigh.Add(highLow)           // mid1 = lowHigh + highLow
	mid2, carry2 := mid1.Add(lowLow.LeftShift(64)) // mid2 = mid1 + (lowLow << 64)

	// Carry propagation to high part
	highResult, carry3 := highHigh.AddDigit(mid2.hi)
	highResult, _ = highResult.AddDigit(carry1 + carry2 + carry3) // Propagate all carries

	return highResult, mid2
}

func (d DoubleDigit) MultiplyDigit(b Digit) (DoubleDigit, Digit) {
	lo := d.lo.Multiply(b)
	hi := d.hi.Multiply(b)

	return DoubleDigit{lo.High() + hi.Low(), lo.Low()}, hi.High()
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
