package main

import (
	"strconv"
)

/*
  You are given a string s that consists of only digits.

  Check if we can split s into two or more non-empty substrings such that the numerical values of the substrings are in descending order and the difference between numerical values of every two adjacent substrings is equal to 1.

    1. For example, the string s = "0090089" can be split into ["0090", "089"] with numerical values [90,89]. The values are in descending order and adjacent values differ by 1, so this way is valid.
    2. Another example, the string s = "001" can be split into ["0", "01"], ["00", "1"], or ["0", "0", "1"]. However all the ways are invalid because they have numerical values [0,1], [0,1], and [0,0,1] respectively, all of which are not in descending order.

  Return true if it is possible to split s​​​​​​ as described above, or false otherwise.

  A substring is a contiguous sequence of characters in a string.

  Example 1:
    Input: s = "1234"
    Output: false
    Explanation: There is no valid way to split s.

  Example 2:
    Input: s = "050043"
    Output: true
    Explanation: s can be split into ["05", "004", "3"] with numerical values [5,4,3].
      The values are in descending order with adjacent values differing by 1.

  Example 3:
    Input: s = "9080701"
    Output: false
    Explanation: There is no valid way to split s.

  Example 4:
    Input: s = "10009998"
    Output: true
    Explanation: s can be split into ["100", "099", "98"] with numerical values [100,99,98].
      The values are in descending order with adjacent values differing by 1.

  Constraints:
    1. 1 <= s.length <= 20
    2. s only consists of digits.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/splitting-a-string-into-descending-consecutive-values
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Backtracking
func splitString(s string) bool {
	return dfs(s, len(s), 0, []int{})
}

func dfs(s string, n, start int, cur []int) bool {
	if start == n {
		return len(cur) > 1
	}
	for i := start + 1; i <= n; i++ {
		num, _ := strconv.Atoi(s[start:i])
		if len(cur) > 0 && cur[len(cur)-1]-1 != num {
			continue
		}
		cur = append(cur, num)
		if dfs(s, n, i, cur) {
			return true
		}
		cur = cur[:len(cur)-1]
	}
	return false
}
