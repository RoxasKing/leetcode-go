package main

// Difficulty:
// Medium

// Tags:
// Math

func superPow(a int, b []int) int {
	a %= MOD
	out := quickMul(a, b[0])
	for _, n := range b[1:] {
		out = quickMul(out, 10) * quickMul(a, n) % MOD
	}
	return out
}

const MOD = 1337

func quickMul(a, n int) int {
	out := 1
	for ; n > 0; n >>= 1 {
		if n&1 == 1 {
			out = (out * a) % MOD
		}
		a = (a * a) % MOD
	}
	return out
}
