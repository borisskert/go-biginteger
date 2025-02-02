package divmod

import "github.com/borisskert/go-biginteger/digits"

func ignoreTrailingZeroBits( // TODO this function is not working
	numerator digits.Digits,
	denominator digits.Digits,
	fn divmodFn,
) (digits.Digits, digits.Digits) {
	length := numerator.Length()

	if length >= 6 {
		trailingZeroBits := min(numerator.TrailingZeros(), denominator.TrailingZeros())

		if trailingZeroBits > 0 {
			q, r := fn(
				numerator.RightShiftBits(trailingZeroBits),
				denominator.RightShiftBits(trailingZeroBits),
			)

			return q.LeftShiftBits(trailingZeroBits), r.LeftShiftBits(trailingZeroBits)
		}
	}

	return fn(numerator, denominator)
}
