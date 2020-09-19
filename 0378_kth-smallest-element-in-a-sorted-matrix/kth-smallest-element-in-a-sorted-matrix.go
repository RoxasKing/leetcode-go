package main

/*
  给定一个 n x n 矩阵，其中每行和每列元素均按升序排序，找到矩阵中第 k 小的元素。
  请注意，它是排序后的第 k 小元素，而不是第 k 个不同的元素。

  提示：
    你可以假设 k 的值永远是有效的，1 ≤ k ≤ n2 。
*/

// Binary Search
func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	l, r := matrix[0][0], matrix[n-1][n-1]
	for l < r {
		m := l + (r-l)>>1
		if !check(matrix, k, m) {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}

func check(matrix [][]int, k, target int) bool {
	n := len(matrix)
	row, col := n-1, 0
	var count int
	for row >= 0 && col < n {
		if matrix[row][col] <= target {
			count += row + 1
			col++
		} else {
			row--
		}
	}
	return count >= k
}
