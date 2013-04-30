package main

import "fmt"
import "math/big"

// x = abs(x) * 10^(8*modifier) * isNeg
// x = sum(x_0 ... x_{len(x)})
// x_i = buffer[i] * primes[i]
type MRSNum struct {
	radixes []uint32
	values []uint32
	modifier int
	isNeg bool
}

// AddPlace(radix uint32)
// Set(place int, value uint32)
// Add(radix, value uint32)

func (n *MRSNum) AddPlace(radix uint32) {
	n.radixes = append(n.radixes, radix)
	n.values = append(n.values, 0)
}

func (n *MRSNum) Set(place int, value uint32) {
	if place >= len(n.values) {
		panic(fmt.Sprintln("Out of bounds access in MRSNum::set"))
	}
	n.values[place] = value
}

func (n *MRSNum) Add(radix, value uint32) {
	n.AddPlace(radix)
	n.Set(len(n.values) - 1, value)
}

func (n *MRSNum) Mod(x uint32) (r uint32) {
	// p_n = prod_{j=0}^{n}(buffer_j.radix) mod x
	// r = sum_{i=0}^{sz_buffer}(buffer_i.value * p_i) mod x
	for i := 0; i < len(n.values); i++ {
		p_i := ProdMod(0, i, x, n.radixes)
		r_i := uint64(n.values[i]) * uint64(p_i) % uint64(x)
		r = uint32((r_i + uint64(r)) % uint64(x))
	}
	return
}

func (n *MRSNum) String() string {
	return n.BigInt().String()
}

func (n *MRSNum) BigInt() (ret *big.Int) {
	ret = big.NewInt(0)
	curRadix := big.NewInt(1)
	for i := 0; i < len(n.values); i++ {
		curRadix.Mul(curRadix, big.NewInt(int64(n.radixes[i])))
		curVal := big.NewInt(int64(n.values[i]))
		curVal.Mul(curVal, curRadix)
		ret.Add(ret, curVal)
	}
	return
}
