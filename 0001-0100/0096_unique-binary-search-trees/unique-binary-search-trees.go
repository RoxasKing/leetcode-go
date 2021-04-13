package main

/*
  Given an integer n, return the number of structurally unique BST's (binary search trees) which has exactly n nodes of unique values from 1 to n.

  Example 1:
    Input: n = 3
    Output: 5

  Example 2:
    Input: n = 1
    Output: 1

  Constraints:
    1 <= n <= 19

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/unique-binary-search-trees
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Dynamic Programming
func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ { // number of nodes
		for j := 1; j <= i; j++ { // root node's index
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}
