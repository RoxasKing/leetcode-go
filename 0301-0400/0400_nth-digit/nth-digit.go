package main

import (
	"strconv"
)

/*
  Given an integer n, return the nth digit of the infinite integer sequence [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, ...].

  Example 1:
    Input: n = 3
    Output: 3

  Example 2:
    Input: n = 11
    Output: 0
    Explanation: The 11th digit of the sequence 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, ... is a 0, which is part of the number 10.

  Constraints:
    1 <= n <= 2^31 - 1

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/nth-digit
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Bit Manipulation
// Math

func findNthDigit(n int) int {
	b, i := 1, 1
	for ; n > i*9*b; b, i = b*10, i+1 {
		n -= i * 9 * b
	}
	x := n / i
	n -= x * i
	if n == 0 {
		return int(strconv.Itoa(b + x - 1)[i-1] - '0')
	}
	return int(strconv.Itoa(b + x)[n-1] - '0')
}
