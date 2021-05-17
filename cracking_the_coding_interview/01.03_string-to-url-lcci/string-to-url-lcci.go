package main

import "unsafe"

/*
  Write a method to replace all spaces in a string with '%20'. You may assume that the string has sufficient space at the end to hold the additional characters,and that you are given the "true" length of the string. (Note: If implementing in Java,please use a character array so that you can perform this operation in place.)

  Example 1:
    Input: "Mr John Smith ", 13
    Output: "Mr%20John%20Smith"

  Example 2:
    Input: "               ", 5
    Output: "%20%20%20%20%20"

  Note:
    0 <= S.length <= 500000

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/string-to-url-lcci
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func replaceSpaces(S string, length int) string {
	chs := make([]byte, 0, len(S))
	for i := 0; i < length; i++ {
		if S[i] == ' ' {
			chs = append(chs, '%', '2', '0')
		} else {
			chs = append(chs, S[i])
		}
	}
	return *(*string)(unsafe.Pointer(&chs))
}
