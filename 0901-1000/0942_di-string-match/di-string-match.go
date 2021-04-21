package main

/*
  A permutation perm of n + 1 integers of all the integers in the range [0, n] can be represented as a string s of length n where:
    1. s[i] == 'I' if perm[i] < perm[i + 1], and
    1. s[i] == 'D' if perm[i] > perm[i + 1].
  Given a string s, reconstruct the permutation perm and return it. If there are multiple valid permutations perm, return any of them.

  Example 1:
    Input: s = "IDID"
    Output: [0,4,1,3,2]

  Example 2:
    Input: s = "III"
    Output: [0,1,2,3]

  Example 3:
    Input: s = "DDI"
    Output: [3,2,0,1]

  Constraints:
    1. 1 <= s.length <= 10^5
    2. s[i] is either 'I' or 'D'.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/di-string-match
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Math
func diStringMatch(s string) []int {
	n := len(s)
	out := make([]int, 0, n+1)
	l, r := 0, n
	for i := range s {
		if s[i] == 'I' {
			out = append(out, l)
			l++
		} else {
			out = append(out, r)
			r--
		}
	}
	out = append(out, l)
	return out
}
