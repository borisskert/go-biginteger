package biginteger

func subtract(a BigInteger, b BigInteger) BigInteger {
	if a.IsEqualTo(b) {
		return zero
	}

	if !a.sign && !b.sign {
		if a.IsLessThan(b) {
			result, _ := subtractUint64Arrays(b.value, a.value, false)

			return BigInteger{
				true,
				result,
			}
		}

		result, _ := subtractUint64Arrays(a.value, b.value, false)

		return BigInteger{
			false,
			result,
		}
	}

	if a.sign && b.sign {
		result := addUint64Arrays(a.value, b.value)

		return BigInteger{
			true,
			result,
		}
	}

	if a.sign {
		return b.Add(a.Abs())
	}

	return a.Add(b.Abs())
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
