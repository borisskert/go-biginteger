package divmod

func Subtract(a, b []uint64) []uint64 {
	result, _ := Wrap(a).Subtract(Wrap(b))
	result.NormalizeInPlace()

	return result.Array()
}
