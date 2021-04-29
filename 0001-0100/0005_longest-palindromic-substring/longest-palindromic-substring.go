package main

/*
  Given a string s, return the longest palindromic substring in s.

  Example 1:
    Input: s = "babad"
    Output: "bab"
    Note: "aba" is also a valid answer.

  Example 2:
    Input: s = "cbbd"
    Output: "bb"

  Example 3:
    Input: s = "a"
    Output: "a"

  Example 4:
    Input: s = "ac"
    Output: "a"

  Constraints:
    1. 1 <= s.length <= 1000
    2. s consist of only digits and English letters (lower-case and/or upper-case),

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/longest-palindromic-substring
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func longestPalindrome(s string) string {
	out := ""
	for i := range s {
		if res := getlongetPalindrome(s, i, i); len(res) > len(out) {
			out = res
		}
		if i+1 == len(s) || s[i+1] != s[i] {
			continue
		}
		if res := getlongetPalindrome(s, i, i+1); len(res) > len(out) {
			out = res
		}
	}
	return out
}

func getlongetPalindrome(s string, l, r int) string {
	for l-1 >= 0 && r+1 < len(s) && s[l-1] == s[r+1] {
		l--
		r++
	}
	return s[l : r+1]
}
