package main

/*
  You are given two strings, word1 and word2. You want to construct a string in the following manner:
    1. Choose some non-empty subsequence subsequence1 from word1.
    2. Choose some non-empty subsequence subsequence2 from word2.
    3. Concatenate the subsequences: subsequence1 + subsequence2, to make the string.
  Return the length of the longest palindrome that can be constructed in the described manner. If no palindromes can be constructed, return 0.

  A subsequence of a string s is a string that can be made by deleting some (possibly none) characters from s without changing the order of the remaining characters.

  A palindrome is a string that reads the same forward as well as backward.

  Example 1:
    Input: word1 = "cacb", word2 = "cbba"
    Output: 5
    Explanation: Choose "ab" from word1 and "cba" from word2 to make "abcba", which is a palindrome.

  Example 2:
    Input: word1 = "ab", word2 = "ab"
    Output: 3
    Explanation: Choose "ab" from word1 and "a" from word2 to make "aba", which is a palindrome.

  Example 3:
    Input: word1 = "aa", word2 = "bb"
    Output: 0
    Explanation: You cannot construct a palindrome from the described method, so return 0.

  Constraints:
    1. 1 <= word1.length, word2.length <= 1000
    2. word1 and word2 consist of lowercase English letters.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/maximize-palindrome-length-from-subsequences
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Dynamic Programming
func longestPalindrome(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	k := m + n
	dp := make([][]int, 2)
	for i := range dp {
		dp[i] = make([]int, k+1)
	}
	out := 0
	word := word1 + word2
	for i := 0; i < k; i++ {
		for j := 0; i+j < k; j++ {
			if word[i] == word[k-1-j] {
				if dp[0][j] != 0 {
					dp[1][j+1] = dp[0][j] + 1
					if i < k-1-j {
						dp[1][j+1]++
					}
				} else if i < m && j < n { // 关键,左端在word1之中，右端在word2之中
					dp[1][j+1] = 2
				}
				out = Max(out, dp[1][j+1])
			} else {
				dp[1][j+1] = Max(dp[1][j], dp[0][j+1])
			}
		}
		dp[0], dp[1] = dp[1], dp[0]
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
