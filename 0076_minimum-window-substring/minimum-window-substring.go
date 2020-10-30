package main

/*
  给你一个字符串 S、一个字符串 T 。请你设计一种算法，可以在 O(n) 的时间复杂度内，从字符串 S 里面找出：包含 T 所有字符的最小子串。

  示例：
    输入：S = "ADOBECODEBANC", T = "ABC"
    输出："BANC"

  提示：
    如果 S 中不存这样的子串，则返回空字符串 ""。
    如果 S 中存在这样的子串，我们保证它是唯一的答案。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/minimum-window-substring
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Sliding Window
func minWindow(s string, t string) string {
	out := ""
	count := len(t)
	dict := [128]int{}
	for i := range t {
		dict[t[i]]++
	}
	for l, r := 0, 0; r < len(s); r++ {
		if dict[s[r]] > 0 {
			count--
		}
		dict[s[r]]--
		for l < r && dict[s[l]] < 0 {
			dict[s[l]]++
			l++
		}
		if count == 0 && (out == "" || r+1-l < len(out)) {
			out = s[l : r+1]
		}
	}
	return out
}
