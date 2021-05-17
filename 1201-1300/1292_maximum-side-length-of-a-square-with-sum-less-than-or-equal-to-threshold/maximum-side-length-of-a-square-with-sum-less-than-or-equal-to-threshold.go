package main

/*
  Given a m x n matrix mat and an integer threshold. Return the maximum side-length of a square with a sum less than or equal to threshold or return 0 if there is no such square.

  Example 1:
    Input: mat = [[1,1,3,2,4,3,2],[1,1,3,2,4,3,2],[1,1,3,2,4,3,2]], threshold = 4
    Output: 2
    Explanation: The maximum side length of square with sum less than 4 is 2 as shown.

  Example 2:
    Input: mat = [[2,2,2,2,2],[2,2,2,2,2],[2,2,2,2,2],[2,2,2,2,2],[2,2,2,2,2]], threshold = 1
    Output: 0

  Example 3:
    Input: mat = [[1,1,1,1],[1,0,0,0],[1,0,0,0],[1,0,0,0]], threshold = 6
    Output: 3

  Example 4:
    Input: mat = [[18,70],[61,1],[25,85],[14,40],[11,96],[97,96],[63,45]], threshold = 40184
    Output: 2

  Constraints:
    1. 1 <= m, n <= 300
    2. m == mat.length
    3. n == mat[i].length
    4. 0 <= mat[i][j] <= 10000
    5. 0 <= threshold <= 10^5

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/maximum-side-length-of-a-square-with-sum-less-than-or-equal-to-threshold
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Binary Search

func maxSideLength(mat [][]int, threshold int) int {
	m, n := len(mat), len(mat[0])
	pSum := make([][]int, m+1)
	for i := range pSum {
		pSum[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			pSum[i+1][j+1] = mat[i][j] + pSum[i][j+1] + pSum[i+1][j] - pSum[i][j]
		}
	}

	out := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			l, r := 0, Min(m-i, n-j)
			for l < r {
				b := (l + r + 1) >> 1
				sum := pSum[i+b][j+b] - pSum[i+b][j] - pSum[i][j+b] + pSum[i][j]
				if sum <= threshold {
					l = b
				} else {
					r = b - 1
				}
			}
			out = Max(out, l)
		}
	}
	return out
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
