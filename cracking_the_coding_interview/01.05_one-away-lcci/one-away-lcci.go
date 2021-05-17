package main

/*
  There are three types of edits that can be performed on strings: insert a character, remove a character, or replace a character. Given two strings, write a function to check if they are one edit (or zero edits) away.

  Example 1:
    Input:
      first = "pale"
      second = "ple"
    Output: True

  Example 2:
    Input:
      first = "pales"
      second = "pal"
    Output: False

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/one-away-lcci
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func oneEditAway(first string, second string) bool {
	n1, n2 := len(first), len(second)
	if n1-n2 > 1 || n1-n2 < -1 {
		return false
	}
	for i := 0; i < n1 && i < n2; i++ {
		if first[i] != second[i] {
			return first[i+1:] == second[i+1:] || first[i+1:] == second[i:] || first[i:] == second[i+1:]
		}
	}
	return true
}
