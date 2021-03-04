package main

/*
  求 1+2+...+n ，要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）。

  示例 1：
    输入: n = 3
    输出: 6

  示例 2：
    输入: n = 9
    输出: 45

  限制：
    1 <= n <= 10000

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/qiu-12n-lcof
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Recursion
func sumNums(n int) int {
	out := 0
	var sum func(int) bool
	sum = func(i int) bool {
		out += i
		return i > 0 && sum(i-1)
	}
	sum(n)
	return out
}

// Bit operation
func sumNums2(n int) int {
	out, A, B := 0, n, n+1

	addGreatZero := func() bool {
		out += A
		return out > 0
	}

	_ = ((B & 1) > 0) && addGreatZero()
	A <<= 1
	B >>= 1

	_ = ((B & 1) > 0) && addGreatZero()
	A <<= 1
	B >>= 1

	_ = ((B & 1) > 0) && addGreatZero()
	A <<= 1
	B >>= 1

	_ = ((B & 1) > 0) && addGreatZero()
	A <<= 1
	B >>= 1

	_ = ((B & 1) > 0) && addGreatZero()
	A <<= 1
	B >>= 1

	_ = ((B & 1) > 0) && addGreatZero()
	A <<= 1
	B >>= 1

	_ = ((B & 1) > 0) && addGreatZero()
	A <<= 1
	B >>= 1

	_ = ((B & 1) > 0) && addGreatZero()
	A <<= 1
	B >>= 1

	_ = ((B & 1) > 0) && addGreatZero()
	A <<= 1
	B >>= 1

	_ = ((B & 1) > 0) && addGreatZero()
	A <<= 1
	B >>= 1

	_ = ((B & 1) > 0) && addGreatZero()
	A <<= 1
	B >>= 1

	_ = ((B & 1) > 0) && addGreatZero()
	A <<= 1
	B >>= 1

	_ = ((B & 1) > 0) && addGreatZero()
	A <<= 1
	B >>= 1

	_ = ((B & 1) > 0) && addGreatZero()
	A <<= 1

	return out >> 1
}
