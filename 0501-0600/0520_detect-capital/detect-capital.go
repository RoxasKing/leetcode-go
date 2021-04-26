package main

/*
  We define the usage of capitals in a word to be right when one of the following cases holds:
    1. All letters in this word are capitals, like "USA".
    2. All letters in this word are not capitals, like "leetcode".
    3. Only the first letter in this word is capital, like "Google".
  Given a string word, return true if the usage of capitals in it is right.

  Example 1:
    Input: word = "USA"
    Output: true

  Example 2:
    Input: word = "FlaG"
    Output: false

  Constraints:
    1. 1 <= word.length <= 100
    2. word consists of lowercase and uppercase English letters.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/detect-capital
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func detectCapitalUse(word string) bool {
	f := word[0]
	word = word[1:]
	if isUpperCase(f) {
		cnt := 0
		for i := range word {
			if isUpperCase(word[i]) {
				cnt++
			}
		}
		if 0 < cnt && cnt < len(word) {
			return false
		}
	} else {
		for i := range word {
			if isUpperCase(word[i]) {
				return false
			}
		}
	}
	return true
}

func isUpperCase(ch byte) bool {
	return 'A' <= ch && ch <= 'Z'
}
