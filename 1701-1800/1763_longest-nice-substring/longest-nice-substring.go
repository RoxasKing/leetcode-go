package main

/*
  A string s is nice if, for every letter of the alphabet that s contains, it appears both in uppercase and lowercase. For example, "abABB" is nice because 'A' and 'a' appear, and 'B' and 'b' appear. However, "abA" is not because 'b' appears, but 'B' does not.

  Given a string s, return the longest substring of s that is nice. If there are multiple, return the substring of the earliest occurrence. If there are none, return an empty string.

  Example 1:
    Input: s = "YazaAay"
    Output: "aAa"
    Explanation: "aAa" is a nice string because 'A/a' is the only letter of the alphabet in s, and both 'A' and 'a' appear.
      "aAa" is the longest nice substring.

  Example 2:
    Input: s = "Bb"
    Output: "Bb"
    Explanation: "Bb" is a nice string because both 'B' and 'b' appear. The whole string is a substring.

  Example 3:
    Input: s = "c"
    Output: ""
    Explanation: There are no nice substrings.

  Example 4:
    Input: s = "dDzeE"
    Output: "dD"
    Explanation: Both "dD" and "eE" are the longest nice substrings.
      As there are multiple longest nice substrings, return "dD" since it occurs earlier.

  Constraints:
    1. 1 <= s.length <= 100
    2. s consists of uppercase and lowercase English letters.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/longest-nice-substring
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func longestNiceSubstring(s string) string {
	out := ""
	n := len(s)
	for i := 0; i < n; i++ {
		f1 := [26]bool{}
		f2 := [26]bool{}
		for j := i; j < n; j++ {
			if 'a' <= s[j] && s[j] <= 'z' {
				f1[s[j]-'a'] = true
			} else {
				f2[s[j]-'A'] = true
			}

			ok := true
			for i := 0; i < 26; i++ {
				if f1[i] && !f2[i] || !f1[i] && f2[i] {
					ok = false
					break
				}
			}

			if ok && j+1-i > len(out) {
				out = s[i : j+1]
			}
		}
	}
	return out
}
