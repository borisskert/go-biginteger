package divmod

import "github.com/borisskert/go-biginteger/digits"

func divModChunked(a digits.Digits, b digits.Digits) (digits.Digits, digits.Digits) {
	n := int64(b.Length())
	m := int64(a.Length())

	if n == 0 {
		panic("division by zero")
	}
	if m < n {
		return digits.Empty(), a // If dividend is smaller, return (0, remainder)
	}

	quotient := digits.Empty()
	remainder := digits.Empty()

	for start := max(0, m-2*n); start >= 0; start -= 2 * n {
		end := min(m, start+2*n)
		size := end - start

		chunkOfA := a.Chunk(uint(start), uint(end)-1)

		// Align previous remainder and combine with current chunk
		chunk := chunkOfA.Append(remainder)

		// Perform division on the chunk
		q, r := divModSelect(chunk.Trim(), b.Trim())

		// Accumulate the quotient correctly
		quotient = q.EnsureLength(int(size)).Append(quotient) // TODO use shift add logic here

		remainder = r
	}

	return quotient.Trim(), remainder.Trim()
}
