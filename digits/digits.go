package digits

import (
	"fmt"
	"github.com/borisskert/go-biginteger/uintArray"
	"github.com/borisskert/go-biginteger/uintUtils"
	"math/bits"
	"strconv"
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
	if a.sign != b.sign {
		return a.Subtract(b.Negate())
	}

	carry := uint64(0)
	sizeA := len(a.value)
	sizeB := len(b.value)
	size := max(sizeA, sizeB)

	result := make([]uint64, size+1) // Extra space for carry

	for i := 0; i < size; i++ {
		ai := uint64(0)
		if i < sizeA {
			ai = a.value[i]
		}

		bi := uint64(0)
		if i < sizeB {
			bi = b.value[i]
		}

		// Perform addition with carry
		result[i], carry = bits.Add64(ai, bi, carry)
	}

	// If there's a carry left, store it
	result[size] = carry

	// Return result (Trim() ensures we don't keep unnecessary leading zeros)
	return Digits{sign: a.sign, value: result}.Trim()
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

func (a Digits) MultiplyDoNotUse(b Digits) Digits { // TODO remove
	//if len(a.value) == 0 || len(b.value) == 0 {
	//	return Digits{false, []uint64{0}}
	//}
	//
	//result := uintArray.
	//	MultiplyUint64Array(a.value, b.value)
	//
	//digits := Digits{a.sign != b.sign, result}
	//
	//return digits.Trim()
	panic("Do not use this function (yet)")
}

func (a Digits) SubtractAndBorrow(b Digits) (Digits, bool) {
	if a.sign != b.sign {
		return a.Add(b.Negate()), false
	}

	if a.IsGreaterThanOrEqual(b) {
		return a.Subtract(b), false
	}

	return b.Subtract(a).Negate(), true
}

func (a Digits) Subtract(b Digits) Digits {
	if a.sign != b.sign {
		return a.Add(b.Negate())
	}

	if a.Abs().IsLessThan(b.Abs()) {
		diff := b.Difference(a)
		return Digits{!a.sign, diff.value}
	}

	diff := a.Difference(b)

	if diff.IsZero() {
		return Zero().AsDigits()
	}

	return Digits{a.sign, diff.value}.Trim()
}

func (a Digits) SubtractExact(b Digits) Digits {
	if a.sign != b.sign {
		return a.Add(b.Negate())
	}

	if a.IsGreaterThanOrEqual(b) {
		return a.Subtract(b)
	}

	panic("Subtraction is not exact")
}

func (a *Digits) SubtractInPlace(b Digits) bool {
	if a.sign != b.sign {
		return a.AddInPlace(b.Negate())
	}

	if a.Abs().IsLessThan(b.Abs()) {
		diff := b.Difference(*a)

		a.value = diff.value
		a.sign = !a.sign

		return true
	}

	diff := a.Difference(b)
	a.value = diff.value

	return false
}

func (a Digits) Difference(b Digits) Digits {
	if a.sign != b.sign {
		// If signs are different, Difference(a, b) = Abs(a + (-b))
		return a.Add(b.Negate()).Abs()
	}

	// Ensure a ≥ b, otherwise swap (since Difference should always be positive)
	if a.Abs().IsLessThan(b.Abs()) {
		return b.Difference(a).Abs()
	}

	carry := uint64(0)
	size := len(a.value) // Since a ≥ b, we use len(a.value)
	result := make([]uint64, size)

	for i := 0; i < size; i++ {
		ai := a.value[i] // Guaranteed valid since size = len(a.value)
		bi := uint64(0)
		if i < len(b.value) {
			bi = b.value[i]
		}

		// Perform subtraction with borrow
		diff, borrow := bits.Sub64(ai, bi, carry)
		result[i] = diff
		carry = borrow
	}

	// Always return a positive number
	digits := Digits{sign: false, value: result}

	return digits.Trim()
}

func (a Digits) SubtractUnderflow(b Digits) (Digits, bool) { // TODO rename to NoBorrow
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
		a.sign = false
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
	if b == 0 {
		return Zero().AsDigits()
	}

	if b == 1 {
		return a.Copy()
	}

	if b == 2 {
		return a.LeftShiftBits(1)
	}

	result := make([]uint64, len(a.value)+2)

	carry := uint64(0)

	for i := range uint(len(a.value)) {
		pi := a.DigitAt(i).Multiply(b)

		lo, carryLo := bits.Add64(result[i], uint64(pi.Low()), carry)
		result[i] = lo

		mid, carryMid := bits.Add64(result[i+1], uint64(pi.High()), carryLo)
		result[i+1] = mid

		hi, carryHi := bits.Add64(result[i+2], carryMid, 0)
		result[i+2] = hi

		carry = carryHi
	}

	return Digits{a.sign, result}.Trim()
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

	carry := uint64(0)

	for i := range uint(len(a.value)) {
		pi := a.DigitAt(i).Multiply(b.Low())
		qi := a.DigitAt(i).Multiply(b.High())

		lo, carryLo := bits.Add64(result[i], uint64(pi.Low()), carry)
		result[i] = lo

		mid, carryMid := uintUtils.AddFour64(result[i+1], uint64(pi.High()), uint64(qi.Low()), carryLo)
		result[i+1] = mid

		hi, carryHi := bits.Add64(result[i+2], uint64(qi.High()), carryMid)
		result[i+2] = hi

		carry = carryHi
	}

	return Digits{a.sign, result}.Trim()
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

func (a Digits) IsGreaterThanOrEqual(b Digits) bool {
	return a.Compare(b) >= 0
}

func (a Digits) Negate() Digits {
	return Digits{!a.sign, a.value}
}

func (a Digits) Negative() Digits {
	return Digits{true, a.value}
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

func (a *Digits) DecrementInPlace() {
	if len(a.value) == 0 {
		a.value = []uint64{1}
		a.sign = true
	}

	a.SubtractInPlace(Digits{false, []uint64{1}})
}

func (a Digits) Decrement() (Digits, bool) {
	result := a.Copy()
	borrowed := result.SubtractUnderflowInPlace(Digits{false, []uint64{1}})

	return result.Trim(), borrowed
}

func (a Digits) Increment() Digits {
	return a.AddDigit(1)
}

func (a Digits) Split2(size uint) (Digits, Digits) {
	if size >= uint(len(a.value)) {
		values := make([]uint64, len(a.value))
		copy(values, a.value)

		return Zero().AsDigits(), Digits{false, values}
	}

	a1 := Digits{false, a.value[size:]}
	a0 := Digits{false, a.value[:size]}

	return a1.Trim(),
		a0.Trim()
}

func (a Digits) Split3(size uint) (Digits, Digits, Digits) {
	if size >= uint(len(a.value)) {
		values := make([]uint64, len(a.value))
		copy(values, a.value)

		return Zero().AsDigits(), Zero().AsDigits(), Digits{false, values}
	}

	a0 := Digits{false, a.value[:size]}.Trim()

	if 2*size >= uint(len(a.value)) {
		return Zero().AsDigits(), Digits{false, a.value[size:]}.Trim(), a0
	}

	a1 := Digits{false, a.value[size : 2*size]}.Trim()
	a2 := Digits{false, a.value[2*size:]}.Trim()

	return a2, a1, a0
}

func (a Digits) SubtractAbs(b Digits) Digits {
	if a.IsGreaterThanOrEqual(b) {
		diff := a.Subtract(b)
		return diff
	}

	diff := b.Subtract(a)
	return diff
}

func (a Digits) Sign(isNegative bool) Digits {
	return Digits{isNegative, a.value}
}

func (a Digits) DivideByDigit(b Digit) (Digits, Digit) {
	if b == 0 {
		panic("Division by zero")
	}

	if a.IsZero() {
		return Zero().AsDigits(), 0
	}

	if b == 1 {
		return a, 0
	}

	if b == 2 {
		return a.RightShiftBits(1), Digit(a.value[0] & 1)
	}

	quotient := Empty()
	remainder := Zero().AsDoubleDigit()

	for i := int64(a.Length()) - 1; i >= 0; i-- {
		di := a.DigitAt(uint(i))
		remainder, _ = remainder.LeftShift(64).AddDigit(di)

		qHat, rHat := remainder.DivideByDigit(b)

		quotient = quotient.LeftShiftDigits(1).AddDoubleDigit(qHat)
		remainder = rHat.AsDoubleDigit()
	}

	if a.IsNegative() {
		quotient = quotient.Negate()
	}

	return quotient.Trim(), remainder.Low()
}

func (a Digits) DivideByDigitNoRemainder(b Digit) Digits {
	if b == 2 {
		return a.RightShiftBits(1)
	}

	quotient := Empty()
	remainder := Zero().AsDoubleDigit()

	for i := int64(a.Length()) - 1; i >= 0; i-- {
		di := a.DigitAt(uint(i))
		remainder, _ = remainder.LeftShift(64).AddDigit(di)

		qHat, rHat := remainder.DivideByDigit(b)

		quotient = quotient.LeftShiftBits(64).AddDoubleDigit(qHat)
		remainder = rHat.AsDoubleDigit()
	}

	if a.IsNegative() {
		quotient = quotient.Negate()
	}

	return quotient.Trim()
}

func (a Digits) DivideByDigitExact(b Digit) Digits {
	quotient, remainder := a.DivideByDigit(b)

	if remainder != 0 {
		panic("Division is not exact")
	}

	return quotient
}

func (a Digits) Abs() Digits {
	return Digits{false, a.value}
}

func (a Digits) IsEven() bool {
	if len(a.value) == 0 {
		return true
	}

	return a.value[0]&1 == 0
}

func (a Digits) IsOdd() bool {
	return !a.IsEven()
}

func (a Digits) String() string {
	if a.IsNegative() {
		return "-" + a.Abs().stringAbs()
	}

	return a.stringAbs()
}

func (a Digits) stringAbs() string {
	var e19 = Digit(1000000000000000000)

	if a.IsLessThan(e19.AsDigits()) {
		return strconv.FormatUint(a.value[0], 10)
	}

	result := ""
	for a.IsGreaterThan(Zero().AsDigits()) {
		quotient, remainder := a.DivideByDigit(e19)
		a = quotient

		if a.IsGreaterThan(Zero().AsDigits()) {
			result = fmt.Sprintf("%018d", uint64(remainder)) + result
		} else {
			result = strconv.FormatUint(uint64(remainder), 10) + result
		}
	}

	return result
}

func (a Digits) IsEqualTo(result Digits) bool {
	if a.sign != result.sign {
		return false
	}

	if len(a.value) != len(result.value) {
		return false
	}

	for i := range a.value {
		if a.value[i] != result.value[i] {
			return false
		}
	}

	return true
}

func OfUint64Array(array []uint64) Digits {
	values := make([]uint64, len(array))
	copy(values, array)

	return Digits{false, values}
}
