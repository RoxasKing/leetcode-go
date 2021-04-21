package main

/*
  Give a string s, count the number of non-empty (contiguous) substrings that have the same number of 0's and 1's, and all the 0's and all the 1's in these substrings are grouped consecutively.

  Substrings that occur multiple times are counted the number of times they occur.

  Example 1:
    Input: "00110011"
    Output: 6
    Explanation: There are 6 substrings that have equal number of consecutive 1's and 0's: "0011", "01", "1100", "10", "0011", and "01".

    Notice that some of these substrings repeat and are counted the number of times they occur.

    Also, "00110011" is not a valid substring because all the 0's (and 1's) are not grouped together.

  Example 2:
    Input: "10101"
    Output: 4
    Explanation: There are 4 substrings: "10", "01", "10", "01" that have equal number of consecutive 1's and 0's.

  Note:
    1. s.length will be between 1 and 50,000.
    2. s will only consist of "0" or "1" characters.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/count-binary-substrings
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Two Pointers
func countBinarySubstrings(s string) int {
	out := 0
	n := len(s)
	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			continue
		}
		for l, r := i, i+1; l >= 0 && s[l] == s[i] && r < n && s[r] == s[i+1]; l, r = l-1, r+1 {
			out++
		}
	}
	return out
}
