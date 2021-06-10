package main

/*
  Given an integer num, repeatedly add all its digits until the result has only one digit, and return it.

  Example 1:
    Input: num = 38
    Output: 2
    Explanation: The process is
      38 --> 3 + 8 --> 11
      11 --> 1 + 1 --> 2
      Since 2 has only one digit, return it.

  Example 2:
    Input: num = 0
    Output: 0

  Constraints:
    0 <= num <= 2^31 - 1

  Follow up: Could you do it without any loop/recursion in O(1) runtime?

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/add-digits
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Math

func addDigits(num int) int {
	return (num-1)%9 + 1
}
