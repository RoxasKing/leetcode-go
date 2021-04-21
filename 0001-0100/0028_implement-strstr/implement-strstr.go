package main

/*
  Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

  Clarification:

  What should we return when needle is an empty string? This is a great question to ask during an interview.

  For the purpose of this problem, we will return 0 when needle is an empty string. This is consistent to C's strstr() and Java's indexOf().

  Example 1:
    Input: haystack = "hello", needle = "ll"
    Output: 2

  Example 2:
    Input: haystack = "aaaaa", needle = "bba"
    Output: -1

  Example 3:
    Input: haystack = "", needle = ""
    Output: 0

  Constraints:
    1. 0 <= haystack.length, needle.length <= 5 * 10^4
    2. haystack and needle consist of only lower-case English characters.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/implement-strstr
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func strStr(haystack string, needle string) int {
	m, n := len(haystack), len(needle)
	for i := 0; i < m-n+1; i++ {
		if haystack[i:i+n] == needle {
			return i
		}
	}
	return -1
}

// Important!

// KMP
func strStr2(haystack string, needle string) int {
	m, n := len(haystack), len(needle)
	if n == 0 {
		return 0
	}
	pi := make([]int, n)
	for i, j := 1, 0; i < n; i++ {
		for j > 0 && needle[i] != needle[j] {
			j = pi[j-1]
		}
		if needle[i] == needle[j] {
			j++
		}
		pi[i] = j
	}
	for i, j := 0, 0; i < m; i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = pi[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == n {
			return i - n + 1
		}
	}
	return -1
}
