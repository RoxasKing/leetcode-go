package My_LeetCode_In_Go

/*
  实现 strStr() 函数。
  给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。
*/

func strStr(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}
	if haystack == needle {
		return 0
	}
	for i := 0; i < len(haystack); i++ {
		if i+len(needle) <= len(haystack) &&
			haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}
