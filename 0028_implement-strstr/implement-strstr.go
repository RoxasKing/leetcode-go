package main

/*
  实现 strStr() 函数。

  给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。

  说明:
  当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。
  对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与C语言的 strstr() 以及 Java的 indexOf() 定义相符。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/implement-strstr
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		if isEqual(haystack[i:i+len(needle)], needle) {
			return i
		}
	}
	return -1
}

func isEqual(a, b string) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
