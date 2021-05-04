package main

/*
  Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes the value to go outside the signed 32-bit integer range [-2^31, 2^31 - 1], then return 0.

  Assume the environment does not allow you to store 64-bit integers (signed or unsigned).

  Example 1:
    Input: x = 123
    Output: 321

  Example 2:
    Input: x = -123
    Output: -321

  Example 3:
    Input: x = 120
    Output: 21

  Example 4:
    Input: x = 0
    Output: 0

  Constraints:
    -2^31 <= x <= 2^31 - 1

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/reverse-integer
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Math
func reverse(x int) int {
	out := 0
	for x != 0 {
		if out < (-1<<31-x%10)/10 || (1<<31-1-x%10)/10 < out {
			return 0
		}
		out = out*10 + x%10
		x /= 10
	}
	return out
}
