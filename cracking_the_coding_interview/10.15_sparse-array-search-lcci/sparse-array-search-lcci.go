package main

/*
  Given a sorted array of strings that is interspersed with empty strings, write a method to find the location of a given string.

  Example1:
    Input: words = ["at", "", "", "", "ball", "", "", "car", "", "","dad", "", ""], s = "ta"
    Output: -1
    Explanation: Return -1 if s is not in words.

  Example2:
    Input: words = ["at", "", "", "", "ball", "", "", "car", "", "","dad", "", ""], s = "ball"
    Output: 4

  Note:
    1 <= words.length <= 1000000

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/sparse-array-search-lcci
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Binary Search

func findString(words []string, s string) int {
	n := len(words)
	l, r := 0, n-1
	for l < r {
		m := (l + r) >> 1
		for l < m && words[m] == "" {
			m--
		}
		if words[m] < s {
			l = m + 1
		} else {
			r = m
		}
	}
	if words[l] == s {
		return l
	}
	return -1
}
