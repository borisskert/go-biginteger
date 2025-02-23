package recursive

import (
	"github.com/borisskert/go-biginteger/digits"
	"github.com/borisskert/go-biginteger/multiply/api"
	"github.com/borisskert/go-biginteger/multiply/earlyExit"
	"github.com/borisskert/go-biginteger/multiply/karatsuba"
	"github.com/borisskert/go-biginteger/multiply/multiplyOneByOne"
	"github.com/borisskert/go-biginteger/multiply/multiplyTwoByOne"
	"github.com/borisskert/go-biginteger/multiply/multiplyTwoByTwo"
	"github.com/borisskert/go-biginteger/multiply/schoenhageStrassen"
	"github.com/borisskert/go-biginteger/multiply/schoolbook"
	"github.com/borisskert/go-biginteger/multiply/toomCook3"
)

type recursiveMultiplyAlgorithm struct {
	multiplyOneByOne   api.MultiplyAlgorithm
	multiplyTwoByOne   api.MultiplyAlgorithm
	multiplyTwoByTwo   api.MultiplyAlgorithm
	schoolbook         api.MultiplyAlgorithm
	karatsuba          api.MultiplyAlgorithm
	toomCook3          api.MultiplyAlgorithm
	schoenhageStrassen api.MultiplyAlgorithm
}

func (r recursiveMultiplyAlgorithm) Multiply(
	multiplicand digits.Digits,
	multiplier digits.Digits,
) (product digits.Digits) {
	m := multiplicand.Length()
	n := multiplier.Length()

	if n > m {
		multiplicand, multiplier = multiplier, multiplicand
	}

	selectSuitableAlgorithm := r.selectSuitableAlgorithm(multiplicand, multiplier)

	return selectSuitableAlgorithm.Multiply(multiplicand, multiplier)
}

func (r recursiveMultiplyAlgorithm) selectSuitableAlgorithm(
	multiplicand digits.Digits,
	multiplier digits.Digits,
) api.MultiplyAlgorithm {
	m := multiplicand.Length()
	n := multiplier.Length()

	if m == 1 && n == 1 {
		return r.multiplyOneByOne
	}

	if m == 2 && n == 1 {
		return r.multiplyTwoByOne
	}

	if m == 2 && n == 2 {
		return r.multiplyTwoByTwo
	}

	if min(m, n) < 4 {
		return r.schoolbook
	}

	if min(m, n) < 12 {
		return r.karatsuba
	}

	if min(m, n) < 150 {
		return r.toomCook3
	}

	return r.schoenhageStrassen
}

func NewRecursiveMultiplyAlgorithm() api.MultiplyAlgorithm {
	r := &recursiveMultiplyAlgorithm{
		multiplyOneByOne:   multiplyOneByOne.NewMultiplyOneByOneAlgorithm(),
		multiplyTwoByOne:   multiplyTwoByOne.NewMultiplyTwoByOne(),
		multiplyTwoByTwo:   multiplyTwoByTwo.NewMultiplyTwoByTwo(),
		schoolbook:         schoolbook.NewSchoolbookMultiplyAlgorithm(),
		schoenhageStrassen: schoenhageStrassen.NewSchoenhageStrassenAlgorithm(),
	}

	earlyExitAlg := earlyExit.DecorateWithEarlyExit(r)

	r.karatsuba = karatsuba.DecorateWithKaratsuba(earlyExitAlg)
	r.toomCook3 = toomCook3.DecorateWithToomCook3(earlyExitAlg)

	return r
}
