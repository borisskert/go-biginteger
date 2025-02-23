package chunkBasedLongDivision

import (
	"github.com/borisskert/go-biginteger/digits"
	"github.com/borisskert/go-biginteger/divmod/api"
)

type chunkBasedLongDivision struct {
	algorithm api.DivisionAlgorithm
}

func (d *chunkBasedLongDivision) DivMod(numerator, denominator digits.Digits) (digits.Digits, digits.Digits) {
	return divideChunkBased(
		numerator,
		denominator,
		d.algorithm.DivMod,
	)
}

func divideChunkBased(
	numerator digits.Digits,
	denominator digits.Digits,
	fn func(numerator digits.Digits, denominator digits.Digits) (digits.Digits, digits.Digits),
) (digits.Digits, digits.Digits) {
	n := int64(denominator.Length())
	m := int64(numerator.Length())

	if n == 0 {
		panic("division by zero")
	}
	if m < n {
		return digits.Empty(), numerator // If dividend is smaller, return (0, remainder)
	}

	quotient := digits.Empty()
	remainder := digits.Empty()

	for start := max(0, m-2*n); start >= 0; start -= 2 * n {
		end := min(m, start+2*n)
		size := end - start

		chunkOfA := numerator.ChunkInclusive(uint(start), uint(end)-1)

		// Align previous remainder and combine with current chunk
		chunk := chunkOfA.Append(remainder)

		// Perform division on the chunk
		q, r := fn(chunk.Trim(), denominator.Trim())

		// Accumulate the quotient correctly
		quotient = q.LeftShiftDigits(uint(size)).Add(quotient)

		remainder = r
	}

	return quotient.Trim(), remainder.Trim()
}

func DecorateWithChunkBasedLongDivision(algorithm api.DivisionAlgorithm) api.DivisionAlgorithm {
	return &chunkBasedLongDivision{
		algorithm: algorithm,
	}
}
