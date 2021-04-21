package main

/*
  You are given two strings word1 and word2. Merge the strings by adding letters in alternating order, starting with word1. If a string is longer than the other, append the additional letters onto the end of the merged string.

  Return the merged string.

  Example 1:
    Input: word1 = "abc", word2 = "pqr"
    Output: "apbqcr"
    Explanation: The merged string will be merged as so:
    word1:  a   b   c
    word2:    p   q   r
    merged: a p b q c r

  Example 2:
    Input: word1 = "ab", word2 = "pqrs"
    Output: "apbqrs"
    Explanation: Notice that as word2 is longer, "rs" is appended to the end.
    word1:  a   b
    word2:    p   q   r   s
    merged: a p b q   r   s

  Example 3:
    Input: word1 = "abcd", word2 = "pq"
    Output: "apbqcd"
    Explanation: Notice that as word1 is longer, "cd" is appended to the end.
    word1:  a   b   c   d
    word2:    p   q
    merged: a p b q c   d

  Constraints:
    1 <= word1.length, word2.length <= 100
    word1 and word2 consist of lowercase English letters.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/merge-strings-alternately
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func mergeAlternately(word1 string, word2 string) string {
	i, j := 0, 0
	out := ""
	flg := 0
	for i < len(word1) && j < len(word2) {
		if flg == 0 {
			out += word1[i : i+1]
			i++
		} else {
			out += word2[j : j+1]
			j++
		}
		flg ^= 1
	}
	out += word1[i:]
	out += word2[j:]
	return out
}
