package main

/*
  Given two strings s and t, return the minimum window in s which will contain all the characters in t. If there is no such window in s that covers all characters in t, return the empty string "".

  Note that If there is such a window, it is guaranteed that there will always be only one unique minimum window in s.

  Example 1:
    Input: s = "ADOBECODEBANC", t = "ABC"
    Output: "BANC"

  Example 2:
    Input: s = "a", t = "a"
    Output: "a"

  Constraints:
    1. 1 <= s.length, t.length <= 10^5
    2. s and t consist of English letters.

  Follow up: Could you find an algorithm that runs in O(n) time?

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/minimum-window-substring
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Sliding Window + Two Pointers
func minWindow(s string, t string) string {
	cnt := [128]int{}
	for i := range t {
		cnt[t[i]]++
	}
	out := ""
	for l, r, include := 0, 0, 0; r < len(s); r++ {
		if cnt[s[r]] > 0 {
			include++
		}
		cnt[s[r]]--
		for l <= r && cnt[s[l]] < 0 {
			cnt[s[l]]++
			l++
		}
		if include == len(t) && (out == "" || len(out) > r+1-l) {
			out = s[l : r+1]
		}
	}
	return out
}
