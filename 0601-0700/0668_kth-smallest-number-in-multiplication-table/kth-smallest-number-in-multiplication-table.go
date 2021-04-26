package main

/*
  Nearly every one have used the Multiplication Table. But could you find out the k-th smallest number quickly from the multiplication table?

  Given the height m and the length n of a m * n Multiplication Table, and a positive integer k, you need to return the k-th smallest number in this table.

  Example 1:
    Input: m = 3, n = 3, k = 5
    Output:
    Explanation:
      The Multiplication Table:
      1	2	3
      2	4	6
      3	6	9
      The 5-th smallest number is 3 (1, 2, 2, 3, 3).

  Example 2:
    Input: m = 2, n = 3, k = 6
    Output:
    Explanation:
      The Multiplication Table:
      1	2	3
      2	4	6
      The 6-th smallest number is 6 (1, 2, 2, 3, 4, 6).

  Note:
    1. The m and n will be in the range [1, 30000].
    2. The k will be in the range [1, m * n]

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/kth-smallest-number-in-multiplication-table
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Binary Search
func findKthNumber(m int, n int, k int) int {
	l, r := 1, m*n
	for l < r {
		limit := (l + r) >> 1
		if countNoBigger(m, n, limit) < k {
			l = limit + 1
		} else {
			r = limit
		}
	}
	return l
}

func countNoBigger(m, n, limit int) int {
	out := 0
	row, col := m, 1
	for row >= 1 && col <= n {
		if row*col <= limit {
			out += row
			col++
		} else {
			row--
		}
	}
	return out
}
