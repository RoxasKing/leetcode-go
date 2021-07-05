package main

// Tags:
// Binary Exponentiation + Iteration
func myPow(x float64, n int) float64 {
	if n < 0 {
		n = -n
		x = 1 / x
	}
	return quickMul(x, n)
}

func quickMul(x float64, n int) float64 {
	out := 1.0
	for n > 0 {
		if n&1 == 1 {
			out *= x
		}
		x *= x
		n >>= 1
	}
	return out
}

// Binary Exponentiation + Recursion
func myPow2(x float64, n int) float64 {
	if n < 0 {
		n = -n
		x = 1 / x
	}
	return quickMul2(x, n)
}

func quickMul2(x float64, N int) float64 {
	if N == 0 {
		return 1
	}
	half := quickMul2(x, N>>1)
	out := half * half
	if N&1 == 1 {
		out *= x
	}
	return out
}
