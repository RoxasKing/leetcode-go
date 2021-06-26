package main

import "sort"

/*
  You are given a binary matrix matrix of size m x n, and you are allowed to rearrange the columns of the matrix in any order.

  Return the area of the largest submatrix within matrix where every element of the submatrix is 1 after reordering the columns optimally.

  Example 1:
    Input: matrix = [[0,0,1],[1,1,1],[1,0,1]]
    Output: 4
    Explanation: You can rearrange the columns as shown above.
      The largest submatrix of 1s, in bold, has an area of 4.

  Example 2:
    Input: matrix = [[1,0,1,0,1]]
    Output: 3
    Explanation: You can rearrange the columns as shown above.
      The largest submatrix of 1s, in bold, has an area of 3.

  Example 3:
    Input: matrix = [[1,1,0],[1,0,1]]
    Output: 2
    Explanation: Notice that you must rearrange entire columns, and there is no way to make a submatrix of 1s larger than an area of 2.

  Example 4:
    Input: matrix = [[0,0],[0,0]]
    Output: 0
    Explanation: As there are no 1s, no submatrix of 1s can be formed and the area is 0.

  Constraints:
    1. m == matrix.length
    2. n == matrix[i].length
    3. 1 <= m * n <= 10^5
    4. matrix[i][j] is 0 or 1.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/largest-submatrix-with-rearrangements
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Greedy Algorithm
// Sort

func largestSubmatrix(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	s := make([][]int, m)
	for i := 0; i < m; i++ {
		s[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				continue
			}
			s[i][j] = 1
			if i > 0 && matrix[i-1][j] == 1 {
				s[i][j] += s[i-1][j]
			}
		}
	}
	out := 0
	for i := 0; i < m; i++ {
		arr := make([]int, n)
		copy(arr, s[i])
		sort.Sort(sort.Reverse(sort.IntSlice(arr)))
		for j := 0; j < n; j++ {
			out = Max(out, arr[j]*(j+1))
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
