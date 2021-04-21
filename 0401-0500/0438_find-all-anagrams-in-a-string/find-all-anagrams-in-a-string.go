package main

/*
  Given two strings s and p, return an array of all the start indices of p's anagrams in s. You may return the answer in any order.

  Example 1:
    Input: s = "cbaebabacd", p = "abc"
    Output: [0,6]
    Explanation:
    The substring with start index = 0 is "cba", which is an anagram of "abc".
    The substring with start index = 6 is "bac", which is an anagram of "abc".

  Example 2:
    Input: s = "abab", p = "ab"
    Output: [0,1,2]
    Explanation:
    The substring with start index = 0 is "ab", which is an anagram of "ab".
    The substring with start index = 1 is "ba", which is an anagram of "ab".
    The substring with start index = 2 is "ab", which is an anagram of "ab".

  Constraints:
    1. 1 <= s.length, p.length <= 3 * 10^4
    2. s and p consist of lowercase English letters.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/find-all-anagrams-in-a-string
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Sliding Window
func findAnagrams(s string, p string) []int {
	cnt := [26]int{}
	for _, c := range p {
		cnt[c-'a']++
	}
	out := []int{}
	for l, r := 0, 0; r < len(s); r++ {
		cnt[s[r]-'a']--
		for l <= r && cnt[s[r]-'a'] < 0 {
			cnt[s[l]-'a']++
			l++
		}
		if r+1-l == len(p) {
			out = append(out, l)
			cnt[s[l]-'a']++
			l++
		}
	}
	return out
}
