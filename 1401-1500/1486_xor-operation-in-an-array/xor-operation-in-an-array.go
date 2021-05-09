package main

/*
  Given an integer n and an integer start.

  Define an array nums where nums[i] = start + 2*i (0-indexed) and n == nums.length.

  Return the bitwise XOR of all elements of nums.

  Example 1:
    Input: n = 5, start = 0
    Output: 8
    Explanation: Array nums is equal to [0, 2, 4, 6, 8] where (0 ^ 2 ^ 4 ^ 6 ^ 8) = 8.
    Where "^" corresponds to bitwise XOR operator.

  Example 2:
    Input: n = 4, start = 3
    Output: 8
    Explanation: Array nums is equal to [3, 5, 7, 9] where (3 ^ 5 ^ 7 ^ 9) = 8.

  Example 3:
    Input: n = 1, start = 7
    Output: 7

  Example 4:
    Input: n = 10, start = 5
    Output: 2

  Constraints:
    1. 1 <= n <= 1000
    2. 0 <= start <= 1000
    3. n == nums.length

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/xor-operation-in-an-array
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Bit Manipulation
func xorOperation(n int, start int) int {
	out := start
	for i := 1; i < n; i++ {
		out ^= start + 2*i
	}
	return out
}
