package divmod

func divModChunked(a Digits, b Digits) (Digits, Digits) {
	n := int64(b.Length())
	m := int64(a.Length())

	if n == 0 {
		panic("division by zero")
	}
	if m < n {
		return Empty(), a // If dividend is smaller, return (0, remainder)
	}

	quotient := Empty()
	remainder := Empty()

	for start := max(0, m-2*n); start >= 0; start -= 2 * n {
		end := min(m, start+2*n)
		size := end - start

		chunkOfA := a.Chunks(uint64(start), uint64(end)) // Extract chunk

		// Align previous remainder and combine with current chunk
		chunk := chunkOfA.Append(remainder)

		// Perform division on the chunk
		q, r := divModSelect(chunk.Trim(), b.Trim())

		// Accumulate the quotient correctly
		quotient = q.Extend(int(size)).Append(quotient) // TODO use shift add logic here

		remainder = r
	}

	return quotient.Trim(), remainder.Trim()
}
