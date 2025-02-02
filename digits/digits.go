package digits

import (
	"fmt"
	"github.com/borisskert/go-biginteger/uintArray"
	"math/bits"
	"strings"
)

type Digits struct {
	sign  bool
	value []uint64
}

func (a Digits) Length() uint {
	return uint(len(a.value))
}

func (a Digits) BitLength() uint {
	if len(a.value) == 0 {
		return 0
	}

	last := uint(bits.Len64(a.value[len(a.value)-1]))

	return uint(len(a.value)-1)*64 + last
}

func (a Digits) LeftShiftDigits(n uint) Digits {
	result := make([]uint64, uint(len(a.value))+n)
	copy(result[n:], a.value)
	return Digits{a.sign, result}
}

func (a Digits) RightShiftDigits(n uint) Digits {
	if n >= uint(len(a.value)) {
		return Zero().AsDigits()
	}

	return Digits{a.sign, a.value[n:]}
}

func (a Digits) LeftShiftBits(n uint) Digits {
	array := uintArray.ShiftLeftBits(a.value, n)
	return Digits{a.sign, array}
}

func (a *Digits) LeftShiftBitsInPlace(n uint) {
	a.value = uintArray.ShiftLeftBits(a.value, n)
}

func (a Digits) RightShiftBits(n uint) Digits {
	result := a.Copy()
	result.RightShiftBitsInPlace(n)

	return result
}

func (a *Digits) RightShiftBitsInPlace(n uint) { // TODO is this working?
	shift := n % 64
	shifts := n / 64
	sizeA := uint(len(a.value))

	if shifts >= sizeA {
		a.value = []uint64{0}
		return
	}

	newSize := sizeA - shifts - 1

	for i := uint(0); i < newSize; i++ {
		oldIndexLo := i + shifts
		oldIndexHi := oldIndexLo + 1

		a1 := a.value[oldIndexLo]
		a2 := uint64(0)

		if oldIndexHi < sizeA {
			a2 = a.value[oldIndexHi]
		}

		a.value[i] = (a1 >> shift) | (a2 << (64 - shift))
	}

	a.value[newSize] = a.value[sizeA-1] >> shift
	a.value = a.value[:sizeA-shifts]
}

func (a Digits) Add(b Digits) Digits {
	result := a.Copy()
	result.AddInPlace(b)

	return result
}

func (a *Digits) AddInPlace(b Digits) bool {
	if a.sign == b.sign {
		return a.addInPlaceAbs(b)
	}

	return a.SubtractInPlace(b.Negate())
}

func (a *Digits) addInPlaceAbs(b Digits) bool {
	sizeA := len(a.value)
	sizeB := len(b.value)

	size := max(sizeA, sizeB)

	if sizeA < sizeB {
		a.value = append(a.value, make([]uint64, sizeB-sizeA)...)
	}

	carry := uint64(0)
	var sum uint64

	for i := 0; i < size; i++ {
		summandA := uint64(0)
		if i < len(a.value) {
			summandA = a.value[i]
		}

		summandB := uint64(0)
		if i < len(b.value) {
			summandB = b.value[i]
		}

		sum, carry = bits.Add64(summandA, summandB, carry)
		a.value[i] = sum
	}

	if carry == 1 {
		a.value = append(a.value, 1)
	}

	return carry == 1
}

func (a Digits) Multiply(b Digits) Digits {
	if len(a.value) == 0 || len(b.value) == 0 {
		return Digits{false, []uint64{0}}
	}

	result := uintArray.MultiplyUint64Array(a.value, b.value)

	digits := Digits{a.sign != b.sign, result}
	digits.NormalizeInPlace()

	return digits
}

func (a Digits) Subtract(b Digits) (Digits, bool) {
	result := a.Copy()
	result.SubtractInPlace(b)

	return result, result.IsNegative()
}

func (a *Digits) SubtractInPlace(b Digits) bool {
	if a.sign != b.sign {
		return a.AddInPlace(b.Negate())
	}

	if a.compareAbs(b) < 0 {
		a.value, b.value = b.value, a.value
		a.sign = !a.sign
	}

	carry := uint64(0)
	size := max(len(a.value), len(b.value))

	for i := 0; i < size; i++ {
		ai := uint64(0)
		if i < len(a.value) {
			ai = a.value[i]
		}

		bi := uint64(0)
		if i < len(b.value) {
			bi = b.value[i]
		}

		diff, borrow := bits.Sub64(ai, bi+carry, 0)

		if i >= len(a.value) {
			a.value = append(a.value, diff)
		} else {
			a.value[i] = diff
		}

		carry = borrow
	}

	return carry == 1
}

func (a Digits) SubtractUnderflow(b Digits) (Digits, bool) {
	result := a.Copy()
	borrowed := result.SubtractUnderflowInPlace(b)

	return result, borrowed
}

func (a *Digits) SubtractUnderflowInPlace(b Digits) bool {
	needComplement := a.Compare(b) < 0

	size := max(len(a.value), len(b.value))
	a.value = append(a.value, make([]uint64, size-len(a.value))...)
	var carry uint64

	for i := 0; i < size; i++ {
		ai := uint64(0)
		if i < len(a.value) {
			ai = a.value[i]
		}

		bi := uint64(0)
		if i < len(b.value) {
			bi = b.value[i]
		}

		diff, borrow := bits.Sub64(ai, bi+carry, 0)

		if i >= len(a.value) {
			a.value = append(a.value, diff)
		} else {
			a.value[i] = diff
		}

		carry = borrow
	}

	return needComplement
}

func (a *Digits) TrimInPlace() {
	size := len(a.value)

	for i := size - 1; i >= 0; i-- {
		if a.value[i] != 0 {
			break
		}

		size--
	}

	if size == 0 {
		a.value = []uint64{0}
		return
	}

	a.value = a.value[:size]
}

func (a *Digits) NormalizeInPlace() {
	if len(a.value) == 0 {
		a.value = []uint64{0}
	}

	a.TrimInPlace()
}

func (a Digits) Concat(b Digits) Digits {
	return a.LeftShiftDigits(b.Length()).Add(b)
}

func (a Digits) IsNegative() bool {
	return a.sign
}

func (a Digits) Quarter() (Digits, Digits, Digits, Digits) {
	lenA := len(a.value)

	quarterLen := lenA / 4

	return Digits{false, a.value[3*quarterLen:]},
		Digits{false, a.value[2*quarterLen : 3*quarterLen]},
		Digits{false, a.value[quarterLen : 2*quarterLen]},
		Digits{false, a.value[:quarterLen]}
}

func (a Digits) Halve() (Digits, Digits) {
	lenA := len(a.value)

	halfLen := lenA / 2

	return Digits{false, a.value[halfLen:]},
		Digits{false, a.value[:halfLen]}
}

func (a Digits) IsZero() bool {
	return len(a.value) == 0 || len(a.value) == 1 && a.value[0] == 0
}

func (a Digits) IsOne() bool {
	return !a.sign && len(a.value) == 1 && a.value[0] == 1
}

func (a Digits) Compare(other Digits) int {
	if a.sign && !other.sign {
		return -1
	}

	if !a.sign && other.sign {
		return 1
	}

	if a.sign && other.sign {
		return other.compareAbs(a)
	}

	return a.compareAbs(other)
}

func (a Digits) compareAbs(other Digits) int {
	return uintArray.Compare(a.value, other.value)
}

func (a Digits) IsLessThan(denominator Digits) bool {
	return a.Compare(denominator) < 0
}

func Wrap(values []uint64) Digits {
	return Digits{false, values}
}

func (a Digits) Copy() Digits {
	copiedArray := make([]uint64, len(a.value))
	copy(copiedArray, a.value)

	return Digits{
		sign:  a.sign,
		value: copiedArray,
	}
}

func (a Digits) Trim() Digits {
	result := a.Copy()
	result.TrimInPlace()

	return result
}

func (a Digits) MultiplyByDigit(b Digit) Digits {
	return a.Multiply(Digits{false, []uint64{uint64(b)}})
}

func (a Digits) LeadingZeros() uint64 {
	if len(a.value) == 0 {
		return 0
	}

	return 64 - uint64(bits.Len64(a.value[len(a.value)-1]))
}

func (a Digits) DigitAt(n uint) Digit {
	if n >= uint(len(a.value)) {
		return 0
	}

	return Digit(a.value[n])
}

func (a Digits) MostSignificantDigit() Digit {
	if len(a.value) == 0 {
		return 0
	}

	return Digit(a.value[len(a.value)-1])
}

func (a *Digits) SetDigitAt(position uint, b Digit) {
	lenA := len(a.value)

	if position >= uint(lenA) {
		a.value = append(a.value, make([]uint64, int(position)-lenA+1)...)
	}

	a.value[position] = uint64(b)
}

func (a Digits) Append(b Digits) Digits {
	result := make([]uint64, len(a.value)+len(b.value))
	copy(result, a.value)
	copy(result[len(a.value):], b.value)

	return Digits{a.sign, result}
}

func (a Digits) TrailingZeros() uint { // TODO this logic is weird
	if len(a.value) == 0 {
		return 0
	}

	zeros := uint(0)

	for i := 0; i < len(a.value); i++ {
		if a.value[i] == 0 {
			zeros += 64
		} else {
			zeros += uint(bits.TrailingZeros64(a.value[i]))
			break
		}
	}

	return zeros
}

func (a Digits) AsArray() []uint64 {
	return a.value
}

func (a Digits) MultiplyByDoubleDigit(b DoubleDigit) Digits {
	result := make([]uint64, len(a.value)+2)

	for i := 0; i < len(a.value); i++ {
		loHi, loLo := bits.Mul64(a.value[i], uint64(b.lo))

		loLo += result[i]
		if loLo < result[i] {
			loHi++
		}
		result[i] = loLo

		hiHi, hiLo := bits.Mul64(a.value[i], uint64(b.hi))

		hiLo += result[i+1]
		if hiLo < result[i+1] {
			hiHi++
		}
		result[i+1] = hiLo

		result[i+2] += loHi + hiHi
	}

	return Digits{a.sign, result}
}

func (a Digits) AsDoubleDigit() DoubleDigit {
	if len(a.value) == 0 {
		return DoubleDigitOf(0, 0)
	}

	if len(a.value) == 1 {
		return DoubleDigitOf(0, Digit(a.value[0]))
	}

	return DoubleDigitOf(Digit(a.value[1]), Digit(a.value[0]))
}

func (a Digits) AsDigit() Digit {
	if len(a.value) == 0 {
		return 0
	}

	return Digit(a.value[0])
}

func (a Digits) AddDoubleDigit(b DoubleDigit) Digits {
	result := a.Copy()
	result.AddDoubleDigitInPlace(b)

	return result
}

func (a *Digits) AddDoubleDigitInPlace(b DoubleDigit) {
	if b.IsZero() {
		return
	}

	if len(a.value) == 0 {
		a.value = []uint64{uint64(b.Low())}
		if b.High() != 0 {
			a.value = append(a.value, uint64(b.High()))
		}
		return
	}

	var carry uint64

	a.value[0], carry = bits.Add64(a.value[0], uint64(b.Low()), 0)

	if len(a.value) > 1 {
		a.value[1], carry = bits.Add64(a.value[1], uint64(b.High()), carry)
	} else if b.High() != 0 || carry != 0 {
		a.value = append(a.value, uint64(b.High())+carry)
		carry = 0
	}

	for i := 2; i < len(a.value) && carry != 0; i++ {
		a.value[i], carry = bits.Add64(a.value[i], 0, carry)
	}

	if carry != 0 {
		a.value = append(a.value, carry)
	}
}

func (a Digits) AddDigit(b Digit) Digits {
	result := a.Copy()
	result.AddDigitInPlace(b)

	return result
}

func (a *Digits) AddDigitInPlace(b Digit) {
	if b == 0 {
		return
	}

	if len(a.value) == 0 {
		a.value = []uint64{uint64(b)}
		return
	}

	var carry uint64

	a.value[0], carry = bits.Add64(a.value[0], uint64(b), 0)

	for i := 1; i < len(a.value); i++ {
		a.value[i], carry = bits.Add64(a.value[i], 0, carry)
		if carry == 0 {
			return // Exit early for efficiency
		}
	}

	if carry != 0 {
		a.value = append(a.value, carry)
	}
}

func (a Digits) DoubleDigitAt(position uint) DoubleDigit {
	if position >= uint(len(a.value)) {
		return DoubleDigitOf(0, 0)
	}

	if position+1 >= uint(len(a.value)) {
		return DoubleDigitOf(0, Digit(a.value[position]))
	}

	return DoubleDigitOf(Digit(a.value[position+1]), Digit(a.value[position]))
}

func (a Digits) ChunkInclusive(start uint, end uint) Digits {
	if start >= uint(len(a.value)) {
		return Zero().AsDigits()
	}

	if end+1 >= uint(len(a.value)) {
		return Digits{false, a.value[start:]}
	}

	return Digits{false, a.value[start : end+1]}
}

func (a *Digits) Replace(start, end uint, b Digits) {
	if start > end {
		a.Replace(end, start, b)
	}

	if end >= uint(len(a.value)) {
		newLen := end + 1
		newValue := make([]uint64, newLen)
		copy(newValue, a.value)
		a.value = newValue
	}

	for i := start; i <= end; i++ {
		if i-start >= uint(len(b.value)) {
			a.value[i] = 0
		} else {
			a.value[i] = b.value[i-start]
		}
	}
}

func (a Digits) IsGreaterThan(other Digits) bool {
	return a.Compare(other) > 0
}

func (a Digits) Negate() Digits {
	return Digits{!a.sign, a.value}
}

func MakeDigits(size uint) Digits {
	return Digits{
		value: make([]uint64, size),
	}
}

func Empty() Digits {
	return Digits{false, []uint64{}}
}

func (a Digits) Hexadecimal() string {
	result := make([]string, len(a.value))

	for i, v := range a.value {
		result[i] = fmt.Sprintf("0x%X", v)
	}

	return strings.Join(result, "")
}

func (a Digits) DecrementInPlace() {
	if len(a.value) == 0 {
		return
	}

	a.SubtractInPlace(Digits{false, []uint64{1}})
}

func (a Digits) Decrement() (Digits, bool) {
	result := a.Copy()
	borrowed := result.SubtractUnderflowInPlace(Digits{false, []uint64{1}})

	return result, borrowed
}
