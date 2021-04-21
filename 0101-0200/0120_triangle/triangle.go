package main

/*
  Given a triangle array, return the minimum path sum from top to bottom.

  For each step, you may move to an adjacent number of the row below. More formally, if you are on index i on the current row, you may move to either index i or index i + 1 on the next row.

  Example 1:
    Input: triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
    Output: 11
    Explanation: The triangle looks like:
       2
      3 4
     6 5 7
    4 1 8 3
    The minimum path sum from top to bottom is 2 + 3 + 5 + 1 = 11 (underlined above).

  Example 2:
    Input: triangle = [[-10]]
    Output: -10

  Constraints:
    1. 1 <= triangle.length <= 200
    2. triangle[0].length == 1
    3. triangle[i].length == triangle[i - 1].length + 1
    4. -10^4 <= triangle[i][j] <= 10^4

  Follow up: Could you do this using only O(n) extra space, where n is the total number of rows in the triangle?

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/triangle
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Dynamic Programming
func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	copy(dp, triangle[n-1])
	for i := n - 2; i >= 0; i-- {
		for j := range triangle[i] {
			dp[j] = Min(dp[j], dp[j+1]) + triangle[i][j]
		}
	}
	return dp[0]
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
