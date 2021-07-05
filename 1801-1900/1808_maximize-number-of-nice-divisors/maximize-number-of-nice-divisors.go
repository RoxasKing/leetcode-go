package main

// Tags:
// Greedy Algorithm
func maxNiceDivisors(primeFactors int) int {
	if primeFactors == 1 {
		return 1
	}
	timesOf3 := primeFactors / 3
	if primeFactors-timesOf3*3 == 1 {
		timesOf3--
	}
	timesOf2 := (primeFactors - timesOf3*3) / 2
	return QuickMul(3, timesOf3) * QuickMul(2, timesOf2) % (1e9 + 7)
}

func QuickMul(x, n int) int {
	out := 1
	for n > 0 {
		if n&1 == 1 {
			out = (out * x) % (1e9 + 7)
		}
		x = (x * x) % (1e9 + 7)
		n >>= 1
	}
	return out
}
