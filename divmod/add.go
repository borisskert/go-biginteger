package divmod

func Add(a, b []uint64) []uint64 {
	result := Wrap(a).Add(Wrap(b))
	result.NormalizeInPlace()

	return result.Trim().Array()
}
