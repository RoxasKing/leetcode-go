package main

/*
  Given an m x n matrix, return all elements of the matrix in spiral order.

  Example 1:
    Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
    Output: [1,2,3,6,9,8,7,4,5]

  Example 2:
    Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
    Output: [1,2,3,4,8,12,11,10,9,5,6,7]

  Constraints:
    1. m == matrix.length
    2. n == matrix[i].length
    3. 1 <= m, n <= 10
    4. -100 <= matrix[i][j] <= 100

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/spiral-matrix
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	out := make([]int, 0, m*n)
	u, d, l, r := 0, m-1, 0, n-1
	forward := 0
	for u <= d && l <= r {
		switch forward {
		case 0:
			for i := l; i <= r; i++ {
				out = append(out, matrix[u][i])
			}
			u++
		case 1:
			for i := u; i <= d; i++ {
				out = append(out, matrix[i][r])
			}
			r--
		case 2:
			for i := r; i >= l; i-- {
				out = append(out, matrix[d][i])
			}
			d--
		case 3:
			for i := d; i >= u; i-- {
				out = append(out, matrix[i][l])
			}
			l++
		}
		forward = (forward + 1) % 4
	}
	return out
}
