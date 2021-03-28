package main

/*
  Given a string s consisting only of characters 'a', 'b', and 'c'. You are asked to apply the following algorithm on the string any number of times:

  Pick a non-empty prefix from the string s where all the characters in the prefix are equal.
  Pick a non-empty suffix from the string s where all the characters in this suffix are equal.
  The prefix and the suffix should not intersect at any index.
  The characters from the prefix and suffix must be the same.
  Delete both the prefix and the suffix.
  Return the minimum length of s after performing the above operation any number of times (possibly zero times).

  Example 1:
    Input: s = "ca"
    Output: 2
    Explanation: You can't remove any characters, so the string stays as is.

  Example 2:
    Input: s = "cabaabac"
    Output: 0
    Explanation: An optimal sequence of operations is:
    - Take prefix = "c" and suffix = "c" and remove them, s = "abaaba".
    - Take prefix = "a" and suffix = "a" and remove them, s = "baab".
    - Take prefix = "b" and suffix = "b" and remove them, s = "aa".
    - Take prefix = "a" and suffix = "a" and remove them, s = "".

  Example 3:
    Input: s = "aabccabba"
    Output: 3
    Explanation: An optimal sequence of operations is:
    - Take prefix = "aa" and suffix = "a" and remove them, s = "bccabb".
    - Take prefix = "b" and suffix = "bb" and remove them, s = "cca".

  Constraints:
    1. 1 <= s.length <= 10^5
    2. s only consists of characters 'a', 'b', and 'c'.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/minimum-length-of-string-after-deleting-similar-ends
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func minimumLength(s string) int {
	n := len(s)
	l, r := 0, n-1
	for l < r {
		if s[l] != s[r] {
			break
		}
		ch := s[l]
		for l <= r && s[l] == ch {
			l++
		}
		for l <= r && s[r] == ch {
			r--
		}
	}
	if l > r {
		return 0
	}
	return r - l + 1
}
