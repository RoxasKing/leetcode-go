package main

/*
  For two strings s and t, we say "t divides s" if and only if s = t + ... + t  (t concatenated with itself 1 or more times)

  Given two strings str1 and str2, return the largest string x such that x divides both str1 and str2.

  Example 1:
    Input: str1 = "ABCABC", str2 = "ABC"
    Output: "ABC"

  Example 2:
    Input: str1 = "ABABAB", str2 = "ABAB"
    Output: "AB"

  Example 3:
    Input: str1 = "LEET", str2 = "CODE"
    Output: ""

  Example 4:
    Input: str1 = "ABCDEF", str2 = "ABC"
    Output: ""

  Constraints:
    1. 1 <= str1.length <= 1000
    2. 1 <= str2.length <= 1000
    3. str1 and str2 consist of English uppercase letters.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/greatest-common-divisor-of-strings
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func gcdOfStrings(str1 string, str2 string) string {
	for str2 != "" {
		if len(str1) < len(str2) {
			str1, str2 = str2, str1
		}
		pre2 := str2
		for str1 != "" && str2 != "" && str1[0] == str2[0] {
			str1 = str1[1:]
			str2 = str2[1:]
		}
		if str2 != "" {
			return ""
		}
		str1, str2 = pre2, str1
	}
	return str1
}
