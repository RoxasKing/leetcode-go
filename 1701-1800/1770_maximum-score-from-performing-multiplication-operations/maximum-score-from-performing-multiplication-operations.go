package main

/*
  You are given two integer arrays nums and multipliers of size n and m respectively, where n >= m. The arrays are 1-indexed.

  You begin with a score of 0. You want to perform exactly m operations. On the ith operation (1-indexed), you will:
    1. Choose one integer x from either the start or the end of the array nums.
    2. Add multipliers[i] * x to your score.
    3. Remove x from the array nums.
  Return the maximum score after performing m operations.

  Example 1:
    Input: nums = [1,2,3], multipliers = [3,2,1]
    Output: 14
    Explanation: An optimal solution is as follows:
    - Choose from the end, [1,2,3], adding 3 * 3 = 9 to the score.
    - Choose from the end, [1,2], adding 2 * 2 = 4 to the score.
    - Choose from the end, [1], adding 1 * 1 = 1 to the score.
    The total score is 9 + 4 + 1 = 14.

  Example 2:
    Input: nums = [-5,-3,-3,-2,7,1], multipliers = [-10,-5,3,4,6]
    Output: 102
    Explanation: An optimal solution is as follows:
    - Choose from the start, [-5,-3,-3,-2,7,1], adding -5 * -10 = 50 to the score.
    - Choose from the start, [-3,-3,-2,7,1], adding -3 * -5 = 15 to the score.
    - Choose from the start, [-3,-2,7,1], adding -3 * 3 = -9 to the score.
    - Choose from the end, [-2,7,1], adding 1 * 4 = 4 to the score.
    - Choose from the end, [-2,7], adding 7 * 6 = 42 to the score.
    The total score is 50 + 15 - 9 + 4 + 42 = 102.

  Constraints:
    1. n == nums.length
    2. m == multipliers.length
    3. 1 <= m <= 10^3
    4. m <= n <= 10^5
    5. -1000 <= nums[i], multipliers[i] <= 1000

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/maximum-score-from-performing-multiplication-operations
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Dynamic Programming
func maximumScore(nums []int, multipliers []int) int {
	n, m := len(nums), len(multipliers)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	for i := 1; i <= m; i++ {
		dp[i][0] = dp[i-1][0] + nums[i-1]*multipliers[i-1]
		dp[0][i] = dp[0][i-1] + nums[n-i]*multipliers[i-1]
	}

	for i := 1; i <= m; i++ {
		for j := 1; i+j <= m; j++ {
			k := i + j - 1
			dp[i][j] = Max(dp[i][j-1]+nums[n-j]*multipliers[k], dp[i-1][j]+nums[i-1]*multipliers[k])
		}
	}

	out := -1 << 31
	for i := 0; i <= m; i++ {
		out = Max(out, dp[i][m-i])
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
