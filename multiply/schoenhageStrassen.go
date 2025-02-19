package multiply

import (
	"github.com/borisskert/go-biginteger/digits"
	"github.com/borisskert/go-schoenhageStrassen/ntt/cooleyTukey"
	"github.com/borisskert/go-schoenhageStrassen/schoenhageStrassen"
)

var schoenhageStrassenAlgorithm = schoenhageStrassen.NewSchoenhageStrassen(
	cooleyTukey.IterativeAlgorithm(),
)

func schoenhageStrassenMultiply(a, b digits.Digits) digits.Digits {
	resultArray := schoenhageStrassenAlgorithm.Multiply64(a.AsArray(), b.AsArray())

	sign := a.IsNegative() != b.IsNegative()

	result := digits.OfUint64Array(resultArray)

	if sign && !result.IsZero() {
		return result.Negate()
	}

	return result
}
