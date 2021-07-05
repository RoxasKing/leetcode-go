package main

// Tags:
// Tags:
// Math

func countGoodNumbers(n int64) int {
	if n&1 == 1 {
		return int(quickMul(5, n/2+1) * quickMul(4, n/2) % mod)
	}
	return int(quickMul(5, n/2) * quickMul(4, n/2) % mod)
}

func quickMul(x, n int64) int64 {
	out := int64(1)
	for ; n > 0; n >>= 1 {
		if n&1 == 1 {
			out = out * x % mod
		}
		x = x * x % mod
	}
	return out
}

var mod = int64(1e9 + 7)
