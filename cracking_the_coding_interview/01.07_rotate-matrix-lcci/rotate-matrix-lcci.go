package main

/*
  Given an image represented by an N x N matrix, where each pixel in the image is 4 bytes, write a method to rotate the image by 90 degrees. Can you do this in place?

  Example 1:
    Given matrix =
    [
      [1,2,3],
      [4,5,6],
      [7,8,9]
    ],
    Rotate the matrix in place. It becomes:
    [
      [7,4,1],
      [8,5,2],
      [9,6,3]
    ]

  Example 2:
    Given matrix =
    [
      [ 5, 1, 9,11],
      [ 2, 4, 8,10],
      [13, 3, 6, 7],
      [15,14,12,16]
    ],
    Rotate the matrix in place. It becomes:
    [
      [15,13, 2, 5],
      [14, 3, 4, 1],
      [12, 6, 8, 9],
      [16, 7,10,11]
    ]

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/rotate-matrix-lcci
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n>>1; i++ {
		for j := i; j < n-1-i; j++ {
			a, b, c, d := matrix[i][j], matrix[j][n-1-i], matrix[n-1-i][n-1-j], matrix[n-1-j][i]
			matrix[i][j], matrix[j][n-1-i], matrix[n-1-i][n-1-j], matrix[n-1-j][i] = d, a, b, c
		}
	}
}
