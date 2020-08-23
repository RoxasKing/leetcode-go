package main

/*
  给定范围 [m, n]，其中 0 <= m <= n <= 2147483647，返回此范围内所有数字的按位与（包含 m, n 两端点）。
*/

func rangeBitwiseAnd(m int, n int) int {
	var shift int
	for m < n {
		m, n = m>>1, n>>1
		shift++
	}
	return m << shift
}

// Brian Kernighan
func rangeBitwiseAnd2(m int, n int) int {
	for m < n {
		n &= (n - 1)
	}
	return n
}
