package uintArray

func Compare(a, b []uint64) int {
	sizeA := len(a)
	sizeB := len(b)

	if sizeA < sizeB {
		for i := sizeA; i < sizeB; i++ {
			if b[i] != 0 {
				return -1
			}
		}
	}

	if sizeA > sizeB {
		for i := sizeB; i < sizeA; i++ {
			if a[i] != 0 {
				return 1
			}
		}
	}

	for k := min(sizeA, sizeB) - 1; k > 0; k-- {
		if a[k] > b[k] {
			return 1
		}

		if a[k] < b[k] {
			return -1
		}
	}

	if a[0] > b[0] {
		return 1
	}

	if a[0] < b[0] {
		return -1
	}

	return 0
}
