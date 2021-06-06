package main

/*
  Given an integer n, count the total number of digit 1 appearing in all non-negative integers less than or equal to n.

  Example 1:
    Input: n = 13
    Output: 6

  Example 2:
    Input: n = 0
    Output: 0

  Constraints:
    0 <= n <= 2 * 10^9

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/number-of-digit-one
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Math

func countDigitOne(n int) int {
	out := 0
	for base := 1; n >= base; base *= 10 {
		l := n / (base * 10)
		m := n % (base * 10) / base
		r := n % base
		out += l * base
		if m == 1 {
			out += r + 1
		} else if m > 1 {
			out += base
		}
	}
	return out
}
