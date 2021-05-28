package main

/*
  The Hamming distance between two integers is the number of positions at which the corresponding bits are different.

  Given two integers x and y, return the Hamming distance between them.

  Example 1:
    Input: x = 1, y = 4
    Output: 2
    Explanation:
      1   (0 0 0 1)
      4   (0 1 0 0)
             ↑   ↑
      The above arrows point to positions where the corresponding bits are different.

  Example 2:
    Input: x = 3, y = 1
    Output: 1

  Constraints:
    0 <= x, y <= 2^31 - 1

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/hamming-distance
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Bit Manipulation
// Brian Kernighan

func hammingDistance(x int, y int) int {
	out := 0
	for xor := x ^ y; xor > 0; xor &= xor - 1 {
		out++
	}
	return out
}
