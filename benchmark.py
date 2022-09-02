#!/usr/bin/env python3
import time

def isPrime(n: int) -> int:
    for i in range(2,n//2+1):
        if n%i == 0:
            return 0

    return 1

if __name__ == "__main__":
    start = time.time()

    numPrimes = 0

    for i in range(2,250001):
        numPrimes+=isPrime(i)

    finish = time.time()
    print(f"Computed {numPrimes} primes in {finish - start}s")