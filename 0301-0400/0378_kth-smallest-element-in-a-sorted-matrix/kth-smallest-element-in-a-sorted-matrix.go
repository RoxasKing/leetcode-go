package main

/*
  Given an n x n matrix where each of the rows and columns are sorted in ascending order, return the kth smallest element in the matrix.

  Note that it is the kth smallest element in the sorted order, not the kth distinct element.

  Example 1:
    Input: matrix = [[1,5,9],[10,11,13],[12,13,15]], k = 8
    Output: 13
    Explanation: The elements in the matrix are [1,5,9,10,11,12,13,13,15], and the 8th smallest number is 13

  Example 2:
    Input: matrix = [[-5]], k = 1
    Output: -5

  Constraints:
    1. n == matrix.length
    2. n == matrix[i].length
    3. 1 <= n <= 300
    4. -10^9 <= matrix[i][j] <= 10^9
    5. All the rows and columns of matrix are guaranteed to be sorted in non-decreasing order.
    6. 1 <= k <= n2

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/kth-smallest-element-in-a-sorted-matrix
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Binary Search
func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	l, r := matrix[0][0], matrix[n-1][n-1]
	for l < r {
		limit := (l + r) >> 1
		if countSmaller(matrix, n, limit) < k {
			l = limit + 1
		} else {
			r = limit
		}
	}
	return l
}

func countSmaller(matrix [][]int, n, limit int) int {
	cnt := 0
	for i, j := n-1, 0; i >= 0 && j < n; {
		if matrix[i][j] <= limit {
			cnt += i + 1
			j++
		} else {
			i--
		}
	}
	return cnt
}
