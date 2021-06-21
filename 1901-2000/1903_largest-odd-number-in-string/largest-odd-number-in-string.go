package main

/*
  You are given a string num, representing a large integer. Return the largest-valued odd integer (as a string) that is a non-empty substring of num, or an empty string "" if no odd integer exists.

  A substring is a contiguous sequence of characters within a string.

  Example 1:
    Input: num = "52"
    Output: "5"
    Explanation: The only non-empty substrings are "5", "2", and "52". "5" is the only odd number.

  Example 2:
    Input: num = "4206"
    Output: ""
    Explanation: There are no odd numbers in "4206".

  Example 3:
    Input: num = "35427"
    Output: "35427"
    Explanation: "35427" is already an odd number.

  Constraints:
    1. 1 <= num.length <= 10^5
    2. num only consists of digits and does not contain any leading zeros.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/largest-odd-number-in-string
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func largestOddNumber(num string) string {
	out := ""
	for i := range num {
		if int(num[i]-'0')&1 == 1 {
			out = num[:i+1]
		}
	}
	return out
}
