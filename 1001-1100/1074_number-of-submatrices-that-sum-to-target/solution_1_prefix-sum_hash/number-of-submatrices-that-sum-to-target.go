package main

/*
  Given a matrix and a target, return the number of non-empty submatrices that sum to target.

  A submatrix x1, y1, x2, y2 is the set of all cells matrix[x][y] with x1 <= x <= x2 and y1 <= y <= y2.

  Two submatrices (x1, y1, x2, y2) and (x1', y1', x2', y2') are different if they have some coordinate that is different: for example, if x1 != x1'.

  Example 1:
    Input: matrix = [[0,1,0],[1,1,1],[0,1,0]], target = 0
    Output: 4
    Explanation: The four 1x1 submatrices that only contain 0.

  Example 2:
    Input: matrix = [[1,-1],[-1,1]], target = 0
    Output: 5
    Explanation: The two 1x2 submatrices, plus the two 2x1 submatrices, plus the 2x2 submatrix.

  Example 3:
    Input: matrix = [[904]], target = 0
    Output: 0

  Constraints:
    1 <= matrix.length <= 100
    1 <= matrix[0].length <= 100
    -1000 <= matrix[i] <= 1000
    -10^8 <= target <= 10^8

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/number-of-submatrices-that-sum-to-target
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Prefix Sum
// Hash

func numSubmatrixSumTarget(matrix [][]int, target int) int {
	m, n := len(matrix), len(matrix[0])
	pSum := make([][]int, m)
	for i := 0; i < m; i++ {
		pSum[i] = make([]int, n)
		for j := 0; j < n; j++ {
			pSum[i][j] = matrix[i][j]
			if i > 0 {
				pSum[i][j] += pSum[i-1][j]
			}
			if j > 0 {
				pSum[i][j] += pSum[i][j-1]
			}
			if i > 0 && j > 0 {
				pSum[i][j] -= pSum[i-1][j-1]
			}
		}
	}

	out := 0

	for r1 := 0; r1 < m; r1++ {
		for r2 := r1; r2 < m; r2++ {
			dict := map[int]int{0: 1}
			for c := 0; c < n; c++ {
				sum := pSum[r2][c]
				if r1 > 0 {
					sum -= pSum[r1-1][c]
				}
				if cnt, ok := dict[sum-target]; ok {
					out += cnt
				}
				dict[sum]++
			}
		}
	}

	return out
}
