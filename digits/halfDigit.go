package digits

type HalfDigit uint32

func MakeDoubleDigitOfDigits(hi Digit, lo Digit) DoubleDigit {
	return DoubleDigitOf(hi, lo)
}

func (a HalfDigit) Multiply(b HalfDigit) Digit {
	return Digit(uint64(a) * uint64(b))
}

func (a HalfDigit) AsDigit() Digit {
	return Digit(a)
}
