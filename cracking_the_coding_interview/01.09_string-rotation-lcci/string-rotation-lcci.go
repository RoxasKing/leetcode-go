package main

import "strings"

/*
  Given two strings, s1 and s2, write code to check if s2 is a rotation of s1 (e.g.,"waterbottle" is a rotation of"erbottlewat"). Can you use only one call to the method that checks if one word is a substring of another?

  Example 1:
    Input: s1 = "waterbottle", s2 = "erbottlewat"
    Output: True

  Example 2:
    Input: s1 = "aa", s2 = "aba"
    Output: False

  Note:
    0 <= s1.length, s2.length <= 100000

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/string-rotation-lcci
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func isFlipedString(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	return strings.Contains(s2+s2, s1)
}
