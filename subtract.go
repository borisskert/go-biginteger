package biginteger

func subtract(minuend BigInteger, subtrahend BigInteger) BigInteger {
	if minuend.IsEqualTo(subtrahend) {
		return zero
	}

	if !minuend.sign && !subtrahend.sign {
		if minuend.IsLessThan(subtrahend) {
			result, _ := subtractUint64Arrays(subtrahend.value, minuend.value, false)

			return BigInteger{
				true,
				result,
			}
		}

		result, _ := subtractUint64Arrays(minuend.value, subtrahend.value, false)

		return BigInteger{
			false,
			result,
		}
	}

	if minuend.sign && subtrahend.sign {
		result := addUint64Arrays(minuend.value, subtrahend.value)

		return BigInteger{
			true,
			result,
		}
	}

	if minuend.sign {
		return subtrahend.Add(minuend.Abs())
	}

	return minuend.Add(subtrahend.Abs())
}

func subtractUint64Arrays(a, b []uint64, borrow bool) ([]uint64, bool) {
	if len(b) > len(a) {
		a, b = b, a
	}

	result := make([]uint64, len(a))
	carry := uint64(0)
	if borrow {
		carry = 1
	}

	for i := 0; i < len(a); i++ {
		ai := a[i]
		bi := uint64(0)
		if i < len(b) {
			bi = b[i]
		}

		diff, newBorrow := subtractUint64(ai, bi+carry)
		result[i] = diff
		carry = 0
		if newBorrow {
			carry = 1
		}
	}

	finalBorrow := carry > 0

	for len(result) > 1 && result[len(result)-1] == 0 {
		result = result[:len(result)-1]
	}

	return result, finalBorrow
}

func subtractUint64(a, b uint64) (uint64, bool) {
	return a - b, a < b
}
