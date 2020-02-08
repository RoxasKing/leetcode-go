package leetcode

/*
  实现 pow(x, n) ，即计算 x 的 n 次幂函数。
*/

func myPow(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}
	var pow func(float64, int) float64
	pow = func(x float64, n int) float64 {
		if n == 0 {
			return 1.0
		}
		half := pow(x, n/2)
		if n%2 == 0 {
			return half * half
		}
		return half * half * x
	}
	return pow(x, n)
}
