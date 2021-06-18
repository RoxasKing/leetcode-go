package main

/*
  You have a list arr of all integers in the range [1, n] sorted in a strictly increasing order. Apply the following algorithm on arr:

    1. Starting from left to right, remove the first number and every other number afterward until you reach the end of the list.
    2. Repeat the previous step again, but this time from right to left, remove the rightmost number and every other number from the remaining numbers.
    3. Keep repeating the steps again, alternating left to right and right to left, until a single number remains.

  Given the integer n, return the last number that remains in arr.

  Example 1:
    Input: n = 9
    Output: 6
    Explanation:
      arr = [1, 2, 3, 4, 5, 6, 7, 8, 9]
      arr = [2, 4, 6, 8]
      arr = [2, 6]
      arr = [6]

  Example 2:
    Input: n = 1
    Output: 1

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/elimination-game
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Math
// DFS

func lastRemaining(n int) int {
	if n == 1 {
		return 1
	}
	return 2 * (n>>1 + 1 - lastRemaining(n>>1))
}
