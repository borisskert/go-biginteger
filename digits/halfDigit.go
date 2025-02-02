package digits

type HalfDigit uint32

func (a HalfDigit) Multiply(b HalfDigit) Digit {
	return Digit(uint64(a) * uint64(b))
}

func (a HalfDigit) AsDigit() Digit {
	return Digit(a)
}
