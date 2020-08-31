package main

/*
  实现 pow(x, n) ，即计算 x 的 n 次幂函数。
  说明:
    -100.0 < x < 100.0
    n 是 32 位有符号整数，其数值范围是 [−231, 231 − 1] 。
*/

// Binary Exponentiation + Iterate
func myPow(x float64, n int) float64 {
	if n < 0 {
		n = -n
		x = 1 / x
	}
	return quickMul(x, n)
}

func quickMul(x float64, N int) float64 {
	out := 1.0
	base := x
	for N > 0 {
		if N%2 == 1 {
			out *= base
		}
		base *= base
		N /= 2
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
