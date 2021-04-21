package main

/*
  Given two strings s1 and s2, write a function to return true if s2 contains the permutation of s1. In other words, one of the first string's permutations is the substring of the second string.

  Example 1:
    Input: s1 = "ab" s2 = "eidbaooo"
    Output: True
    Explanation: s2 contains one permutation of s1 ("ba").

  Example 2:
    Input:s1= "ab" s2 = "eidboaoo"
    Output: False

  Constraints:
    The input strings only contain lower case letters.
    The length of both given strings is in range [1, 10,000].

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/permutation-in-string
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Sliding Window
func checkInclusion(s1 string, s2 string) bool {
	cnt := [26]int{}
	for i := range s1 {
		cnt[s1[i]-'a']++
	}
	for l, r := 0, 0; r < len(s2); r++ {
		cnt[s2[r]-'a']--
		for l <= r && cnt[s2[r]-'a'] < 0 {
			cnt[s2[l]-'a']++
			l++
		}
		if r+1-l == len(s1) {
			return true
		}
	}
	return false
}
