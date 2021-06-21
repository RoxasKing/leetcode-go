package main

/*
  Given a positive integer n, you can apply one of the following operations:

    1. If n is even, replace n with n / 2.
    2. If n is odd, replace n with either n + 1 or n - 1.

  Return the minimum number of operations needed for n to become 1.

  Example 1:
    Input: n = 8
    Output: 3
    Explanation: 8 -> 4 -> 2 -> 1

  Example 2:
    Input: n = 7
    Output: 4
    Explanation: 7 -> 8 -> 4 -> 2 -> 1
      or 7 -> 6 -> 3 -> 2 -> 1

  Example 3:
    Input: n = 4
    Output: 2

  Constraints:
    1 <= n <= 2^31 - 1

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/integer-replacement
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Bit Manipulation

func integerReplacement(n int) int {
	out := 0
	for ; n > 1; out++ {
		if n&1 == 0 {
			n >>= 1
		} else if n > 0b11 && n&0b11 == 0b11 {
			n++
		} else {
			n--
		}
	}
	return out
}
