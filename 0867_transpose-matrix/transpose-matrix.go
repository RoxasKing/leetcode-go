package main

/*
  Given a 2D integer array matrix, return the transpose of matrix.

  The transpose of a matrix is the matrix flipped over its main diagonal, switching the matrix's row and column indices.

  Example 1:
    Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
    Output: [[1,4,7],[2,5,8],[3,6,9]]

  Example 2:
    Input: matrix = [[1,2,3],[4,5,6]]
    Output: [[1,4],[2,5],[3,6]]

  Constraints:
    1. m == matrix.length
    2. n == matrix[i].length
    3. 1 <= m, n <= 1000
    4. 1 <= m * n <= 10^5
    5. -10^9 <= matrix[i][j] <= 10^9

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/transpose-matrix
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func transpose(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
	out := make([][]int, n)
	for i := range out {
		out[i] = make([]int, m)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			out[j][i] = matrix[i][j]
		}
	}
	return out
}
