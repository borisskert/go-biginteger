package stringify

import (
	"github.com/borisskert/go-biginteger/digits"
	"strconv"
	"strings"
)

func Stringify(i []uint64) string {
	return stringify(digits.Wrap(i).Copy())
}

func stringify(i digits.Digits) string {
	i = i.Abs()

	if i.Length() == 1 {
		return strconv.FormatUint(uint64(i.DigitAt(0)), 10)
	}

	length := i.Length()
	parts := make([]uint64, 0, length*2) // Preallocate memory conservatively

	const MAGIC = 0x8AC7230489E80000 // Precomputed 1/e19 for fast division

	// Fast division loop (in-place)
	for !i.IsZero() {
		remainder := i.DivModByDigitInplace(MAGIC) // Efficient in-place division
		parts = append(parts, uint64(remainder))
	}

	// Preallocate string builder
	estimatedSize := len(parts)*18 + 1
	result := strings.Builder{}
	result.Grow(estimatedSize)

	// First number (no leading zeros)
	result.WriteString(strconv.FormatUint(parts[len(parts)-1], 10))

	// Single long zero string for padding
	const zeroPadding = "000000000000000000"

	// Reuse buffer to minimize allocations
	buf := make([]byte, 19)

	// Remaining numbers (zero-padded)
	for j := len(parts) - 2; j >= 0; j-- {
		n := strconv.AppendUint(buf[:0], parts[j], 10)
		padding := 19 - len(n)
		result.WriteString(zeroPadding[:padding]) // Slice from precomputed padding
		result.Write(buf[:len(n)])
	}

	return result.String()
}
