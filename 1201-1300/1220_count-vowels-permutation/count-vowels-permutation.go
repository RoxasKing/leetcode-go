package main

// Difficulty:
// Hard

// Tags:
// Dynamic Programming

var mod int = 1e9 + 7

func countVowelPermutation(n int) int {
	f0, f1 := [5]int{}, [5]int{}
	for i := 0; i < 5; i++ {
		f0[i] = 1
	}
	for k := 1; k < n; k++ {
		f1[0] = (f0[1] + f0[2] + f0[4]) % mod
		f1[1] = (f0[0] + f0[2]) % mod
		f1[2] = (f0[1] + f0[3]) % mod
		f1[3] = f0[2] % mod
		f1[4] = (f0[2] + f0[3]) % mod
		f0, f1 = f1, f0
	}
	var out int
	for i := 0; i < 5; i++ {
		out += f0[i]
		out %= mod
	}
	return out
}
