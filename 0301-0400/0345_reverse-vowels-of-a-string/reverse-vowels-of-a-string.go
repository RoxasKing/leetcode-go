package main

/*
  Given a string s, reverse only all the vowels in the string and return it.

  The vowels are 'a', 'e', 'i', 'o', and 'u', and they can appear in both cases.

  Example 1:
    Input: s = "hello"
    Output: "holle"

  Example 2:
    Input: s = "leetcode"
    Output: "leotcede"

  Constraints:
    1. 1 <= s.length <= 3 * 10^5
    2. s consist of printable ASCII characters.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/reverse-vowels-of-a-string
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func reverseVowels(s string) string {
	chs := []byte(s)
	n := len(chs)
	l, r := 0, n-1
	for l < r {
		if isVowel(chs[l]) && isVowel(chs[r]) {
			chs[l], chs[r] = chs[r], chs[l]
			l++
			r--
		}
		for l < r && !isVowel(chs[l]) {
			l++
		}
		for l < r && !isVowel(chs[r]) {
			r--
		}
	}
	return string(chs)
}

func isVowel(ch byte) bool {
	return ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' ||
		ch == 'A' || ch == 'E' || ch == 'I' || ch == 'O' || ch == 'U'
}
