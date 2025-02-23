package schoenhageStrassen

import (
	"github.com/borisskert/go-biginteger/digits"
	"github.com/borisskert/go-biginteger/multiply/api"
	"github.com/borisskert/go-schoenhageStrassen/ntt/cooleyTukey"
	"github.com/borisskert/go-schoenhageStrassen/schoenhageStrassen"
)

type schoenhageStrassenMultiplyAlgorithm struct {
	schoenhageStrassenAlgorithm *schoenhageStrassen.SchoenhageStrassen
}

func (s schoenhageStrassenMultiplyAlgorithm) Multiply(
	multiplicand digits.Digits, multiplier digits.Digits,
) (product digits.Digits) {
	resultArray := s.schoenhageStrassenAlgorithm.Multiply64(multiplicand.AsArray(), multiplier.AsArray())

	sign := multiplicand.IsNegative() != multiplier.IsNegative()

	result := digits.OfUint64Array(resultArray)

	if sign && !result.IsZero() {
		return result.Negate()
	}

	return result
}

func NewSchoenhageStrassenAlgorithm() api.MultiplyAlgorithm {
	return &schoenhageStrassenMultiplyAlgorithm{
		schoenhageStrassenAlgorithm: schoenhageStrassen.NewSchoenhageStrassen(
			cooleyTukey.IterativeAlgorithm(),
		),
	}
}
