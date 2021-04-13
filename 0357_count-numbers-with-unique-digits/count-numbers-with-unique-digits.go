package main

/*
  Given an integer n, return the count of all numbers with unique digits, x, where 0 <= x < 10n.

  Example 1:
    Input: n = 2
    Output: 91
    Explanation: The answer should be the total numbers in the range of 0 ≤ x < 100, excluding 11,22,33,44,55,66,77,88,99

  Example 2:
    Input: n = 0
    Output: 1

  Constraints:
    0 <= n <= 8

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/count-numbers-with-unique-digits
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Math
func countNumbersWithUniqueDigits(n int) int {
	out := 1
	cur := 1
	base := 9
	for i := 1; i <= n; i++ {
		if i > 2 {
			base--
		}
		cur *= base
		out += cur
	}
	return out
}
