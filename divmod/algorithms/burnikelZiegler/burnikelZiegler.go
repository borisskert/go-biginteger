package burnikelZiegler

import (
	"github.com/borisskert/go-biginteger/digits"
	"github.com/borisskert/go-biginteger/divmod/common"
)

// A Decorator for any DivideAlgorithm to use the Burnikel-Ziegler algorithm.
type burnikelZiegler struct {
	algorithm common.DivisionAlgorithm
}

func (bz burnikelZiegler) DivMod(
	numerator digits.Digits, denominator digits.Digits,
) (
	quotient digits.Digits, remainder digits.Digits,
) {
	return divModBurnikelZiegler(numerator, denominator, bz.algorithm.DivMod)
}

func divThreeLongHalvesByTwo(
	a1, a2, a3, b1, b2 digits.Digits,
	fn func(digits.Digits, digits.Digits) (digits.Digits, digits.Digits),
) (digits.Digits, digits.Digits) {
	b := b1.Concat(b2)
	a := a1.Concat(a2)

	q, _ := fn(a.Trim(), b1.Trim())
	c, _ := a.Subtract(q.Multiply(b1))

	d := q.Multiply(b2)
	r := c.LeftShiftDigits(a3.Length()).Add(a3)
	r.SubtractInPlace(d)

	if r.IsNegative() {
		q.DecrementInPlace()
		r.AddInPlace(b)

		if r.IsNegative() {
			q.DecrementInPlace()
			r.AddInPlace(b)

			if r.IsNegative() {
				panic("This should never happen: r was negative twice")
			}
		}
	}

	return q, r
}

func divModBurnikelZiegler(
	a digits.Digits, b digits.Digits,
	fn func(digits.Digits, digits.Digits) (digits.Digits, digits.Digits),
) (digits.Digits, digits.Digits) {
	n := b.Length()

	if a.Length() != 2*n {
		panic("Burnikel-Ziegler's precondition not met: " +
			"a's length must be 2n, b's length must be n")
	}

	a1, a2, a3, a4 := a.Quarter()
	b1, b2 := b.Halve()

	q1, r := divThreeLongHalvesByTwo(a1, a2, a3, b1, b2, fn)

	r1, r2 := r.Halve()

	q2, s := divThreeLongHalvesByTwo(r1, r2, a4, b1, b2, fn)

	q := q1.LeftShiftDigits(n / 2).Add(q2)

	return q.Trim(), s.Trim()
}

func DecorateWithBurnikelZiegler(algorithm common.DivisionAlgorithm) common.DivisionAlgorithm {
	return burnikelZiegler{algorithm}
}
