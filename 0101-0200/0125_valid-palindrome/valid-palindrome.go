package main

/*
  Given a string s, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.

  Example 1:
    Input: s = "A man, a plan, a canal: Panama"
    Output: true
    Explanation: "amanaplanacanalpanama" is a palindrome.

  Example 2:
    Input: s = "race a car"
    Output: false
    Explanation: "raceacar" is not a palindrome.

  Constraints:
    1. 1 <= s.length <= 2 * 10^5
    2. s consists only of printable ASCII characters.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/valid-palindrome
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Two Pointers
func isPalindrome(s string) bool {
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		for l < r && !isValid(s[l]) {
			l++
		}
		for l < r && !isValid(s[r]) {
			r--
		}
		if lowerCase(s[l]) != lowerCase(s[r]) {
			return false
		}
	}
	return true
}

func isValid(ch byte) bool {
	return 'A' <= ch && ch <= 'Z' || 'a' <= ch && ch <= 'z' || '0' <= ch && ch <= '9'
}

func lowerCase(ch byte) byte {
	if 'A' <= ch && ch <= 'Z' {
		return ch - 'A' + 'a'
	}
	return ch
}
