package main

/*
  Given a positive integer n, generate an n x n matrix filled with elements from 1 to n2 in spiral order.

  Example 1:
    Input: n = 3
    Output: [[1,2,3],[8,9,4],[7,6,5]]

  Example 2:
    Input: n = 1
    Output: [[1]]

  Constraints:
    1 <= n <= 20

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/spiral-matrix-ii
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func generateMatrix(n int) [][]int {
	out := make([][]int, n)
	for i := range out {
		out[i] = make([]int, n)
	}
	u, d, l, r := 0, n-1, 0, n-1
	num, move := 1, 0
	for u <= d && l <= r {
		switch move {
		case 0:
			for i := l; i <= r; i++ {
				out[u][i] = num
				num++
			}
			u++
		case 1:
			for i := u; i <= d; i++ {
				out[i][r] = num
				num++
			}
			r--
		case 2:
			for i := r; i >= l; i-- {
				out[d][i] = num
				num++
			}
			d--
		case 3:
			for i := d; i >= u; i-- {
				out[i][l] = num
				num++
			}
			l++
		}
		move = (move + 1) % 4
	}
	return out
}
