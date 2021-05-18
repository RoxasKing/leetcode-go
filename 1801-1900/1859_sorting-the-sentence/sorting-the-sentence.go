package main

import "strings"

/*
  A sentence is a list of words that are separated by a single space with no leading or trailing spaces. Each word consists of lowercase and uppercase English letters.

  A sentence can be shuffled by appending the 1-indexed word position to each word then rearranging the words in the sentence.

  For example, the sentence "This is a sentence" can be shuffled as "sentence4 a3 is2 This1" or "is2 sentence4 This1 a3".

  Given a shuffled sentence s containing no more than 9 words, reconstruct and return the original sentence.

  Example 1:
    Input: s = "is2 sentence4 This1 a3"
    Output: "This is a sentence"
    Explanation: Sort the words in s to their original positions "This1 is2 a3 sentence4", then remove the numbers.

  Example 2:
    Input: s = "Myself2 Me1 I4 and3"
    Output: "Me Myself and I"
    Explanation: Sort the words in s to their original positions "Me1 Myself2 and3 I4", then remove the numbers.

  Constraints:
    1. 2 <= s.length <= 200
    2. s consists of lowercase and uppercase English letters, spaces, and digits from 1 to 9.
    3. The number of words in s is between 1 and 9.
    4. The words in s are separated by a single space.
    5. s contains no leading or trailing spaces.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/sorting-the-sentence
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func sortSentence(s string) string {
	strs := strings.Split(s, " ")
	out := make([]string, len(strs))
	for _, str := range strs {
		idx := int(str[len(str)-1]-'0') - 1
		out[idx] = str[:len(str)-1]
	}
	return strings.Join(out, " ")
}