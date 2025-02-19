package burnikelZiegler

import (
	"github.com/borisskert/go-biginteger/digits"
	"github.com/borisskert/go-biginteger/divmod/api"
	"github.com/borisskert/go-biginteger/multiply"
)

// DecorateWithBurnikelZiegler A Decorator to embed any DivideAlgorithm into Burnikel-Ziegler's Fast Recursive Division.
// See Christoph Burnikel and Joachim Ziegler: Fast Recursive Division, October 1998
func DecorateWithBurnikelZiegler(algorithm api.DivisionAlgorithm) api.DivisionAlgorithm {
	return burnikelZiegler{algorithm}
}

type burnikelZiegler struct {
	algorithm api.DivisionAlgorithm
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
	c := a.Subtract(multiply.MultiplySwitch(q, b1))

	d := multiply.MultiplySwitch(q, b2)

	var r digits.Digits
	if c.IsZero() {
		r = a3.Copy()
	} else {
		r = c.Concat(a3)
	}

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
