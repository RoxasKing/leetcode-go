package leetcode

/*
  假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
  每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
  注意：给定 n 是一个正整数。
*/

func climbStairs(n int) int {
	if n == 0 {
		return 0
	}
	dp1, dp2 := 1, 1
	for i := 2; i <= n; i++ {
		dp1, dp2 = dp2, dp1+dp2
	}
	return dp2
}
