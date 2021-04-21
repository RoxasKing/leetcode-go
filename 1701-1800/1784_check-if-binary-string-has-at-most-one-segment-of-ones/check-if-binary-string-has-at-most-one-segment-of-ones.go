package main

/*
  Given a binary string s ​​​​​without leading zeros, return true​​​ if s contains at most one contiguous segment of ones. Otherwise, return false.

  Example 1:
    Input: s = "1001"
    Output: false
    Explanation: The ones do not form a contiguous segment.

  Example 2:
    Input: s = "110"
    Output: true

  Constraints:
    1. 1 <= s.length <= 100
    2. s[i]​​​​ is either '0' or '1'.
    3. s[0] is '1'.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/check-if-binary-string-has-at-most-one-segment-of-ones
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func checkOnesSegment(s string) bool {
	cnt := 0
	for s != "" {
		for s != "" && s[0] == '0' {
			s = s[1:]
		}
		count := 0
		for s != "" && s[0] == '1' {
			s = s[1:]
			count++
		}
		if count > 0 {
			cnt++
		}
		if cnt > 1 {
			return false
		}
	}
	return cnt == 1
}
