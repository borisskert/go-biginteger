package biginteger

import "strconv"

func parseToBigInteger(s string) (*BigInteger, error) {
	if len(s) == 0 {
		return &Zero, nil
	}

	result := Zero

	for len(s) > 0 {
		first := firstRunes(s, 1)
		parsedDigit, err := strconv.ParseUint(first, 10, 64)
		if err != nil {
			return nil, err
		}

		digit := OfUint64(parsedDigit)
		result = result.Multiply(Ten).Add(digit)
		s = lastRunes(s, len(s)-1)
	}

	return &result, nil
}

func firstRunes(s string, n int) string {
	if n <= 0 {
		return ""
	}

	if n >= len(s) {
		return s
	}

	runes := []rune(s)
	if len(runes) < n {
		return s
	}

	return string(runes[:n])
}

func lastRunes(s string, n int) string {
	runes := []rune(s)
	if len(runes) < n {
		return s
	}

	return string(runes[len(runes)-n:])
}
