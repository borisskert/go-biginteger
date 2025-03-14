package toomCook3

import (
	"github.com/borisskert/go-biginteger/digits"
	"github.com/borisskert/go-biginteger/multiply/api"
)

// DecorateWithToomCook3 A Decorator to embed any MultiplyAlgorithm into Toom-Cook-3's Fast Multiplication.
// See André Toom: The Complexity of a Scheme of Functional Elements Realizing the Multiplication of Integers, 1963
// and Robert E. Cook: On the Minimum Computation Time of Functions, 1966
func DecorateWithToomCook3(algorithm api.MultiplyAlgorithm) api.MultiplyAlgorithm {
	return &toomCook3Algorithm{
		algorithm: algorithm,
	}
}

type toomCook3Algorithm struct {
	algorithm api.MultiplyAlgorithm
}

func (t toomCook3Algorithm) Multiply(
	multiplicand digits.Digits, multiplier digits.Digits,
) (product digits.Digits) {
	return toomCook3Multiply(
		multiplicand,
		multiplier,
		t.algorithm.Multiply,
	)
}

func toomCook3Multiply(a, b digits.Digits, fn func(digits.Digits, digits.Digits) digits.Digits) digits.Digits {
	n := max(a.Length(), b.Length())

	k := max((n+2)/3, 1)

	m2, m1, m0 := a.Split3(k)
	n2, n1, n0 := b.Split3(k)

	p0, p1, pm1, pm2, pinf := polynoms(m2, m1, m0)
	q0, q1, qm1, qm2, qinf := polynoms(n2, n1, n0)

	v0 := fn(p0, q0)
	v1 := fn(p1, q1)
	vm1 := fn(pm1, qm1)
	vm2 := fn(pm2, qm2)
	vinf := fn(pinf, qinf)

	r0 := v0
	r4 := vinf
	r3 := vm2.Subtract(v1).DivideByDigitExact(3)
	r1 := v1.Subtract(vm1).DivideByDigitExact(2)
	r2 := vm1.Subtract(v0)
	r3 = r2.Subtract(r3).DivideByDigitExact(2).
		Add(vinf.LeftShiftBits(1))
	r2 = r2.Add(r1).Subtract(r4)
	r1 = r1.Subtract(r3)

	result := r4.LeftShiftDigits(k * 4).
		Add(r3.LeftShiftDigits(k * 3)).
		Add(r2.LeftShiftDigits(k * 2)).
		Add(r1.LeftShiftDigits(k)).
		Add(r0)

	if a.IsNegative() != b.IsNegative() && !result.IsZero() {
		result = result.Negate()
	}

	return result.Trim()
}

func polynoms(x2, x1, x0 digits.Digits) (p0, p1, pm1, pm2, pInf digits.Digits) {
	t0 := x0.Add(x2)

	p0 = x0
	p1 = t0.Add(x1)
	pm1 = t0.Subtract(x1)
	pm2 = pm1.
		Subtract(x1).
		Add(x2.LeftShiftBits(1)).
		Add(x2)
	pInf = x2

	return
}
