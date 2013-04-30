package main

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
func nextPrime() uint32 {
	if len(primes) == 0 {
		return 3
	}
	for i := primes[len(primes) - 1] + 2; true; i += 2 {
		if isPrime(i) {
			return i
		}
	}
	panic("No next prime.")
}

func growPrimesToValOf(min int) {
	for primeMax() < uint32(min) {
		primes = append(primes, nextPrime())
	}
}

func growPrimesToSizeOf(size int) {
	for len(primes) < size {
		primes = append(primes, nextPrime())
	}
}
