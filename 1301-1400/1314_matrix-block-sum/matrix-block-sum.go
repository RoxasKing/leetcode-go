package main

/*
  Given a m * n matrix mat and an integer K, return a matrix answer where each answer[i][j] is the sum of all elements mat[r][c] for i - K <= r <= i + K, j - K <= c <= j + K, and (r, c) is a valid position in the matrix.

  Example 1:
    Input: mat = [[1,2,3],[4,5,6],[7,8,9]], K = 1
    Output: [[12,21,16],[27,45,33],[24,39,28]]

  Example 2:
    Input: mat = [[1,2,3],[4,5,6],[7,8,9]], K = 2
    Output: [[45,45,45],[45,45,45],[45,45,45]]

  Constraints:
    1. m == mat.length
    2. n == mat[i].length
    3. 1 <= m, n, K <= 100
    4. 1 <= mat[i][j] <= 100

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/matrix-block-sum
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Dynamic Programming
func matrixBlockSum(mat [][]int, K int) [][]int {
	m, n := len(mat), len(mat[0])
	sum := make([][]int, m)
	out := make([][]int, m)
	for i := range out {
		sum[i] = make([]int, n)
		out[i] = make([]int, n)
	}
	for i := range mat {
		for j := range mat[i] {
			sum[i][j] = mat[i][j]
			if i > 0 {
				sum[i][j] += sum[i-1][j]
			}
			if j > 0 {
				sum[i][j] += sum[i][j-1]
			}
			if i > 0 && j > 0 {
				sum[i][j] -= sum[i-1][j-1]
			}
		}
	}
	for i := range out {
		for j := range out[i] {
			a, b := Max(0, i-K), Max(0, j-K)
			c, d := Min(m-1, i+K), Min(n-1, j+K)
			out[i][j] = sum[c][d]
			if a > 0 {
				out[i][j] -= sum[a-1][d]
			}
			if b > 0 {
				out[i][j] -= sum[c][b-1]
			}
			if a > 0 && b > 0 {
				out[i][j] += sum[a-1][b-1]
			}
		}
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
