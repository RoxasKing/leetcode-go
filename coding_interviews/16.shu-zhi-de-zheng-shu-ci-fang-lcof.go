package codinginterviews

/*
  实现函数double Power(double base, int exponent)，求base的exponent次方。不得使用库函数，同时不需要考虑大数问题。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/shu-zhi-de-zheng-shu-ci-fang-lcof
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func myPow(x float64, n int) float64 {
	if x == 0.0 && n < 0 {
		return 0.0
	}
	absN := uint32(n)
	if n < 0 {
		absN = uint32(-n)
	}
	out := pow(x, absN)
	if n < 0 {
		out = 1.0 / out
	}
	return out
}

func pow(base float64, exponent uint32) float64 {
	if exponent == 0 {
		return 1
	}
	if exponent == 1 {
		return base
	}
	res := pow(base, exponent>>1)
	res *= res
	if exponent&0x1 == 1 {
		res *= base
	}
	return res
}
