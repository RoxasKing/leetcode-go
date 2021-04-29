package main

/*
  Given a string s, find the length of the longest substring without repeating characters.

  Example 1:
    Input: s = "abcabcbb"
    Output: 3
    Explanation: The answer is "abc", with the length of 3.

  Example 2:
    Input: s = "bbbbb"
    Output: 1
    Explanation: The answer is "b", with the length of 1.

  Example 3:
    Input: s = "pwwkew"
    Output: 3
    Explanation: The answer is "wke", with the length of 3.
      Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

  Example 4:
    Input: s = ""
    Output: 0

  Constraints:
    1. 0 <= s.length <= 5 * 10^4
    2. s consists of English letters, digits, symbols and spaces.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Sliding Window
func lengthOfLongestSubstring(s string) int {
	out := 0
	cnt := [128]int{}
	for l, r := 0, 0; r < len(s); r++ {
		cnt[s[r]]++
		for l < r && cnt[s[r]] > 1 {
			cnt[s[l]]--
			l++
		}
		if r+1-l > out {
			out = r + 1 - l
		}
	}
	return out
}
