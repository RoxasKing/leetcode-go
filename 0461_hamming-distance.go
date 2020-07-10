package leetcode

/*
  两个整数之间的汉明距离指的是这两个数字对应二进制位不同的位置的数目。
  给出两个整数 x 和 y，计算它们之间的汉明距离。

  注意：
    0 ≤ x, y < 231.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/hamming-distance
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Bit Operation
func hammingDistance(x int, y int) int {
	var count int
	xor := x ^ y
	for xor != 0 {
		count++
		xor = xor & (xor - 1)
	}
	return count
}
