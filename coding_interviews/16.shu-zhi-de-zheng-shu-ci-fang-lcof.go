package codinginterviews

/*
  实现函数double Power(double base, int exponent)，求base的exponent次方。不得使用库函数，同时不需要考虑大数问题。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/shu-zhi-de-zheng-shu-ci-fang-lcof
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
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
		return 1.0
	}
	half := pow(base, exponent>>1)
	out := half * half
	if exponent&1 == 1 {
		out *= base
	}
	return out
}
