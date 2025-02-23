package schoenhageStrassen

import (
	"github.com/borisskert/go-biginteger/digits"
	"github.com/borisskert/go-biginteger/multiply/api"
	"github.com/borisskert/go-schoenhageStrassen/ntt/cooleyTukey"
	"github.com/borisskert/go-schoenhageStrassen/schoenhageStrassen"
)

// NewSchoenhageStrassenAlgorithm creates a new instance of the Schoenhage-Strassen multiply algorithm.
// See Schönhage, A., & Strassen, V. (1971). Schnelle Multiplikation großer Zahlen. Computing, 7(3), 281–292.
// It uses the Cooley-Tukey algorithm for the NTT.
// See Cooley, J. W., & Tukey, J. W. (1965). An algorithm for the machine calculation of complex Fourier series.
func NewSchoenhageStrassenAlgorithm() api.MultiplyAlgorithm {
	return &schoenhageStrassenMultiplyAlgorithm{
		schoenhageStrassenAlgorithm: schoenhageStrassen.NewSchoenhageStrassen(
			cooleyTukey.IterativeAlgorithm(),
		),
	}
}

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
