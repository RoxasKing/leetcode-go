package main

/*
  Write an efficient algorithm that searches for a target value in an m x n integer matrix. The matrix has the following properties:
    1. Integers in each row are sorted in ascending from left to right.
    2. Integers in each column are sorted in ascending from top to bottom.

  Example 1:
    Input: matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 5
    Output: true

  Example 2:
    Input: matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 20
    Output: false

  Constraints:
    1. m == matrix.length
    2. n == matrix[i].length
    3. 1 <= n, m <= 300
    4. -10^9 <= matix[i][j] <= 10^9
    5. All the integers in each row are sorted in ascending order.
    6. All the integers in each column are sorted in ascending order.
    7. -10^9 <= target <= 10^9

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/search-a-2d-matrix-ii
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	i, j := m-1, 0
	for i >= 0 && j < n {
		if matrix[i][j] < target {
			j++
		} else if matrix[i][j] > target {
			i--
		} else {
			return true
		}
	}
	return false
}
