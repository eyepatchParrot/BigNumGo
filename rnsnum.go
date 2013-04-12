package main

import "fmt"

var primes []uint32

func primeMax() (max uint32) {
	max = 1
	for i := 0; i < len(primes); i++ {
		max = max * primes[i]
	}
	return
}

func isPrime(n uint32) bool {
	for i := uint32(2); i * i <= n; i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}

// doesn't know what to do if there's no next prime
func nextPrime(min uint32) uint32 {
	for i := min + 2; true; i++ {
		if isPrime(i) {
			return i
		}
	}
	panic("No next prime.")
}

func growPrimesTo(min uint32) {
	for primeMax() < min {
		var newPrime uint32
		if len(primes) == 0 {
			newPrime = 3
		} else {
			newPrime = nextPrime(primes[len(primes) - 1])
		}
		primes = append(primes, newPrime)
	}
}

type RNSNum struct {
	buffer []uint32
	modifier int
	isNeg bool
}

// set(int)
// Copy() *RNSNum
// Assign(int)
// Add(*RNSNum) *RNSNum

func (n *RNSNum) set(val int) {
	if val < 0 {
		val = val * -1
		n.isNeg = true
	}
	growPrimesTo(uint32(val))
	if len(primes) > len(n.buffer) {
		n.buffer = make([]uint32, len(primes))
	}
	for i := 0; i < len(n.buffer); i++ {
		n.buffer[i] = uint32(val) % primes[i]
	}
}

func (n *RNSNum) Copy() (ret *RNSNum) {
	ret = &RNSNum{buffer: make([]uint32, len(n.buffer))}
	copy(ret.buffer, n.buffer)
	return
}

func (n *RNSNum) Assign(val int) {
	n.set(val)
}

func (n *RNSNum) Add(val *RNSNum) (ret *RNSNum) {
	if len(n.buffer) != len(val.buffer) {
		panic("Nums not the same size")
	}
	ret = n.Copy()
	for i := 0; i < len(ret.buffer); i++ {
		ret.buffer[i] = (ret.buffer[i] + val.buffer[i]) % primes[i]
	}
	return
}
