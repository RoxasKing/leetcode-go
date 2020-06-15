package codinginterviews

/*
  求 1+2+...+n ，要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）。

  限制：
    1 <= n <= 10000
*/

// Bit operation
func sumNums(n int) int {
	ans, A, B := 0, n, n+1
	addGreatZero := func() bool {
		ans += A
		return ans > 0
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
	return ans >> 1
}
