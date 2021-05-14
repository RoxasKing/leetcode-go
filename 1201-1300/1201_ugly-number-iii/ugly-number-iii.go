package main

/*
  An ugly number is a positive integer that is divisible by a, b, or c.

  Given four integers n, a, b, and c, return the nth ugly number.

  Example 1:
    Input: n = 3, a = 2, b = 3, c = 5
    Output: 4
    Explanation: The ugly numbers are 2, 3, 4, 5, 6, 8, 9, 10... The 3rd is 4.

  Example 2:
    Input: n = 4, a = 2, b = 3, c = 4
    Output: 6
    Explanation: The ugly numbers are 2, 3, 4, 6, 8, 9, 10, 12... The 4th is 6.

  Example 3:
    Input: n = 5, a = 2, b = 11, c = 13
    Output: 10
    Explanation: The ugly numbers are 2, 4, 6, 8, 10, 11, 12, 13... The 5th is 10.

  Example 4:
    Input: n = 1000000000, a = 2, b = 217983653, c = 336916467
    Output: 1999999984

  Constraints:
    1. 1 <= n, a, b, c <= 10^9
    2. 1 <= a * b * c <= 10^18
    3. It is guaranteed that the result will be in range [1, 2 * 10^9].

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/ugly-number-iii
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Math + Binary Search

func nthUglyNumber(n int, a int, b int, c int) int {
	d, e, f := a/Gcd(a, b)*b, b/Gcd(b, c)*c, a/Gcd(a, c)*c
	g := a / Gcd(a, e) * e
	l, r := 1, 4*int(1e9)+1
	for l < r {
		m := (l + r) >> 1
		if m/a+m/b+m/c-m/d-m/e-m/f+m/g < n {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}

func Gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
