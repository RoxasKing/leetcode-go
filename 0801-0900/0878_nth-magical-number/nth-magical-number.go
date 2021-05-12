package main

/*
  A positive integer is magical if it is divisible by either a or b.

  Given the three integers n, a, and b, return the nth magical number. Since the answer may be very large, return it modulo 109 + 7.

  Example 1:
    Input: n = 1, a = 2, b = 3
    Output: 2

  Example 2:
    Input: n = 4, a = 2, b = 3
    Output: 6

  Example 3:
    Input: n = 5, a = 2, b = 4
    Output: 10

  Example 4:
    Input: n = 3, a = 6, b = 4
    Output: 8

  Constraints:
    1. 1 <= n <= 10^9
    2. 2 <= a, b <= 4 * 10^4

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/nth-magical-number
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Math + Binary Search
func nthMagicalNumber(n int, a int, b int) int {
	c := a / Gcd(a, b) * b
	l, r := 1, int(1e15)
	for l < r {
		m := l + (r-l)>>1
		if m/a+m/b-m/c < n {
			l = m + 1
		} else {
			r = m
		}
	}
	return l % (1e9 + 7)
}

func Gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return Gcd(b, a%b)
}
