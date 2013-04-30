package main

import "fmt"

// modifier is coefficient of 10^(8*modifier)
type RNSNum struct {
	buffer []uint32
	modifier int
	isNeg bool
}

// size() int
// neededCapacityFor(int) int
// growBufferToValOf(int)
// growBufferToSizeOf(uint)
// Copy() *RNSNum
// AssignTo(int)
// Plus(*RNSNum) *RNSNum
// Equals(int) bool
// String() string
// ToMRSNum() (MRSNum, int)

func (n *RNSNum) size() int {
	return len(n.buffer)
}

func neededCapacityFor(val int) (r int) {
	r = 1
	for maxVal := int(primes[r - 1]); maxVal < val; r++ {
		maxVal *= int(primes[r-1])
	}
	return
}

func (n *RNSNum) growBufferToValOf(val int) {
	if val < 0 {
		val = val * -1
	}
	growPrimesToValOf(val)
	n.growBufferToSizeOf(neededCapacityFor(val))
}

func (n *RNSNum) growBufferToSizeOf(size int) {
	if n.size() < size {
		t, _ := n.ToMRSNum()
		growPrimesToSizeOf(size)
		for n.size() < size {
			n.buffer = append(n.buffer, t.Mod(primes[n.size()]))
		}
	}
}

func (n *RNSNum) Copy() (ret *RNSNum) {
	ret = &RNSNum{buffer: make([]uint32, len(n.buffer))}
	copy(ret.buffer, n.buffer)
	return
}

func (n *RNSNum) AssignTo(val int) {
	if val < 0 {
		val = val * -1
		n.isNeg = true
	}
	n.growBufferToValOf(val)
	for i := 0; i < n.size(); i++ {
		n.buffer[i] = uint32(val) % primes[i]
	}
}

func (n *RNSNum) Plus(val *RNSNum) (ret *RNSNum) {
	// Plus() assumes that MaxPrime() < max value of int
	maxSize := max(val.size(), n.size()) + 1
	n.growBufferToSizeOf(maxSize)
	val.growBufferToSizeOf(maxSize)
	ret = n.Copy()
	for i := 0; i < ret.size(); i++ {
		ret.buffer[i] = (ret.buffer[i] + val.buffer[i]) % primes[i]
	}
	return
}

func (n *RNSNum) Equals(val int) bool {
	if val < 0 {
		if n.isNeg == false {
			return false
		}
		val = val * -1
	}
	for i := 0; i < n.size(); i++ {
		if uint32(val) % primes[i] != n.buffer[i] {
			return false
		}
	}
	return true
}

func (n *RNSNum) String() (ret string) {
	i, _ := n.ToMRSNum()
	return fmt.Sprint(&i)
}

func (n *RNSNum) ToMRSNum() (ret MRSNum, numGuesses int) {
	numGuesses = 0
	if len(n.buffer) > 0 {
		ret.Add(1, n.buffer[0])
		fmt.Println("len(n.buffer) : ", len(n.buffer), " len(primes) : ", len(primes))
		for i := 1; i < len(n.buffer); i++ {
			ret.Add(primes[i - 1], 0)
			for ret.Mod(primes[i]) != n.buffer[i] {
				ret.Set(i, ret.values[i] + 1)
				numGuesses++
			}
		}
	}
	return ret, numGuesses
}
