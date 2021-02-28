package main

/*
  Given a string s and an integer k, return the length of the longest substring of s such that the frequency of each character in this substring is greater than or equal to k.

  Example 1:
    Input: s = "aaabb", k = 3
    Output: 3
    Explanation: The longest substring is "aaa", as 'a' is repeated 3 times.

  Example 2:
    Input: s = "ababbc", k = 2
    Output: 5
    Explanation: The longest substring is "ababb", as 'a' is repeated 2 times and 'b' is repeated 3 times.

  Constraints:
    1. 1 <= s.length <= 10^4
    2. s consists of only lowercase English letters.
    3. 1 <= k <= 10^5

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/longest-substring-with-at-least-k-repeating-characters
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Sliding Window
func longestSubstring(s string, k int) int {
	out := 0
	for t := 1; t <= 26; t++ { // limit on the number of character types
		cnt := [26]int{}
		typ := 0 // current number of character types
		lss := 0 // frequency > 0 && < k 's character
		for l, r := 0, 0; r < len(s); r++ {
			c := s[r] - 'a'
			if cnt[c] == 0 {
				typ++
				lss++
			}
			cnt[c]++
			if cnt[c] == k {
				lss--
			}
			for typ > t {
				c := s[l] - 'a'
				if cnt[c] == k {
					lss++
				}
				cnt[c]--
				if cnt[c] == 0 {
					typ--
					lss--
				}
				l++
			}
			if lss == 0 {
				out = Max(out, r+1-l)
			}
		}
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
