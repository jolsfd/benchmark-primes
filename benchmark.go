package main

import (
	"fmt"
	"time"
)

func isPrime(n int) bool {
	for i := 2; i < n/2+1; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	start := time.Now()

	var numPrimes int

	for i := 2; i < 250001; i++ {
		if isPrime(i) {
			numPrimes++
		}
	}

	fmt.Printf("Computed %v primes in %v\n", numPrimes, time.Since(start))
}
