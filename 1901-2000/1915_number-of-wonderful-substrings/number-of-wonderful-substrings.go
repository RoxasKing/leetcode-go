package main

/*
  A wonderful string is a string where at most one letter appears an odd number of times.

    For example, "ccjjc" and "abab" are wonderful, but "ab" is not.

  Given a string word that consists of the first ten lowercase English letters ('a' through 'j'), return the number of wonderful non-empty substrings in word. If the same substring appears multiple times in word, then count each occurrence separately.

  A substring is a contiguous sequence of characters in a string.

  Example 1:
    Input: word = "aba"
    Output: 4
    Explanation: The four wonderful substrings are underlined below:
      - "aba" -> "a"
      - "aba" -> "b"
      - "aba" -> "a"
      - "aba" -> "aba"

  Example 2:
    Input: word = "aabb"
    Output: 9
    Explanation: The nine wonderful substrings are underlined below:
      - "aabb" -> "a"
      - "aabb" -> "aa"
      - "aabb" -> "aab"
      - "aabb" -> "aabb"
      - "aabb" -> "a"
      - "aabb" -> "abb"
      - "aabb" -> "b"
      - "aabb" -> "bb"
      - "aabb" -> "b"

  Example 3:
    Input: word = "he"
    Output: 2
    Explanation: The two wonderful substrings are underlined below:
      - "he" -> "h"
      - "he" -> "e"

  Constraints:
    1. 1 <= word.length <= 10^5
    2. word consists of lowercase English letters from 'a' to 'j'.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/number-of-wonderful-substrings
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Bit Manipulation
// Prefix Sum

func wonderfulSubstrings(word string) int64 {
	cnt := map[int64]int64{0: 1}
	var out, sum int64
	for i := range word {
		sum ^= 1 << int(word[i]-'a')
		out += cnt[sum]
		for j := 0; j < 10; j++ {
			out += cnt[sum^(1<<j)]
		}
		cnt[sum]++
	}
	return out
}
