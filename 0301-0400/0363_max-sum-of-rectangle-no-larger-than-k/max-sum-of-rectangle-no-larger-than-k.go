package main

/*
  Given an m x n matrix matrix and an integer k, return the max sum of a rectangle in the matrix such that its sum is no larger than k.

  It is guaranteed that there will be a rectangle with a sum no larger than k.

  Example 1:
    Input: matrix = [[1,0,1],[0,-2,3]], k = 2
    Output: 2
    Explanation: Because the sum of the blue rectangle [[0, 1], [-2, 3]] is 2, and 2 is the max number no larger than k (k = 2).

  Example 2:
    Input: matrix = [[2,2,-1]], k = 3
    Output: 3

  Constraints:
    1. m == matrix.length
    2. n == matrix[i].length
    3. 1 <= m, n <= 100
    4. -100 <= matrix[i][j] <= 100
    5. -10^5 <= k <= 10^5

  Follow up: What if the number of rows is much larger than the number of columns?

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/max-sum-of-rectangle-no-larger-than-k
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func maxSumSubmatrix(matrix [][]int, k int) int {
	out := -1 << 31
	m, n := len(matrix), len(matrix[0])
	for u := 0; u < m; u++ { // upper boundary
		sumC := make([]int, n)
		for i := u; i < m; i++ { // lower boundary
			for j := 0; j < n; j++ {
				sumC[j] += matrix[i][j]
			}

			sum := sumC[0]
			maxS := sum // max sum of current rectangle
			for j := 1; j < n; j++ {
				if sum > 0 {
					sum += sumC[j]
				} else {
					sum = sumC[j]
				}
				maxS = Max(maxS, sum)
			}
			if maxS < k {
				out = Max(out, maxS)
				continue
			} else if maxS == k {
				return k
			}

			for l := 0; l < n; l++ { // left boundary
				sum := 0
				for j := l; j < n; j++ { // right boundary
					sum += sumC[j]
					if sum < k {
						out = Max(out, sum)
					} else if sum == k {
						return k
					}
				}
			}
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
