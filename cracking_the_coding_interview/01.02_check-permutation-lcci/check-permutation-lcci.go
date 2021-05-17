package main

import "sort"

/*
  Given two strings,write a method to decide if one is a permutation of the other.

  Example 1:
    Input: s1 = "abc", s2 = "bca"
    Output: true

  Example 2:
    Input: s1 = "abc", s2 = "bad"
    Output: false

  Note:
    0 <= len(s1) <= 100
    0 <= len(s2) <= 100

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/check-permutation-lcci
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func CheckPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	chs1 := []rune(s1)
	chs2 := []rune(s2)
	sort.Slice(chs1, func(i, j int) bool { return chs1[i] < chs1[j] })
	sort.Slice(chs2, func(i, j int) bool { return chs2[i] < chs2[j] })
	for i := 0; i < len(chs1); i++ {
		if chs1[i] != chs2[i] {
			return false
		}
	}
	return true
}
