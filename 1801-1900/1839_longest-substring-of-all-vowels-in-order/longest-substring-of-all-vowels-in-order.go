package main

/*
  A string is considered beautiful if it satisfies the following conditions:
    1. Each of the 5 English vowels ('a', 'e', 'i', 'o', 'u') must appear at least once in it.
    2. The letters must be sorted in alphabetical order (i.e. all 'a's before 'e's, all 'e's before 'i's, etc.).

  For example, strings "aeiou" and "aaaaaaeiiiioou" are considered beautiful, but "uaeio", "aeoiu", and "aaaeeeooo" are not beautiful.

  Given a string word consisting of English vowels, return the length of the longest beautiful substring of word. If no such substring exists, return 0.

  A substring is a contiguous sequence of characters in a string.

  Example 1:
    Input: word = "aeiaaioaaaaeiiiiouuuooaauuaeiu"
    Output: 13
    Explanation: The longest beautiful substring in word is "aaaaeiiiiouuu" of length 13.

  Example 2:
    Input: word = "aeeeiiiioooauuuaeiou"
    Output: 5
    Explanation: The longest beautiful substring in word is "aeiou" of length 5.

  Example 3:
    Input: word = "a"
    Output: 0
    Explanation: There is no beautiful substring, so return 0.

  Constraints:
    1. 1 <= word.length <= 5 * 10^5
    2. word consists of characters 'a', 'e', 'i', 'o', and 'u'.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/longest-substring-of-all-vowels-in-order
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Two Pointers
func longestBeautifulSubstring(word string) int {
	n := len(word)
	arr := make([]int, n)
	for i := range word {
		if word[i] == 'a' {
			arr[i] = 0
		} else if word[i] == 'e' {
			arr[i] = 1
		} else if word[i] == 'i' {
			arr[i] = 2
		} else if word[i] == 'o' {
			arr[i] = 3
		} else {
			arr[i] = 4
		}
	}

	out := 0

	for l, r := 0, 0; r < n; {
		if arr[l] != 0 {
			l, r = l+1, l+1
			continue
		}
		if r+1 < n && (arr[r+1] == arr[r] || arr[r+1] == arr[r]+1) {
			r++
			if arr[r] == 4 && r+1-l > out {
				out = r + 1 - l
			}
		} else {
			l, r = r+1, r+1
		}
	}

	return out
}
