package divmod

import "github.com/borisskert/go-biginteger/digits"

func Subtract(a, b []uint64) []uint64 {
	result, _ := digits.Wrap(a).Subtract(digits.Wrap(b))
	result.NormalizeInPlace()

	return result.AsArray()
}
