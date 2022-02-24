package main

// Difficulty:
// Hard

// Tags:
// Math
// Dynamic Programming

var mod int = 1e9 + 7
var primes = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}

func quickPow(x, n int) int {
	out := 1
	for ; n > 0; n >>= 1 {
		if n&1 == 1 {
			out = out * x % mod
		}
		x = x * x % mod
	}
	return out
}

func numberOfGoodSubsets(nums []int) int {
	freq := [31]int{}
	for _, num := range nums {
		freq[num]++
	}
	f := make([]int, 1<<len(primes))
	f[0] = quickPow(2, freq[1])
LOOP:
	for i := 2; i <= 30; i++ {
		if freq[i] == 0 {
			continue
		}
		subSet := 0
		for j, prime := range primes {
			if i%(prime*prime) == 0 {
				continue LOOP
			}
			if i%prime == 0 {
				subSet |= 1 << j
			}
		}
		for mask := 1 << len(primes); mask > 0; mask-- {
			if mask&subSet == subSet {
				f[mask] = (f[mask] + f[mask^subSet]*freq[i]) % mod
			}
		}
	}
	out := 0
	for _, cnt := range f[1:] {
		out = (out + cnt) % mod
	}
	return out
}
