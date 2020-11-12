package main

/*
  编写一个程序，找出第 n 个丑数。
  丑数就是质因数只包含 2, 3, 5 的正整数。

  示例:
    输入: n = 10
    输出: 12
    解释: 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 是前 10 个丑数。

  说明:
    1 是丑数。
    n 不超过1690。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/ugly-number-ii
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Dynamic Programming
func nthUglyNumber(n int) int {
	dp := make([]int, n)
	dp[0] = 1
	mul2, mul3, mul5 := 0, 0, 0
	for i := 1; i < n; i++ {
		dp[i] = Min(Min(dp[mul2]*2, dp[mul3]*3), dp[mul5]*5)
		if dp[mul2]*2 <= dp[i] {
			mul2++
		}
		if dp[mul3]*3 <= dp[i] {
			mul3++
		}
		if dp[mul5]*5 <= dp[i] {
			mul5++
		}
	}
	return dp[n-1]
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
