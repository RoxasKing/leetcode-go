package main

/*
  In MATLAB, there is a very useful function called 'reshape', which can reshape a matrix into a new one with different size but keep its original data.

  You're given a matrix represented by a two-dimensional array, and two positive integers r and c representing the row number and column number of the wanted reshaped matrix, respectively.

  The reshaped matrix need to be filled with all the elements of the original matrix in the same row-traversing order as they were.

  If the 'reshape' operation with given parameters is possible and legal, output the new reshaped matrix; Otherwise, output the original matrix.

  Example 1:
    Input:
      nums =
      [[1,2],
       [3,4]]
      r = 1, c = 4
    Output:
      [[1,2,3,4]]
    Explanation:
      The row-traversing of nums is [1,2,3,4]. The new reshaped matrix is a 1 * 4 matrix, fill it row by row by using the previous list.

  Example 2:
    Input:
      nums =
      [[1,2],
       [3,4]]
      r = 2, c = 4
    Output:
      [[1,2],
       [3,4]]
    Explanation:
      There is no way to reshape a 2 * 2 matrix to a 2 * 4 matrix. So output the original matrix.

  Note:
    1. The height and width of the given matrix is in range [1, 100].
    2. The given r and c are all positive.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/reshape-the-matrix
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func matrixReshape(nums [][]int, r int, c int) [][]int {
	m, n := len(nums), len(nums[0])
	if m*n != r*c {
		return nums
	}

	out := make([][]int, r)
	for i := range out {
		out[i] = make([]int, c)
	}

	i, j := 0, 0
	for _, arr := range nums {
		for _, num := range arr {
			out[i][j] = num
			j++
			if j == c {
				i++
				j = 0
			}
		}
	}
	return out
}
