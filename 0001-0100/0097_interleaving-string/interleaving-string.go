package main

/*
  Given strings s1, s2, and s3, find whether s3 is formed by an interleaving of s1 and s2.

  An interleaving of two strings s and t is a configuration where they are divided into non-empty substrings such that:
    1. s = s1 + s2 + ... + sn
    2. t = t1 + t2 + ... + tm
    3. |n - m| <= 1
    4. The interleaving is s1 + t1 + s2 + t2 + s3 + t3 + ... or t1 + s1 + t2 + s2 + t3 + s3 + ...
  Note: a + b is the concatenation of strings a and b.

  Example 1:
    Input: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
    Output: true

  Example 2:
    Input: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
    Output: false

  Example 3:
    Input: s1 = "", s2 = "", s3 = ""
    Output: true

  Constraints:
    1. 0 <= s1.length, s2.length <= 100
    2. 0 <= s3.length <= 200
    3. s1, s2, and s3 consist of lowercase English letters.

  Follow up: Could you solve it using only O(s2.length) additional memory space?

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/interleaving-string
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Dynamic Programming
func isInterleave(s1 string, s2 string, s3 string) bool {
	m, n, k := len(s1), len(s2), len(s3)
	if m+n != k {
		return false
	}
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 0; i < m && dp[i][0]; i++ {
		dp[i+1][0] = s1[i] == s3[i]
	}
	for j := 0; j < n && dp[0][j]; j++ {
		dp[0][j+1] = s2[j] == s3[j]
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dp[i+1][j+1] = dp[i][j+1] && s1[i] == s3[i+j+1] || dp[i+1][j] && s2[j] == s3[i+j+1]
		}
	}
	return dp[m][n]
}

// Dynamic Programming (O(s2.length) additional memory space)
func isInterleave2(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	m, n := len(s1), len(s2)
	dp := make([]bool, n+1)
	dp[0] = true
	for j := 0; j < n && dp[j]; j++ {
		dp[j+1] = s2[j] == s3[j]
	}
	for i := 0; i < m; i++ {
		dp[0] = dp[0] && s1[i] == s3[i]
		for j := 0; j < n; j++ {
			dp[j+1] = dp[j+1] && s1[i] == s3[i+j+1] || dp[j] && s2[j] == s3[i+j+1]
		}
	}
	return dp[n]
}
