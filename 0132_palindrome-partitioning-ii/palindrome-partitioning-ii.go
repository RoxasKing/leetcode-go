package main

/*
  Given a string s, partition s such that every substring of the partition is a palindrome.
  Return the minimum cuts needed for a palindrome partitioning of s.

  Example 1:
    Input: s = "aab"
    Output: 1
    Explanation: The palindrome partitioning ["aa","b"] could be produced using 1 cut.

  Example 2:
    Input: s = "a"
    Output: 0

  Example 3:
    Input: s = "ab"
    Output: 1

  Constraints:
    1. 1 <= s.length <= 2000
    2. s consists of lower-case English letters only.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/palindrome-partitioning-ii
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Dynamic Programming
func minCut(s string) int {
	n := len(s)
	isPalindrome := make([][]bool, n)
	for r := range isPalindrome {
		isPalindrome[r] = make([]bool, n)
		for l := 0; l <= r; l++ {
			if s[l] == s[r] && (r-l < 2 || isPalindrome[l+1][r-1]) {
				isPalindrome[l][r] = true
			}
		}
	}

	dp := make([]int, n)
	for i := range dp {
		dp[i] = i
	}

	for r := 1; r < n; r++ {
		if isPalindrome[0][r] {
			dp[r] = 0
			continue
		}
		for l := 1; l <= r; l++ {
			if isPalindrome[l][r] {
				dp[r] = Min(dp[r], dp[l-1]+1)
			}
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
