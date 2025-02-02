package recursive

import (
	"github.com/borisskert/go-biginteger/digits"
	"github.com/borisskert/go-biginteger/divmod/algorithms/burnikelZiegler"
	"github.com/borisskert/go-biginteger/divmod/algorithms/chunkedDivision"
	"github.com/borisskert/go-biginteger/divmod/algorithms/divideDigitByDigit"
	"github.com/borisskert/go-biginteger/divmod/algorithms/divideDigitsByDigit"
	"github.com/borisskert/go-biginteger/divmod/algorithms/divideDigitsByDoubleDigit"
	"github.com/borisskert/go-biginteger/divmod/algorithms/divideDoubleDigitByDigit"
	"github.com/borisskert/go-biginteger/divmod/algorithms/divisionShortcut"
	"github.com/borisskert/go-biginteger/divmod/algorithms/donaldKnuth"
	"github.com/borisskert/go-biginteger/divmod/algorithms/trailingZeroReduction"
	"github.com/borisskert/go-biginteger/divmod/common"
)

type recursiveDivisionAlgorithm struct {
	divideOneByOne         divideDigitByDigit.DivideDigitByDigit
	divideTwoByOne         divideDoubleDigitByDigit.DivideDoubleDigitByDigit
	divideManyByOne        divideDigitsByDigit.DivideDigitsByDigit
	divideManyByTwo        divideDigitsByDoubleDigit.DivideDigitsByDoubleDigit
	donaldKnuthsAlgorithmD donaldKnuth.DonaldKnuthsAlgorithmD
}

func (f *recursiveDivisionAlgorithm) DivMod(
	numerator digits.Digits, denominator digits.Digits,
) (quotient digits.Digits, remainder digits.Digits) {
	algorithm := f.selectSuitableDivideAlgorithm(numerator, denominator)
	algorithm = divisionShortcut.DecorateWithShortcut(algorithm)
	algorithm = trailingZeroReduction.DecorateWithTrailingZeroReduction(algorithm)

	return algorithm.DivMod(numerator, denominator)
}

func (f *recursiveDivisionAlgorithm) selectSuitableDivideAlgorithm(
	numerator,
	denominator digits.Digits,
) common.DivisionAlgorithm {
	n := denominator.Length()
	m := numerator.Length()

	if n == 1 && m == 1 {
		return f.divideOneByOne
	}

	if n == 1 && m == 2 {
		return f.divideTwoByOne
	}

	if n == 1 {
		return f.divideManyByOne
	}

	if n == 2 {
		return f.divideManyByTwo
	}

	defaultAlgorithmFn := func() common.DivisionAlgorithm {
		return divisionShortcut.DecorateWithShortcut(f)
	}

	if m < 40 {
		return &f.donaldKnuthsAlgorithmD
	}

	if m == 2*n {
		return burnikelZiegler.DecorateWithBurnikelZiegler(defaultAlgorithmFn)
	}

	if m > 2*n {
		return chunkedDivision.DecorateWithChunkedDivision(defaultAlgorithmFn)
	}

	return &f.donaldKnuthsAlgorithmD
}

func NewRecursiveDivisionAlgorithm() common.DivisionAlgorithm {
	return &recursiveDivisionAlgorithm{}
}
