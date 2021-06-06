package main

/*
  You are given an array of binary strings strs and two integers m and n.

  Return the size of the largest subset of strs such that there are at most m 0's and n 1's in the subset.

  A set x is a subset of a set y if all elements of x are also elements of y.

  Example 1:
    Input: strs = ["10","0001","111001","1","0"], m = 5, n = 3
    Output: 4
    Explanation: The largest subset with at most 5 0's and 3 1's is {"10", "0001", "1", "0"}, so the answer is 4.
      Other valid but smaller subsets include {"0001", "1"} and {"10", "1", "0"}.
      {"111001"} is an invalid subset because it contains 4 1's, greater than the maximum of 3.

  Example 2:
    Input: strs = ["10","0","1"], m = 1, n = 1
    Output: 2
    Explanation: The largest subset is {"0", "1"}, so the answer is 2.

  Constraints:
    1. 1 <= strs.length <= 600
    2. 1 <= strs[i].length <= 100
    3. strs[i] consists only of digits '0' and '1'.
    4. 1 <= m, n <= 100

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/ones-and-zeroes
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Dynamic Programming

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for _, s := range strs {
		cnt := [2]int{}
		for k := range s {
			cnt[s[k]-'0']++
		}
		for i := m; i >= cnt[0]; i-- {
			for j := n; j >= cnt[1]; j-- {
				dp[i][j] = Max(dp[i][j], dp[i-cnt[0]][j-cnt[1]]+1)
			}
		}
	}
	return dp[m][n]
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
