package main

/*
  Write an algorithm such that if an element in an MxN matrix is 0, its entire row and column are set to 0.

  Example 1:
    Input:
      [
        [1,1,1],
        [1,0,1],
        [1,1,1]
      ]
    Output:
      [
        [1,0,1],
        [0,0,0],
        [1,0,1]
      ]

  Example 2:
    Input:
      [
        [0,1,2,0],
        [3,4,5,2],
        [1,3,1,5]
      ]
    Output:
      [
        [0,0,0,0],
        [0,4,5,0],
        [0,3,1,0]
      ]

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/zero-matrix-lcci
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	col0 := false
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			col0 = true
		}
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
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
