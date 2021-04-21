package main

/*
  Given a positive integer num, write a function which returns True if num is a perfect square else False.

  Follow up: Do not use any built-in library function such as sqrt.

  Example 1:
    Input: num = 16
    Output: true

  Example 2:
    Input: num = 14
    Output: false

  Constraints:
    1 <= num <= 2^31 - 1

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/valid-perfect-square
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Binary Search
func isPerfectSquare(num int) bool {
	l, r := 0, num
	for l < r {
		m := (l + r) >> 1
		if m*m < num {
			l = m + 1
		} else {
			r = m
		}
	}
	return l*l == num
}
