package crackingthecodingintervew

/*
  求 1+2+...+n ，要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）。

  限制：
    1 <= n <= 10000
*/

// Recursive
func sumNums(n int) int {
	if n == 0 {
		return 0
	}
	return sumNums(n-1) + n
}

// Russian peasant multiplication
func sumNums2(n int) int {
	ans, A, B := 0, n, n+1
	addGreatZero := func() bool {
		ans += A
		return true
	}

	_ = B&1 > 0 && addGreatZero()
	A <<= 1
	B >>= 1

	_ = B&1 > 0 && addGreatZero()
	A <<= 1
	B >>= 1

	_ = B&1 > 0 && addGreatZero()
	A <<= 1
	B >>= 1

	_ = B&1 > 0 && addGreatZero()
	A <<= 1
	B >>= 1

	_ = B&1 > 0 && addGreatZero()
	A <<= 1
	B >>= 1

	_ = B&1 > 0 && addGreatZero()
	A <<= 1
	B >>= 1

	_ = B&1 > 0 && addGreatZero()
	A <<= 1
	B >>= 1

	_ = B&1 > 0 && addGreatZero()
	A <<= 1
	B >>= 1

	_ = B&1 > 0 && addGreatZero()
	A <<= 1
	B >>= 1

	_ = B&1 > 0 && addGreatZero()
	A <<= 1
	B >>= 1

	_ = B&1 > 0 && addGreatZero()
	A <<= 1
	B >>= 1

	_ = B&1 > 0 && addGreatZero()
	A <<= 1
	B >>= 1

	_ = B&1 > 0 && addGreatZero()
	A <<= 1
	B >>= 1

	_ = B&1 > 0 && addGreatZero()
	A <<= 1

	return ans >> 1
}
