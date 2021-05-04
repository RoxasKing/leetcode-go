package main

import "strings"

/*
  Given an input string s, reverse the order of the words.

  A word is defined as a sequence of non-space characters. The words in s will be separated by at least one space.

  Return a string of the words in reverse order concatenated by a single space.

  Note that s may contain leading or trailing spaces or multiple spaces between two words. The returned string should only have a single space separating the words. Do not include any extra spaces.

  Example 1:
    Input: s = "the sky is blue"
    Output: "blue is sky the"

  Example 2:
    Input: s = "  hello world  "
    Output: "world hello"
    Explanation: Your reversed string should not contain leading or trailing spaces.

  Example 3:
    Input: s = "a good   example"
    Output: "example good a"
    Explanation: You need to reduce multiple spaces between two words to a single space in the reversed string.

  Example 4:
    Input: s = "  Bob    Loves  Alice   "
    Output: "Alice Loves Bob"

  Example 5:
    Input: s = "Alice does not even like bob"
    Output: "bob like even not does Alice"

  Constraints:
    1. 1 <= s.length <= 10^4
    2. s contains English letters (upper-case and lower-case), digits, and spaces ' '.
    3. There is at least one word in s.

  Follow up: Could you solve it in-place with O(1) extra space?

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/reverse-words-in-a-string
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func reverseWords(s string) string {
	arr := []string{}
	n := len(s)

	for i := 0; i < n; i++ {
		if s[i] == ' ' {
			continue
		}
		j := i
		for j+1 < n && s[j+1] != ' ' {
			j++
		}
		arr = append(arr, s[i:j+1])
		i = j
	}

	for l, r := 0, len(arr)-1; l < r; l, r = l+1, r-1 {
		arr[l], arr[r] = arr[r], arr[l]
	}

	return strings.Join(arr, " ")
}
