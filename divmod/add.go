package divmod

import "github.com/borisskert/go-biginteger/digits"

func Add(a, b []uint64) []uint64 {
	result := digits.Wrap(a).Add(digits.Wrap(b))
	result.NormalizeInPlace()

	return result.Trim().AsArray()
}
