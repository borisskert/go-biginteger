package biginteger

var zero = BigInteger{value: []uint64{0}}
var one = BigInteger{value: []uint64{1}}
var two = BigInteger{value: []uint64{2}}
var ten = BigInteger{value: []uint64{10}}
var e19 = BigInteger{value: []uint64{1000000000000000000}}

func Zero() BigInteger {
	return zero
}

func One() BigInteger {
	return one
}

func Two() BigInteger {
	return two
}

func Ten() BigInteger {
	return ten
}
