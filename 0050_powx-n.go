package leetcode

/*
  实现 pow(x, n) ，即计算 x 的 n 次幂函数。
*/

func myPow(x float64, n int) float64 {
	if n < 0 {
		n = -n
		x = 1 / x
	}
	return pow(x, n)
}

func pow(base float64, exponent int) float64 {
	if exponent == 0 {
		return 1
	}
	half := pow(base, exponent>>1)
	out := half * half
	if exponent&1 == 1 {
		out *= base
	}
	return out
}
