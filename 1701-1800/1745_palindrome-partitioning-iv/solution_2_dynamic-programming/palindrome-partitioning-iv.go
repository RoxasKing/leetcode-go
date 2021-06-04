package main

/*
  Given a string s, return true if it is possible to split the string s into three non-empty palindromic substrings. Otherwise, return false.​​​​​

  A string is said to be palindrome if it the same string when reversed.

  Example 1:
    Input: s = "abcbdd"
    Output: true
    Explanation: "abcbdd" = "a" + "bcb" + "dd", and all three substrings are palindromes.

  Example 2:
    Input: s = "bcbddxy"
    Output: false
    Explanation: s cannot be split into 3 palindromes.

  Constraints:
    1. 3 <= s.length <= 2000
    2. s​​​​​​ consists only of lowercase English letters.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/palindrome-partitioning-iv
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Dynamic Programming

func checkPartitioning(s string) bool {
	n := len(s)
	f := make([][]bool, n)
	for r := 0; r < n; r++ {
		f[r] = make([]bool, n)
		for l := 0; l <= r; l++ {
			if s[l] == s[r] && (r-l < 2 || f[l+1][r-1]) {
				f[l][r] = true
			}
		}
	}

	for i := 1; i < n-1; i++ {
		for j := i; j < n-1; j++ {
			if f[0][i-1] && f[i][j] && f[j+1][n-1] {
				return true
			}
		}
	}
	return false
}
