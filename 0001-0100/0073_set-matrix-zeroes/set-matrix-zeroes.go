package main

/*
  Given an m x n matrix. If an element is 0, set its entire row and column to 0. Do it in-place.

  Follow up:
    1. A straight forward solution using O(mn) space is probably a bad idea.
    2. A simple improvement uses O(m + n) space, but still not the best solution.
    3. Could you devise a constant space solution?

  Example 1:
    Input: matrix = [[1,1,1],[1,0,1],[1,1,1]]
    Output: [[1,0,1],[0,0,0],[1,0,1]]

  Example 2:
    Input: matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
    Output: [[0,0,0,0],[0,4,5,0],[0,3,1,0]]

  Constraints:
    1. m == matrix.length
    2. n == matrix[0].length
    3. 1 <= m, n <= 200
    4. -2^31 <= matrix[i][j] <= 2^31 - 1

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/set-matrix-zeroes
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	col0 := false
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 { // 记录第一列是否要置0
			col0 = true
		}
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0 // 第一列用来记录当前行是否要置0
				matrix[0][j] = 0 // 第一行用来记录当前列是否要置0
			}
		}
	}
	for i := m - 1; i >= 0; i-- {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
		if col0 {
			matrix[i][0] = 0
		}
	}
}
