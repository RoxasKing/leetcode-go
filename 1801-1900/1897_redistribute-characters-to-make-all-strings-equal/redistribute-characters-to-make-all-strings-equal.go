package main

/*
  You are given an array of strings words (0-indexed).

  In one operation, pick two distinct indices i and j, where words[i] is a non-empty string, and move any character from words[i] to any position in words[j].

  Return true if you can make every string in words equal using any number of operations, and false otherwise.

  Example 1:
    Input: words = ["abc","aabc","bc"]
    Output: true
    Explanation: Move the first 'a' in words[1] to the front of words[2],
      to make words[1] = "abc" and words[2] = "abc".
      All the strings are now equal to "abc", so return true.

  Example 2:
    Input: words = ["ab","a"]
    Output: false
    Explanation: It is impossible to make all the strings equal using the operation.

  Constraints:
    1. 1 <= words.length <= 100
    2. 1 <= words[i].length <= 100
    3. words[i] consists of lowercase English letters.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/redistribute-characters-to-make-all-strings-equal
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func makeEqual(words []string) bool {
	n := len(words)
	cnt := [26]int{}
	for _, word := range words {
		for i := range word {
			cnt[word[i]-'a']++
		}
	}
	for i := 0; i < 26; i++ {
		if cnt[i]%n > 0 {
			return false
		}
	}
	return true
}
