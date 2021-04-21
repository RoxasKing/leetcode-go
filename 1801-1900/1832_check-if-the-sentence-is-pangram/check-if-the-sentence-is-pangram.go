package main

/*
  A pangram is a sentence where every letter of the English alphabet appears at least once.

  Given a string sentence containing only lowercase English letters, return true if sentence is a pangram, or false otherwise.

  Example 1:
    Input: sentence = "thequickbrownfoxjumpsoverthelazydog"
    Output: true
    Explanation: sentence contains at least one of every letter of the English alphabet.

  Example 2:
    Input: sentence = "leetcode"
    Output: false

  Constraints:
    1. 1 <= sentence.length <= 1000
    2. sentence consists of lowercase English letters.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/check-if-the-sentence-is-pangram
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func checkIfPangram(sentence string) bool {
	mark := [26]bool{}
	cnt := 0
	for i := range sentence {
		idx := int(sentence[i] - 'a')
		if mark[idx] {
			continue
		}
		cnt++
		mark[idx] = true
	}
	return cnt == 26
}
