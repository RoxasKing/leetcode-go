package main

/*
  Given a non-empty string s, you may delete at most one character. Judge whether you can make it a palindrome.

  Example 1:
    Input: "aba"
    Output: True

  Example 2:
    Input: "abca"
    Output: True
    Explanation: You could delete the character 'c'.

  Note:
    The string will only contain lowercase characters a-z. The maximum length of the string is 50000.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/valid-palindrome-ii
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Two Pointers
func validPalindrome(s string) bool {
	n := len(s)
	l, r := 0, n-1
	for l < r {
		if s[l] != s[r] {
			return isPalindrome(s, l+1, r) || isPalindrome(s, l, r-1)
		}
		l++
		r--
	}
	return true
}

func isPalindrome(s string, l, r int) bool {
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
