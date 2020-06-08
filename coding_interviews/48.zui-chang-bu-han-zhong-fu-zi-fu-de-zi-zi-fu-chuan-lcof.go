package codinginterviews

/*
  请从字符串中找出一个最长的不包含重复字符的子字符串，计算该最长子字符串的长度。

  提示：
    s.length <= 40000
*/

func lengthOfLongestSubstring(s string) int {
	var max int
	count := [128]int{}
	l, r := 0, 0
	for r = range s {
		count[s[r]]++
		for l < r && count[s[r]] > 1 {
			count[s[l]]--
			l++
		}
		max = Max(max, r+1-l)
	}
	return max
}
