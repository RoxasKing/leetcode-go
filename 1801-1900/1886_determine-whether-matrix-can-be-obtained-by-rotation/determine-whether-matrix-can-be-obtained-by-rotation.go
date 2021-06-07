package main

/*
  Given two n x n binary matrices mat and target, return true if it is possible to make mat equal to target by rotating mat in 90-degree increments, or false otherwise.

  Example 1:
    Input: mat = [[0,1],[1,0]], target = [[1,0],[0,1]]
    Output: true
    Explanation: We can rotate mat 90 degrees clockwise to make mat equal target.

  Example 2:
    Input: mat = [[0,1],[1,1]], target = [[1,0],[0,1]]
    Output: false
    Explanation: It is impossible to make mat equal to target by rotating mat.

  Example 3:
    Input: mat = [[0,0,0],[0,1,0],[1,1,1]], target = [[1,1,1],[0,1,0],[0,0,0]]
    Output: true
    Explanation: We can rotate mat 90 degrees clockwise two times to make mat equal target.

  Constraints:
    1. n == mat.length == target.length
    2. n == mat[i].length == target[i].length
    3. 1 <= n <= 10
    4. mat[i][j] and target[i][j] are either 0 or 1.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/determine-whether-matrix-can-be-obtained-by-rotation
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func findRotation(mat [][]int, target [][]int) bool {
	n := len(mat)
	for i := 0; i < 4; i++ {
		if equal(mat, target, n) {
			return true
		}
		rotate(mat, n)
	}
	return false
}

func rotate(matrix [][]int, n int) {
	for i := 0; i < n-1-i; i++ {
		for j := i; j < n-1-i; j++ {
			a, b, c, d := matrix[i][j], matrix[j][n-1-i], matrix[n-1-i][n-1-j], matrix[n-1-j][i]
			matrix[i][j], matrix[j][n-1-i], matrix[n-1-i][n-1-j], matrix[n-1-j][i] = d, a, b, c
		}
	}
}

func equal(mat, target [][]int, n int) bool {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] != target[i][j] {
				return false
			}
		}
	}
	return true
}
