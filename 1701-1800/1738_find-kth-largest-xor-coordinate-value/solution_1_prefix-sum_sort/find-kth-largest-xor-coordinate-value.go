package main

import "sort"

/*
  You are given a 2D matrix of size m x n, consisting of non-negative integers. You are also given an integer k.

  The value of coordinate (a, b) of the matrix is the XOR of all matrix[i][j] where 0 <= i <= a < m and 0 <= j <= b < n (0-indexed).

  Find the kth largest value (1-indexed) of all the coordinates of matrix.

  Example 1:
    Input: matrix = [[5,2],[1,6]], k = 1
    Output: 7
    Explanation: The value of coordinate (0,1) is 5 XOR 2 = 7, which is the largest value.

  Example 2:
    Input: matrix = [[5,2],[1,6]], k = 2
    Output: 5
    Explanation: The value of coordinate (0,0) is 5 = 5, which is the 2nd largest value.

  Example 3:
    Input: matrix = [[5,2],[1,6]], k = 3
    Output: 4
    Explanation: The value of coordinate (1,0) is 5 XOR 1 = 4, which is the 3rd largest value.

  Example 4:
    Input: matrix = [[5,2],[1,6]], k = 4
    Output: 0
    Explanation: The value of coordinate (1,1) is 5 XOR 2 XOR 1 XOR 6 = 0, which is the 4th largest value.

  Constraints:
    1. m == matrix.length
    2. n == matrix[i].length
    3. 1 <= m, n <= 1000
    4. 0 <= matrix[i][j] <= 10^6
    5. 1 <= k <= m * n

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/find-kth-largest-xor-coordinate-value
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Bit Manipulation
// Prefix Sum
// Sort

func kthLargestValue(matrix [][]int, k int) int {
	m, n := len(matrix), len(matrix[0])
	cnt := map[int]int{}
	f := make([][]int, m)
	for i := 0; i < m; i++ {
		f[i] = make([]int, n)
		for j := 0; j < n; j++ {
			f[i][j] = matrix[i][j]
			if i > 0 {
				f[i][j] ^= f[i-1][j]
			}
			if j > 0 {
				f[i][j] ^= f[i][j-1]
			}
			if i > 0 && j > 0 {
				f[i][j] ^= f[i-1][j-1]
			}
			cnt[f[i][j]]++
		}
	}

	arr := make([]int, 0, len(cnt))
	for num := range cnt {
		arr = append(arr, num)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))

	var out int
	for k > 0 {
		out = arr[0]
		k -= cnt[arr[0]]
		arr = arr[1:]
	}
	return out
}
