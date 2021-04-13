package main

import "strconv"

/*
  Implement a method to perform basic string compression using the counts of repeated characters. For example, the string aabcccccaaa would become a2blc5a3. If the "compressed" string would not become smaller than the original string, your method should return the original string. You can assume the string has only uppercase and lowercase letters (a - z).

  Example 1:
    Input: "aabcccccaaa"
    Output: "a2b1c5a3"

  Example 2:
    Input: "abbccd"
    Output: "abbccd"
    Explanation:
    The compressed string is "a1b2c2d1", which is longer than the original string.

  Note:
    0 <= S.length <= 50000

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/compress-string-lcci
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func compressString(S string) string {
	if len(S) == 0 {
		return ""
	}
	compressed := ""
	cur, cnt := "", 0
	for i := range S {
		if S[i:i+1] == cur {
			cnt++
		} else if i > 0 {
			compressed += cur + strconv.Itoa(cnt)
			cur, cnt = S[i:i+1], 1
		} else {
			cur, cnt = S[i:i+1], 1
		}
	}
	compressed += cur + strconv.Itoa(cnt)
	if len(compressed) < len(S) {
		return compressed
	}
	return S
}
