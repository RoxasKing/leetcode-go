package main

/*
  Given an m x n matrix mat, return an array of all the elements of the array in a diagonal order.

  Example 1:
    Input: mat = [[1,2,3],[4,5,6],[7,8,9]]
    Output: [1,2,4,7,5,3,6,8,9]

  Example 2:
    Input: mat = [[1,2],[3,4]]
    Output: [1,2,3,4]

  Constraints:
    1. m == mat.length
    2. n == mat[i].length
    3. 1 <= m, n <= 10^4
    4. 1 <= m * n <= 10^4
    5. -10^5 <= mat[i][j] <= 10^5

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/diagonal-traverse
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

func findDiagonalOrder(mat [][]int) []int {
	m, n := len(mat), len(mat[0])
	out := make([]int, 0, m*n)
	i, j := 0, 0
	x := 1
	for len(out) < m*n {
		out = append(out, mat[i][j])
		if 0 <= i-x && i-x < m && 0 <= j+x && j+x < n {
			i += -x
			j += x
			continue
		}
		x = -x
		if i == 0 {
			if j < n-1 {
				j++
			} else {
				i++
			}
		} else {
			if i < m-1 {
				i++
			} else {
				j++
			}
		}
	}
	return out
}
