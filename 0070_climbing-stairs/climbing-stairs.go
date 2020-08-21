package main

/*
  假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
  每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
  注意：给定 n 是一个正整数。
*/

func climbStairs(n int) int {
	pre, cur := 1, 1
	for i := 2; i <= n; i++ {
		pre, cur = cur, pre+cur
	}
	return cur
}
