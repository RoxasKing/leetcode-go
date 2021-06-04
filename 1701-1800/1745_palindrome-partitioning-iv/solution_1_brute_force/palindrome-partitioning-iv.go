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

// Brute Force

func checkPartitioning(s string) bool {
	n := len(s)
	for i := 0; i < n-2; i++ {
		if !isPalindrome(s[:i+1]) {
			continue
		}
		for j := i + 1; j < n-1; j++ {
			if !isPalindrome(s[i+1 : j+1]) {
				continue
			}
			if isPalindrome(s[j+1:]) {
				return true
			}
		}
	}
	return false
}

func isPalindrome(s string) bool {
	n := len(s)
	for l, r := 0, n-1; l < r; l, r = l+1, r-1 {
		if s[l] != s[r] {
			return false
		}
	}
	return true
}
