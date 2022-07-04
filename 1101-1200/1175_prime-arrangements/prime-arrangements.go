package main

// Difficulty:
// Easy

// Tags:
// Math

func numPrimeArrangements(n int) int {
	const mod int = 1e9 + 7
	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}
	c := 0
	for i := 2; i <= n; i++ {
		if isPrime[i] {
			c++
			for j := i << 1; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}
	o := 1
	for i := 2; i <= c; i++ {
		o = o * i % mod
	}
	for i := 2; i <= n-c; i++ {
		o = o * i % mod
	}
	return o
}
