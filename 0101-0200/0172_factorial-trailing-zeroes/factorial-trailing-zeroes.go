package main

/*
  给定一个整数 n，返回 n! 结果尾数中零的数量。

  说明: 你算法的时间复杂度应为 O(log n) 。
*/

func trailingZeroes(n int) int {
	var out int
	for n != 0 {
		n /= 5
		out += n
	}
	return out
}
