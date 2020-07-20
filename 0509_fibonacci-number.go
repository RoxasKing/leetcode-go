package leetcode

/*
  斐波那契数，通常用 F(n) 表示，形成的序列称为斐波那契数列。该数列由 0 和 1 开始，后面的每一项数字都是前面两项数字的和。也就是：

  提示：
    0 ≤ N ≤ 30
*/

func fib(N int) int {
	dp0, dp1 := 0, 1
	for i := 1; i <= N; i++ {
		dp0, dp1 = dp1, dp0+dp1
	}
	return dp0
}
